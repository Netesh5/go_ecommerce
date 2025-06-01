package services

import (
	"errors"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient
var serviceID *string

func CreateTwiiloClient(username string, password string, serviceId string) *twilio.RestClient {
	serviceID = &serviceId
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})
	return client
}

func TwilioSendOTP(email string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(email)
	params.SetChannel("email")
	res, err := client.VerifyV2.CreateVerification(*serviceID, params)
	if err != nil {
		return "", err
	}

	return *res.Sid, nil
}

func TwilioVerifyOTP(email string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(email)
	params.SetCode(code)
	res, err := client.VerifyV2.CreateVerificationCheck(*serviceID, params)
	if err != nil {
		return err
	}

	if *res.Status != "approved" {
		return errors.New("not a valid code")
	}
	return nil

}
