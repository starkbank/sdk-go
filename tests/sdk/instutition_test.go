package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	Institution "github.com/starkbank/sdk-go/starkbank/institution"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInstitutionQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	institutions := Institution.Query(params, nil)

	for institution := range institutions {
		assert.NotNil(t, institution.SpiCode)
	}
}
