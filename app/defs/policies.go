package defs

import (
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"gotham/policies"
)

var PoliciesDefs = []dingo.Def{
	{
		Name:  "user-policy",
		Scope: di.App,
		Build: func() (s policies.IUserPolicy, err error) {
			return policies.UserPolicy{}, nil
		},
	},
}
