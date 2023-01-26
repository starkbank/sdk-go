package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Key "github.com/starkinfra/core-go/starkcore/key"
	"testing"
)

func TestCreatePrivateKey(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	privateKey, publicKey := Key.Create("")
	fmt.Println("PRIVATE KEY", privateKey)
	fmt.Println("PUBLIC KEY", publicKey)

}

func TestPathPrivateKey(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	privateKey, publicKey := Key.Create("sample")
	fmt.Println("PRIVATE KEY", privateKey)
	fmt.Println("PUBLIC KEY", publicKey)

}
