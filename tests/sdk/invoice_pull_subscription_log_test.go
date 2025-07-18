package sdk

import (
	"fmt"
	"testing"
	"github.com/starkbank/sdk-go/starkbank"
	InvoicePullSubscriptionLog "github.com/starkbank/sdk-go/starkbank/invoicepullsubscription/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestInvoicePullSubscriptionLogQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	params := map[string]interface{}{"limit": 10}
	logs := InvoicePullSubscriptionLog.Query(params, nil)
	var firstLogId string
	for log := range logs {
		assert.NotNil(t, log.Id)
		if firstLogId == "" {
			firstLogId = log.Id
		}
	}
	getLog, err := InvoicePullSubscriptionLog.Get(firstLogId, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.Equal(t, firstLogId, getLog.Id)
}

func TestInvoicePullSubscriptionLogPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	ids := make(map[string]bool)
	cursor := ""
	count := 0
	for i := 0; i < 2; i++ {
		params := map[string]interface{}{"limit": 5}
		if cursor != "" {
			params["cursor"] = cursor
		}
		page, nextCursor, err := InvoicePullSubscriptionLog.Page(params, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
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
