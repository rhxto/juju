// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package vault

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
	"github.com/juju/errors"
	"github.com/juju/loggo"
	"github.com/juju/names/v4"
	vault "github.com/mittwald/vaultgo"

	"github.com/juju/juju/secrets/provider"
)

var logger = loggo.GetLogger("juju.secrets.vault")

const (
	// BackendType is the type of the Vault secrets backend.
	BackendType = "vault"
)

// NewProvider returns a Vault secrets provider.
func NewProvider() provider.SecretBackendProvider {
	return vaultProvider{}
}

type vaultProvider struct {
}

func (p vaultProvider) Type() string {
	return BackendType
}

// Initialise sets up a kv store mounted on the model uuid.
func (p vaultProvider) Initialise(m provider.Model) error {
	cfg, err := p.adminConfig(m)
	if err != nil {
		return errors.Trace(err)
	}
	client, err := p.newBackend(cfg)
	if err != nil {
		return errors.Trace(err)
	}
	sys := client.client.Sys()
	ctx := context.Background()

	mounts, err := sys.ListMountsWithContext(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	logger.Debugf("kv mounts: %v", mounts)
	modelUUID := m.UUID()
	if _, ok := mounts[modelUUID]; !ok {
		err = sys.MountWithContext(ctx, modelUUID, &api.MountInput{
			Type:    "kv",
			Options: map[string]string{"version": "1"},
		})
		if !isAlreadyExists(err, "path is already in use") {
			return errors.Trace(err)
		}
	}
	return nil
}

// CleanupModel deletes all secrets and policies associated with the model.
func (p vaultProvider) CleanupModel(m provider.Model) error {
	cfg, err := p.adminConfig(m)
	if err != nil {
		return errors.Trace(err)
	}
	k, err := p.newBackend(cfg)
	if err != nil {
		return errors.Trace(err)
	}
	sys := k.client.Sys()

	// First remove any policies.
	ctx := context.Background()
	policies, err := sys.ListPoliciesWithContext(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	for _, p := range policies {
		if strings.HasPrefix(p, "model-"+k.modelUUID) {
			if err := sys.DeletePolicyWithContext(ctx, p); err != nil {
				if isNotFound(err) {
					continue
				}
				return errors.Annotatef(err, "deleting policy %q", p)
			}
		}
	}

	// Now remove any secrets.
	s, err := k.client.Logical().ListWithContext(ctx, k.modelUUID)
	if err != nil {
		return errors.Trace(err)
	}
	if s == nil || s.Data == nil {
		return nil
	}
	keys, ok := s.Data["keys"].([]interface{})
	if !ok {
		return nil
	}
	for _, id := range keys {
		err = k.client.KVv1(k.modelUUID).Delete(ctx, fmt.Sprintf("%s", id))
		if err != nil && !isNotFound(err) {
			return errors.Annotatef(err, "deleting secret %q", id)
		}
	}
	return nil
}

// CleanupSecrets removes policies associated with the removed secrets.
func (p vaultProvider) CleanupSecrets(m provider.Model, tag names.Tag, removed provider.SecretRevisions) error {
	cfg, err := p.adminConfig(m)
	if err != nil {
		return errors.Trace(err)
	}
	client, err := p.newBackend(cfg)
	if err != nil {
		return errors.Trace(err)
	}
	sys := client.client.Sys()

	isRelevantPolicy := func(p string) bool {
		for id := range removed {
			if strings.HasPrefix(p, fmt.Sprintf("model-%s-%s-", m.UUID(), id)) {
				return true
			}
		}
		return false
	}

	ctx := context.Background()
	policies, err := sys.ListPoliciesWithContext(ctx)
	if err != nil {
		return errors.Trace(err)
	}
	for _, p := range policies {
		if isRelevantPolicy(p) {
			if err := sys.DeletePolicyWithContext(ctx, p); err != nil {
				if isNotFound(err) {
					continue
				}
				return errors.Annotatef(err, "deleting policy %q", p)
			}
		}
	}
	return nil
}

// adminConfig returns the config needed to create a vault secrets backend client
// with full admin rights.
func (p vaultProvider) adminConfig(m provider.Model) (*provider.BackendConfig, error) {
	b, err := m.GetSecretBackend()
	if err != nil {
		return nil, errors.Trace(err)
	}
	backendCfg := &provider.BackendConfig{
		ControllerUUID: m.ControllerUUID(),
		ModelUUID:      m.UUID(),
		ModelName:      m.Name(),
		BackendType:    BackendType,
		Config:         make(map[string]interface{}),
	}
	for k, v := range b.Config {
		backendCfg.Config[k] = v
	}
	return backendCfg, nil
}

// BackendConfig returns the config needed to create a vault secrets backend client.
func (p vaultProvider) BackendConfig(m provider.Model, tag names.Tag, owned provider.SecretRevisions, read provider.SecretRevisions) (*provider.BackendConfig, error) {
	adminUser := tag == nil
	// Get an admin backend client so we can set up the policies.
	backendCfg, err := p.adminConfig(m)
	if err != nil {
		return nil, errors.Trace(err)
	}
	backend, err := p.newBackend(backendCfg)
	if err != nil {
		return nil, errors.Trace(err)
	}
	sys := backend.client.Sys()

	ctx := context.Background()
	modelUUID := backendCfg.ModelUUID
	var policies []string
	if adminUser {
		// For admin users, all secrets for the model can be read.
		rule := fmt.Sprintf(`path "%s/*" {capabilities = ["read"]}`, modelUUID)
		policyName := fmt.Sprintf("model-%s-read", modelUUID)
		err = sys.PutPolicyWithContext(ctx, policyName, rule)
		if err != nil {
			return nil, errors.Annotatef(err, "creating read policy for model %q", modelUUID)
		}
		policies = append(policies, policyName)
	} else {
		// Agents can create new secrets in the model.
		rule := fmt.Sprintf(`path "%s/*" {capabilities = ["create"]}`, modelUUID)
		policyName := fmt.Sprintf("model-%s-create", modelUUID)
		err = sys.PutPolicyWithContext(ctx, policyName, rule)
		if err != nil {
			return nil, errors.Annotatef(err, "creating create policy for model %q", modelUUID)
		}
		policies = append(policies, policyName)
	}
	// Any secrets owned by the agent can be updated/deleted etc.
	logger.Debugf("owned secrets: %#v", owned)
	for id := range owned {
		rule := fmt.Sprintf(`path "%s/%s-*" {capabilities = ["create", "read", "update", "delete", "list"]}`, modelUUID, id)
		policyName := fmt.Sprintf("model-%s-%s-owner", modelUUID, id)
		err = sys.PutPolicyWithContext(ctx, policyName, rule)
		if err != nil {
			return nil, errors.Annotatef(err, "creating owner policy for %q", id)
		}
		policies = append(policies, policyName)
	}

	// Any secrets consumed by the agent can be read etc.
	logger.Debugf("consumed secrets: %#v", read)
	for id := range read {
		rule := fmt.Sprintf(`path "%s/%s-*" {capabilities = ["read"]}`, modelUUID, id)
		policyName := fmt.Sprintf("model-%s-%s-read", modelUUID, id)
		err = sys.PutPolicyWithContext(ctx, policyName, rule)
		if err != nil {
			return nil, errors.Annotatef(err, "creating read policy for %q", id)
		}
		policies = append(policies, policyName)
	}
	s, err := backend.client.Auth().Token().Create(&api.TokenCreateRequest{
		TTL:             "10m", // 10 minutes for now, can configure later.
		NoDefaultPolicy: true,
		Policies:        policies,
	})
	if err != nil {
		return nil, errors.Annotate(err, "creating secret access token")
	}
	backendCfg.Config[TokenKey] = s.Auth.ClientToken

	return backendCfg, nil
}

// NewVaultClient is patched for testing.
var NewVaultClient = vault.NewClient

// NewBackend returns a vault backed secrets backend client.
func (p vaultProvider) NewBackend(cfg *provider.BackendConfig) (provider.SecretsBackend, error) {
	return p.newBackend(cfg)
}

func (p vaultProvider) newBackend(cfg *provider.BackendConfig) (*vaultBackend, error) {
	validCfg, err := newConfig(cfg.Config)
	if err != nil {
		return nil, errors.Annotatef(err, "invalid vault config")
	}

	var clientCertPath, clientKeyPath string
	clientCert := validCfg.clientCert()
	clientKey := validCfg.clientKey()
	if clientCert != "" {
		clientCertFile, err := os.CreateTemp("", "client-cert")
		if err != nil {
			return nil, errors.Annotate(err, "creating client cert file")
		}
		defer func() { _ = clientCertFile.Close() }()
		clientCertPath = clientCertFile.Name()
		if _, err := clientCertFile.Write([]byte(clientCert)); err != nil {
			return nil, errors.Annotate(err, "writing client cert file")
		}

		clientKeyFile, err := os.CreateTemp("", "client-key")
		if err != nil {
			return nil, errors.Annotate(err, "creating client key file")
		}
		defer func() { _ = clientKeyFile.Close() }()
		clientKeyPath = clientKeyFile.Name()
		if _, err := clientKeyFile.Write([]byte(clientKey)); err != nil {
			return nil, errors.Annotate(err, "writing client key file")
		}
	}

	tlsConfig := vault.TLSConfig{
		TLSConfig: &api.TLSConfig{
			CACertBytes:   []byte(validCfg.caCert()),
			ClientCert:    clientCertPath,
			ClientKey:     clientKeyPath,
			TLSServerName: validCfg.tlsServerName(),
		},
	}
	c, err := NewVaultClient(validCfg.endpoint(),
		&tlsConfig,
		vault.WithAuthToken(validCfg.token()),
	)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if ns := validCfg.namespace(); ns != "" {
		c.SetNamespace(ns)
	}
	return &vaultBackend{modelUUID: cfg.ModelUUID, client: c}, nil
}
