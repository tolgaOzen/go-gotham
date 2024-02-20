package defs

import (
	"gotham/server"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
)

var ServersDefs = []dingo.Def{
	{
		Name:  "echo-server",
		Scope: di.App,
		Build: func() (ES *server.EchoServer, err error) {
			return &server.EchoServer{}, nil
		},
	},
}
