package shared

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
)

// Helper function to send an email
func SendEmail(to string, subject string, body string) error {
	// set up sendgrid
	request := sendgrid.GetRequest(config.SendgridKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	// get from email
	from := config.FromEmail

	// no double quotes allowed
	to = strings.ReplaceAll(to, "\"", "")
	from = strings.ReplaceAll(from, "\"", "")
	subject = strings.ReplaceAll(subject, "\"", "")
	body = strings.ReplaceAll(body, "\"", "")

	// build request body
	request.Body = []byte(
		fmt.Sprintf(
			` {
				"personalizations": [
					{
						"to": [
							{
								"email": "%v"
							}
						],
						"subject": "%v"
					}
				],
				"from": {
					"email": "%v"
				},
				"content": [
					{
						"type": "text/html",
						"value": "%v"
					}
				]
			}`,
			to, subject, from, body,
		),
	)

	// send email
	res, err := sendgrid.API(request)

	// check for error
	if err != nil {
		// log error
		TeeLog(
			fmt.Sprintf("Error sending email from %v to %v: %v", from, to, err.Error()),
		)
	}

	// check for error in response itself
	if res.StatusCode != 200 && res.StatusCode != 202 {
		// log error
		TeeLog(
			fmt.Sprintf("Error sending email [2] from %v to %v: %v {%v}", from, to, res.Body, res.StatusCode),
		)

		// store generic error
		err = errors.New("unable to send email")
	}

	// return error, if any
	return err
}

// Helper function to create the body of a password reset token email
func FillPasswordResetTokenEmailBody(frontendDomain string, uid uuid.UUID, token uuid.UUID) string {
	// fill body
	b := fmt.Sprintf(
		`
			<div>
				Click <a href='%v#/auth/reset?u=%v&t=%v'>here</a> to reset your Project Helia account password.
			</div>
			<div>
				If you did not request a password reset, please contact support at <a href="mailto:contact@projecthelia.com">contact@projecthelia.com</a>.
			</div>
		`,
		frontendDomain,
		uid,
		token,
	)

	// strip special characters
	b = strings.ReplaceAll(b, "\t", "")
	b = strings.ReplaceAll(b, "\n", "")

	// return result
	return b
}

// Helper function to create the body of a password reset success email
func FillPasswordResetSuccessEmailBody(ip string) string {
	// fill body
	b := fmt.Sprintf(
		`
			<div>
				Your password for Project Helia was reset from the IP address %v.
			</div>
			<div>
				If you did not request a password reset, please contact support at <a href="mailto:contact@projecthelia.com">contact@projecthelia.com</a> immediately.
			</div>
		`,
		ip,
	)

	// strip special characters
	b = strings.ReplaceAll(b, "\t", "")
	b = strings.ReplaceAll(b, "\n", "")

	// return result
	return b
}
