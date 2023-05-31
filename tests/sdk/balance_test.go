package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/balance"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalanceGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	balance := balance.Get(nil)
	assert.NotNil(t, balance.Amount)
	assert.NotNil(t, balance.Id)
}
