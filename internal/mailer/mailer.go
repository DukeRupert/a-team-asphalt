package mailer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Config holds mailer settings loaded from environment.
type Config struct {
	Token     string // Postmark server token
	From      string // Sender email (must be verified in Postmark)
	To        string // Recipient email for estimate notifications
}

// Mailer sends transactional email via the Postmark API.
type Mailer struct {
	cfg    Config
	client *http.Client
}

// New creates a Mailer. Returns nil if token is empty (allows graceful degradation).
func New(cfg Config) *Mailer {
	if cfg.Token == "" {
		return nil
	}
	return &Mailer{
		cfg: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// EstimateRequest holds the form data for an estimate submission.
type EstimateRequest struct {
	Name        string
	Phone       string
	Email       string
	ProjectType string
	Description string
}

type postmarkMessage struct {
	From     string `json:"From"`
	To       string `json:"To"`
	Subject  string `json:"Subject"`
	HtmlBody string `json:"HtmlBody"`
	TextBody string `json:"TextBody"`
	ReplyTo  string `json:"ReplyTo,omitempty"`
}

// SendEstimateNotification sends a notification email for a new estimate request.
func (m *Mailer) SendEstimateNotification(req EstimateRequest) error {
	subject := fmt.Sprintf("New estimate request: %s — %s", req.Name, req.ProjectType)

	textBody := fmt.Sprintf(
		"New estimate request from the website:\n\n"+
			"Name: %s\n"+
			"Phone: %s\n"+
			"Email: %s\n"+
			"Project type: %s\n"+
			"Description:\n%s\n",
		req.Name, req.Phone, req.Email, req.ProjectType, req.Description,
	)

	htmlBody := fmt.Sprintf(
		`<h2>New Estimate Request</h2>
<table style="border-collapse:collapse;font-family:sans-serif;font-size:14px;">
<tr><td style="padding:6px 12px 6px 0;font-weight:bold;vertical-align:top;">Name</td><td style="padding:6px 0;">%s</td></tr>
<tr><td style="padding:6px 12px 6px 0;font-weight:bold;vertical-align:top;">Phone</td><td style="padding:6px 0;"><a href="tel:%s">%s</a></td></tr>
<tr><td style="padding:6px 12px 6px 0;font-weight:bold;vertical-align:top;">Email</td><td style="padding:6px 0;"><a href="mailto:%s">%s</a></td></tr>
<tr><td style="padding:6px 12px 6px 0;font-weight:bold;vertical-align:top;">Project</td><td style="padding:6px 0;">%s</td></tr>
<tr><td style="padding:6px 12px 6px 0;font-weight:bold;vertical-align:top;">Details</td><td style="padding:6px 0;">%s</td></tr>
</table>`,
		req.Name, req.Phone, req.Phone, req.Email, req.Email, req.ProjectType, req.Description,
	)

	msg := postmarkMessage{
		From:     m.cfg.From,
		To:       m.cfg.To,
		Subject:  subject,
		HtmlBody: htmlBody,
		TextBody: textBody,
	}
	if req.Email != "" {
		msg.ReplyTo = req.Email
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal email: %w", err)
	}

	httpReq, err := http.NewRequest("POST", "https://api.postmarkapp.com/email", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-Postmark-Server-Token", m.cfg.Token)

	resp, err := m.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send email: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("postmark returned %d", resp.StatusCode)
	}

	return nil
}
