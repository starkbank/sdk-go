package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/balance"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"testing"
)

func TestBalanceGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	balance := balance.Get(nil)
	fmt.Printf("%+v", balance)
}
