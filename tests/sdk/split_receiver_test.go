package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	SplitReceiver "github.com/starkbank/sdk-go/starkbank/splitreceiver"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSplitReceiverCreate(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	receivers := []SplitReceiver.SplitReceiver{
		{
			Name:          "Arya Stark",
			TaxId:         "01234567890",
			BankCode:      "20018183",
			BranchCode:    "1357-9",
			AccountNumber: "876543-2",
			AccountType:   "checking",
			Tags:          []string{"split-receiver"},
		},
	}

	result, err := SplitReceiver.Create(receivers, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, receiver := range result {
		assert.NotNil(t, receiver.Id)
	}
}

func TestSplitReceiverGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var receiverList []SplitReceiver.SplitReceiver

	receivers, errorChannel := SplitReceiver.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case receiver, ok := <-receivers:
			if !ok {
				break loop
			}
			receiverList = append(receiverList, receiver)
		}
	}

	receiver, err := SplitReceiver.Get(receiverList[rand.Intn(len(receiverList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	assert.NotNil(t, receiver.Id)
}

func TestSplitReceiverQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var receiverList []SplitReceiver.SplitReceiver

	receivers, errorChannel := SplitReceiver.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case receiver, ok := <-receivers:
			if !ok {
				break loop
			}
			receiverList = append(receiverList, receiver)
		}
	}

	assert.Equal(t, limit, len(receiverList))
}

func TestSplitReceiverPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	receivers, cursor, err := SplitReceiver.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, receiver := range receivers {
		ids = append(ids, receiver.Id)
		assert.NotNil(t, receiver.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
