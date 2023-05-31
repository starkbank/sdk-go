package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	CorporatePurchase "github.com/starkbank/sdk-go/starkbank/corporatepurchase"
	"github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCorporatePurchaseQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	purchases := CorporatePurchase.Query(params, nil)
	for purchase := range purchases {
		assert.NotNil(t, purchase.Id)
	}
}

func TestCorporatePurchasePage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	purchases, cursor, err := CorporatePurchase.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, purchase := range purchases {
		assert.NotNil(t, purchase.Id)
	}
	assert.NotNil(t, cursor)
}

func TestCorporatePurchaseGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var purchaseList []CorporatePurchase.CorporatePurchase
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	purchases := CorporatePurchase.Query(paramsQuery, nil)
	for purchase := range purchases {
		purchaseList = append(purchaseList, purchase)
	}

	purchase, err := CorporatePurchase.Get(purchaseList[rand.Intn(len(purchaseList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, purchase)
}

func TestCorporatePurchaseParseRight(t *testing.T) {

	starkbank.User = utils.ExampleProject

	content := "{\"acquirerId\": \"236090\", \"amount\": 100, \"cardId\": \"5671893688385536\", \"cardTags\": [], \"endToEndId\": \"2fa7ef9f-b889-4bae-ac02-16749c04a3b6\", \"holderId\": \"5917814565109760\", \"holderTags\": [], \"isPartialAllowed\": false, \"issuerAmount\": 100, \"issuerCurrencyCode\": \"BRL\", \"merchantAmount\": 100, \"merchantCategoryCode\": \"bookStores\", \"merchantCountryCode\": \"BRA\", \"merchantCurrencyCode\": \"BRL\", \"merchantFee\": 0, \"merchantId\": \"204933612653639\", \"merchantName\": \"COMPANY 123\", \"methodCode\": \"token\", \"purpose\": \"purchase\", \"score\": null, \"tax\": 0, \"walletId\": \"\"}"
	validSignature := "MEUCIBxymWEpit50lDqFKFHYOgyyqvE5kiHERi0ZM6cJpcvmAiEA2wwIkxcsuexh9BjcyAbZxprpRUyjcZJ2vBAjdd7o28Q="

	parsed := CorporatePurchase.Parse(content, validSignature, nil)
	assert.NotNil(t, parsed)
}

func TestCorporatePurchaseParseWrong(t *testing.T) {

	starkbank.User = utils.ExampleProject

	content := "{\"acquirerId\": \"236090\", \"amount\": 100, \"cardId\": \"5671893688385536\", \"cardTags\": [], \"endToEndId\": \"2fa7ef9f-b889-4bae-ac02-16749c04a3b6\", \"holderId\": \"5917814565109760\", \"holderTags\": [], \"isPartialAllowed\": false, \"issuerAmount\": 100, \"issuerCurrencyCode\": \"BRL\", \"merchantAmount\": 100, \"merchantCategoryCode\": \"bookStores\", \"merchantCountryCode\": \"BRA\", \"merchantCurrencyCode\": \"BRL\", \"merchantFee\": 0, \"merchantId\": \"204933612653639\", \"merchantName\": \"COMPANY 123\", \"methodCode\": \"token\", \"purpose\": \"purchase\", \"score\": null, \"tax\": 0, \"walletId\": \"\"}"
	invalidSignature := "MEUCIQDOpo1j+V40pNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="

	parsed := CorporatePurchase.Parse(content, invalidSignature, nil)
	assert.NotNil(t, parsed)
}

func TestCorporatePurchaseResponseApproved(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var approved = map[string]interface{}{}
	approved["status"] = "approved"
	approved["amount"] = 10000
	approved["tags"] = []string{"tony", "stark"}

	response := CorporatePurchase.Response(approved)
	assert.NotNil(t, response)
}

func TestCorporatePurchaseResponseDenied(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var denied = map[string]interface{}{}
	denied["status"] = "denied"
	denied["reason"] = "other"
	denied["tags"] = []string{"tony", "stark"}

	response := CorporatePurchase.Response(denied)
	assert.NotNil(t, response)
}
