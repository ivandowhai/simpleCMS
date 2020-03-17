package services

import (
	"../core"
	"log"
	"net/smtp"
)

func SendConfirmationEmail(to string, code string) {
	config := core.GetSettings()
	body := "Your confirmation link " + config.SiteDomain + "/confirm?email=" + to + "&code=" + code
	sendEmail(to, body)
}

func sendEmail(to string, body string) {
	config := core.GetSettings()
	auth := smtp.PlainAuth("", config.SMTPAccount, config.SMTPPassword, config.SMTPServer)
	err := smtp.SendMail(config.SMTPServer+":"+config.SMTPPort, auth, config.SMTPAccount, []string{to}, []byte(body))
	if err != nil {
		log.Printf("smtp error: %s", err)
	}
}
