package infrastructures

import (
	"fmt"
	"github.com/jordan-wright/email"
	"gotham/config"
	"net/smtp"
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
	Config  *config.EmailConfig
}

func NewEmailService(emailConfig *config.EmailConfig) IEmailService {
	return &EmailService{
		Config:  emailConfig,
	}
}

/**
 * Send
 *
 */
func (e EmailService) Send(Context email.Email) error {
	return Context.Send(fmt.Sprintf("%v:%v", e.Config.Host, e.Config.Port), smtp.PlainAuth("", e.Config.From, e.Config.Password, e.Config.Host))
}

