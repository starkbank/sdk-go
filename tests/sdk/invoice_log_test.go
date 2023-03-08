package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	InvoiceLog "github.com/starkbank/sdk-go/starkbank/invoice/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestInvoiceLogGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var logList []InvoiceLog.Log
	var params = map[string]interface{}{}
	params["after"] = "2021-04-01"
	params["before"] = "2021-04-30"

	logs := InvoiceLog.Query(params, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	log, err := InvoiceLog.Get(logList[0].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	fmt.Printf("%+v", log)
}

func TestInvoiceLogQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 201

	logs := InvoiceLog.Query(params, nil)

	for log := range logs {
		assert.NotNil(t, log.Id)
		i++
	}
	assert.Equal(t, 201, i)
}

func TestInvoiceLogPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	logs, cursor, err := InvoiceLog.Page(params, nil)
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

func TestInvoiceLogPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var logList []InvoiceLog.Log
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)
	params["types"] = "paid"

	logs := InvoiceLog.Query(params, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	pdf, err := InvoiceLog.Pdf(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "invoice-log", "5155165527080960")
	errFile := ioutil.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}
}
