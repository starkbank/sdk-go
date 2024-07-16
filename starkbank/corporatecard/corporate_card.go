package corporatecard

import (
	"encoding/json"
	"fmt"
	CorporateRule "github.com/starkbank/sdk-go/starkbank/corporaterule"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"time"
)

//	CorporateCard struct
//
//	The CorporateCard struct displays the information of the cards created in your Workspace.
//	Sensitive information will only be returned when the "expand" parameter is used, to avoid security concerns.
//
//	When you initialize a CorporateCard, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the objects
//	to the Stark Bank API and returns the created object.
//
//	Parameters (required):
//	- HolderId [string]: Card holder unique id. ex: "5656565656565656"
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when CorporateCard is created. ex: "5656565656565656"
//	- HolderName [string]: Card holder name. ex: "Tony Stark"
//	- DisplayName [string]: Card displayed name. ex: "ANTHONY STARK"
//	- Rules [slice of CorporateRule struct]: [EXPANDABLE] Slice of card spending rules.
//	- Tags [slice of strings]: Slice of strings for tagging. ex: []string{"travel", "food"}
//	- StreetLine1 [string, default sub-issuer street line 1]: Card holder main address. ex: "Av. Paulista, 200"
//	- StreetLine2 [string, default sub-issuer street line 2]: Card holder address complement. ex: "Apto. 123"
//	- District [string, default sub-issuer district]: Card holder address district / neighbourhood. ex: "Bela Vista"
//	- City [string, default sub-issuer city]: Card holder address city. ex: "Rio de Janeiro"
//	- StateCode [string, default sub-issuer state code]: Card holder address state. ex: "GO"
//	- ZipCode [string, default sub-issuer zip code]: Card holder address zip code. ex: "01311-200"
//	- Type [string]: Card type. ex: "virtual"
//	- Status [string]: Current CorporateCard status. ex: "active", "blocked", "canceled", "expired".
//	- Number [string]: [EXPANDABLE] Masked card number. Expand to unmask the value. ex: "123".
//	- SecurityCode [string]: [EXPANDABLE] Masked card verification value (cvv). Expand to unmask the value. ex: "123".
//	- Expiration [time.Time]: [EXPANDABLE] Masked card expiration datetime. Expand to unmask the value.
//	- Updated [time.Time]: Latest update datetime for the CorporateCard. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the CorporateCard. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type CorporateCard struct {
	HolderId     string                        `json:",omitempty"`
	HolderName   string                        `json:",omitempty"`
	DisplayName  string                        `json:",omitempty"`
	Rules        []CorporateRule.CorporateRule `json:",omitempty"`
	Tags         []string                      `json:",omitempty"`
	StreetLine1  string                        `json:",omitempty"`
	StreetLine2  string                        `json:",omitempty"`
	District     string                        `json:",omitempty"`
	City         string                        `json:",omitempty"`
	StateCode    string                        `json:",omitempty"`
	ZipCode      string                        `json:",omitempty"`
	Id           string                        `json:",omitempty"`
	Type         string                        `json:",omitempty"`
	Status       string                        `json:",omitempty"`
	Number       string                        `json:",omitempty"`
	SecurityCode string                        `json:",omitempty"`
	Expiration   string                        `json:",omitempty"`
	Updated      *time.Time                    `json:",omitempty"`
	Created      *time.Time                    `json:",omitempty"`
}

var resource = map[string]string{"name": "CorporateCard"}

func Create(card CorporateCard, expand map[string]interface{}, user user.User) (CorporateCard, Error.StarkErrors) {
	//	Create CorporateCard
	//
	// 	Send a CorporateCard struct for creation in the Stark Bank API
	// 	If the CorporateCard was not used in the last purchase, this resource will return it.
	//
	//	Parameters (required):
	//	- card [CorporateCard struct]: CorporateCard struct to be created in the API
	//
	//	Parameters (optional):
	//	- expand [slice of strings, default nil]: Fields to expand information. ex: []string{"rules", "securityCode", "number", "expiration"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- CorporateCard struct with updated attributes
	data := map[string][]map[string]interface{}{}
	cardResource:= api.ApiJson(card, resource)
	path := fmt.Sprintf("%v/%v", api.Endpoint(resource), "token")
	raw, err := utils.PostRaw(path, cardResource, user, expand, "", true)
	unmarshalErrorRaw := json.Unmarshal(raw.Content, &data)
	if unmarshalErrorRaw != nil {
		return card, err
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	unmarshalError := json.Unmarshal(jsonBytes, &card)
	if unmarshalError != nil {
		return card, err
	}
	return card, err
}

func Get(id string, expand map[string]interface{}, user user.User) (CorporateCard, Error.StarkErrors) {
	//	Retrieve a specific CorporateCard by its id
	//
	// 	Receive a single CorporateCard struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- expand [slice of strings, default nil]: Fields to expand information. ex: []string{"rules", "securityCode", "number", "expiration"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call.
	//
	//	Return:
	//	- CorporateCard struct that corresponds to the given id.
	var corporateCard CorporateCard
	get, err := utils.Get(resource, id, expand, user)
	unmarshalError := json.Unmarshal(get, &corporateCard)
	if unmarshalError != nil {
		return corporateCard, err
	}
	return corporateCard, err
}

func Query(params map[string]interface{}, user user.User) chan CorporateCard {
	//	Retrieve CorporateCard structs
	//
	//	Receive a channel of CorporateCards structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"active", "blocked", "canceled", "expired"}
	//		- types [slice of strings, default nil]: Card type. ex: []string{"virtual"}
	//		- holderIds [slice of strings, default nil]: Card holder IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- expand [slice of strings, default nil]: Fields to expand information. ex: []string{"rules", "securityCode", "number", "expiration"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- channel of CorporateCard structs with updated attributes
	var corporateCard CorporateCard
	cards := make(chan CorporateCard)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &corporateCard)
			if err != nil {
				print(err.Error())
			}
			cards <- corporateCard
		}
		close(cards)
	}()
	return cards
}

func Page(params map[string]interface{}, user user.User) ([]CorporateCard, string, Error.StarkErrors) {
	//	Retrieve paged CorporateCards
	//
	//	Receive a slice of up to 100 CorporateCard structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"active", "blocked", "canceled", "expired"}
	//		- types [slice of strings, default nil]: Card type. ex: []string{"virtual"}
	//		- holderIds [slice of strings, default nil]: Card holder IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- expand [slice of strings, default nil]: Fields to expand information. ex: []string{"rules", "securityCode", "number", "expiration"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of CorporateCards structs with updated attributes
	//	- cursor to retrieve the next page of CorporateCards structs
	var corporateCards []CorporateCard
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &corporateCards)
	if unmarshalError != nil {
		return corporateCards, cursor, err
	}
	return corporateCards, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (CorporateCard, Error.StarkErrors) {
	//	Update CorporateCard entity
	//
	//	Update a CorporateCard by passing its id.
	//
	//	Parameters (required):
	//	- id [string]: CorporateCard id. ex: '5656565656565656'
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"amount": 9090}
	//		Parameters (optional):
	//		- status [string, default nil]: You may block the CorporateCard by passing 'blocked' or activate by passing 'active' in the status
	//		- pin [string, default nil]: You may unlock your physical card by passing its PIN. This is also the PIN you use to authorize a purchase.
	//		- displayName [string, default nil]: Card displayed name. ex: "ANTHONY EDWARD"
	//		- rules [slice of maps, default nil]: Slice of maps with "amount": int, "currencyCode": string, "id": string, "interval": string, "name": string pairs.
	//		- tags [slice of strings]: Slice of strings for tagging
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- target CorporateCard with updated attributes
	var corporateCard CorporateCard
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &corporateCard)
	if unmarshalError != nil {
		return corporateCard, err
	}
	return corporateCard, err
}

func Cancel(id string, user user.User) (CorporateCard, Error.StarkErrors) {
	//	Cancel a CorporateCard entity
	//
	//	Cancel a CorporateCard entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: CorporateCard unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- canceled CorporateCard struct
	var corporateCard CorporateCard
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &corporateCard)
	if unmarshalError != nil {
		return corporateCard, err
	}
	return corporateCard, err
}
