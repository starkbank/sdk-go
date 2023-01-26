package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	DynamicBrcode "github.com/starkbank/sdk-go/starkbank/dynamicbrcode"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestDynamicBrcodePost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	brcodes, err := DynamicBrcode.Create(Example.DynamicBrcode(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, brcode := range brcodes {
		assert.NotNil(t, brcode.Uuid)
	}
}

func TestDynamicBrcodeGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var brcodeList []DynamicBrcode.DynamicBrcode
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	brcodes := DynamicBrcode.Query(params, nil)
	for brcode := range brcodes {
		brcodeList = append(brcodeList, brcode)
	}

	brcode, err := DynamicBrcode.Get(brcodeList[rand.Intn(len(brcodeList))].Uuid, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, brcode.Uuid)
}

func TestDynamicBrcodeQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)
	brcodes := DynamicBrcode.Query(params, nil)

	for brcode := range brcodes {
		assert.NotNil(t, brcode.Uuid)
	}
}

func TestDynamicBrcodePage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	brcodes, cursor, err := DynamicBrcode.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, brcode := range brcodes {
		ids = append(ids, brcode.Uuid)
		assert.NotNil(t, brcode.Uuid)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}
