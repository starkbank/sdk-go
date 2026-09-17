package sdk

import (
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkbank/sdk-go/starkbank/transfer/rule"
	VerifiedAccount "github.com/starkbank/sdk-go/starkbank/verifiedaccount"
	VerifiedTransfer "github.com/starkbank/sdk-go/starkbank/verifiedtransfer"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifiedTransferCreate(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	// The API caps VerifiedAccount creation at three per tax id per 24h, and the
	// fixtures use a fixed tax id, so reuse an account that is already active.
	var accountId string
	var params = map[string]interface{}{"limit": 1, "status": "active"}
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
			accountId = account.Id
		}
	}

	if accountId == "" {
		created, err := VerifiedAccount.Create(Example.VerifiedAccount(), nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
			return
		}
		accountId = created[0].Id
	}

	transfers, err := VerifiedTransfer.Create(Example.VerifiedTransfer(accountId), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	transfer := transfers[0]
	assert.NotNil(t, transfer.Id)

	var resendingLimit rule.Rule
	for _, r := range transfer.Rules {
		if r.Key == "resendingLimit" {
			resendingLimit = r
			break
		}
	}
	assert.Equal(t, "resendingLimit", resendingLimit.Key)
	assert.EqualValues(t, 5, resendingLimit.Value)
}
