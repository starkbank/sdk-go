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

	categories := MerchantCategory.Query(nil, nil)
	for category := range categories {
		assert.NotNil(t, category.Code)
	}
}
