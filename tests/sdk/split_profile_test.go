package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	SplitProfile "github.com/starkbank/sdk-go/starkbank/splitprofile"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSplitProfilePut(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	profiles := []SplitProfile.SplitProfile{
		{
			Interval: "week",
			Delay:    0,
		},
	}

	result, err := SplitProfile.Put(profiles, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, profile := range result {
		assert.NotNil(t, profile.Id)
	}
}

func TestSplitProfileGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	var profileList []SplitProfile.SplitProfile

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
			profileList = append(profileList, profile)
		}
	}

	profile, err := SplitProfile.Get(profileList[rand.Intn(len(profileList))].Id, nil)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}
	assert.NotNil(t, profile.Id)
}

func TestSplitProfileQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	var profileList []SplitProfile.SplitProfile

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
			profileList = append(profileList, profile)
		}
	}

	assert.Equal(t, limit, len(profileList))
}

func TestSplitProfilePage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	limit := 1
	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = limit

	profiles, cursor, err := SplitProfile.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, profile := range profiles {
		ids = append(ids, profile.Id)
		assert.NotNil(t, profile.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, limit)
}
