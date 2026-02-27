package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	MerchantPurchase "github.com/starkbank/sdk-go/starkbank/merchantpurchase"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantPurchaseCreate(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	params := map[string]interface{}{
		"limit": 1,
		"status": "confirmed",
	}

	var purchaseList []MerchantPurchase.MerchantPurchase

	merchantPurchases, errorChannel := MerchantPurchase.Query(params, starkbank.User)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case purchase, ok := <-merchantPurchases:
			if !ok {
				break loop
			}
			purchaseList = append(purchaseList, purchase)
		}
	}

	confirmedMerchantPurchase := purchaseList[0]

	merchantPurchase := MerchantPurchase.MerchantPurchase{
		Amount: 1000,
		FundingType: "credit",
		CardId: confirmedMerchantPurchase.CardId,
		ChallengeMode: "disabled",
		HolderId: "5746894506843",
	}

	createdMerchantPurchase, err := MerchantPurchase.Create(merchantPurchase, starkbank.User)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, createdMerchantPurchase.Id)
}

func TestMerchantPurchaseGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var purchaseList []MerchantPurchase.MerchantPurchase

	purchases, errorChannel := MerchantPurchase.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case purchase, ok := <-purchases:
			if !ok {
				break loop
			}
			purchaseList = append(purchaseList, purchase)
		}
	}

	purchase, err := MerchantPurchase.Get(purchaseList[0].Id, nil)
	
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, purchase.Id)
}

func TestMerchantPurchaseQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["after"] = "2020-04-01"
	params["before"] = "2020-04-30"

	purchases, errorChannel := MerchantPurchase.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case purchase, ok := <-purchases:
			if !ok {
				break loop
			}
			assert.NotNil(t, purchase.Id)
		}
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
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
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

	var params = map[string]interface{}{}
	params["status"] = "approved"
	
	var purchaseList []MerchantPurchase.MerchantPurchase

	purchases, errorChannel := MerchantPurchase.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case purchase, ok := <-purchases:
			if !ok {
				break loop
			}
			purchaseList = append(purchaseList, purchase)
		}
	}

	var patchData = map[string]interface{}{}
	patchData["amount"] = 0	
	patchData["status"] = "canceled"

	updated, err := MerchantPurchase.Update(purchaseList[0].Id, patchData, nil)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.Equal(t, updated.Amount, patchData["amount"])
}
