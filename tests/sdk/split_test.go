package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	Split "github.com/starkbank/sdk-go/starkbank/split"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSplitGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var splitList []Split.Split

	splits, errorChannel := Split.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case split, ok := <-splits:
			if !ok {
				break loop
			}
			splitList = append(splitList, split)
		}
	}

	split, err := Split.Get(splitList[rand.Intn(len(splitList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	assert.NotNil(t, split.Id)
}

func TestSplitQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var splitList []Split.Split

	splits, errorChannel := Split.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case split, ok := <-splits:
			if !ok {
				break loop
			}
			splitList = append(splitList, split)
		}
	}

	assert.Equal(t, limit, len(splitList))
}

func TestSplitPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	splits, cursor, err := Split.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, split := range splits {
		ids = append(ids, split.Id)
		assert.NotNil(t, split.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
