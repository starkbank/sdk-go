package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Key "github.com/starkinfra/core-go/starkcore/key"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePrivateKey(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	privateKey, publicKey := Key.Create("")
	assert.NotNil(t, privateKey)
	assert.NotNil(t, publicKey)

}

func TestPathPrivateKey(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	privateKey, publicKey := Key.Create("sample")
	assert.NotNil(t, privateKey)
	assert.NotNil(t, publicKey)

}
