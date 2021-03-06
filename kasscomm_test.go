package kasscomm

import (
	"encoding/json"
	"testing"
)

var test_auth_token string = "kass_test_auth_token"

var base_request Request = Request{
	Amount:      2199,
	Description: "Kass bolur",
	Image_Url:   "https://photos.kassapi.is/kass/kass-bolur.jpg",
	Order:       "ABC123",
	Recipient:   "1001000",
	Terminal:    1,
	Expires_In:  90,
	Notify_Url:  "",
}

var new_request Request

func initTestEnv() {
	SetAuthToken(test_auth_token)
	SetDev()
	new_request = base_request
}

//Expects successful run
func TestInitiatePayment(t *testing.T) {
	initTestEnv()
	resp, err := InitiatePayment(&base_request)

	if err != nil {
		t.Errorf("Expected err == nil but got %s", err)
	}

	if !resp.Success {
		out, _ := json.Marshal(resp)
		t.Errorf("Expected Response.Success == true but got Response = %s", string(out))
	}
}

//Expects an error
func TestInitiatePaymentMissingRecipient(t *testing.T) {
	initTestEnv()
	new_request.Recipient = ""
	_, err := InitiatePayment(&new_request)

	if err == nil {
		t.Error("Expected err != nil but got err == nil")
	}
}

//Expects an error
func TestInitiatePaymentInvalidAmount(t *testing.T) {
	initTestEnv()
	new_request.Amount = -1
	_, err := InitiatePayment(&new_request)
	if err == nil {
		t.Error("Expected err != nil but got err == nil")
	}
}

//Expects an error
func TestInitiatePaymentEmptyRequest(t *testing.T) {
	initTestEnv()
	_, err := InitiatePayment(&Request{})
	if err == nil {
		t.Error("Expected err != nil but got err == nil")
	}
}

//Expects a success==false from Api
func TestInitiatePaymentInvalidRecipient(t *testing.T) {
	initTestEnv()
	new_request.Recipient = "123"
	resp, err := InitiatePayment(&new_request)
	if err != nil {
		t.Errorf("Expected err == nil but got err == %s", err)
	}

	if resp.Success {
		out, _ := json.Marshal(resp)
		t.Errorf("Expected respones.Success == false but got response == %s", out)
	}
}

//Expects an error
func TestInitiatePaymentNoAuthToken(t *testing.T) {
	SetDev()
	SetAuthToken("")
	_, err := InitiatePayment(&base_request)

	if err == nil {
		t.Error("Expected err != nil but got err == nil")
	}
}
