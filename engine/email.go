package engine

import (
	"fmt"
	"helia/shared"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
)

// Helper function to send an email
func sendEmail(from string, to string, subject string, body string) *error {
	// set up sendgrid
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

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
						"type": "text/plain",
						"value": "%v"
					}
				]
			}`,
			to, subject, from, body,
		),
	)

	// send email
	_, err := sendgrid.API(request)

	// check for error
	if err != nil {
		// log error
		shared.TeeLog(
			fmt.Sprintf("Error sending email from %v to %v: %v", from, to, err.Error()),
		)
	}

	// return error, if any
	return &err
}
