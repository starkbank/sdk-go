package corporatewithdrawal

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	CorporateWithdrawal struct
//
//	The CorporateWithdrawal structs created in your Workspace return cash from your Corporate balance to your
//	Banking balance.
//
//	Parameters (required):
//	- Amount [int]: CorporateWithdrawal value in cents. Minimum = 0 (any value will be accepted). ex: 1234 (= R$ 12.34)
//	- ExternalId [string] CorporateWithdrawal external ID. ex: "12345"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"tony", "stark"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when CorporateWithdrawal is created. ex: "5656565656565656"
//	- TransactionId [string]: Stark Bank ledger transaction ids linked to this CorporateWithdrawal
//	- CorporateTransactionId [string]: Corporate ledger transaction ids linked to this CorporateWithdrawal
//	- Updated [time.Time]: Latest update datetime for the CorporateWithdrawal. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the CorporateWithdrawal. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type CorporateWithdrawal struct {
	Amount                 int        `json:",omitempty"`
	ExternalId             string     `json:",omitempty"`
	Tags                   []string   `json:",omitempty"`
	Id                     string     `json:",omitempty"`
	TransactionId          string     `json:",omitempty"`
	CorporateTransactionId string     `json:",omitempty"`
	Updated                *time.Time `json:",omitempty"`
	Created                *time.Time `json:",omitempty"`
}

var object CorporateWithdrawal
var objects []CorporateWithdrawal
var resource = map[string]string{"name": "CorporateWithdrawal"}

func Create(withdrawal CorporateWithdrawal, user user.User) (CorporateWithdrawal, Error.StarkErrors) {
	//	Create a CorporateWithdrawal
	//
	//	Send a single CorporateWithdrawal struct for creation at the Stark Bank API
	//
	//	Parameters (required):
	//	- withdrawal [CorporateWithdrawal struct]: CorporateWithdrawal struct to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- corporateWithdrawal struct with updated attributes
	create, err := utils.Single(resource, withdrawal, user)
	unmarshalError := json.Unmarshal(create, &withdrawal)
	if unmarshalError != nil {
		return withdrawal, err
	}
	return withdrawal, err
}

func Get(id string, user user.User) (CorporateWithdrawal, Error.StarkErrors) {
	//	Retrieve a specific CorporateWithdrawal by its id
	//
	//	Receive a single CorporateWithdrawal struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- corporateWithdrawal struct that corresponds to the given id.
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan CorporateWithdrawal {
	//	Retrieve CorporateWithdrawal structs
	//
	//	Receive a channel of CorporateWithdrawal structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- channel of CorporateWithdrawal structs with updated attributes
	withdrawals := make(chan CorporateWithdrawal)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err.Error())
			}
			withdrawals <- object
		}
		close(withdrawals)
	}()
	return withdrawals
}

func Page(params map[string]interface{}, user user.User) ([]CorporateWithdrawal, string, Error.StarkErrors) {
	//	Retrieve paged CorporateWithdrawal structs
	//
	//	Receive a slice of up to 100 CorporateWithdrawal structs previously created in the Stark Bank API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of CorporateWithdrawal structs with updated attributes
	//	- cursor to retrieve the next page of CorporateWithdrawal structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
