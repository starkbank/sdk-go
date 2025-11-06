package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	MerchantSession "github.com/starkbank/sdk-go/starkbank/merchantsession"
	Purchase "github.com/starkbank/sdk-go/starkbank/merchantsession"
	AllowedInstallment "github.com/starkbank/sdk-go/starkbank/merchantsession/allowedinstallment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantSessionCreate(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	merchantSession := MerchantSession.MerchantSession{
		AllowedFundingTypes: []string{"credit"},
		AllowedIps:          []string{"192.168.0.1"},
		AllowedInstallments: []AllowedInstallment.AllowedInstallment{
			{Count: 1, TotalAmount: 0},
			{Count: 2, TotalAmount: 120},
			{Count: 12, TotalAmount: 180},
		},
		Expiration:   		 60,
		ChallengeMode: 		 "disabled",
		Tags:          		 []string{"test"},
	}

	createdSession, err := MerchantSession.Create(merchantSession, starkbank.User)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, createdSession.Id)
}

func TestMerchantSessionGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var sessionList []MerchantSession.MerchantSession

	sessions, errorChannel := MerchantSession.Query(params, starkbank.User)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case session, ok := <-sessions:
			if !ok {
				break loop
			}
			sessionList = append(sessionList, session)
		}
	}

	session, err := MerchantSession.Get(sessionList[0].Id, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, session.Id)
}

func TestMerchantSessionQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "created"

	sessions, errorChannel := MerchantSession.Query(params, starkbank.User)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case session, ok := <-sessions:
			if !ok {
				break loop
			}
			assert.NotNil(t, session.Id)
		}
	}
}

func TestMerchantSessionPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	sessions, cursor, err := MerchantSession.Page(params, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, session := range sessions {
		ids = append(ids, session.Id)
		assert.NotNil(t, session.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestMerchantSessionPurchase(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	merchantSession := MerchantSession.MerchantSession{
		AllowedFundingTypes: []string{"credit"},
		AllowedInstallments: []AllowedInstallment.AllowedInstallment{
			{Count: 1, TotalAmount: 0},
			{Count: 2, TotalAmount: 120},
			{Count: 12, TotalAmount: 180},
		},
		Expiration:   		 60,
		ChallengeMode: 		 "disabled",
		Tags:          		 []string{"test"},
	}

	createdSession, err := MerchantSession.Create(merchantSession, starkbank.User)

	purchase := Purchase.Purchase{
		Amount:            180,
		InstallmentCount:  12,
		CardExpiration:    "2035-01",
		CardNumber:        "5102589999999913",
		CardSecurityCode:  "123",
		HolderName:        "Holder Name",
		HolderEmail:       "holdeName@email.com",
		HolderPhone:       "11111111111",
		FundingType:       "credit",
		BillingCountryCode: "BRA",
		BillingCity:       "São Paulo",
		BillingStateCode:  "SP",
		BillingStreetLine1: "Rua do Holder Name, 123",
		BillingStreetLine2: "casa",
		BillingZipCode:    "11111-111",
		Metadata: map[string]interface{}{
			"userAgent":      "Postman",
			"userIp":         "255.255.255.255",
			"language":       "pt-BR",
			"timezoneOffset": 3,
			"extraData":      "extraData",
		},
	}

	createdPurchase, err := MerchantSession.PostPurchase(createdSession.Uuid, purchase, starkbank.User)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, createdPurchase)
}
