package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Event "github.com/starkbank/sdk-go/starkbank/event"
	Attempt "github.com/starkbank/sdk-go/starkbank/event/attempt"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestEventGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var eventList []Event.Event
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	events := Event.Query(params, nil)
	for event := range events {
		eventList = append(eventList, event)
	}

	event, err := Event.Get(eventList[rand.Intn(len(eventList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, event)
}

func TestEventQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["isDelivered"] = true
	params["limit"] = 150

	events := Event.Query(params, nil)

	for event := range events {
		assert.NotNil(t, event)
		assert.NotNil(t, event.Id)
	}
}

func TestEventPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	events, cursor, err := Event.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, event := range events {
		ids = append(ids, event.Id)
		assert.NotNil(t, event.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestEventDelete(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var eventList []Event.Event
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	events := Event.Query(params, nil)
	for event := range events {
		eventList = append(eventList, event)
	}

	event, err := Event.Delete(eventList[rand.Intn(len(eventList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, event.Id)
}

func TestEvenUpdate(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var eventList []Event.Event
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	events := Event.Query(params, nil)
	for event := range events {
		eventList = append(eventList, event)
	}

	patchData := map[string]interface{}{}
	patchData["isDelivered"] = true

	event, err := Event.Update(eventList[rand.Intn(len(eventList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.Equal(t, event.IsDelivered, true)
}

func TestEventAttemptQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	attempts := Attempt.Query(params, nil)

	for attempt := range attempts {
		assert.NotNil(t, attempt.Id)
	}
}

func TestEventAttemptGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var attemptList []Attempt.Attempt
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	attempts := Attempt.Query(params, nil)
	for attempt := range attempts {
		attemptList = append(attemptList, attempt)
	}

	attempt, err := Attempt.Get(attemptList[rand.Intn(len(attemptList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, attempt.Id)
}

func TestEventAttemptPage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 4

	attempts, cursor, err := Attempt.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, attempt := range attempts {
		ids = append(ids, attempt.Id)
		assert.NotNil(t, attempt.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 4)
}

func TestEventRightParse(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	content := "{\"event\": {\"created\": \"2021-04-26T20:16:51.866857+00:00\", \"id\": \"5415223380934656\", \"log\": {\"created\": \"2021-04-26T20:16:50.927706+00:00\", \"errors\": [], \"id\": \"4687457496858624\", \"invoice\": {\"amount\": 256, \"brcode\": \"00020101021226890014br.gov.bcb.pix2567invoice-h.sandbox.starkbank.com/v2/afdf94b770b0458a8440a335daf77c4c5204000053039865802BR5915Stark Bank S.A.6009Sao Paulo62070503***6304CC32\", \"created\": \"2021-04-26T20:16:50.886319+00:00\", \"descriptions\": [{\"key\": \"Field1\", \"value\": \"Something\"}], \"discountAmount\": 0, \"discounts\": [{\"due\": \"2021-05-07T09:43:15+00:00\", \"percentage\": 10.0}], \"due\": \"2021-05-09T19:11:39+00:00\", \"expiration\": 123456789, \"fee\": 0, \"fine\": 2.5, \"fineAmount\": 0, \"id\": \"5941925571985408\", \"interest\": 1.3, \"interestAmount\": 0, \"link\": \"https://cdottori.sandbox.starkbank.com/invoicelink/afdf94b770b0458a8440a335daf77c4c\", \"name\": \"Oscar Cartwright\", \"nominalAmount\": 256, \"pdf\": \"https://invoice-h.sandbox.starkbank.com/pdf/afdf94b770b0458a8440a335daf77c4c\", \"status\": \"created\", \"tags\": [\"war supply\", \"invoice #1234\"], \"taxId\": \"337.451.076-08\", \"transactionIds\": [], \"updated\": \"2021-04-26T20:16:51.442989+00:00\"}, \"type\": \"created\"}, \"subscription\": \"invoice\", \"workspaceId\": \"5078376503050240\"}}"
	validSignature := "MEUCIG69+s7bcS9pvvbwN0Rx9xtsVQcIuavfdJvAi2wtyHMdAiEAh/vtDWJjI76IcJvci1BNw10iM2qV57Jb5VUOLcQAZmY="

	parsed := Event.Parse(content, validSignature, nil)
	assert.NotNil(t, parsed)
}

func TestEventWrongParse(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	content := "{\"event\": {\"created\": \"2021-04-26T20:16:51.866857+00:00\", \"id\": \"5415223380934656\", \"log\": {\"created\": \"2021-04-26T20:16:50.927706+00:00\", \"errors\": [], \"id\": \"4687457496858624\", \"invoice\": {\"amount\": 256, \"brcode\": \"00020101021226890014br.gov.bcb.pix2567invoice-h.sandbox.starkbank.com/v2/afdf94b770b0458a8440a335daf77c4c5204000053039865802BR5915Stark Bank S.A.6009Sao Paulo62070503***6304CC32\", \"created\": \"2021-04-26T20:16:50.886319+00:00\", \"descriptions\": [{\"key\": \"Field1\", \"value\": \"Something\"}], \"discountAmount\": 0, \"discounts\": [{\"due\": \"2021-05-07T09:43:15+00:00\", \"percentage\": 10.0}], \"due\": \"2021-05-09T19:11:39+00:00\", \"expiration\": 123456789, \"fee\": 0, \"fine\": 2.5, \"fineAmount\": 0, \"id\": \"5941925571985408\", \"interest\": 1.3, \"interestAmount\": 0, \"link\": \"https://cdottori.sandbox.starkbank.com/invoicelink/afdf94b770b0458a8440a335daf77c4c\", \"name\": \"Oscar Cartwright\", \"nominalAmount\": 256, \"pdf\": \"https://invoice-h.sandbox.starkbank.com/pdf/afdf94b770b0458a8440a335daf77c4c\", \"status\": \"created\", \"tags\": [\"war supply\", \"invoice #1234\"], \"taxId\": \"337.451.076-08\", \"transactionIds\": [], \"updated\": \"2021-04-26T20:16:51.442989+00:00\"}, \"type\": \"created\"}, \"subscription\": \"invoice\", \"workspaceId\": \"5078376503050240\"}}"
	invalidSignature := "MEUCIQDOpo1j+V40DNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="

	parsed := Event.Parse(content, invalidSignature, nil)
	assert.NotNil(t, parsed)
}

func TestEventMalFormedParse(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	content := "{\"event\": {\"created\": \"2021-04-26T20:16:51.866857+00:00\", \"id\": \"5415223380934656\", \"log\": {\"created\": \"2021-04-26T20:16:50.927706+00:00\", \"errors\": [], \"id\": \"4687457496858624\", \"invoice\": {\"amount\": 256, \"brcode\": \"00020101021226890014br.gov.bcb.pix2567invoice-h.sandbox.starkbank.com/v2/afdf94b770b0458a8440a335daf77c4c5204000053039865802BR5915Stark Bank S.A.6009Sao Paulo62070503***6304CC32\", \"created\": \"2021-04-26T20:16:50.886319+00:00\", \"descriptions\": [{\"key\": \"Field1\", \"value\": \"Something\"}], \"discountAmount\": 0, \"discounts\": [{\"due\": \"2021-05-07T09:43:15+00:00\", \"percentage\": 10.0}], \"due\": \"2021-05-09T19:11:39+00:00\", \"expiration\": 123456789, \"fee\": 0, \"fine\": 2.5, \"fineAmount\": 0, \"id\": \"5941925571985408\", \"interest\": 1.3, \"interestAmount\": 0, \"link\": \"https://cdottori.sandbox.starkbank.com/invoicelink/afdf94b770b0458a8440a335daf77c4c\", \"name\": \"Oscar Cartwright\", \"nominalAmount\": 256, \"pdf\": \"https://invoice-h.sandbox.starkbank.com/pdf/afdf94b770b0458a8440a335daf77c4c\", \"status\": \"created\", \"tags\": [\"war supply\", \"invoice #1234\"], \"taxId\": \"337.451.076-08\", \"transactionIds\": [], \"updated\": \"2021-04-26T20:16:51.442989+00:00\"}, \"type\": \"created\"}, \"subscription\": \"invoice\", \"workspaceId\": \"5078376503050240\"}}"
	malformedSignature := "something is definitely wrong"

	parsed := Event.Parse(content, malformedSignature, nil)
	assert.NotNil(t, parsed)
}
