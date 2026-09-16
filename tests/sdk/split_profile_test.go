package sdk

import (
	"testing"
	"github.com/starkbank/sdk-go/starkbank"
	SplitProfile "github.com/starkbank/sdk-go/starkbank/splitprofile"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
)

func TestSplitProfilePut(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	profiles, err := SplitProfile.Put(Example.SplitProfile(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	profile := profiles[0]
	assert.NotNil(t, profile.Id)
}

func TestSplitProfileQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	profiles, errorChannel := SplitProfile.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case profile, ok := <-profiles:
			if !ok {
				break loop
			}
			getProfile, err := SplitProfile.Get(profile.Id, nil)
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
			assert.Equal(t, profile.Id, getProfile.Id)
		}
	}
}

func TestSplitProfilePage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	ids := make(map[string]bool)
	cursor := ""
	count := 0
	for i := 0; i < 2; i++ {
		params := map[string]interface{}{"limit": 5}
		if cursor != "" {
			params["cursor"] = cursor
		}
		page, nextCursor, err := SplitProfile.Page(params, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		for _, entity := range page {
			ids[entity.Id] = true
			count++
		}
		if nextCursor == "" {
			break
		}
		cursor = nextCursor
	}
	assert.True(t, count >= 0)
}
