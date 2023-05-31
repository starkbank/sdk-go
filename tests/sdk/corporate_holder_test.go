package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	CorporateHolder "github.com/starkbank/sdk-go/starkbank/corporateholder"
	"github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCorporateHolderPost(t *testing.T) {

	starkbank.User = utils.ExampleProject

	holders, err := CorporateHolder.Create(Example.CorporateHolder(), nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, holder := range holders {
		assert.NotNil(t, holder.Id)
	}
}

func TestCorporateHolderQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	holders := CorporateHolder.Query(paramsQuery, nil)

	for holder := range holders {
		assert.NotNil(t, holder.Id)
	}
}

func TestCorporateHolderPage(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	holders, cursor, err := CorporateHolder.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, holder := range holders {
		assert.NotNil(t, holder.Id)
	}

	assert.NotNil(t, cursor)
}

func TestCorporateHolderGet(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var holderList []CorporateHolder.CorporateHolder
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	holders := CorporateHolder.Query(paramsQuery, nil)
	for holder := range holders {
		holderList = append(holderList, holder)
	}

	holder, err := CorporateHolder.Get(holderList[rand.Intn(len(holderList))].Id, nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, holder.Id)
}

func TestCorporateHolderUpdate(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var holderList []CorporateHolder.CorporateHolder
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "active"

	holders := CorporateHolder.Query(paramsQuery, nil)
	for holder := range holders {
		holderList = append(holderList, holder)
	}

	var patchData = map[string]interface{}{}
	patchData["name"] = "Tony Starkoso"

	holder, err := CorporateHolder.Update(holderList[rand.Intn(len(holderList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, holder.Id)
}

func TestCorporateHolderCancel(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var holderList []CorporateHolder.CorporateHolder
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "active"

	holders := CorporateHolder.Query(paramsQuery, nil)
	for holder := range holders {
		holderList = append(holderList, holder)
	}

	holder, err := CorporateHolder.Cancel(holderList[rand.Intn(len(holderList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, holder.Id)
}
