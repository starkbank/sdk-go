package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	VerifiedAccount "github.com/starkbank/sdk-go/starkbank/verifiedaccount"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifiedAccountCreateAndCancel(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	accounts, err := VerifiedAccount.Create(Example.VerifiedAccountPixKey(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	account := accounts[0]
	assert.NotNil(t, account.Id)

	canceledAccount, err := VerifiedAccount.Cancel(account.Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.Equal(t, canceledAccount.Id, account.Id)
	assert.Equal(t, "canceled", canceledAccount.Status)
}

func TestVerifiedAccountQuery(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var accountsList []VerifiedAccount.VerifiedAccount

	accounts, errorChannel := VerifiedAccount.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case account, ok := <-accounts:
			if !ok {
				break loop
			}
			assert.NotNil(t, account.Id)
			accountsList = append(accountsList, account)
		}
	}

	assert.Equal(t, limit, len(accountsList))
}

func TestVerifiedAccountQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	accounts, errorChannel := VerifiedAccount.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case account, ok := <-accounts:
			if !ok {
				break loop
			}
			getAccount, err := VerifiedAccount.Get(account.Id, nil)
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
			assert.Equal(t, account.Id, getAccount.Id)
		}
	}
}

func TestVerifiedAccountPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	ids := make(map[string]bool)
	cursor := ""
	count := 0
	for i := 0; i < 2; i++ {
		params := map[string]interface{}{"limit": 5}
		if cursor != "" {
			params["cursor"] = cursor
		}
		page, nextCursor, err := VerifiedAccount.Page(params, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		for _, entity := range page {
			assert.False(t, ids[entity.Id])
			ids[entity.Id] = true
			count++
		}
		if nextCursor == "" {
			break
		}
		cursor = nextCursor
	}
	assert.Equal(t, 10, count)
}
