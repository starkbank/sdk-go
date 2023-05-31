package corporatetransaction

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	CorporateTransaction struct
//
//	The CorporateTransaction structs created in your Workspace to represent each balance shift.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when CorporateTransaction is created. ex: "5656565656565656"
//	- Amount [int]: CorporateTransaction value in cents. ex: 1234 (= R$ 12.34)
//	- Balance [int]: Balance amount of the Workspace at the instant of the Transaction in cents. ex: 200 (= R$ 2.00)
//	- Description [string]: CorporateTransaction description. ex: "Buying food"
//	- Source [string]: Source of the transaction. ex: "corporate-purchase/5656565656565656"
//	- Tags [slice of string]: Slice of strings inherited from the source resource. ex: []string{"tony", "stark"}
//	- Created [time.Time]: Creation datetime for the CorporateTransaction. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type CorporateTransaction struct {
	Id          string     `json:",omitempty"`
	Amount      int        `json:",omitempty"`
	Balance     int        `json:",omitempty"`
	Description string     `json:",omitempty"`
	Source      string     `json:",omitempty"`
	Tags        []string   `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
}

var object CorporateTransaction
var objects []CorporateTransaction
var resource = map[string]string{"name": "CorporateTransaction"}

func Get(id string, user user.User) (CorporateTransaction, Error.StarkErrors) {
	//	Retrieve a specific CorporateTransaction by its id
	//
	//	Receive a single CorporateTransaction struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- corporateTransaction struct that corresponds to the given id.
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan CorporateTransaction {
	//	Retrieve CorporateTransaction structs
	//
	//	Receive a channel of CorporateTransaction structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "approved", "canceled", "denied", "confirmed" or "voided"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: Purchase IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- channel of CorporateTransaction structs with updated attributes
	transactions := make(chan CorporateTransaction)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err.Error())
			}
			transactions <- object
		}
		close(transactions)
	}()
	return transactions
}

func Page(params map[string]interface{}, user user.User) ([]CorporateTransaction, string, Error.StarkErrors) {
	//	Retrieve paged CorporateTransaction structs
	//
	//	Receive a slice of up to 100 CorporateTransaction structs previously created in the Stark Bank API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "approved", "canceled", "denied", "confirmed" or "voided"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: Purchase IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of CorporateTransaction structs with updated attributes
	//	- cursor to retrieve the next page of CorporatePurchase structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
