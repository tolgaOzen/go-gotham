package provider

import (
	"github.com/sarulabs/dingo/v4"
	"gotham/app/defs"
)

type Provider struct {
	dingo.BaseProvider
}

/**
 * Load
 * All the definitions are combined and gathered under one provider. When you create a service definition you need to add here like DatabaseServiceDefs
 */
func (p *Provider) Load() error {
	if err := p.AddDefSlice(defs.DatabaseServiceDefs); err != nil {
		return err
	}

	if err := p.AddDefSlice(defs.UserServiceDefs); err != nil {
		return err
	}

	if err := p.AddDefSlice(defs.AuthServiceDefs); err != nil {
		return err
	}

	if err := p.AddDefSlice(defs.ControllerDefs); err != nil {
		return err
	}

	if err := p.AddDefSlice(defs.MiddlewareDefs); err != nil {
		return err
	}

	return nil
}

