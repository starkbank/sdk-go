package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	MerchantCard "github.com/starkbank/sdk-go/starkbank/merchantcard"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMerchantCardQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var cardList []MerchantCard.MerchantCard
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	cards := MerchantCard.Query(params, starkbank.User)
	for card := range cards {
		cardList = append(cardList, card)
	}

	card, err := MerchantCard.Get(cardList[rand.Intn(len(cardList))].Id, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, card.Id)
}


func TestMerchantCardPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)
	
	cards, cursor, err := MerchantCard.Page(params, starkbank.User)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, cards)
	assert.NotNil(t, cursor)
	assert.Greater(t, len(cards), 0)
	for _, card := range cards {
		assert.NotNil(t, card.Id)
	}
}
