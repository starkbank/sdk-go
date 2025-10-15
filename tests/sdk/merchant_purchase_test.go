package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	MerchantPurchase "github.com/starkbank/sdk-go/starkbank/merchantpurchase"
	MerchantSession "github.com/starkbank/sdk-go/starkbank/merchantsession"
	Purchase "github.com/starkbank/sdk-go/starkbank/merchantsession"
	AllowedInstallment "github.com/starkbank/sdk-go/starkbank/merchantsession/allowedinstallment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMerchantPurchaseCreate(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	merchantSession := MerchantSession.MerchantSession{
		AllowedFundingTypes: []string{"credit"},
		AllowedInstallments: []AllowedInstallment.AllowedInstallment{
			{Count: 1, TotalAmount: 1000},
		},
		Expiration:   		 60,
		ChallengeMode: 		 "disabled",
		Tags:          		 []string{"test"},
	}

	createdSession, err := MerchantSession.Create(merchantSession, starkbank.User)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	purchase := Purchase.Purchase{
		Amount:           	1000,
		FundingType: 		"credit",
		CardExpiration: 	"2035-01",
		CardNumber: 		"36490101441625",
		CardSecurityCode: 	"123",
		HolderName: 		"Margaery Tyrell",
	}

	createdPurchase, err := MerchantSession.PostPurchase(createdSession.Uuid, purchase, starkbank.User)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	merchantPurchase := MerchantPurchase.MerchantPurchase{
		Amount:           	1000,
		FundingType: 		"credit",
		CardId: 		 	createdPurchase.CardId,
		ChallengeMode: 		"disabled",
	}

	createdMerchantPurchase, err := MerchantPurchase.Create(merchantPurchase, nil)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, createdMerchantPurchase.Id)
}

func TestMerchantPurchaseGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var purchaseList []MerchantPurchase.MerchantPurchase
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	purchases := MerchantPurchase.Query(params, nil)
	for purchase := range purchases {
		purchaseList = append(purchaseList, purchase)
	}

	purchase, err := MerchantPurchase.Get(purchaseList[rand.Intn(len(purchaseList))].Id, nil)
	
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, purchase.Id)
}

func TestMerchantPurchaseQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["after"] = "2020-04-01"
	params["before"] = "2020-04-30"

	purchases := MerchantPurchase.Query(params, nil)

	for purchase := range purchases {
		assert.NotNil(t, purchase.Id)
	}
}

func TestMerchantPurchasePage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	purchases, cursor, err := MerchantPurchase.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, purchase := range purchases {
		ids = append(ids, purchase.Id)
		assert.NotNil(t, purchase.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestMerchantPurchasePatch(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var purchaseList []MerchantPurchase.MerchantPurchase
	var params = map[string]interface{}{}
	params["status"] = "approved"

	purchases := MerchantPurchase.Query(params, nil)
	for purchase := range purchases {
		purchaseList = append(purchaseList, purchase)
	}

	var patchData = map[string]interface{}{}
	patchData["amount"] = 0	
	patchData["status"] = "canceled"

	updated, err := MerchantPurchase.Update(purchaseList[rand.Intn(len(purchaseList))].Id, patchData, nil)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.Equal(t, updated.Amount, patchData["amount"])
}
