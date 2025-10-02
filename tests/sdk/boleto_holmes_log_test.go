package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	HolmesLog "github.com/starkbank/sdk-go/starkbank/boletoholmes/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestHolmesLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var sherlock []HolmesLog.Log

	logs, err := HolmesLog.Query(params, nil)

	loop:
	for {
		select {
		case err := <-err:
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

	holmes, errGet := HolmesLog.Get(sherlock[rand.Intn(len(sherlock))].Id, nil)
	if errGet.Errors != nil {
		for _, erro := range errGet.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	assert.NotNil(t, holmes.Id)
}

func TestHolmesLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logList []HolmesLog.Log

	logs, err := HolmesLog.Query(params, nil)

	loop:
	for {
		select {
		case err := <-err:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			logList = append(logList, log)
		}
	}
	assert.Equal(t, limit, len(logList))
}

func TestHolmesLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := HolmesLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, log := range logs {
		ids = append(ids, log.Id)
		assert.NotNil(t, log.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
