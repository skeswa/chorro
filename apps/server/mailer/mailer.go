package mailer

import (
	"fmt"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"github.com/pkg/errors"
	"github.com/skeswa/chorro/apps/server/config"
)

// Interface between the server and the service it uses to send emails.
type Mailer struct {
	// Name of the "person" that emails sent by this server will appear to come
	// from.
	fromName string
	// Email address of the "person" that emails sent by this server will appear
	// to come from.
	fromEmailAddress string
	// Client used authenticated against the SMTP server.
	saslClient sasl.Client
	// Address of the SMTP server.
	smtpServerAddress string
}

// Creates and initializes a new Mailer based on the specified config.
func New(config *config.Config) *Mailer {
	return &Mailer{
		fromName:         config.Mail.UserName,
		fromEmailAddress: config.Mail.User,
		saslClient: sasl.NewPlainClient(
			"",
			config.Mail.User,
			config.Mail.Password,
		),
		smtpServerAddress: fmt.Sprintf("%s:%d", config.Mail.Host, config.Mail.Port),
	}
}

// Options configuring Mailer#Send(...).
type SendOptions struct {
	// Body of the email to be sent.
	Message string
	// Subject of the email to be sent.
	Subject string
	// Email address of the receipient of the email to be sent.
	ToEmailAddress string
}

// Sends an email with the specified contents to the specified email address.
func (mailer *Mailer) Send(options SendOptions) error {
	to := []string{options.ToEmailAddress}
	email := fmt.Sprintf(
		"From: \"%s\" <%s>\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n"+
			"%s\r\n",
		mailer.fromName,
		mailer.fromEmailAddress,
		options.ToEmailAddress,
		options.Subject,
		options.Message,
	)

	if err := smtp.SendMail(
		mailer.smtpServerAddress,
		mailer.saslClient,
		mailer.fromEmailAddress,
		to,
		strings.NewReader(email),
	); err != nil {
		return errors.Wrap(err, "Failed to send email")
	}

	return nil
}
