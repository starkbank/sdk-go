package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	MerchantCategory "github.com/starkbank/sdk-go/starkbank/merchantcategory"
	"github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantCategoryQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	categories, errorChannel := MerchantCategory.Query(nil, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case category, ok := <-categories:
			if !ok {
				break loop
			}
			assert.NotNil(t, category.Code)
		}
	}
}
