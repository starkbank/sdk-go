package merchantcard

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

// MerchantCard struct
//
// The MerchantCard resource stores information about cards used in approved purchases, so they can be reused in new purchases without creating a new Merchant Session.
//
// Attributes (return-only):
// - Id [string]: unique id returned when the card is created. ex: "5656565656565656"
// - Ending [string]: last 4 digits of the card number.
// - FundingType [string]: funding type. Options: "credit", "debit"
// - HolderName [string]: name of the card holder.
// - Network [string]: card network.
// - Status [string]: current card status. Options: "active", "expired", "canceled", "blocked"
// - Expiration [string]: card expiration date. ex: "2025-06"
// - Tags [slice of strings]: tags associated with the card.
// - Created [time.Time]: creation datetime for the card.
// - Updated [time.Time]: latest update datetime for the card.

type MerchantCard struct {
	Id          string     `json:",omitempty"`
	Ending      string     `json:",omitempty"`
	FundingType string     `json:",omitempty"`
	HolderName  string     `json:",omitempty"`
	Network     string     `json:",omitempty"`
	Status      string     `json:",omitempty"`
	Tags        []string   `json:",omitempty"`
	Expiration  string     `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
	Updated     *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantCard"}

func Get(id string, user user.User) (MerchantCard, Error.StarkErrors){
	// Retrieve a specific MerchantCard by its id
	//
	// Receive a single MerchantCard struct previously created in the Stark Bank API by its id
	//
	// Parameters (required):
	// - id [string]: MerchantCard unique id. ex: "5656565656565656"
	//
	// Parameters (optional):
	// - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - MerchantCard struct that corresponds to the given id.
	var merchantCard MerchantCard
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &merchantCard)
	if unmarshalError != nil {
		return merchantCard, err
	}
	return merchantCard, err
}

func Query(params map[string]interface{}, user user.User) (chan MerchantCard, chan Error.StarkErrors) {
	// Retrieve MerchantCard structs
	//
	// Receive a channel of MerchantCard structs previously created in the Stark Bank API
	//
	// Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "active" or "blocked"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Channel of MerchantCard structs with updated attributes
	merchantCards := make(chan MerchantCard)
	merchantCardsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var merchantCard MerchantCard
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &merchantCard)
			if err != nil {
				merchantCardsError <- Error.UnknownError(err.Error())
				continue
			}
			merchantCards <- merchantCard
		}
		for err := range errorChannel {
			merchantCardsError <- err
		}
		close(merchantCards)
		close(merchantCardsError)
	}()
	return merchantCards, merchantCardsError
}

func Page(params map[string]interface{}, user user.User) ([]MerchantCard, string, Error.StarkErrors) {
	// Retrieve paged MerchantCard structs
	//
	// Receive a slice of up to 100 MerchantCard structs previously created in the Stark Bank API and the cursor to the next page.
	// Use this function instead of query if you want to manually page your requests.
	//
	// Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "active" or "blocked"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Slice of MerchantCard structs with updated attributes
	// - Cursor to retrieve the next page of MerchantCard structs
	var merchantCards []MerchantCard
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &merchantCards)
	if unmarshalError != nil {
		return merchantCards, cursor, err
	}
	return merchantCards, cursor, err
}
