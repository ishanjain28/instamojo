package instamojo

import (
	"time"
)

type Config struct {
	ApiKey      string
	AuthToken   string
	SandboxMode bool
	endpoint    string
}

type PaymentURLRequest struct {
	Purpose               string `json:"purpose"`
	Amount                int    `json:"amount"`
	Phone                 string `json:"phone"`
	BuyerName             string `json:"buyer_name"`
	RedirectURL           string `json:"redirect_url"`
	SendEmail             bool   `json:"send_email"`
	Webhook               string `json:"webhook"`
	SendSms               bool   `json:"send_sms"`
	Email                 string `json:"email"`
	AllowRepeatedPayments bool   `json:"allow_repeated_payments"`
}

type unauthorized struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type badrequest struct {
	Success bool                     `json:"success"`
	Message map[string][]interface{} `json:"message"`
}

func (u unauthorized) Error() string {
	return u.Message
}

func (b badrequest) Error() string {

	for _, v := range b.Message {
		return v[0].(string)
	}
	return "instamojo: bad request"
}

type paymenturlok struct {
	PaymentRequest struct {
		ID                    string    `json:"id"`
		Phone                 string    `json:"phone"`
		Email                 string    `json:"email"`
		BuyerName             string    `json:"buyer_name"`
		Amount                string    `json:"amount"`
		Purpose               string    `json:"purpose"`
		Status                string    `json:"status"`
		SendSms               bool      `json:"send_sms"`
		SendEmail             bool      `json:"send_email"`
		SmsStatus             string    `json:"sms_status"`
		EmailStatus           string    `json:"email_status"`
		Shorturl              string    `json:"shorturl"`
		Longurl               string    `json:"longurl"`
		RedirectURL           string    `json:"redirect_url"`
		Webhook               string    `json:"webhook"`
		CreatedAt             time.Time `json:"created_at"`
		ModifiedAt            time.Time `json:"modified_at"`
		AllowRepeatedPayments bool      `json:"allow_repeated_payments"`
	} `json:"payment_request"`
	Success bool `json:"success"`
}

type WebhookResponse struct {
	PaymentID        string `json:"payment_id"`
	Status           string `json:"status"`
	Shorturl         string `json:"shorturl"`
	Longurl          string `json:"longurl"`
	Purpose          string `json:"purpose"`
	Amount           string `json:"amount"`
	Fees             string `json:"fees"`
	Currency         string `json:"currency"`
	Buyer            string `json:"buyer"`
	BuyerName        string `json:"buyer_name"`
	BuyerPhone       string `json:"buyer_phone"`
	PaymentRequestID string `json:"payment_request_id"`
	Mac              string `json:"mac"`
}
