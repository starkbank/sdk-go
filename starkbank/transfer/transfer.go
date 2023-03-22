package transfer

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/transfer/rule"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	Transfer struct
//
//	When you initialize a Transfer, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the slice of created structs.
//
//	Parameters (required):
//	- Amount [int]: amount in cents to be transferred. ex: 1234 (= R$ 12.34)
//	- Name [string]: receiver full name. ex: "Anthony Edward Stark"
//	- TaxId [string]: receiver tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- BankCode [string]: code of the receiver bank institution in Brazil. If an ISPB (8 digits) is informed, a Pix transfer will be created, else a TED will be issued. ex: "20018183" or "341"
//	- BranchCode [string]: receiver bank account branch. Use '-' in case there is a verifier digit. ex: "1357-9"
//	- AccountNumber [string]: receiver bank account number. Use '-' before the verifier digit. ex: "876543-2"
//
//	Parameters (optional):
//	- AccountType [string, default "checking"]: Receiver bank account type. This parameter only has effect on Pix Transfers. ex: "checking", "savings", "salary" or "payment"
//	- ExternalId [string, default nil]: url safe string that must be unique among all your transfers. Duplicated external_ids will cause failures. By default, this parameter will block any transfer that repeats amount and receiver information on the same date. ex: "my-internal-id-123456"
//	- Scheduled [time.Time, default now]: date when the transfer will be processed. May be pushed to next business day if necessary. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC) or ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- Description [string, default nil]: optional description to override default description to be shown in the bank statement. ex: "Payment for service #1234"
//	- Tags [slice of strings, default nil]: slice of strings for reference when searching for transfers. ex: []string{"John", "Paul"}
//	- Rules [slice of Transfer.Rule structs, default nil]: slice of Transfer.Rule structs for modifying transfer behavior. ex: []rule.Rule{{Key: "resendingLimit", Value: 5}},
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the transfer is created. ex: "5656565656565656"
//	- Fee [int]: fee charged when the Transfer is processed. ex: 200 (= R$ 2.00)
//	- Status [string]: current transfer status. ex: "success" or "failed"
//	- TransactionIds [slice of strings]: ledger Transaction IDs linked to this Transfer (if there are two, the second is the chargeback). ex: []string{"19827356981273"}
//	- Metadata [map[string]interface{}]: object used to store additional information about the Transfer struct.
//	- Created [time.Time]: creation datetime for the transfer. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: latest update datetime for the transfer. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Transfer struct {
	Id             string                 `json:",omitempty"`
	Amount         int                    `json:",omitempty"`
	Name           string                 `json:",omitempty"`
	TaxId          string                 `json:",omitempty"`
	BankCode       string                 `json:",omitempty"`
	BranchCode     string                 `json:",omitempty"`
	AccountNumber  string                 `json:",omitempty"`
	AccountType    string                 `json:",omitempty"`
	ExternalId     string                 `json:",omitempty"`
	Scheduled      *time.Time             `json:",omitempty"`
	Description    string                 `json:",omitempty"`
	Tags           []string               `json:",omitempty"`
	Rules          []rule.Rule            `json:",omitempty"`
	Fee            int                    `json:",omitempty"`
	Status         string                 `json:",omitempty"`
	TransactionIds []string               `json:",omitempty"`
	Metadata       map[string]interface{} `json:",omitempty"`
	Created        *time.Time             `json:",omitempty"`
	Updated        *time.Time             `json:",omitempty"`
}

var Object Transfer
var objects []Transfer
var resource = map[string]string{"name": "Transfer"}

func Create(transfers []Transfer, user user.User) ([]Transfer, Error.StarkErrors) {
	//	Create Transfers
	//
	//	Send a slice of Transfer structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- transfers [slice of Transfer structs]: slice of Transfer structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Transfer structs with updated attributes
	create, err := utils.Multi(resource, transfers, nil, user)
	unmarshalError := json.Unmarshal(create, &transfers)
	if unmarshalError != nil {
		return transfers, err
	}
	return transfers, err
}

func Get(id string, user user.User) (Transfer, Error.StarkErrors) {
	//	Retrieve a specific Transfer
	//
	//	Receive a single Transfer struct previously created by the Stark Bank API by passing its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Transfer struct with updated attributes
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &Object)
	if unmarshalError != nil {
		return Object, err
	}
	return Object, err
}

func Delete(id string, user user.User) (Transfer, Error.StarkErrors) {
	//	Delete a Transfer entity
	//
	//	Delete a Transfer entity previously created in the Stark Bank API
	//
	// 	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- deleted Transfer struct
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &Object)
	if unmarshalError != nil {
		return Object, err
	}
	return Object, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific Transfer .pdf file
	//
	//	Receive a single Transfer pdf receipt file generated in the Stark Bank API by its id.
	//	Only valid for transfers with "processing" and "success" status.
	//
	// 	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Transfer .pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}

func Query(params map[string]interface{}, user user.User) chan Transfer {
	//	Retrieve Transfer structs
	//
	//	Receive a channel of Transfer structs previously created by this user in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: date filter for structs created or updated only after specified date.
	//		- before [string, default nil]: date filter for structs created or updated only before specified date.
	//		- transactionIds [slice of strings, default nil]: slice of transaction IDs linked to the desired transfers. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success" or "failed"
	//		- taxId [string, default nil]: filter for transfers sent to the specified tax ID. ex: "012.345.678-90"
	//		- sort [string, default "-created"]: sort order considered in response. Valid options are "created", "-created", "updated" or "-updated".
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	 - Channel of Transfer objects with updated attributes
	transfers := make(chan Transfer)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &Object)
			if err != nil {
				panic(err)
			}
			transfers <- Object
		}
		close(transfers)
	}()
	return transfers
}

func Page(params map[string]interface{}, user user.User) ([]Transfer, string, Error.StarkErrors) {
	//	Retrieve paged Transfer structs
	//
	//	Receive a slice of up to 100 Transfer structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	// 	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: date filter for structs created or updated only after specified date.
	//		- before [string, default nil]: date filter for structs created or updated only before specified date.
	//		- transactionIds [slice of strings, default nil]: slice of transaction IDs linked to the desired transfers. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success" or "failed"
	//		- taxId [string, default nil]: filter for transfers sent to the specified tax ID. ex: "012.345.678-90"
	//		- sort [string, default "-created"]: sort order considered in response. Valid options are "created", "-created", "updated" or "-updated".
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Transfer structs with updated attributes
	//	- Cursor to retrieve the next page of Transfer structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
