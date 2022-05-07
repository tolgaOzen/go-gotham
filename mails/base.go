package mails

import (
	"github.com/jordan-wright/email"
)

type IMailRenderer interface {
	Render(data map[string]interface{}, to []string) (context email.Email, err error)
}
