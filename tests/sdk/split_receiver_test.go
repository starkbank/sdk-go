package sdk

import (
	"testing"
	"github.com/starkbank/sdk-go/starkbank"
	SplitReceiver "github.com/starkbank/sdk-go/starkbank/splitreceiver"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
)

func TestSplitReceiverCreate(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	receivers, err := SplitReceiver.Create(Example.SplitReceiver(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	receiver := receivers[0]
	assert.NotNil(t, receiver.Id)
}

func TestSplitReceiverQuery(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	var receiversList []SplitReceiver.SplitReceiver

	receivers, errorChannel := SplitReceiver.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case receiver, ok := <-receivers:
			if !ok {
				break loop
			}
			assert.NotNil(t, receiver.Id)
			receiversList = append(receiversList, receiver)
		}
	}

	assert.Equal(t, limit, len(receiversList))
}

func TestSplitReceiverQueryAndGet(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	receivers, errorChannel := SplitReceiver.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case receiver, ok := <-receivers:
			if !ok {
				break loop
			}
			getReceiver, err := SplitReceiver.Get(receiver.Id, nil)
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
			assert.Equal(t, receiver.Id, getReceiver.Id)
		}
	}
}

func TestSplitReceiverPage(t *testing.T) {
	starkbank.User = Utils.ExampleProject

	ids := make(map[string]bool)
	cursor := ""
	count := 0
	for i := 0; i < 2; i++ {
		params := map[string]interface{}{"limit": 5}
		if cursor != "" {
			params["cursor"] = cursor
		}
		page, nextCursor, err := SplitReceiver.Page(params, nil)
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
