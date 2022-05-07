package mails

import (
	"bytes"
	"fmt"

	"github.com/alecthomas/template"
	"github.com/jordan-wright/email"
)

/**
 * Welcome
 *
 * struct
 */
type Welcome struct {
	Type    string
	Context email.Email
}

/**
 * NewWelcome
 *
 * @return *UserWelcome
 */
func NewWelcome(context email.Email) Welcome {
	return Welcome{
		Type:    "-",
		Context: context,
	}
}

/**
 * Render
 *
 * @return infrastructures.IEmailService
 */
func (w Welcome) Render(data map[string]interface{}, to []string) (context email.Email, err error) {
	var t *template.Template
	t, err = template.ParseFiles("views/welcome.html")
	if err != nil {
		return email.Email{}, err
	}
	var body bytes.Buffer
	err = t.Execute(&body, struct {
		Url interface{}
	}{
		Url: data["url"],
	})
	w.Context.From = "Gotham <example@go-gotham.com>"
	w.Context.To = to
	w.Context.Subject = fmt.Sprintf("Welcome to Gotham")
	w.Context.HTML = body.Bytes()
	return w.Context, err
}
