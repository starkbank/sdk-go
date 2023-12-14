package log

import (
	"encoding/json"
	Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	Invoice.Log struct
//
//	Every time a Invoice entity is updated, a corresponding Invoice.Log
//	is generated for the entity. This log is never generated by the
//	user, but it can be retrieved to check additional information
//	on the Invoice.
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the log is created. ex: "5656565656565656"
//	- Invoice [Invoice struct]: Invoice entity to which the log refers to.
//	- Errors [slice of strings]: list of errors linked to this Invoice event
//	- Type [string]: type of the Invoice event which triggered the log creation. ex: "registered" or "paid"
//	- Created [time.Time]: creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Log struct {
	Id      string          `json:",omitempty"`
	Invoice Invoice.Invoice `json:",omitempty"`
	Errors  []string        `json:",omitempty"`
	Type    string          `json:",omitempty"`
	Created *time.Time      `json:",omitempty"`
}

var Object Log
var objects []Log
var resource = map[string]string{"name": "InvoiceLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific Invoice.Log by its id
	//
	//	Receive a single Invoice.Log struct previously created by the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Invoice.Log struct that corresponds to the given id.
	var Object Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &Object)
	if unmarshalError != nil {
		return Object, err
	}
	return Object, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve Invoice.Log structs
	//
	//	Receive a channel of Invoice.Log structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- types [slice of strings, default nil]: filter for log event types. []string{"paid", "registered"}
	//		- boletoIds [slice of strings, default nil]: list of Invoice ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of boleto.Log structs with updated attributes
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
	//	Retrieve paged Invoice.Log structs
	//
	//	Receive a slice of up to 100 Invoice.Log structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- types [slice of strings, default nil]: filter for log event types. []string{"paid", "registered"}
	//		- boletoIds [slice of strings, default nil]: list of Invoice ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Invoice.Log structs with updated attributes
	//	- Cursor to retrieve the next page of Invoice.Log structs
	var objects []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific Invoice .pdf file
	//
	//	Receive a reversed Invoice.Log pdf receipt file generated in the Stark Bank API by its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Invoice .pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}
