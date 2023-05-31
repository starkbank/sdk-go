package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	CorporateCard "github.com/starkbank/sdk-go/starkbank/corporatecard"
	"github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCorporateCardPost(t *testing.T) {

	starkbank.User = utils.ExampleProject

	card, err := CorporateCard.Create(Example.CorporateCard(), nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, card)
}

func TestCorporateCardQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	cards := CorporateCard.Query(params, nil)
	for card := range cards {
		assert.NotNil(t, card.Id)
	}
}

func TestCorporateCardPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	cards, cursor, err := CorporateCard.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, card := range cards {
		assert.NotNil(t, card.Id)
	}

	assert.NotNil(t, cursor)
}

func TestCorporateCardGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var cardList []CorporateCard.CorporateCard
	var paramsQuery = map[string]interface{}{}

	cards := CorporateCard.Query(paramsQuery, nil)
	for card := range cards {
		cardList = append(cardList, card)
	}

	card, err := CorporateCard.Get(cardList[rand.Intn(len(cardList))].Id, nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, card.Id)
}

func TestCorporateCardDelete(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var cardList []CorporateCard.CorporateCard
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "active"

	cards := CorporateCard.Query(paramsQuery, nil)
	for card := range cards {
		cardList = append(cardList, card)
	}

	card, err := CorporateCard.Cancel(cardList[rand.Intn(len(cardList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, card.Id)
}

func TestCorporateCardUpdate(t *testing.T) {

	starkbank.User = utils.ExampleProject

	card, err := CorporateCard.Create(Example.CorporateCard(), nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, card)

	var patchData = map[string]interface{}{}
	patchData["displayName"] = "ANTHONY EDWARD"

	updatedCard, err := CorporateCard.Update(card.Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, updatedCard.Id)
	assert.NotNil(t, card.DisplayName)
}
