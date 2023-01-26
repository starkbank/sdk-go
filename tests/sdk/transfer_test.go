package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Transfer "github.com/starkbank/sdk-go/starkbank/transfer"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestTransferPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	transfers, err := Transfer.Create(Example.Transfer(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, transfer := range transfers {
		assert.NotNil(t, transfer.Id)
	}
}

func TestTransferGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var transferList []Transfer.Transfer
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	transfers := Transfer.Query(params, nil)
	for transfer := range transfers {
		transferList = append(transferList, transfer)
	}

	transfer, err := Transfer.Get(transferList[rand.Intn(len(transferList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, transfer.Id)
}

func TestTransferPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var transferList []Transfer.Transfer
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)
	params["status"] = "success"

	transfers := Transfer.Query(params, nil)
	for transfer := range transfers {
		transferList = append(transferList, transfer)
	}

	pdf, err := Transfer.Pdf(transferList[rand.Intn(len(transferList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "transfer", transferList[rand.Intn(len(transferList))].Id)
	errFile := ioutil.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}
	assert.NotNil(t, pdf)
}

func TestTransferQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "success"
	params["limit"] = 10

	transfers := Transfer.Query(params, nil)

	for transfer := range transfers {
		assert.Equal(t, transfer.Status, "success")
	}
}

func TestTransferPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	transfers, cursor, err := Transfer.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, transfer := range transfers {
		ids = append(ids, transfer.Id)
		assert.NotNil(t, transfer.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestTransferCancel(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	transfers, errCreate := Transfer.Create(Example.Transfer(), nil)
	if errCreate.Errors != nil {
		for _, erro := range errCreate.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	canceled, err := Transfer.Delete(transfers[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, canceled.Id)
}
