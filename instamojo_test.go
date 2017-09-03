package instamojo_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/ishanjain28/instamojo"
)

func ExampleParseWebhookResponse() {

	values := url.Values{
		"fees":               []string{"125.00"},
		"buyer":              []string{"abc@xyz.com"},
		"buyer_name":         []string{"John Doe"},
		"buyer_phone":        []string{"9999999999"},
		"status":             []string{"Credit"},
		"amount":             []string{"2500.00"},
		"longurl":            []string{"https://www.instamojo.com/@portrack/077a7ff202f94d3e86ffe64511efa8a4"},
		"currency":           []string{"INR"},
		"mac":                []string{"1ddf3b78f84d071324c0bf1d3f095898267d60ee"},
		"payment_id":         []string{"MOJO5a06005J21512197"},
		"payment_request_id": []string{"d66cb29dd059482e8072999f995c4eef"},
		"purpose":            []string{"FIFA 16"},
		"shorturl":           []string{"https://imjo.in/NNxHg"},
	}
	fmt.Println(instamojo.ParseWebhookResponse(values))
}

func TestParseWebhookResponse(t *testing.T) {

	values := url.Values{
		"fees":               []string{"125.00"},
		"buyer":              []string{"abc@xyz.com"},
		"buyer_name":         []string{"John Doe"},
		"buyer_phone":        []string{"9999999999"},
		"status":             []string{"Credit"},
		"amount":             []string{"2500.00"},
		"longurl":            []string{"https://www.instamojo.com/@portrack/077a7ff202f94d3e86ffe64511efa8a4"},
		"currency":           []string{"INR"},
		"mac":                []string{"1ddf3b78f84d071324c0bf1d3f095898267d60ee"},
		"payment_id":         []string{"MOJO5a06005J21512197"},
		"payment_request_id": []string{"d66cb29dd059482e8072999f995c4eef"},
		"purpose":            []string{"FIFA 16"},
		"shorturl":           []string{"https://imjo.in/NNxHg"},
	}

	got := *instamojo.ParseWebhookResponse(values)

	want := instamojo.WebhookResponse{
		PaymentID:        "MOJO5a06005J21512197",
		Status:           "Credit",
		Shorturl:         "https://imjo.in/NNxHg",
		Longurl:          "https://www.instamojo.com/@portrack/077a7ff202f94d3e86ffe64511efa8a4",
		Purpose:          "FIFA 16",
		Amount:           "2500.00",
		Fees:             "125.00",
		Currency:         "INR",
		Buyer:            "abc@xyz.com",
		BuyerName:        "John Doe",
		BuyerPhone:       "9999999999",
		PaymentRequestID: "d66cb29dd059482e8072999f995c4eef",
		Mac:              "1ddf3b78f84d071324c0bf1d3f095898267d60ee",
	}

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
