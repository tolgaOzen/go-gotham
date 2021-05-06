package defs

import (
	"github.com/jordan-wright/email"
	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"
	"gotham/mails"
)

var MailsDefs = []dingo.Def{
	{
		Name:  "user-welcome-mail",
		Scope: di.App,
		Build: func() (welcome mails.IMailRenderer, err error) {
			return mails.NewWelcome(*email.NewEmail()), nil
		},
	},
}
