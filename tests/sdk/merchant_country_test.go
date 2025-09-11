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

	countries, errorChannel := MerchantCountry.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case country, ok := <-countries:
			if !ok {
				break loop
			}
			assert.NotNil(t, country.Code)
		}
	}
}
