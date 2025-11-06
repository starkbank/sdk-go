package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	MerchantCard "github.com/starkbank/sdk-go/starkbank/merchantcard"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantCardQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var cardList []MerchantCard.MerchantCard

	cards, errorChannel := MerchantCard.Query(params, starkbank.User)
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

	card, err := MerchantCard.Get(cardList[0].Id, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, card.Id)
}


func TestMerchantCardPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 5
	
	cards, cursor, err := MerchantCard.Page(params, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, cards)
	assert.NotNil(t, cursor)
	assert.Greater(t, len(cards), 0)
	for _, card := range cards {
		assert.NotNil(t, card.Id)
	}
}
