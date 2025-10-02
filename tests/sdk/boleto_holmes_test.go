package sdk

import (
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
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}

	sherlock, err := Holmes.Create(Example.Holmes(boletos[0].Id), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, holmes := range sherlock {
		assert.NotNil(t, holmes.Id)
		assert.Equal(t, boletos[0].Id, holmes.BoletoId)
	}
}

func TestHolmesGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var sherlock []Holmes.BoletoHolmes

	logs, errorChannel := Holmes.Query(params, nil)
	
	loop: 
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			sherlock = append(sherlock, log)
		}
	}

	holmes, err := Holmes.Get(sherlock[rand.Intn(len(sherlock))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, holmes.Id)
}

func TestHolmesQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var holmesList []Holmes.BoletoHolmes

	sherlock, errorChannel := Holmes.Query(params, nil)

	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case holmes, ok := <-sherlock:
			if !ok {
				break loop
			}
			holmesList = append(holmesList, holmes)
		}
	}
	assert.Equal(t, limit, len(holmesList))
}

func TestHolmesPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	sherlock, cursor, err := Holmes.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, holmes := range sherlock {
		ids = append(ids, holmes.Id)
		assert.NotNil(t, holmes.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
