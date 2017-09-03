// Package instamojo aims to provide a Wrapper for instamojo.com's API
// It is a work in progress and all remaining endpoints shall be added soon
package instamojo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Init initialises a new Config from the provided settings
func Init(c *Config) (*Config, error) {
	if c.APIKey == "" || c.AuthToken == "" {
		return nil, fmt.Errorf("invalid tokens")
	}

	if c.SandboxMode {
		c.endpoint = "https://test.instamojo.com"
	} else {
		c.endpoint = "https://www.instamojo.com"
	}

	return c, nil
}

// ParseWebhookResponse parses the urlencoded response that instamojo sends to the webhook
func ParseWebhookResponse(u url.Values) *WebhookResponse {

	return &WebhookResponse{
		Fees:             u.Get("fees"),
		Buyer:            u.Get("buyer"),
		Status:           u.Get("status"),
		Amount:           u.Get("amount"),
		Longurl:          u.Get("longurl"),
		Purpose:          u.Get("purpose"),
		Currency:         u.Get("currency"),
		Shorturl:         u.Get("shorturl"),
		PaymentID:        u.Get("payment_id"),
		BuyerName:        u.Get("buyer_name"),
		BuyerPhone:       u.Get("buyer_phone"),
		PaymentRequestID: u.Get("payment_request_id"),
	}

}

func (c *Config) makeRequest(m, url string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(m, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Key", c.APIKey)
	req.Header.Set("X-Auth-Token", c.AuthToken)

	if m == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Handle the irrecoverable errors here
	switch resp.StatusCode {
	case 404:
		return nil, fmt.Errorf("404 Not Found")
	case 500, 502, 504:
		return nil, fmt.Errorf("internal server error")
	case 403:
		return nil, fmt.Errorf("insufficient permissions")
	}

	return resp, nil

}

// CreatePaymentURL creates a new Payment URL
func (c *Config) CreatePaymentURL(p *PaymentURLRequest) (*PaymentURLResponse, error) {

	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("error in marshalling PaymentURLRequest: %v", err)
	}

	resp, err := c.makeRequest("POST", fmt.Sprintf("%s/api/1.1/payment-requests/", c.endpoint), strings.NewReader(string(b)))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 201:
		s := &PaymentURLResponse{}
		err := json.NewDecoder(resp.Body).Decode(s)
		if err != nil {
			return nil, err
		}

		return s, nil
	case 400:
		return nil, badrequest(resp)
	case 401:
		return nil, unauthorized(resp)
	default:
		return nil, defaultResponse(resp)
	}

	return nil, nil
}

//ListRequests returns a array of all the lists created so far
func (c *Config) ListRequests() (*RequestsList, error) {

	resp, err := c.makeRequest("GET", fmt.Sprintf("%s/api/1.1/payment-requests/", c.endpoint), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		r := &RequestsList{}
		err := json.NewDecoder(resp.Body).Decode(r)
		if err != nil {
			return nil, err
		}

		return r, nil

	case 400:
		return nil, badrequest(resp)
	case 401:
		return nil, unauthorized(resp)
	default:
		return nil, defaultResponse(resp)
	}
	return nil, nil
}

// PaymentDetails fetches details about a payment ID
func (c *Config) PaymentRequestDetails(id string) (*PaymentRequestDetails, error) {

	resp, err := c.makeRequest("GET", fmt.Sprintf("%s/api/1.1/payment-requests/%s", c.endpoint, id), nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		prd := &PaymentRequestDetails{}
		err := json.NewDecoder(resp.Body).Decode(prd)
		if err != nil {
			return nil, err
		}
		return prd, nil

	case 400:
		return nil, badrequest(resp)
	case 401:
		return nil, unauthorized(resp)
	default:
		return nil, defaultResponse(resp)
	}

	return nil, nil
}

func (c *Config) CreateRefundRequest(r *CreateRefundRequest) (*CreateRefundResponse, error) {
	resp, err := c.makeRequest("POST", fmt.Sprintf("%s/api/1.1/refunds"), strings.NewReader())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 201:
		crr := &CreateRefundResponse{}
		err := json.NewDecoder(resp.Body).Decode(crr)
		if err != nil {
			return nil, err
		}

		return crr, nil
	case 400:
		return nil, badrequest(resp)
	case 401:
		return nil, unauthorized(resp)
	default:
		return nil, defaultResponse(resp)
	}
	return nil, nil
}

func (c *Config) ListRefunds() (*RefundsList, error) {
	resp, err := c.makeRequest("GET", fmt.Sprintf("%s/api/1.1/refunds", c.endpoint), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		rl := &RefundsList{}
		err := json.NewDecoder(resp.Body).Decode(rl)

		if err != nil {
			return nil, err
		}
		return rl, nil
	case 400:
		return nil, badrequest(resp)
	case 401:
		return nil, unauthorized(resp)
	default:
		return nil, defaultResponse(resp)
	}

	return nil, nil
}

func (c *Config) RefundDetails(id string) (*RefundDetails, error) {
	resp, err := c.makeRequest("GET", fmt.Sprintf("%s/api/1.1/refunds/%s", c.endpoint, id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		rd := &RefundDetails{}
		err := json.NewDecoder(resp.Body).Decode(rd)
		if err != nil {
			return nil, err
		}
		return rd, nil
	case 400:
		return nil, badrequest(resp)
	case 401:
		return nil, unauthorized(resp)
	default:
		return nil, defaultResponse(resp)
	}

	return nil, nil
}

func (c *Config) PaymentDetails(id string) (*PaymentDetails, error) {

	resp, err := c.makeRequest("GET", fmt.Sprintf("%s/api/1.1/payments/%s", c.endpoint, id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		pd := &PaymentDetails{}
		err := json.NewDecoder(resp.Body).Decode(pd)
		if err != nil {
			return nil, err
		}

		return pd, nil
	case 400:
		return nil, badrequest(resp)
	case 401:
		return nil, unauthorized(resp)
	default:
		return nil, defaultResponse(resp)
	}

	return nil, nil
}

func (c *Config) DisableRequest(id string) (*successResponse, error) {

	resp, err := c.makeRequest("POST", fmt.Sprintf("%s/api/1.1/payment-requests/%s/disable", c.endpoint, id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		sr := &successResponse{}
		err := json.NewDecoder(resp.Body).Decode(sr)
		if err != nil {
			return nil, err
		}

		return sr, nil
	case 400:
		return nil, badrequest(resp)
	case 401:
		return nil, unauthorized(resp)
	default:
		return nil, defaultResponse(resp)
	}
	return nil, nil
}

func (c *Config) EnableRequest(id string) (*successResponse, error) {

	resp, err := c.makeRequest("POST", fmt.Sprintf("%s/api/1.1/payment-requests/%s/enable", c.endpoint, id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		sr := &successResponse{}
		err := json.NewDecoder(resp.Body).Decode(sr)
		if err != nil {
			return nil, err
		}

		return sr, nil
	case 400:
		return nil, badrequest(resp)
	case 401:
		return nil, unauthorized(resp)
	default:
		return nil, defaultResponse(resp)
	}
	return nil, nil
}

func badrequest(resp *http.Response) error {
	br := &BadRequest{}
	err := json.NewDecoder(resp.Body).Decode(br)
	if err != nil {
		return err
	}
	return br
}

func unauthorized(resp *http.Response) error {
	u := &Unauthorized{}
	err := json.NewDecoder(resp.Body).Decode(u)
	if err != nil {
		return err
	}
	return u
}

func defaultResponse(resp *http.Response) error {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return fmt.Errorf("unrecognized response from instamojo: %s", string(b))

}
