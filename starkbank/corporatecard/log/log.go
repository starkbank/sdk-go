package log

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/corporatecard"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	CorporateCard.Log struct
//
//	Every time a CorporateCard entity is updated, a corresponding CorporateCard.Log
//	is generated for the entity. This log is never generated by the
//	user, but it can be retrieved to check additional information
//	on the CorporateCard.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the log is created. ex: "5656565656565656"
//	- Card [CorporateCard]: CorporateCard entity to which the log refers to.
//	- Type [string]: Type of the CorporateCard event which triggered the log creation. ex: "blocked", "canceled", "created", "expired", "unblocked", "updated"
//	- Created [time.Time]: Creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Log struct {
	Id      string                      `json:",omitempty"`
	Card    corporatecard.CorporateCard `json:",omitempty"`
	Type    string                      `json:",omitempty"`
	Created *time.Time                  `json:",omitempty"`
}

var Object Log
var objects []Log
var resource = map[string]string{"name": "CorporateCardLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific CorporateCard by its id
	//
	//	Receive a single CorporateCard.Log struct previously created by the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- corporateCard.Log struct that corresponds to the given id.
	var Object Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &Object)
	if unmarshalError != nil {
		return Object, err
	}
	return Object, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve CorporateCard.Log
	//
	//	Receive a channel of CorporateCard.Log structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"blocked", "canceled", "created", "expired", "unblocked", "updated"}
	//		- cardIds [slice of strings, default nil]: Slice of CorporateCard ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- channel of CorporateCard.Log structs with updated attributes
	var Object Log
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &Object)
			if err != nil {
				print(err.Error())
			}
			logs <- Object
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	//	Retrieve paged CorporateCard.Log
	//
	//	Receive a slice of up to 100 CorporateCard.Log structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"blocked", "canceled", "created", "expired", "unblocked", "updated"}
	//		- cardIds [slice of strings, default nil]: Slice of CorporateCard ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of CorporateCard.Log structs with updated attributes
	//	- cursor to retrieve the next page of CorporateCard.Log structs
	var objects []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
