package providers

import (
	"github.com/sarulabs/dingo/v4"
	"gotham/app/defs"
)

type Provider struct {
	dingo.BaseProvider
}

func (p *Provider) Load() error {
	if err := p.AddDefSlice(defs.DatabaseServiceDefs); err != nil {
		return err
	}
	return nil
}

