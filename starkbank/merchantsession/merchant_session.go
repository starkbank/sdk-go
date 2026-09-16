package merchantsession

import (
	"encoding/json"
	"time"
	AllowedInstallment "github.com/starkbank/sdk-go/starkbank/merchantsession/allowedinstallment"
	"github.com/starkbank/sdk-go/starkbank/utils"
	"github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

// MerchantSession struct
//
// The MerchantSession resource is used to register a card for future purchases in a secure way, guiding the holder verification (3DS) challenge when required. Once a session's purchase is approved, the card can be saved as a MerchantCard and charged directly through MerchantPurchase without going through the challenge again.
//
// Parameters (optional):
// - AllowedFundingTypes [slice of strings, default nil]: funding types accepted in this session. Options: "credit", "debit"
// - AllowedInstallments [slice of AllowedInstallment structs, default nil]: installment plans accepted in this session.
// - AllowedIps [slice of strings, default nil]: list of IPs that are allowed to use this session.
// - ChallengeMode [string, default "enabled"]: whether holder verification (3DS) is used. Options: "enabled", "disabled"
// - Tags [slice of strings, default nil]: slice of strings for tagging.
//
// Attributes (return-only):
// - Id [string]: unique id returned when the session is created. ex: "5656565656565656"
// - Status [string]: current session status.
// - Uuid [string]: unique uuid returned when the session is created, used by the client-side SDK to reference this session.
// - HolderId [string]: id of the holder linked to this session.
// - SoftDescriptor [string]: text that will be shown in the holder's bank statement.
// - Created [time.Time]: creation datetime for the session.
// - Updated [time.Time]: latest update datetime for the session.

type MerchantSession struct {
	Id                  string                                  `json:",omitempty"`
	AllowedFundingTypes []string                                `json:",omitempty"`
	AllowedInstallments []AllowedInstallment.AllowedInstallment `json:",omitempty"`
	AllowedIps          []string                                `json:",omitempty"`
	ChallengeMode       string                                  `json:",omitempty"`
	Created             *time.Time                              `json:",omitempty"`
	Expiration          int                                     `json:",omitempty"`
	Status              string                                  `json:",omitempty"`
	Tags                []string                                `json:",omitempty"`
	Updated             *time.Time                              `json:",omitempty"`
	Uuid                string                                  `json:",omitempty"`
	HolderId            string                                  `json:",omitempty"`
	SoftDescriptor      string                                  `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantSession"}

func Create(merchantSession MerchantSession, user user.User) (MerchantSession, error.StarkErrors) {
	// Create a MerchantSession
	//
	// Send a MerchantSession struct for creation in the Stark Bank API. Use the returned Uuid to initialize the client-side SDK that collects the card data and drives the holder verification (3DS) challenge, then call PostPurchase with that Uuid to complete the charge.
	//
	// Parameters (required):
	// - merchantSession [MerchantSession struct]: MerchantSession struct to be created in the API
	//
	// Parameters (optional):
	// - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - MerchantSession struct with updated attributes
	create, err := utils.Single(resource, merchantSession, user)
	unmarshalError := json.Unmarshal(create, &merchantSession)
	if unmarshalError != nil {
		return merchantSession, err
	}
	return merchantSession, err
}

func Get(id string, user user.User) (MerchantSession, error.StarkErrors) {
	// Retrieve a specific MerchantSession by its id
	//
	// Receive a single MerchantSession struct previously created in the Stark Bank API by its id
	//
	// Parameters (required):
	// - id [string]: MerchantSession unique id. ex: "5656565656565656"
	//
	// Parameters (optional):
	// - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - MerchantSession struct that corresponds to the given id.
	var merchantSession MerchantSession
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &merchantSession)
	if unmarshalError != nil {
		return merchantSession, err
	}
	return merchantSession, err
}

func Query(params map[string]interface{}, user user.User) (chan MerchantSession, chan error.StarkErrors) {
	// Retrieve MerchantSession structs
	//
	// Receive a channel of MerchantSession structs previously created in the Stark Bank API
	//
	// Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs.
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Channel of MerchantSession structs with updated attributes
	var merchantSession MerchantSession
	merchantSessions := make(chan MerchantSession)
	merchantSessionsError := make(chan error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &merchantSession)
			if err != nil {
				merchantSessionsError <- error.UnknownError(err.Error())
				continue
			}
			merchantSessions <- merchantSession
		}
		for err := range errorChannel {
			merchantSessionsError <- err
		}
		close(merchantSessions)
		close(merchantSessionsError)
	}()
	return merchantSessions, merchantSessionsError
}

func Page(params map[string]interface{}, user user.User) ([]MerchantSession, string, error.StarkErrors) {
	// Retrieve paged MerchantSession structs
	//
	// Receive a slice of up to 100 MerchantSession structs previously created in the Stark Bank API and the cursor to the next page.
	// Use this function instead of query if you want to manually page your requests.
	//
	// Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs.
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Slice of MerchantSession structs with updated attributes
	// - Cursor to retrieve the next page of MerchantSession structs
	var merchantSessions []MerchantSession
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &merchantSessions)
	if unmarshalError != nil {
		return merchantSessions, cursor, err
	}
	return merchantSessions, cursor, err
}

func PostPurchase(uuid string, payload Purchase, user user.User) (Purchase, error.StarkErrors) {
	// Complete a MerchantSession Purchase
	//
	// Send the card data and purchase details collected through the client-side SDK for the MerchantSession identified by uuid, completing the holder verification (3DS) challenge when the session requires it. Send a Purchase struct for creation in the Stark Bank API.
	//
	// Parameters (required):
	// - uuid [string]: MerchantSession unique uuid returned on creation. ex: "5656565656565656"
	// - payload [Purchase struct]: Purchase struct with the card data and purchase details to be created in the API
	//
	// Parameters (optional):
	// - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Purchase struct with updated attributes
	post, err := utils.PostSubResource(resource, payload, uuid, user, SubResourcePurchase)
	unmarshalError := json.Unmarshal(post, &purchase)
	if unmarshalError != nil {
		return purchase, err
	}
	return purchase, err
}
