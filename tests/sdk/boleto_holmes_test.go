package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Boleto "github.com/starkbank/sdk-go/starkbank/boleto"
	Holmes "github.com/starkbank/sdk-go/starkbank/boletoholmes"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestHolmesPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	boletos, errCreate := Boleto.Create(Example.Boleto(), nil)
	if errCreate.Errors != nil {
		for _, erro := range errCreate.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}

	sherlock, err := Holmes.Create(Example.Holmes(boletos[0].Id), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, holmes := range sherlock {
		assert.NotNil(t, holmes.Id)
		assert.Equal(t, boletos[0].Id, holmes.BoletoId)
	}
}

func TestHolmesGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var sherlock []Holmes.BoletoHolmes
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	logs := Holmes.Query(params, nil)
	for log := range logs {
		sherlock = append(sherlock, log)
	}

	holmes, err := Holmes.Get(sherlock[rand.Intn(len(sherlock))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, holmes.Id)
}

func TestHolmesQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 201

	sherlock := Holmes.Query(params, nil)

	for holmes := range sherlock {
		assert.NotNil(t, holmes.Id)
		i++
	}
	assert.Equal(t, 201, i)
}

func TestHolmesPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	sherlock, cursor, err := Holmes.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, holmes := range sherlock {
		ids = append(ids, holmes.Id)
		assert.NotNil(t, holmes.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
