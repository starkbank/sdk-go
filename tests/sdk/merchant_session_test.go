package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	MerchantSession "github.com/starkbank/sdk-go/starkbank/merchantsession"
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
			{Count: 1, TotalAmount: 500},
			{Count: 2, TotalAmount: 1000},
			{Count: 12, TotalAmount: 6000},
		},
		Expiration:   		 60,
		ChallengeMode: 		 "disabled",
		Tags:          		 []string{"test"},
		HolderId:            "5656565656565656",
		SoftDescriptor:      "softDescriptor",
	}

	createdSession, err := MerchantSession.Create(merchantSession, starkbank.User)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
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

	_, err := MerchantSession.PostPurchase("0bb894a2697d41d99fe02cad2c00c9bc", MerchantSession.Purchase{}, nil)
	for _, e := range err.Errors {
		assert.Equal(t, "Unknown exception encountered: Function deprecated since v1.7.0", e.Message)
	}
}
