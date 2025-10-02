package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/transfer"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"os"
	"math/rand"
	"testing"
)

func TestTransferPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	transfers, err := transfer.Create(Example.Transfer(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, transfer := range transfers {
		assert.NotNil(t, transfer)
	}
}

func TestTransferGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var transferList []transfer.Transfer

	transfers, errorChannel := transfer.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transfer, ok := <-transfers:
			if !ok {
				break loop
			}
			transferList = append(transferList, transfer)
		}
	}

	transfer, err := transfer.Get(transferList[rand.Intn(len(transferList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, transfer)
}

func TestTransferPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["status"] = "success"

	var transferList []transfer.Transfer
	
	transfers, errorChannel := transfer.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transfer, ok := <-transfers:
			if !ok {
				break loop
			}
			transferList = append(transferList, transfer)
		}
	}

	pdf, err := transfer.Pdf(transferList[rand.Intn(len(transferList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "transfer", transferList[rand.Intn(len(transferList))].Id)
	errFile := os.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		t.Errorf("error writing file: %v", errFile)
	}
	assert.NotNil(t, pdf)
}

func TestTransferQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["status"] = "success"
	params["limit"] = limit

	var transferList []transfer.Transfer

	transfers, errorChannel := transfer.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transfer, ok := <-transfers:
			if !ok {
				break loop
			}
			transferList = append(transferList, transfer)
		}
	}

	for _, transfer := range transferList {
		assert.Equal(t, transfer.Status, "success")
	}
}

func TestTransferPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	transfers, cursor, err := transfer.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, transfer := range transfers {
		ids = append(ids, transfer.Id)
		assert.NotNil(t, transfer.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestTransferDelete(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	transfers, errCreate := transfer.Create(Example.Transfer(), nil)
	if errCreate.Errors != nil {
		for _, erro := range errCreate.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	canceled, err := transfer.Delete(transfers[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, canceled.Id)
}
