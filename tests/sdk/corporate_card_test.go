package sdk

import (
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
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, card)
}

func TestCorporateCardQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var cardList []CorporateCard.CorporateCard

	cards, errorChannel := CorporateCard.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case card, ok := <-cards:
			if !ok {
				break loop
			}
			cardList = append(cardList, card)
		}
	}

	assert.Equal(t, limit, len(cardList))
}

func TestCorporateCardPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	cards, cursor, err := CorporateCard.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, card := range cards {
		assert.NotNil(t, card.Id)
	}

	assert.NotNil(t, cursor)
}

func TestCorporateCardGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var cardList []CorporateCard.CorporateCard

	cards, errorChannel := CorporateCard.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case card, ok := <-cards:
			if !ok {
				break loop
			}
			cardList = append(cardList, card)
		}
	}


	card, err := CorporateCard.Get(cardList[rand.Intn(len(cardList))].Id, nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, card.Id)
}

func TestCorporateCardDelete(t *testing.T) {

	starkbank.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "active"
	
	var cardList []CorporateCard.CorporateCard

	cards, errorChannel := CorporateCard.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case card, ok := <-cards:
			if !ok {
				break loop
			}
			cardList = append(cardList, card)
		}
	}

	if len(cardList) == 0 {
		t.Skip("No Card with status 'active' found")
	}

	card, err := CorporateCard.Cancel(cardList[rand.Intn(len(cardList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, card.Id)
}

func TestCorporateCardUpdate(t *testing.T) {

	starkbank.User = utils.ExampleProject

	card, err := CorporateCard.Create(Example.CorporateCard(), nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, card)

	var patchData = map[string]interface{}{}
	patchData["displayName"] = "ANTHONY EDWARD"

	updatedCard, err := CorporateCard.Update(card.Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, updatedCard.Id)
	assert.NotNil(t, card.DisplayName)
}
