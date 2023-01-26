package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Utility "github.com/starkbank/sdk-go/starkbank/utilitypayment"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestUtilityPost(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	utilities, err := Utility.Create(Example.Utility(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, utility := range utilities {
		assert.NotNil(t, utility.Id)
	}
}

func TestUtilityGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var utilityList []Utility.UtilityPayment
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	utilitys := Utility.Query(params, nil)
	for utility := range utilitys {
		utilityList = append(utilityList, utility)
	}

	utility, err := Utility.Get(utilityList[rand.Intn(len(utilityList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, utility.Id)
}

func TestUtilityPdf(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var utilityList []Utility.UtilityPayment
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)
	params["status"] = "success"

	utilitys := Utility.Query(params, nil)
	for utility := range utilitys {
		utilityList = append(utilityList, utility)
	}

	pdf, err := Utility.Pdf(utilityList[rand.Intn(len(utilityList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "utility", utilityList[rand.Intn(len(utilityList))].Id)
	errFile := ioutil.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		fmt.Print(errFile)
	}
	assert.NotNil(t, pdf)
}

func TestUtilityQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "success"
	params["limit"] = 4

	utilities := Utility.Query(params, nil)
	for utility := range utilities {
		assert.Equal(t, utility.Status, "success")
	}
}

func TestUtilityPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	utilities, cursor, err := Utility.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	for _, utility := range utilities {
		ids = append(ids, utility.Id)
		assert.NotNil(t, utility.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestUtilityCancel(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	utilities, err := Utility.Create(Example.Utility(), nil)

	canceled, err := Utility.Delete(utilities[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, canceled.Id)
}
