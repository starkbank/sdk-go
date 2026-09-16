package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	DepositLog "github.com/starkbank/sdk-go/starkbank/deposit/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
)

func TestDepositLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var logIds, depositsIds []string

	logs, errorChannel := DepositLog.Query(params, nil)
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
			logIds = append(logIds, log.Id)
		}
	}

	for _, ids := range logIds {
		deposit, err := DepositLog.Get(ids, nil)
		if err.Errors != nil {
			for _, erro := range err.Errors {
				t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
			}
		}
		assert.NotNil(t, deposit)
		depositsIds = append(depositsIds, deposit.Id)
	}
	assert.Equal(t, logIds, depositsIds)
}

func TestDepositLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logList []DepositLog.Log

	logs, errorChannel := DepositLog.Query(params, nil)
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
			logList = append(logList, log)
		}
	}

	assert.Equal(t, limit, len(logList))
}

func TestDepositLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := DepositLog.Page(params, nil)
	for _, log := range logs {
		if err.Errors != nil {
			for _, erro := range err.Errors {
				t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
			}
		}
		ids = append(ids, log.Id)
		assert.NotNil(t, log.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestDepositLogPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["types"] = "reversed"

	var logList []DepositLog.Log

	logs, errorChannel := DepositLog.Query(params, nil)
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
			logList = append(logList, log)
		}
	}

	pdf, err := DepositLog.Pdf(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "deposit-log", "5155165527080960")
	errFile := os.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}
}
