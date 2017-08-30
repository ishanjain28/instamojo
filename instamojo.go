// This package aims to provide a Wrapper for instamojo.com's API
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
	if c.ApiKey == "" || c.AuthToken == "" {
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
	req.Header.Set("X-Api-Key", c.ApiKey)
	req.Header.Set("X-Auth-Token", c.AuthToken)

	if m == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
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
		s := &BadRequest{}

		err := json.NewDecoder(resp.Body).Decode(s)
		if err != nil {
			return nil, err
		}

		return nil, s

	case 401:
		s := &Unauthorized{}
		err := json.NewDecoder(resp.Body).Decode(s)
		if err != nil {
			return nil, err
		}
		return nil, s
	default:

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("unrecognized response from instamojo: %s", string(b))
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

		return nil, r

	case 401:
		u := &Unauthorized{}
		err := json.NewDecoder(resp.Body).Decode(u)
		if err != nil {
			return nil, err
		}

		return nil, u
	default:

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("unrecognized response from instamojo: %s", string(b))
	}
	return nil, nil
}
