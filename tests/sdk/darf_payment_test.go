package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	DarfPayment "github.com/starkbank/sdk-go/starkbank/darfpayment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestDarfPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	darfs, err := DarfPayment.Create(Example.Darf(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, darf := range darfs {
		assert.NotNil(t, darf.Id)
	}
}

func TestDarfGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var darfList []DarfPayment.DarfPayment
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	darfs := DarfPayment.Query(params, nil)
	for darf := range darfs {
		darfList = append(darfList, darf)
	}

	darf, err := DarfPayment.Get(darfList[rand.Intn(len(darfList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, darf.Id)
}

func TestDarfPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var darfList []DarfPayment.DarfPayment
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)
	params["status"] = "success"

	darfs := DarfPayment.Query(params, nil)
	for darf := range darfs {
		darfList = append(darfList, darf)
	}

	pdf, err := DarfPayment.Pdf(darfList[rand.Intn(len(darfList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "darf", darfList[rand.Intn(len(darfList))].Id)
	errFile := ioutil.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}
	assert.NotNil(t, pdf)
}

func TestDarfQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 201

	darfs := DarfPayment.Query(params, nil)
	for darf := range darfs {
		assert.NotNil(t, darf.Id)
	}
}

func TestDarfPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	darfs, cursor, err := DarfPayment.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, darf := range darfs {
		ids = append(ids, darf.Id)
		assert.NotNil(t, darf.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestDarfCancel(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	payments, errCreate := DarfPayment.Create(Example.Darf(), nil)
	if errCreate.Errors != nil {
		for _, e := range errCreate.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	payment, err := DarfPayment.Delete(payments[rand.Intn(len(payments))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, payment)
}
