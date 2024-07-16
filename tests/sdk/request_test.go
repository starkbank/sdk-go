package sdk

import (
	"encoding/json"
	"strconv"
	"time"
	"fmt"
	"os"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject
	data := map[string]interface{}{}
	var path string
	var query = map[string]interface{}{}

	path = "/invoice/"
	query["limit"] = 2

	response, err := Request.Get(
		path,
		query,
		nil,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	invoicesData, ok1 := data["invoices"].([]interface{})
	if !ok1 {
        fmt.Println("Erro ao converter os tipos content")
        return
    }
	for _, invoice := range invoicesData{
		invoiceMap, ok2 := invoice.(map[string]interface{})
        if !ok2 {
            fmt.Println("Erro ao converter item de list 'invoices' para map[string]interface{}")
            continue
        }
        id, ok3 := invoiceMap["id"].(string)
        if !ok3 {
            fmt.Println("Erro ao converter list 'id' para string")
            continue
        }
		path = "invoice/" + id
		for k := range data {
			delete(data, k)
		}
		response, err := Request.Get(
			path,
			nil,
			Utils.ExampleProject,
		)
	
		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		invoiceData, ok4 := data["invoice"].(map[string]interface{})
        if !ok4 {
            fmt.Println("Erro ao converter 'id' para string")
            continue
        }
		getId, ok5 := invoiceData["id"].(string)
		if !ok5 {
            fmt.Println("Erro ao converter 'id' para string")
            continue
        }
		assert.Equal(t, id, getId)
	}
}

func TestRequestGetFile(t *testing.T) {
	starkbank.User = Utils.ExampleProject
	data := map[string]interface{}{}
	var path string
	var query = map[string]interface{}{}

	path = "/invoice/"
	query["limit"] = 2
	query["status"] = "paid"

	response, err := Request.Get(
		path,
		query,
		Utils.ExampleProject,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	invoicesData, ok1 := data["invoices"].([]interface{})
	if !ok1 {
        fmt.Println("Erro ao converter os tipos content")
        return
    }
	for _, invoice := range invoicesData{
		invoiceMap, ok2 := invoice.(map[string]interface{})
        if !ok2 {
            fmt.Println("Erro ao converter item de list 'invoices' para map[string]interface{}")
            continue
        }
        id, ok3 := invoiceMap["id"].(string)
        if !ok3 {
            fmt.Println("Erro ao converter list 'id' para string")
            continue
        }
		path = "invoice/" + id + "/pdf"
		for k := range data {
			delete(data, k)
		}
		pdf, err := Request.Get(
			path,
			nil,
			Utils.ExampleProject,
		)
	
		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
			}
		}
		filename := fmt.Sprintf("%v%v.pdf", "invoice", id)
		errFile := os.WriteFile(filename, pdf.Content, 0666)
		if errFile != nil {
		  fmt.Print(errFile)
		}
		assert.NotNil(t, pdf)

		path = "invoice/" + id + "/qrcode"
		for k := range data {
			delete(data, k)
		}
		qrCode, err := Request.Get(
			path,
			nil,
			Utils.ExampleProject,
		)
	
		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
			}
		}
		qrCodefilename := fmt.Sprintf("%v%v.png", "invoice", id)
		qrErrFile := os.WriteFile(qrCodefilename, qrCode.Content, 0666)
		if qrErrFile != nil {
			fmt.Print(errFile)
  		}
	}
}

func TestRequestPost(t *testing.T) {
	starkbank.User = Utils.ExampleProject
	data := map[string]interface{}{}
	var path string
	body := map[string][]map[string]interface{}{
        "invoices": {
            {
				"amount": 996699999,
				"name":   "Tony Stark",
				"taxId":  "38.446.231/0001-04",
			},
        },
    }


	path = "/invoice/"

	response, err := Request.Post(
		path,
		body,
		nil,
		Utils.ExampleProject,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	assert.NotNil(t, data)
}

func TestRequestPatch(t *testing.T) {
	starkbank.User = Utils.ExampleProject
	data := map[string]interface{}{}
	var path string
	var query = map[string]interface{}{}

	path = "/invoice/"
	query["limit"] = 2
	query["status"] = "paid"

	response, err := Request.Get(
		path,
		query,
		nil,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	invoicesData, ok1 := data["invoices"].([]interface{})
	if !ok1 {
        fmt.Println("Erro ao converter os tipos content")
        return
    }
	for _, invoice := range invoicesData{
		invoiceMap, _ := invoice.(map[string]interface{})
        id, _ := invoiceMap["id"].(string)
		path = "invoice/" + id
		for k := range data {
			delete(data, k)
		}
		body := map[string]interface{}{
			"amount" : 0,
		}

		response, err := Request.Patch(
			path,
			body,
			nil,
			Utils.ExampleProject,
		)
	
		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		invoiceData, _ := data["invoice"].(map[string]interface{})
		amount, _ := invoiceData["amount"].(int)
		assert.Equal(t, 0, amount)
	}
}

func TestRequestPut(t *testing.T) {
	starkbank.User = Utils.ExampleProject
	data := map[string]interface{}{}
	body := map[string][]map[string]interface{}{
        "profiles": {
            {
				"interval": "day",
				"delay": 0,
			},
        },
    }
	path := "split-profile/"
	response, err := Request.Put(
		path,
		body,
		nil,
		Utils.ExampleProject,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	assert.NotNil(t, data)
}

func TestRequestDelete(t *testing.T) {
	starkbank.User = Utils.ExampleProject
	data := map[string]interface{}{}
    
	now := time.Now()
    futureDate := now.AddDate(0, 0, 10).Format("2006-01-02")
    milliseconds := now.UnixNano() / int64(time.Millisecond)
    timestamp := strconv.FormatInt(milliseconds, 10)

	body := map[string][]map[string]interface{}{
        "transfers": {
            {
				"amount": 10000,
				"name": "Steve Rogers",
				"taxId": "330.731.970-10",
				"bankCode": "001",
				"branchCode": "1234",
				"accountNumber": "123456-0",
				"accountType": "checking",
				"scheduled": futureDate,
				"externalId": timestamp,
			},
        },
    }

	path := "transfer/"

	response, err := Request.Post(
		path,
		body,
		nil,
		Utils.ExampleProject,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		panic(unmarshalError)
	}
	transfersData, ok1 := data["transfers"].([]interface{})
	if !ok1 {
        fmt.Println("Erro ao converter os tipos content")
        return
    }
	for _, transfer := range transfersData{
		transferMap, ok2 := transfer.(map[string]interface{})
        if !ok2 {
            fmt.Println("Erro ao converter item de list 'invoices' para map[string]interface{}")
            continue
        }
        id, ok3 := transferMap["id"].(string)
        if !ok3 {
            fmt.Println("Erro ao converter list 'id' para string")
            continue
        }
		path = "transfer/" + id
		for k := range data {
			delete(data, k)
		}
		response, err := Request.Delete(
			path,
			Utils.ExampleProject,
		)
	
		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		assert.NotNil(t, data)
	}
}