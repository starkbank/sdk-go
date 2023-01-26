package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	DictKey "github.com/starkbank/sdk-go/starkbank/dictkey"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDictKeyGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var keyList []DictKey.DictKey
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	keys := DictKey.Query(params, nil)
	for key := range keys {
		keyList = append(keyList, key)
	}

	key, err := DictKey.Get(keyList[rand.Intn(len(keyList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, key.Id)
}

func TestDictKeyQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 1

	keys := DictKey.Query(params, nil)

	for key := range keys {
		assert.NotNil(t, key.Id)
		i++
	}
	assert.Equal(t, params["limit"], i)
}

func TestDictKeyPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 1

	keys, cursor, err := DictKey.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, key := range keys {
		ids = append(ids, key.Id)
		assert.NotNil(t, key.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 1)
}
