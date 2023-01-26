package event

import (
	"encoding/json"
	BoletoLog "github.com/starkbank/sdk-go/starkbank/boleto/log"
	HolmesLog "github.com/starkbank/sdk-go/starkbank/boletoholmes/log"
	BoletoPaymentLog "github.com/starkbank/sdk-go/starkbank/boletopayment/log"
	BrcodePaymentlog "github.com/starkbank/sdk-go/starkbank/brcodepayment/log"
	DarfPaymentlog "github.com/starkbank/sdk-go/starkbank/darfpayment/log"
	Depositlog "github.com/starkbank/sdk-go/starkbank/deposit/log"
	InvoiceLog "github.com/starkbank/sdk-go/starkbank/invoice/log"
	TaxPaymentlog "github.com/starkbank/sdk-go/starkbank/taxpayment/log"
	Transferlog "github.com/starkbank/sdk-go/starkbank/transfer/log"
	UtilityPaymentlog "github.com/starkbank/sdk-go/starkbank/utilitypayment/log"
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

var object Event
var objects []Event
var resource = map[string]string{"name": "Event"}

func Get(id string, user user.User) (Event, Error.StarkErrors) {
	//	Retrieve a specific notification Event by its id
	//
	//	Receive a single notification Event struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- Event struct that corresponds to the given id
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object.ParseLog(), err
	}
	return object.ParseLog(), err
}

func Query(params map[string]interface{}, user user.User) chan Event {
	//	Retrieve notification Event struct
	//
	//	Receive a generator of notification Event structs previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Parameters (optional):
	//	- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//	- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//	- isDelivered [bool, default nil]: Bool to filter successfully delivered events. ex: True or False
	//
	//	Return:
	//	- Generator of Event structs with updated attributes
	events := make(chan Event)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			events <- object.ParseLog()
		}
		close(events)
	}()
	return events
}

func Page(params map[string]interface{}, user user.User) ([]Event, string, Error.StarkErrors) {
	//	Retrieve paged Event structs
	//
	//	Receive a list of up to 100 Event structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (required):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Parameters (optional):
	//	- cursor [string, default nil]: Cursor returned on the previous page function call
	//	- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//	- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//	- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//	- isDelivered [bool, default nil]: Bool to filter successfully delivered events. ex: True or False
	//
	//	Return:
	//	- List of Event structs with updated attributes
	//	- Cursor to retrieve the next page of Event structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return ParseEvents(objects), cursor, err
	}
	return ParseEvents(objects), cursor, err
}

func Delete(id string, user user.User) (Event, Error.StarkErrors) {
	//	Delete a webhook Event entity
	//
	//	Delete a of notification Event entity previously created in the Stark Bank API by its ID
	//
	//	Parameters (required):
	//	- id [string]: Event unique id. ex: "5656565656565656"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- Deleted Event struct
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &object)
	if unmarshalError != nil {
		return object.ParseLog(), err
	}
	return object.ParseLog(), err
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
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- Target Event with updated attributes
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &object)
	if unmarshalError != nil {
		return object.ParseLog(), err
	}
	return object.ParseLog(), err
}

func Parse(content string, signature string, key string, user user.User) interface{} {
	//	Create single notification Event from a content string
	//
	//	Create a single Event struct received from event listening at subscribed user endpoint.
	//	If the provided digital signature does not check out with the StarkBank public key, a
	//	starkbank.error.InvalidSignatureError will be raised.
	//
	//	Parameters (required):
	//	- content [string]: Response content from request received at user endpoint (not parsed)
	//	- signature [string]: Base-64 digital signature received at response header "Digital-Signature"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- Parsed Event struct
	return utils.ParseAndVerify(content, signature, key, user)
}

func (e Event) ParseLog() Event {
	if e.Subscription == "invoice" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &InvoiceLog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = InvoiceLog.Object
		return e
	}
	if e.Subscription == "boleto" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &BoletoLog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = BoletoLog.Object
		return e
	}
	if e.Subscription == "boleto-holmes" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &HolmesLog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = HolmesLog.Object
		return e
	}
	if e.Subscription == "boleto-payment" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &BoletoPaymentLog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = BoletoPaymentLog.Object
		return e
	}
	if e.Subscription == "brcode-payment" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &BrcodePaymentlog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = BrcodePaymentlog.Object
		return e
	}
	if e.Subscription == "darf-payment" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &DarfPaymentlog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = DarfPaymentlog.Object
		return e
	}
	if e.Subscription == "deposit" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &Depositlog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = Depositlog.Object
		return e
	}
	if e.Subscription == "tax-payment" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &TaxPaymentlog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = TaxPaymentlog.Object
		return e
	}
	if e.Subscription == "transfer" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &Transferlog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = Transferlog.Object
		return e
	}
	if e.Subscription == "utility-payment" {
		marshal, _ := json.Marshal(e.Log)
		unmarshalError := json.Unmarshal(marshal, &UtilityPaymentlog.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Log = UtilityPaymentlog.Object
		return e
	}
	return e
}

func ParseEvents(events []Event) []Event {
	for i := 0; i < len(events); i++ {
		events[i] = events[i].ParseLog()
	}
	return events
}
