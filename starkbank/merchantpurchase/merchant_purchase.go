package merchantpurchase

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

// MerchantPurchase struct
//
// The MerchantPurchase resource is used to charge customers with credit or debit cards. A card that has never
// been used before must first go through an approved Merchant Session Purchase before it can be charged directly
// through MerchantPurchase.
//
// Parameters (required):
// - Amount [int]: amount in cents to be received. ex: 100 (= R$1.00)
// - CardId [string]: ID of the Merchant Card to be used for the purchase.
// - FundingType [string]: funding type used for the purchase. Options: "credit", "debit"
//
// Parameters (conditionally required, when ChallengeMode is "enabled"):
// - BillingCity, BillingCountryCode, BillingStateCode, BillingStreetLine1, BillingStreetLine2, BillingZipCode [string]: card holder's billing address.
// - HolderEmail, HolderPhone [string]: card holder's contact information.
// - Metadata [map[string]interface{}]: must include userAgent, timezoneOffset, userIp and language for the buyer's device when 3DS is enabled.
//
// Parameters (optional):
// - ChallengeMode [string, default "enabled"]: whether holder verification (3DS) is used. Options: "enabled", "disabled"
// - InstallmentCount [int, default 1]: number of purchase installments.
//
// Attributes (return-only):
// - Status [string]: current purchase status. Options: "created", "approved", "denied", "confirmed", "paid", "pending", "canceled", "voided", "failed"

type MerchantPurchase struct {
	Id                 string                 `json:",omitempty"`
	Amount             int                    `json:",omitempty"`
	InstallmentCount   int                    `json:",omitempty"`
	CardExpiration     string                 `json:",omitempty"`
	CardNumber         string                 `json:",omitempty"`
	CardSecurityCode   string                 `json:",omitempty"`
	HolderName         string                 `json:",omitempty"`
	HolderEmail        string                 `json:",omitempty"`
	HolderPhone        string                 `json:",omitempty"`
	HolderId           string                 `json:",omitempty"`
	SoftDescriptor     string                 `json:",omitempty"`	
	FundingType        string                 `json:",omitempty"`
	BillingCountryCode string                 `json:",omitempty"`
	BillingCity        string                 `json:",omitempty"`
	BillingStateCode   string                 `json:",omitempty"`
	BillingStreetLine1 string                 `json:",omitempty"`
	BillingStreetLine2 string                 `json:",omitempty"`
	BillingZipCode     string                 `json:",omitempty"`
	Metadata           map[string]interface{} `json:",omitempty"`
	CardEnding         string                 `json:",omitempty"`
	CardId             string                 `json:",omitempty"`
	ChallengeMode      string                 `json:",omitempty"`
	ChallengeUrl       string                 `json:",omitempty"`
	Created            *time.Time             `json:",omitempty"`
	CurrencyCode       string                 `json:",omitempty"`
	EndToEndId         string                 `json:",omitempty"`
	Fee                int                    `json:",omitempty"`
	Network            string                 `json:",omitempty"`
	Source             string                 `json:",omitempty"`
	Status             string                 `json:",omitempty"`
	Tags               []string               `json:",omitempty"`
	Updated            *time.Time             `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantPurchase"}

func Create(merchantPurchase MerchantPurchase, user user.User) (MerchantPurchase, Error.StarkErrors) {
	// Create a MerchantPurchase
	//
	// Charge a card that has already been saved and approved through a Merchant Session Purchase. Send a MerchantPurchase struct for creation in the Stark Bank API.
	//
	// Parameters (required):
	// - merchantPurchase [MerchantPurchase struct]: MerchantPurchase struct to be created in the API
	//
	// Parameters (optional):
	// - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - MerchantPurchase struct with updated attributes
	create, err := utils.Single(resource, merchantPurchase, user)
	unmarshalError := json.Unmarshal(create, &merchantPurchase)
	if unmarshalError != nil {
		return merchantPurchase, err
	}
	return merchantPurchase, err
}

func Get(id string, user user.User) (MerchantPurchase, Error.StarkErrors) {
	// Retrieve a specific MerchantPurchase by its id
	//
	// Receive a single MerchantPurchase struct previously created in the Stark Bank API by its id
	//
	// Parameters (required):
	// - id [string]: MerchantPurchase unique id. ex: "5656565656565656"
	//
	// Parameters (optional):
	// - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - MerchantPurchase struct that corresponds to the given id.
	var merchantPurchase MerchantPurchase
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &merchantPurchase)
	if unmarshalError != nil {
		return merchantPurchase, err
	}
	return merchantPurchase, err
}

func Query(params map[string]interface{}, user user.User) (chan MerchantPurchase, chan Error.StarkErrors) {
	// Retrieve MerchantPurchase structs
	//
	// Receive a channel of MerchantPurchase structs previously created in the Stark Bank API
	//
	// Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "approved"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Channel of MerchantPurchase structs with updated attributes
	merchantPurchases := make(chan MerchantPurchase)
	merchantPurchasesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var merchantPurchase MerchantPurchase
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &merchantPurchase)
			if err != nil {
				merchantPurchasesError <- Error.UnknownError(err.Error())
				continue
			}
			merchantPurchases <- merchantPurchase
		}
		for err := range errorChannel {
			merchantPurchasesError <- err
		}
		close(merchantPurchases)
		close(merchantPurchasesError)
	}()
	return merchantPurchases, merchantPurchasesError
}

func Page(params map[string]interface{}, user user.User) ([]MerchantPurchase, string, Error.StarkErrors) {
	// Retrieve paged MerchantPurchase structs
	//
	// Receive a slice of up to 100 MerchantPurchase structs previously created in the Stark Bank API and the cursor to the next page.
	// Use this function instead of query if you want to manually page your requests.
	//
	// Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "approved"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Slice of MerchantPurchase structs with updated attributes
	// - Cursor to retrieve the next page of MerchantPurchase structs
	var merchantPurchases []MerchantPurchase
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &merchantPurchases)
	if unmarshalError != nil {
		return merchantPurchases, cursor, err
	}
	return merchantPurchases, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (MerchantPurchase, Error.StarkErrors) {
	// Update a MerchantPurchase entity
	//
	// Update a MerchantPurchase by passing its id. An "approved" purchase can only have its status set to "canceled" with amount set to 0, which cancels the authorization. A "confirmed" purchase can have its status set to "reversed" with a lower amount, which debits the difference and reverses the purchase partially or totally: a partial reversal keeps status "confirmed", a full reversal moves status to "voided".
	//
	// Parameters (required):
	// - id [string]: MerchantPurchase unique id. ex: "5656565656565656"
	// - patchData [map[string]interface{}]: map containing the attributes to be updated
	//     Parameters (optional):
	//     - amount [int]: new amount for the purchase; 0 to cancel an approved purchase, or a lower value to reverse a confirmed purchase. ex: 200 (R$2.00)
	//     - status [string]: "canceled" or "reversed" to cancel or reverse the purchase
	//
	// Parameters (optional):
	// - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Target MerchantPurchase with updated attributes
	var purchase MerchantPurchase
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &purchase)
	if unmarshalError != nil {
		return purchase, err
	}
	return purchase, err
}
