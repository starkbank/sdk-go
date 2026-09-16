package sdk

import (
	"testing"
	"github.com/starkbank/sdk-go/starkbank"
	SplitReceiverLog "github.com/starkbank/sdk-go/starkbank/splitreceiver/log"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestSplitReceiverLogQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	logs, errorChannel := SplitReceiverLog.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			getLog, err := SplitReceiverLog.Get(log.Id, nil)
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
			assert.Equal(t, log.Id, getLog.Id)
		}
	}
}

func TestSplitReceiverLogPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	ids := make(map[string]bool)
	cursor := ""
	count := 0
	for i := 0; i < 2; i++ {
		params := map[string]interface{}{"limit": 5}
		if cursor != "" {
			params["cursor"] = cursor
		}
		page, nextCursor, err := SplitReceiverLog.Page(params, nil)
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
