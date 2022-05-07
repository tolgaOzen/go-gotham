package infrastructures

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"

	"gotham/config"
)

// Email Service

/**
 * EmailService
 *
 * interface
 */
type IEmailService interface {
	Send(Context email.Email) error
}

/**
 * SendGridService
 *
 */
type EmailService struct {
	Config *config.Email
}

func NewEmailService(emailConfig *config.Email) IEmailService {
	return &EmailService{
		Config: emailConfig,
	}
}

/**
 * Send
 *
 */
func (e EmailService) Send(Context email.Email) error {
	return Context.Send(fmt.Sprintf("%v:%v", e.Config.Host, e.Config.Port), smtp.PlainAuth("", e.Config.From, e.Config.Password, e.Config.Host))
}
