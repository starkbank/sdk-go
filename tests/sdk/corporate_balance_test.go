package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/corporatebalance"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCorporateBalanceGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	balance, err := corporatebalance.Get(nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, balance)
}
