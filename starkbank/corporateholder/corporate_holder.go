package corporateholder

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/corporateholder/permission"
	CorporateRule "github.com/starkbank/sdk-go/starkbank/corporaterule"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	CorporateHolder struct
//
//	The CorporateHolder describes a card holder that may group several cards.
//
//	When you initialize a CorporateHolder, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the objects
//	to the Stark Bank API and returns the created object.
//
//	Parameters (required):
//	- Name [string]: Card holder's name. ex: Jannie Lanister
//
//	Parameters (optional):
//	- CenterId [string, default nil]: target cost center ID. ex: "5656565656565656"
//	- Permissions [slice of Permission object, default nil]: slice of Permission object representing access granted to an user for a particular cardholder
//	- Rules [slice of CorporateRule object, default nil]: [EXPANDABLE] slice of holder spending rules
//	- Tags [slice of strings, default nil]: list of strings for tagging. ex: []string{"travel", "food"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when CorporateHolder is created. ex: "5656565656565656"
//	- Status [string]: Current CorporateHolder status. ex: "active", "blocked", "canceled"
//	- Updated [time.Time]: Latest update datetime for the CorporateHolder. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the CorporateHolder. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type CorporateHolder struct {
	Name        string                        `json:",omitempty"`
	CenterId    string                        `json:",omitempty"`
	Rules       []CorporateRule.CorporateRule `json:",omitempty"`
	Permissions []permission.Permission       `json:",omitempty"`
	Tags        []string                      `json:",omitempty"`
	Id          string                        `json:",omitempty"`
	Status      string                        `json:",omitempty"`
	Updated     *time.Time                    `json:",omitempty"`
	Created     *time.Time                    `json:",omitempty"`
}

var resource = map[string]string{"name": "CorporateHolder"}

func Create(holders []CorporateHolder, expand map[string]interface{}, user user.User) ([]CorporateHolder, Error.StarkErrors) {
	//	Create CorporateHolders
	//
	//	Send a slice of CorporateHolder structs for creation at the Stark Bank API
	//
	//	Parameters (required):
	//	- holders [slice of CorporateHolder structs]: Slice of CorporateHolder structs to be created in the API
	//
	//	Parameters (optional):
	//	- expand [slice of strings, default nil]: Fields to expand information. ex: []string{"rules"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of CorporateHolder structs with updated attributes
	var corporateHolders []CorporateHolder
	create, err := utils.Multi(resource, holders, expand, user)
	unmarshalError := json.Unmarshal(create, &corporateHolders)
	if unmarshalError != nil {
		return corporateHolders, err
	}
	return corporateHolders, err
}

func Get(id string, expand map[string]interface{}, user user.User) (CorporateHolder, Error.StarkErrors) {
	//	Retrieve a specific CorporateHolder by its id
	//
	//	Receive a single CorporateHolder struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//	- expand [slice of strings, default nil]: Fields to expand information. ex: []string{"rules"}
	//
	//	Return:
	//	- corporateHolder struct that corresponds to the given id.
	var corporateHolder CorporateHolder
	get, err := utils.Get(resource, id, expand, user)
	unmarshalError := json.Unmarshal(get, &corporateHolder)
	if unmarshalError != nil {
		return corporateHolder, err
	}
	return corporateHolder, err
}

func Query(params map[string]interface{}, user user.User) (chan CorporateHolder, chan Error.StarkErrors) {
	//	Retrieve CorporateHolders
	//
	//	Receive a channel of CorporateHolder structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"active", "blocked", "canceled"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- expand [string, default nil]: Fields to expand information. ex: "rules"
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- channel of CorporateHolder structs with updated attributes
	var corporateHolder CorporateHolder
	holders := make(chan CorporateHolder)
	holdersError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &corporateHolder)
			if err != nil {
				holdersError <- Error.UnknownError(err.Error())
				continue
			}
			holders <- corporateHolder
		}
		for err := range errorChannel {
			holdersError <- err
		}
		close(holders)
		close(holdersError)
	}()
	return holders, holdersError
}

func Page(params map[string]interface{}, user user.User) ([]CorporateHolder, string, Error.StarkErrors) {
	//	Retrieve CorporateHolders
	//
	//	Receive a slice of up to 100 CorporateHolder structs previously created in the Stark Bank API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"active", "blocked", "canceled"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- expand [string, default nil]: Fields to expand information. ex: "rules"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of CorporateHolder structs with updated attributes
	//	- cursor to retrieve the next page of CorporateHolder structs
	var corporateHolder []CorporateHolder
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &corporateHolder)
	if unmarshalError != nil {
		return corporateHolder, cursor, err
	}
	return corporateHolder, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (CorporateHolder, Error.StarkErrors) {
	//	Update CorporateHolder entity
	//
	//	Update a CorporateHolder by passing id, if it hasn't been paid yet.
	//
	//	Parameters (required):
	//	- id [string]: CorporateHolder id. ex: '5656565656565656'
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"amount": 9090}
	//		Parameters (optional):
	//		- centerId [string, default nil]: target cost center ID. ex: "5656565656565656"
	//		- permissions [slice of Permission corporateHolder, default nil]: slice of Permission corporateHolder representing access granted to an user for a particular cardholder.
	//		- status [string, default nil]: You may block the CorporateHolder by passing 'blocked' in the status. ex: "blocked"
	//		- name [string, default nil]: card holder name. ex: "Jaime Lannister"
	//		- tags [slice of strings, default nil]: Slice of strings for tagging
	//		- rules [slice of maps, default nil]: Slice of maps with "amount": int, "currencyCode": string, "id": string, "interval": string, "name": string pairs
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- target CorporateHolder with updated attributes
	var corporateHolder CorporateHolder
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &corporateHolder)
	if unmarshalError != nil {
		return corporateHolder, err
	}
	return corporateHolder, err
}

func Cancel(id string, user user.User) (CorporateHolder, Error.StarkErrors) {
	//	Cancel a CorporateHolder entity
	//
	//	Cancel a CorporateHolder entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: CorporateHolder unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- canceled CorporateHolder struct
	var corporateHolder CorporateHolder
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &corporateHolder)
	if unmarshalError != nil {
		return corporateHolder, err
	}
	return corporateHolder, err
}
