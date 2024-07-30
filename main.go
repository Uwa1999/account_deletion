package main

import (
	"log"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
)

// EmailRequest struct to parse the incoming request
type EmailRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	CID   string `json:"cid"`
	Store string `json:"store"`
}

// sendEmail sends an email using Gmail SMTP
func sendEmail(from, password, to, subject, body string) error {
	smtpServer := "smtp.gmail.com"
	port := "587"

	auth := smtp.PlainAuth("", from, password, smtpServer)

	message := []byte("Subject: " + subject + "\n\n" + body)

	return smtp.SendMail(smtpServer+":"+port, auth, from, []string{to}, message)
}

// emailHandler handles the email sending requests
func emailHandler(c *fiber.Ctx) error {
	var emailReq EmailRequest
	if err := c.BodyParser(&emailReq); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	from := "card.ottokonek@fortress-asya.com" // Replace with your Gmail address
	appPassword := "rmstspbabcewhqei"          // Replace with your app-specific password
	to := "joshua.arban@fortress-asya.com"     // Replace with recipient email address
	subject := "New Submission from " + emailReq.Name
	body := "Name: " + emailReq.Name + "\nPhone: " + emailReq.Phone + "\nCID: " + emailReq.CID + "\nStore: " + emailReq.Store

	if err := sendEmail(from, appPassword, to, subject, body); err != nil {
		log.Printf("Error sending email: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to send email")
	}

	return c.SendString("Email sent successfully")
}

func main() {
	app := fiber.New()

	app.Post("/send-email", emailHandler)

	log.Fatal(app.Listen(":8080"))
}
