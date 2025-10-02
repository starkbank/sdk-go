package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	DictKey "github.com/starkbank/sdk-go/starkbank/dictkey"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDictKeyGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 1
	var keyList []DictKey.DictKey
	var params = map[string]interface{}{}
	params["status"] = "registered"
	params["limit"] = limit
	params["type"] = "evp"

	keys, errorChannel := DictKey.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case key, ok := <-keys:
			if !ok {
				break loop
			}
			keyList = append(keyList, key)
		}
	}

	key, err := DictKey.Get(keyList[rand.Intn(len(keyList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, key.Id)
}

func TestDictKeyQuery(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	var keyList []DictKey.DictKey

	keys, errorChannel := DictKey.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case key, ok := <-keys:
			if !ok {
				break loop
			}
			keyList = append(keyList, key)
		}
	}

	assert.Equal(t, limit, len(keyList))
}

func TestDictKeyPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 1

	keys, cursor, err := DictKey.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	for _, key := range keys {
		ids = append(ids, key.Id)
		assert.NotNil(t, key.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 1)
}
