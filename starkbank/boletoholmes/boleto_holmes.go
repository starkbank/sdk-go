package boletoholmes

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	BoletoHolmes struct
//
//	When you initialize a BoletoHolmes, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the list of created structs.
//
//	Parameters (required):
//	- BoletoId [string]: Investigated boleto entity ID. ex: "5656565656565656"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"Edward", "Stark"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when Holmes is created. ex: "5656565656565656"
//	- Status [string]: Current holmes status. ex: "solving" or "solved"
//	- Result [string]: Result of boleto status investigation. ex: "paid" or "cancelled"
//	- Created [time.Time]: Creation datetime for the holmes. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Updated [time.Time]: Latest update datetime for the holmes. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type BoletoHolmes struct {
	Id       string     `json:",omitempty"`
	BoletoId string     `json:",omitempty"`
	Tags     []string   `json:",omitempty"`
	Status   string     `json:",omitempty"`
	Result   string     `json:",omitempty"`
	Created  *time.Time `json:",omitempty"`
	Updated  *time.Time `json:",omitempty"`
}

var object BoletoHolmes
var objects []BoletoHolmes
var resource = map[string]string{"name": "BoletoHolmes"}

func Create(holmes []BoletoHolmes, user user.User) ([]BoletoHolmes, Error.StarkErrors) {
	//	Create BoletoHolme structs
	//
	//	Send a list of BoletoHolmes structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- holmes [slice of BoletoHolmes struct]: List of BoletoHolmes structs to be created in the API
	//
	//  Parameters (optional)
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of BoletoHolmes structs with updated attributes
	var objects []BoletoHolmes
	create, err := utils.Multi(resource, holmes, nil, user)
	unmarshalError := json.Unmarshal(create, &objects)
	if unmarshalError != nil {
		return objects, err
	}
	return objects, err
}

func Get(id string, user user.User) (BoletoHolmes, Error.StarkErrors) {
	//	Retrieve a specific BoletoHolmes by its id
	//
	//	Receive a single BoletoHolmes struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//  Parameters (optional)
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- BoletoHolmes struct that corresponds to the given id
	var object BoletoHolmes
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan BoletoHolmes {
	//	Retrieve BoletoHolmes structs
	//
	//	Receive a channel of BoletoHolmes structs previously created in the Stark Bank API
	//
	//	Parameters (required):
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//		- boletoId [string, default nil]: Filter for holmes that investigate a specific boleto by its ID. ex: "5656565656565656"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of BoletoHolmes structs with updated attributes
	var object BoletoHolmes
	holmes := make(chan BoletoHolmes)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			holmes <- object
		}
		close(holmes)
	}()
	return holmes
}

func Page(params map[string]interface{}, user user.User) ([]BoletoHolmes, string, Error.StarkErrors) {
	//	Retrieve paged BoletoHolmes structs
	//
	//	Receive a slice of up to 100 BoletoHolmes structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//		- boletoId [string, default nil]: Filter for holmes that investigate a specific boleto by its ID. ex: "5656565656565656"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of BoletoHolmes structs with updated attributes
	//	- Cursor to retrieve the next page of BoletoHolmes structs
	var objects []BoletoHolmes
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
