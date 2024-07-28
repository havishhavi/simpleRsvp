package handlers

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error Handling env variables")
// 	}
// }

// SMTP_HOST=smtp.example.com
// SMTP_PORT=587
// SMTP_USER=RevanthSindhuHome@gmail.com
// SMTP_PASS=wcab erwi rvyb rowl
// SendEmail sends an email to the specified recipient
// SendEmail sends an email using the configured SMTP server
func SendEmail(to, subject, body string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	// Prepare the email message
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))

	// Set up authentication information.
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	// Set up the TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	// Connect to the SMTP server.
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to dial SMTP server: %w", err)
	}
	defer conn.Close()

	// Create a new SMTP client from the connection.
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Close()

	// Authenticate to the SMTP server.
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate to SMTP server: %w", err)
	}

	// Send the email.
	if err = client.Mail(smtpUser); err != nil {
		return fmt.Errorf("failed to set sender email: %w", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient email: %w", err)
	}

	// Write the email data.
	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to get email data writer: %w", err)
	}
	if _, err = writer.Write(msg); err != nil {
		return fmt.Errorf("failed to write email data: %w", err)
	}
	if err = writer.Close(); err != nil {
		return fmt.Errorf("failed to close email data writer: %w", err)
	}

	return nil
}

// SendEmail sends an email using the configured SMTP server
// func SendEmail(to, subject, body string) error {
// 	smtpHost := os.Getenv("SMTP_HOST")
// 	smtpPort := os.Getenv("SMTP_PORT")
// 	smtpUser := os.Getenv("SMTP_USER")
// 	smtpPass := os.Getenv("SMTP_PASS")

// 	// Prepare the email message
// 	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body))

// 	// Set up authentication information.
// 	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

// 	// Convert port to an integer
// 	port, err := strconv.Atoi(smtpPort)
// 	if err != nil {
// 		return fmt.Errorf("invalid SMTP port: %v", err)
// 	}

// 	// Connect to the SMTP server without TLS
// 	addr := fmt.Sprintf("%s:%d", smtpHost, port)
// 	conn, err := smtp.Dial(addr)
// 	if err != nil {
// 		return fmt.Errorf("failed to dial SMTP server: %w", err)
// 	}
// 	defer conn.Close()

// 	// Authenticate to the SMTP server
// 	if err = conn.Auth(auth); err != nil {
// 		return fmt.Errorf("failed to authenticate to SMTP server: %w", err)
// 	}

// 	// Set the sender and recipient
// 	if err = conn.Mail(smtpUser); err != nil {
// 		return fmt.Errorf("failed to set sender email: %w", err)
// 	}
// 	if err = conn.Rcpt(to); err != nil {
// 		return fmt.Errorf("failed to set recipient email: %w", err)
// 	}

// 	// Write the email data
// 	writer, err := conn.Data()
// 	if err != nil {
// 		return fmt.Errorf("failed to get email data writer: %w", err)
// 	}
// 	if _, err = writer.Write(msg); err != nil {
// 		return fmt.Errorf("failed to write email data: %w", err)
// 	}
// 	if err = writer.Close(); err != nil {
// 		return fmt.Errorf("failed to close email data writer: %w", err)
// 	}

// 	// Close the SMTP client connection
// 	return conn.Quit()
// }
