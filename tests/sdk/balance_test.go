package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Balance "github.com/starkbank/sdk-go/starkbank/balance"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBalanceGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	balance := Balance.Get(nil)
	fmt.Println(balance)
	assert.NotNil(t, balance)
}
