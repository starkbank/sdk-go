package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	HolmesLog "github.com/starkbank/sdk-go/starkbank/boletoholmes/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestHolmesLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var sherlock []HolmesLog.Log
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	logs := HolmesLog.Query(params, nil)
	for log := range logs {
		sherlock = append(sherlock, log)
	}

	holmes, err := HolmesLog.Get(sherlock[rand.Intn(len(sherlock))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	assert.NotNil(t, holmes.Id)
}

func TestHolmesLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 201

	logs := HolmesLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
		i++
	}
	assert.Equal(t, 201, i)
}

func TestHolmesLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := HolmesLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, log := range logs {
		ids = append(ids, log.Id)
		assert.NotNil(t, log.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
