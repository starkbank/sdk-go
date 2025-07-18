package sdk

import (
	"fmt"
	"testing"
	"github.com/starkbank/sdk-go/starkbank"
	InvoicePullRequestLog "github.com/starkbank/sdk-go/starkbank/invoicepullrequest/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestInvoicePullRequestLogQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	params := map[string]interface{}{"limit": 10}
	logs := InvoicePullRequestLog.Query(params, nil)
	var firstLogId string
	for log := range logs {
		assert.NotNil(t, log.Id)
		if firstLogId == "" {
			firstLogId = log.Id
		}
	}
	getLog, err := InvoicePullRequestLog.Get(firstLogId, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.Equal(t, firstLogId, getLog.Id)
}

func TestInvoicePullRequestLogPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	ids := make(map[string]bool)
	cursor := ""
	count := 0
	for i := 0; i < 2; i++ {
		params := map[string]interface{}{"limit": 5}
		if cursor != "" {
			params["cursor"] = cursor
		}
		page, nextCursor, err := InvoicePullRequestLog.Page(params, nil)
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
