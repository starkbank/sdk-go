package log

import (
	"encoding/json"
	TaxPayment "github.com/starkbank/sdk-go/starkbank/taxpayment"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	TaxPayment.Log struct
//
//	Every time a TaxPayment entity is modified, a corresponding taxpayment.Log
//	is generated for the entity. This log is never generated by the user, but it can
//	be retrieved to check additional information on the TaxPayment.
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the log is created. ex: "5656565656565656"
//	- Payment [TaxPayment struct]: TaxPayment entity to which the log refers to.
//	- Errors [slice of strings]: slice of errors linked to this TaxPayment event.
//	- Type [string]: type of the TaxPayment event which triggered the log creation. ex: "processing" or "success"
//	- Created [time.Time]: creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Log struct {
	Id      string                `json:",omitempty"`
	Payment TaxPayment.TaxPayment `json:",omitempty"`
	Errors  []string              `json:",omitempty"`
	Type    string                `json:",omitempty"`
	Created *time.Time            `json:",omitempty"`
}

var Object Log
var objects []Log
var resource = map[string]string{"name": "TaxPaymentLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific taxpayment.Log
	//
	//	Receive a single taxpayment.Log struct previously created by the Stark Bank API by passing its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- taxpayment.Log struct with updated attributes
	var Object Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &Object)
	if unmarshalError != nil {
		return Object, err
	}
	return Object, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve taxpayment.Log structs
	//
	//	Receive a channel of taxpayment.Log structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: date filter for structs created only after specified date.
	//		- before [string, default nil]: date filter for structs created only before specified date.
	//		- types [slice of strings, default nil]: filter retrieved structs by event types. ex: []string{"processing", "success"}
	//		- paymentIds [slice of strings, default nil]: slice of TaxPayment ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of taxpayment.Log structs with updated attributes
	var Object Log
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &Object)
			if err != nil {
				panic(err)
			}
			logs <- Object
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	//	Retrieve paged taxpayment.Logs
	//
	//	Receive a slice of up to 100 taxpayment.Log structs previously created in the Stark Bank API and the cursor to the next page
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: date filter for structs created only after specified date.
	//		- before [string, default nil]: date filter for structs created only before specified date.
	//		- types [slice of strings, default nil]: filter retrieved structs by types. ex: []string{"success", "failed"}
	//		- paymentIds [slice of strings, default nil]: slice of TaxPayment ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of taxpayment.Log structs with updated attributes
	//	- cursor to retrieve the next page of taxpayment.Log structs
	var objects []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
