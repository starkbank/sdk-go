package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	MerchantCountry "github.com/starkbank/sdk-go/starkbank/merchantcountry"
	"github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantCountryQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["search"] = "brazil"

	countries := MerchantCountry.Query(params, nil)
	for country := range countries {
		assert.NotNil(t, country.Code)
	}
}
