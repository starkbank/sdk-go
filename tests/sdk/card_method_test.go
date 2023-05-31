package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	CardMethod "github.com/starkbank/sdk-go/starkbank/cardmethod"
	"github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardMethodQuery(t *testing.T) {

	starkbank.User = utils.ExampleProject

	methods := CardMethod.Query(nil, nil)
	for method := range methods {
		assert.NotNil(t, method.Code)
	}
}
