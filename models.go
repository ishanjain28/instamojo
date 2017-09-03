package instamojo

import (
	"time"
)

// Unauthorized Response from Instamojo
type Unauthorized struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// BadRequest Response from Instamojo
type BadRequest struct {
	Success bool                     `json:"success"`
	Message map[string][]interface{} `json:"message"`
}

func (u Unauthorized) Error() string {
	return u.Message
}

func (b BadRequest) Error() string {

	for _, v := range b.Message {
		return v[0].(string)
	}
	return "instamojo: bad request"
}

// Config is the configuration struct that is used in initialising the package
type Config struct {
	APIKey      string
	AuthToken   string
	SandboxMode bool
	endpoint    string
}

// PaymentURLRequest is the information that you need to provide when creating a Payment URL
// Note that all fields are not mandatory. Take a look at Instamojo Documentation for more
// Information at https://docs.instamojo.com
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

// PaymentURLResponse is returned when creating a new payment URL
type PaymentURLResponse struct {
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

// WebhookResponse is the data that Instamojo sends to the webhook
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

//RequestsList is the list of all the requests created so far
type RequestsList struct {
	Success         bool `json:"success"`
	PaymentRequests []struct {
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
	} `json:"payment_requests"`
}

// PaymentRequestDetails is the response that has complete details about a Payment ID
type PaymentRequestDetails struct {
	PaymentRequest struct {
		ID          string `json:"id"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
		BuyerName   string `json:"buyer_name"`
		Amount      string `json:"amount"`
		Purpose     string `json:"purpose"`
		Status      string `json:"status"`
		SendSms     bool   `json:"send_sms"`
		SendEmail   bool   `json:"send_email"`
		SmsStatus   string `json:"sms_status"`
		EmailStatus string `json:"email_status"`
		Shorturl    string `json:"shorturl"`
		Longurl     string `json:"longurl"`
		RedirectURL string `json:"redirect_url"`
		Webhook     string `json:"webhook"`
		Payments    []struct {
			PaymentID         string        `json:"payment_id"`
			Quantity          int           `json:"quantity"`
			Status            string        `json:"status"`
			LinkSlug          string        `json:"link_slug"`
			LinkTitle         string        `json:"link_title"`
			BuyerName         string        `json:"buyer_name"`
			BuyerPhone        string        `json:"buyer_phone"`
			BuyerEmail        string        `json:"buyer_email"`
			Currency          string        `json:"currency"`
			UnitPrice         string        `json:"unit_price"`
			Amount            string        `json:"amount"`
			Fees              string        `json:"fees"`
			ShippingAddress   string        `json:"shipping_address"`
			ShippingCity      string        `json:"shipping_city"`
			ShippingState     string        `json:"shipping_state"`
			ShippingZip       string        `json:"shipping_zip"`
			ShippingCountry   string        `json:"shipping_country"`
			DiscountCode      string        `json:"discount_code"`
			DiscountAmountOff interface{}   `json:"discount_amount_off"`
			Variants          []interface{} `json:"variants"`
			CustomFields      struct {
			} `json:"custom_fields"`
			AffiliateID         interface{} `json:"affiliate_id"`
			AffiliateCommission string      `json:"affiliate_commission"`
			CreatedAt           time.Time   `json:"created_at"`
			PaymentRequest      string      `json:"payment_request"`
		} `json:"payments"`
		CreatedAt             time.Time `json:"created_at"`
		ModifiedAt            time.Time `json:"modified_at"`
		AllowRepeatedPayments bool      `json:"allow_repeated_payments"`
	} `json:"payment_request"`
	Success bool `json:"success"`
}

// CreateRefundResponse is the data required to create a new Refund request.
// All fields are not necessary, Head over to instamojo docs for more more information
type CreateRefundRequest struct {
	PaymentID    string `json:"payment_id"`
	Type         string `json:"type"`
	RefundAmount string `json:"refund_amount"`
	Body         string `json:"body"`
}

// CreateRefundResponse is the response that is returned when a refund request is created successfully
type CreateRefundResponse struct {
	Refund struct {
		ID           string    `json:"id"`
		PaymentID    string    `json:"payment_id"`
		Status       string    `json:"status"`
		Type         string    `json:"type"`
		Body         string    `json:"body"`
		RefundAmount string    `json:"refund_amount"`
		TotalAmount  string    `json:"total_amount"`
		CreatedAt    time.Time `json:"created_at"`
	} `json:"refund"`
	Success bool `json:"success"`
}

// RefundsList is the list of all the refunds made
type RefundsList struct {
	Refunds []struct {
		ID           string    `json:"id"`
		PaymentID    string    `json:"payment_id"`
		Status       string    `json:"status"`
		Type         string    `json:"type"`
		Body         string    `json:"body"`
		RefundAmount string    `json:"refund_amount"`
		TotalAmount  string    `json:"total_amount"`
		CreatedAt    time.Time `json:"created_at"`
	} `json:"refunds"`
	Success bool `json:"success"`
}

// RefundDetails is details of a Refund
type RefundDetails struct {
	Refund struct {
		ID           string    `json:"id"`
		PaymentID    string    `json:"payment_id"`
		Status       string    `json:"status"`
		Type         string    `json:"type"`
		Body         string    `json:"body"`
		RefundAmount string    `json:"refund_amount"`
		TotalAmount  string    `json:"total_amount"`
		CreatedAt    time.Time `json:"created_at"`
	} `json:"refund"`
	Success bool `json:"success"`
}

// PaymentDetails is detailed information about a successfull payment
type PaymentDetails struct {
	Payment struct {
		PaymentID         string        `json:"payment_id"`
		Quantity          int           `json:"quantity"`
		Status            string        `json:"status"`
		LinkSlug          string        `json:"link_slug"`
		LinkTitle         string        `json:"link_title"`
		BuyerName         string        `json:"buyer_name"`
		BuyerPhone        string        `json:"buyer_phone"`
		BuyerEmail        string        `json:"buyer_email"`
		Currency          string        `json:"currency"`
		UnitPrice         string        `json:"unit_price"`
		Amount            string        `json:"amount"`
		Fees              string        `json:"fees"`
		ShippingAddress   string        `json:"shipping_address"`
		ShippingCity      string        `json:"shipping_city"`
		ShippingState     string        `json:"shipping_state"`
		ShippingZip       string        `json:"shipping_zip"`
		ShippingCountry   string        `json:"shipping_country"`
		DiscountCode      interface{}   `json:"discount_code"`
		DiscountAmountOff interface{}   `json:"discount_amount_off"`
		Variants          []interface{} `json:"variants"`
		CustomFields      struct {
		} `json:"custom_fields"`
		AffiliateID         interface{} `json:"affiliate_id"`
		AffiliateCommission string      `json:"affiliate_commission"`
		CreatedAt           time.Time   `json:"created_at"`
		PaymentRequest      string      `json:"payment_request"`
	} `json:"payment"`
	Success bool `json:"success"`
}

type successResponse struct {
	Success bool `json:"success"`
}
