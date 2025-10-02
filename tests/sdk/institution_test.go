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
	params["search"] = "stark"

	institutions, errorChannel := Institution.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case institution, ok := <-institutions:
			if !ok {
				break loop
			}
			assert.NotNil(t, institution)
		}
	}
}

func TestInstitutionQueryFail(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["invalid"] = "stark"

	institutions, errorChannel := Institution.Query(params, nil)
	
	select {
	case err := <-errorChannel:
		if err.Errors != nil {
			for _, e := range err.Errors {
				assert.NotNil(t, e)
			}
		}
	case institution := <-institutions:
		assert.Nil(t, institution)
	}

	assert.Equal(t, 0, len(institutions))
}
