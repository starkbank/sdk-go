package event

import (
	"encoding/json"
	BoletoLog "github.com/starkbank/sdk-go/starkbank/boleto/log"
	HolmesLog "github.com/starkbank/sdk-go/starkbank/boletoholmes/log"
	BoletoPaymentLog "github.com/starkbank/sdk-go/starkbank/boletopayment/log"
	BrcodePaymentLog "github.com/starkbank/sdk-go/starkbank/brcodepayment/log"
	DarfPaymentLog "github.com/starkbank/sdk-go/starkbank/darfpayment/log"
	DepositLog "github.com/starkbank/sdk-go/starkbank/deposit/log"
	InvoiceLog "github.com/starkbank/sdk-go/starkbank/invoice/log"
	TaxPaymentLog "github.com/starkbank/sdk-go/starkbank/taxpayment/log"
	TransferLog "github.com/starkbank/sdk-go/starkbank/transfer/log"
	UtilityPaymentLog "github.com/starkbank/sdk-go/starkbank/utilitypayment/log"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	Webhook Event struct
//
//	An Event is the notification received from the subscription to the Webhook.
//	Events cannot be created, but may be retrieved from the Stark Bank API to
//	list all generated updates on entities.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the event is created. ex: "5656565656565656"
//	- Log [Log]: A Log struct from one of the subscribed services (TransferLog, InvoiceLog, DepositLog, BoletoLog, BoletoHolmesLog, BrcodePaymentLog, BoletoPaymentLog, UtilityPaymentLog, TaxPaymentLog or DarfPaymentLog)
//	- Created [string]: Creation datetime for the notification event. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- IsDelivered [bool]: True if the event has been successfully delivered to the user url. ex: False
//	- Subscription [string]: Service that triggered this event. ex: "transfer", "utility-payment"
//	- WorkspaceId [string]: Id of the Workspace that generated this event. Mostly used when multiple Workspaces have Webhooks registered to the same endpoint. ex: "4545454545454545"

type Event struct {
	Id           string      `json:",omitempty"`
	Log          interface{} `json:",omitempty"`
	Created      *time.Time  `json:",omitempty"`
	IsDelivered  bool        `json:",omitempty"`
	Subscription string      `json:",omitempty"`
	WorkspaceId  string      `json:",omitempty"`
}

var resource = map[string]string{"name": "Event"}

func Get(id string, user user.User) (Event, Error.StarkErrors) {
	//	Retrieve a specific notification Event by its id
	//
	//	Receive a single notification Event struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Event struct that corresponds to the given id
	var event Event
	get, err := utils.Get(resource, id, nil, user)
	if err.Errors != nil {
		return event, err
	}
	
	unmarshalError := json.Unmarshal(get, &event)
	if unmarshalError != nil {
		return event, Error.UnknownError(unmarshalError.Error())
	}

	parsedEvent, err := event.ParseLog()
	if err.Errors != nil {
		return event, err
	}
	return parsedEvent, Error.StarkErrors{}
}

func Query(params map[string]interface{}, user user.User) (chan Event, chan Error.StarkErrors) {
	//	Retrieve notification Event struct
	//
	//	Receive a channel of notification Event structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- isDelivered [bool, default nil]: Bool to filter successfully delivered events. ex: True or False
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of Event structs with updated attributes
	var event Event
	events := make(chan Event)
	eventsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &event)
			if err != nil {
				eventsError <- Error.UnknownError(err.Error())
				continue
			}
			parsedEvent, parseErr := event.ParseLog()
			if parseErr.Errors != nil {
				eventsError <- parseErr
				continue
			}
			events <- parsedEvent
		}
		for err := range errorChannel {
			eventsError <- err
		}
		close(events)
		close(eventsError)
	}()
	return events, eventsError
}

func Page(params map[string]interface{}, user user.User) ([]Event, string, Error.StarkErrors) {
	//	Retrieve paged Event structs
	//
	//	Receive a slice of up to 100 Event structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- isDelivered [bool, default nil]: Bool to filter successfully delivered events. ex: True or False
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Event structs with updated attributes
	//	- Cursor to retrieve the next page of Event structs
	var events []Event
	page, cursor, err := utils.Page(resource, params, user)
	if err.Errors != nil {
		return nil, "", err
	}
	
	unmarshalError := json.Unmarshal(page, &events)
	if unmarshalError != nil {
		return nil, "", Error.UnknownError(unmarshalError.Error())
	}

	parsedEvents, err := ParseEvents(events)
	if err.Errors != nil {
		return nil, "", err
	}
	return parsedEvents, cursor, Error.StarkErrors{}
}

func Delete(id string, user user.User) (Event, Error.StarkErrors) {
	//	Delete a webhook Event entity
	//
	//	Delete a of notification Event entity previously created in the Stark Bank API by its ID
	//
	//	Parameters (required):
	//	- id [string]: Event unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Deleted Event struct
	var event Event
	deleted, err := utils.Delete(resource, id, user)
	if err.Errors != nil {
		return event, err
	}

	unmarshalError := json.Unmarshal(deleted, &event)
	if unmarshalError != nil {
		return event, Error.UnknownError(unmarshalError.Error())
	}

	parsedEvent, err := event.ParseLog()
	if err.Errors != nil {
		return event, err
	}
	return parsedEvent, Error.StarkErrors{}
}

func Update(id string, patchData map[string]interface{}, user user.User) (Event, Error.StarkErrors) {
	//	Update notification Event entity
	//
	//	Update notification Event by passing id.
	//	If isDelivered is True, the event will no longer be returned on queries with is_delivered=False.
	//
	//	Parameters (required):
	//	- id [slice of strings]: Event unique ids. ex: "5656565656565656"
	//	- patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"isDelivered": true}
	//		Parameters (required):
	//		- isDelivered [bool]: If True and event hasn't been delivered already, event will be set as delivered. ex: true
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Target Event with updated attributes
	var event Event
	update, err := utils.Patch(resource, id, patchData, user)
	if err.Errors != nil {
		return event, err
	}

	unmarshalError := json.Unmarshal(update, &event)
	if unmarshalError != nil {
		return event, Error.UnknownError(unmarshalError.Error())
	}

	parsedEvent, err := event.ParseLog()
	if err.Errors != nil {
		return event, err
	}
	return parsedEvent, Error.StarkErrors{}
}

func Parse(content string, signature string, user user.User) (interface{}, Error.StarkErrors) {
	//	Create single notification Event from a content string
	//
	//	Create a single Event struct received from event listening at subscribed user endpoint.
	//	If the provided digital signature does not check out with the StarkBank public key, an
	//	error.InvalidSignatureError will be raised.
	//
	//	Parameters (required):
	//	- content [string]: Response content from request received at user endpoint (not parsed)
	//	- signature [string]: Base-64 digital signature received at response header "Digital-Signature"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Parsed Event struct
	return utils.ParseAndVerify(content, signature, "event", user)
}

func (e Event) ParseLog() (Event, Error.StarkErrors) {
	if e.Subscription == "invoice" {
		var log InvoiceLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	if e.Subscription == "boleto" {
		var log BoletoLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	if e.Subscription == "boleto-holmes" {
		var log HolmesLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)

		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	if e.Subscription == "boleto-payment" {
		var log BoletoPaymentLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	if e.Subscription == "brcode-payment" {
		var log BrcodePaymentLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	if e.Subscription == "darf-payment" {
		var log DarfPaymentLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	if e.Subscription == "deposit" {
		var log DepositLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	if e.Subscription == "tax-payment" {
		var log TaxPaymentLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	if e.Subscription == "transfer" {
		var log TransferLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	if e.Subscription == "utility-payment" {
		var log UtilityPaymentLog.Log
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &log)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Log = log
		return e, Error.StarkErrors{}
	}
	return e, Error.StarkErrors{}
}

func ParseEvents(events []Event) ([]Event, Error.StarkErrors) {
	var err Error.StarkErrors
	for i := 0; i < len(events); i++ {
		events[i], err = events[i].ParseLog()
		if err.Errors != nil {
			return nil, err
		}
	}
	return events, Error.StarkErrors{}
}
