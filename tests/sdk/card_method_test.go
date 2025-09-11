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

	methods, errorChannel := CardMethod.Query(nil, nil)
	
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case method, ok := <-methods:
			if !ok {
				break loop
			}
			assert.NotNil(t, method.Code)
		}
	}
}
