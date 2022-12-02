// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secretbackends_test

import (
	"time"

	"github.com/golang/mock/gomock"
	"github.com/juju/cmd/v3/cmdtesting"
	jujutesting "github.com/juju/testing"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	apisecretbackends "github.com/juju/juju/api/client/secretbackends"
	"github.com/juju/juju/cmd/juju/secretbackends"
	"github.com/juju/juju/cmd/juju/secretbackends/mocks"
	"github.com/juju/juju/jujuclient"
)

type ListSuite struct {
	jujutesting.IsolationSuite
	store             *jujuclient.MemStore
	secretBackendsAPI *mocks.MockListSecretBackendsAPI
}

var _ = gc.Suite(&ListSuite{})

func (s *ListSuite) SetUpTest(c *gc.C) {
	s.IsolationSuite.SetUpTest(c)
	store := jujuclient.NewMemStore()
	store.Controllers["mycontroller"] = jujuclient.ControllerDetails{}
	store.CurrentControllerName = "mycontroller"
	s.store = store
}

func (s *ListSuite) setup(c *gc.C) *gomock.Controller {
	ctrl := gomock.NewController(c)

	s.secretBackendsAPI = mocks.NewMockListSecretBackendsAPI(ctrl)

	return ctrl
}

func ptr[T any](v T) *T {
	return &v
}

func (s *ListSuite) TestListTabular(c *gc.C) {
	defer s.setup(c).Finish()

	s.secretBackendsAPI.EXPECT().ListSecretBackends(false).Return(
		[]apisecretbackends.SecretBackend{{
			Name:                "myvault",
			BackendType:         "vault",
			TokenRotateInterval: ptr(666 * time.Minute),
			Config:              map[string]interface{}{"endpoint": "http://vault"},
			NumSecrets:          666,
		}, {
			Name:        "internal",
			BackendType: "controller",
			NumSecrets:  668,
		}}, nil)
	s.secretBackendsAPI.EXPECT().Close().Return(nil)

	ctx, err := cmdtesting.RunCommand(c, secretbackends.NewListCommandForTest(s.store, s.secretBackendsAPI))
	c.Assert(err, jc.ErrorIsNil)
	out := cmdtesting.Stdout(ctx)
	c.Assert(out, gc.Equals, `
Name      Type        Secrets
internal  controller  668  
myvault   vault       666  
`[1:])
}

func (s *ListSuite) TestListYAML(c *gc.C) {
	defer s.setup(c).Finish()

	s.secretBackendsAPI.EXPECT().ListSecretBackends(true).Return(
		[]apisecretbackends.SecretBackend{{
			Name:                "myvault",
			BackendType:         "vault",
			TokenRotateInterval: ptr(666 * time.Minute),
			Config:              map[string]interface{}{"endpoint": "http://vault"},
			NumSecrets:          666,
		}, {
			Name:        "internal",
			BackendType: "controller",
			NumSecrets:  668,
		}}, nil)

	s.secretBackendsAPI.EXPECT().Close().Return(nil)

	ctx, err := cmdtesting.RunCommand(c, secretbackends.NewListCommandForTest(s.store, s.secretBackendsAPI), "--reveal", "--format", "yaml")
	c.Assert(err, jc.ErrorIsNil)
	out := cmdtesting.Stdout(ctx)
	c.Assert(out, gc.Equals, `
internal:
  name: internal
  backend: controller
  secrets: 668
myvault:
  name: myvault
  backend: vault
  token-rotate-interval: 11h6m0s
  config:
    endpoint: http://vault
  secrets: 666
`[1:])
}

func (s *ListSuite) TestListJSON(c *gc.C) {
	defer s.setup(c).Finish()

	s.secretBackendsAPI.EXPECT().ListSecretBackends(true).Return(
		[]apisecretbackends.SecretBackend{{
			Name:        "internal",
			BackendType: "controller",
			NumSecrets:  668,
		}}, nil)
	s.secretBackendsAPI.EXPECT().Close().Return(nil)

	ctx, err := cmdtesting.RunCommand(c, secretbackends.NewListCommandForTest(s.store, s.secretBackendsAPI), "--reveal", "--format", "json")
	c.Assert(err, jc.ErrorIsNil)
	out := cmdtesting.Stdout(ctx)
	c.Assert(out, gc.Equals, `
{"internal":{"name":"internal","backend":"controller","secrets":668}}
`[1:])
}