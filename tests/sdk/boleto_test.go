package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Boleto "github.com/starkbank/sdk-go/starkbank/boleto"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestBoletoPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	boletos, err := Boleto.Create(Example.Boleto(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, boleto := range boletos {
		assert.NotNil(t, boleto.Id)
	}
}

func TestBoletoQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "registered"
	params["limit"] = rand.Intn(100)

	boletos := Boleto.Query(params, nil)
	for boleto := range boletos {
		fmt.Println(boleto)
		assert.Equal(t, boleto.Status, "registered")
	}
}

func TestBoletoPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	boletos, cursor, err := Boleto.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, boleto := range boletos {
		ids = append(ids, boleto.Id)
		assert.NotNil(t, boleto.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestBoletoPostAndDelete(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	boletos, errCreate := Boleto.Create(Example.Boleto(), nil)
	if errCreate.Errors != nil {
		for _, erro := range errCreate.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	canceled, err := Boleto.Delete(boletos[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, canceled.Id)
}

func TestBoletoGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var boletoList []Boleto.Boleto
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	boletos := Boleto.Query(params, nil)
	for boleto := range boletos {
		boletoList = append(boletoList, boleto)
	}

	boleto, err := Boleto.Get(boletoList[rand.Intn(len(boletoList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, boleto.Id)
}

func TestBoletoPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var boletoList []Boleto.Boleto
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "paid"

	boletos := Boleto.Query(paramsQuery, nil)
	for boleto := range boletos {
		boletoList = append(boletoList, boleto)
	}

	var params = map[string]interface{}{}
	params["layout"] = "booklet"

	pdf, err := Boleto.Pdf(boletoList[rand.Intn(len(boletoList))].Id, params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "boleto", boletoList[rand.Intn(len(boletoList))].Id)
	errFile := ioutil.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}
	assert.NotNil(t, pdf)
}
