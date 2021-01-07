package providers

import (
	"github.com/sarulabs/dingo/v4"
	"gotham/services"
)

type Provider struct {
	dingo.BaseProvider
}

func (p *Provider) Load() error {
	if err := p.AddDefSlice(services.DatabaseServiceDefs); err != nil {
		return err
	}

	// add providers

	//if err := p.AddDefSlice(services.ExampleServiceDefs); err != nil {
	//	return err
	//}


	return nil
}

