package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/champion19/Flighthours_backend/config"
)

type ResendNotifier interface {
	SendVerificationEmail(email, verificationLink string) error
}

type resendNotifier struct {
	apiKey    string
	fromEmail string
	client    *http.Client
}

type EmailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	HTML    string   `json:"html"`
}

func NewResendNotifier(cfg config.ResendConfig) ResendNotifier {
	return &resendNotifier{
		apiKey:    cfg.APIKey,
		fromEmail: cfg.FromEmail,
		client:    &http.Client{},
	}
}

func (r *resendNotifier) SendVerificationEmail(email, verificationLink string) error {
	htmlContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Verificación de Email</title>
		</head>
		<body>
			<h2>Verificación de Email</h2>
			<p>Hola,</p>
			<p>Gracias por registrarte en Flight Hours. Para completar tu registro, por favor verifica tu dirección de email haciendo clic en el siguiente enlace:</p>
			<p><a href="%s" style="background-color: #4CAF50; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px;">Verificar Email</a></p>
			<p>Si no puedes hacer clic en el botón, copia y pega el siguiente enlace en tu navegador:</p>
			<p>%s</p>
			<p>Este enlace expirará en 24 horas.</p>
			<p>Si no creaste esta cuenta, puedes ignorar este email.</p>
			<p>Saludos,<br>El equipo de Flight Hours</p>
		</body>
		</html>
	`, verificationLink, verificationLink)

	emailReq := EmailRequest{
		From:    r.fromEmail,
		To:      []string{email},
		Subject: "Verifica tu dirección de email",
		HTML:    htmlContent,
	}

	jsonData, err := json.Marshal(emailReq)
	if err != nil {
		return fmt.Errorf("failed to marshal email request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.resend.com/emails", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.apiKey)

	resp, err := r.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	defer func() {
    if cerr := resp.Body.Close(); cerr != nil {
        log.Printf("failed to close response body: %v", cerr)
    }
}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("resend API returned status: %d", resp.StatusCode)
	}

	return nil
}
