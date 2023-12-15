package log

import (
	"encoding/json"
	Holmes "github.com/starkbank/sdk-go/starkbank/boletoholmes"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	Boleto.Log struct
//
//	Every time a BoletoHolmes entity is updated, a corresponding BoletoHolmes.Log
//	is generated for the entity. This log is never generated by the
//	user, but it can be retrieved to check additional information
//	on the BoletoHolmes.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the log is created. ex: "5656565656565656"
//	- Holmes [BoletoHolmes struct]: BoletoHolmes entity to which the log refers to.
//	- Errors [slice of strings]: List of errors linked to this BoletoHolmes event
//	- Type [string]: Type of the BoletoHolmes event which triggered the log creation. ex: "solving" or "solved"
//	- Created [time.Time]: Creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Log struct {
	Id      string              `json:",omitempty"`
	Holmes  Holmes.BoletoHolmes `json:",omitempty"`
	Errors  []string            `json:",omitempty"`
	Type    string              `json:",omitempty"`
	Created *time.Time          `json:",omitempty"`
	Updated *time.Time          `json:",omitempty"`
}

var resource = map[string]string{"name": "BoletoHolmesLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific BoletoHolmes.Log by its id
	//
	//	Receive a single BoletoHolmes.Log struct previously created by the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- BoletoHolmes.Log struct that corresponds to the given id
	var boletoHolmesLog Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &boletoHolmesLog)
	if unmarshalError != nil {
		return boletoHolmesLog, err
	}
	return boletoHolmesLog, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve BoletoHolmes.Log structs
	//
	//	Receive a channel of BoletoHolmes.Log structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"solving", "solved"}
	//		- holmesIds [slice of strings, default nil]: List of BoletoHolmes ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of BoletoHolmes.Log structs with updated attributes
	var boletoHolmesLog Log
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &boletoHolmesLog)
			if err != nil {
				panic(err)
			}
			logs <- boletoHolmesLog
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	//	Retrieve paged BoletoHolmes.Log structs
	//
	//	Receive a slice of up to 100 BoletoHolmes.Log structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"solving", "solved"}
	//		- holmesIds [slice of strings, default nil]: List of BoletoHolmes ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of BoletoHolmes.Log structs with updated attributes
	//	- Cursor to retrieve the next page of BoletoHolmes.Log structs
	var boletoHolmesLogs []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &boletoHolmesLogs)
	if unmarshalError != nil {
		return boletoHolmesLogs, cursor, err
	}
	return boletoHolmesLogs, cursor, err
}
