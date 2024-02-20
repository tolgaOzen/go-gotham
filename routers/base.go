package routers

import (
	"gotham/app"
	"gotham/app/flags"
)

func Initialize() {
	if *flags.Server {
		echoServer := app.Application.Container.GetEchoServer()
		echo := echoServer.New()
		routes := GetRoute(echo)
		echoServer.StartServer(routes)
		return
	}
}
