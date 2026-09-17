package verifiedaccount

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	VerifiedAccount struct
//
//	When you initialize a VerifiedAccount, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends a slice of VerifiedAccount
//	structs to the Stark Bank API and returns the slice of created structs.
//
//	Parameters (required):
//	- TaxId [string]: receiver tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//
//	Parameters (conditionally required):
//	- BankCode [string]: code of the receiver bank institution in Brazil. If an ISPB (8 digits) is informed, a Pix transfer will be created, else a TED will be issued. Required if verifying with bank details. ex: "20018183" or "341"
//	- BranchCode [string]: receiver bank account branch. Use '-' in case there is a verifier digit. Required if verifying with bank details. ex: "1357-9"
//	- KeyId [string]: pix key identifier. Required if verifying with a Pix key. ex: "tony@starkbank.com" or "012.345.678-90"
//	- Name [string]: receiver full name. Required if verifying with bank details. ex: "Anthony Edward Stark"
//	- Number [string]: receiver bank account number. Use '-' before the verifier digit. Required if verifying with bank details. ex: "876543-2"
//	- Type [string]: verified account type. Required if verifying with bank details. ex: "checking", "savings", "salary" or "payment"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: slice of strings for reference when searching for verified accounts. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the VerifiedAccount is created. ex: "5656565656565656"
//	- BankName [string]: bank name associated with the verified account. ex: "Stark Bank"
//	- Status [string]: current verified account status. ex: "creating", "created", "processing", "active", "failed" or "canceled"
//	- Created [time.Time]: creation datetime for the VerifiedAccount. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC)
//	- Updated [time.Time]: latest update datetime for the VerifiedAccount. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC)

type VerifiedAccount struct {
	Id         string     `json:",omitempty"`
	TaxId      string     `json:",omitempty"`
	BankCode   string     `json:",omitempty"`
	BranchCode string     `json:",omitempty"`
	KeyId      string     `json:",omitempty"`
	Name       string     `json:",omitempty"`
	Number     string     `json:",omitempty"`
	Type       string     `json:",omitempty"`
	Tags       []string   `json:",omitempty"`
	BankName   string     `json:",omitempty"`
	Status     string     `json:",omitempty"`
	Created    *time.Time `json:",omitempty"`
	Updated    *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "VerifiedAccount"}

func Create(verifiedAccounts []VerifiedAccount, user user.User) ([]VerifiedAccount, Error.StarkErrors) {
	//	Create VerifiedAccounts
	//
	//	Send a slice of VerifiedAccount structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- verifiedAccounts [slice of VerifiedAccount structs]: slice of VerifiedAccount structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of VerifiedAccount structs with updated attributes
	create, err := utils.Multi(resource, verifiedAccounts, nil, user)
	unmarshalError := json.Unmarshal(create, &verifiedAccounts)
	if unmarshalError != nil {
		return verifiedAccounts, err
	}
	return verifiedAccounts, err
}

func Get(id string, user user.User) (VerifiedAccount, Error.StarkErrors) {
	//	Retrieve a specific VerifiedAccount
	//
	//	Receive a single VerifiedAccount struct previously created in the Stark Bank API by passing its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- VerifiedAccount struct with updated attributes
	var verifiedAccount VerifiedAccount
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &verifiedAccount)
	if unmarshalError != nil {
		return verifiedAccount, err
	}
	return verifiedAccount, err
}

func Cancel(id string, user user.User) (VerifiedAccount, Error.StarkErrors) {
	//	Cancel a VerifiedAccount entity
	//
	//	Cancel a VerifiedAccount entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: VerifiedAccount unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- canceled VerifiedAccount struct
	var verifiedAccount VerifiedAccount
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &verifiedAccount)
	if unmarshalError != nil {
		return verifiedAccount, err
	}
	return verifiedAccount, err
}

func Query(params map[string]interface{}, user user.User) (chan VerifiedAccount, chan Error.StarkErrors) {
	//	Retrieve VerifiedAccounts
	//
	//	Receive a channel of VerifiedAccount structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: date filter for structs created or updated only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: date filter for structs created or updated only before specified date. ex: "2020-03-10"
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "creating", "created", "processing", "active", "failed" or "canceled"
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of VerifiedAccount structs with updated attributes
	verifiedAccounts := make(chan VerifiedAccount)
	verifiedAccountsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var verifiedAccount VerifiedAccount
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &verifiedAccount)
			if err != nil {
				verifiedAccountsError <- Error.UnknownError(err.Error())
				continue
			}
			verifiedAccounts <- verifiedAccount
		}
		for err := range errorChannel {
			verifiedAccountsError <- err
		}
		close(verifiedAccounts)
		close(verifiedAccountsError)
	}()
	return verifiedAccounts, verifiedAccountsError
}

func Page(params map[string]interface{}, user user.User) ([]VerifiedAccount, string, Error.StarkErrors) {
	//	Retrieve paged VerifiedAccounts
	//
	//	Receive a slice of up to 100 VerifiedAccount structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: date filter for structs created or updated only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: date filter for structs created or updated only before specified date. ex: "2020-03-10"
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "creating", "created", "processing", "active", "failed" or "canceled"
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of VerifiedAccount structs with updated attributes
	//	- Cursor to retrieve the next page of VerifiedAccount structs
	var verifiedAccounts []VerifiedAccount
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &verifiedAccounts)
	if unmarshalError != nil {
		return verifiedAccounts, cursor, err
	}
	return verifiedAccounts, cursor, err
}
