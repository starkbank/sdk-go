package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Boleto "github.com/starkbank/sdk-go/starkbank/boleto"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"os"
	"math/rand"
	"testing"
)

func TestBoletoPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	boletos, err := Boleto.Create(Example.Boleto(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, boleto := range boletos {
		assert.NotNil(t, boleto.Id)
	}
}

func TestBoletoQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit

	var boletoList []Boleto.Boleto

	boletos, errorChannel := Boleto.Query(params, nil)
	
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case boleto, ok := <-boletos:
			if !ok {
				break loop
			}
			boletoList = append(boletoList, boleto)
		}
	}
	assert.Equal(t, limit, len(boletoList))
}

func TestBoletoPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	boletos, cursor, err := Boleto.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
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
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	canceled, err := Boleto.Delete(boletos[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, canceled.Id)
}

func TestBoletoGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 5
	var params = map[string]interface{}{}
	params["limit"] = limit
	
	var boletoList []Boleto.Boleto

	boletos, errorChannel := Boleto.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case boleto, ok := <-boletos:
			if !ok {
				break loop
			}
			boletoList = append(boletoList, boleto)
		}
	}

	boleto, err := Boleto.Get(boletoList[rand.Intn(len(boletoList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, boleto.Id)
}

func TestBoletoPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "paid"
	
	var boletoList []Boleto.Boleto

	boletos, errorChannel := Boleto.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case boleto, ok := <-boletos:
			if !ok {
				break loop
			}
			boletoList = append(boletoList, boleto)
		}
	}

	var params = map[string]interface{}{}
	params["layout"] = "booklet"

	if len(boletoList) == 0 {
		t.Skip("No Boleto with status 'paid' found")
	}

	pdf, err := Boleto.Pdf(boletoList[rand.Intn(len(boletoList))].Id, params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "boleto", boletoList[rand.Intn(len(boletoList))].Id)
	errFile := os.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		t.Errorf("error writing file: %v", errFile)
	}
	assert.NotNil(t, pdf)
}
