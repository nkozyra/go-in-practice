package main

import (
	"bytes"
	"net/smtp"
	"strconv"

	"text/template" // # A
)

type EmailMessage struct { // # B
	From, Subject, Body string   // # B
	To                  []string // # B
} // # B
type EmailCredentials struct {
	Username, Password, Server string
	Port                       int
}

const emailTemplate = `From: {{.From}} // # C
To: {{.To}} // # C
Subject {{.Subject}} // # C
{{.Body}} // # C
` // # C
var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(emailTemplate)
}
func main() {
	message := &EmailMessage{ // # D
		From:    "me@example.com",            // # D
		To:      []string{"you@example.com"}, // # D
		Subject: "A test",                    // # D
		Body:    "Just saying hi",            // # D
	} // # D
	var body bytes.Buffer           // # E
	t.Execute(&body, message)       // # E
	authCreds := &EmailCredentials{ // # F
		Username: "myUsername",       // # F
		Password: "myPass",           // # F
		Server:   "smtp.example.com", // # F
		Port:     25,                 // # F
	} // # F
	auth := smtp.PlainAuth("", // # F
		authCreds.Username, // # F
		authCreds.Password, // # F
		authCreds.Server,   // # F
	) // # F
	smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port), // # G
		auth,
		message.From,
		message.To,
		body.Bytes()) // # H
}
