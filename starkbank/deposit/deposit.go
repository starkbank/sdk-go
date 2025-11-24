package deposit

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	Deposit struct
//
//	Deposits represent passive cash-in received by your account from external transfers
//
//	Attributes (return-only):
//	- Id [string]: Unique id associated with a Deposit when it is created. ex: "5656565656565656"
//	- Name [string]: Payer name. ex: "Iron Bank S.A."
//	- TaxId [string]: Payer tax ID (CPF or CNPJ). ex: "012.345.678-90" or "20.018.183/0001-80"
//	- BankCode [string]: Payer bank code in Brazil. ex: "20018183" or "341"
//	- BranchCode [string]: Payer bank account branch. ex: "1357-9"
//	- AccountNumber [string]: Payer bank account number. ex: "876543-2"
//	- AccountType [string]: Payer bank account type. ex: "checking"
//	- Amount [int]: Deposit value in cents. ex: 1234 (= R$ 12.34)
//	- Type [string]: Type of settlement that originated the deposit. ex: "pix" or "ted"
//	- Status [string]: Current Deposit status. ex: "created"
//	- Tags [slice of strings]: Slice of strings that are tagging the deposit. ex: []string{"reconciliationId", "txId"}
//	- DisplayDescription [string, default nil]: optional description to be shown in the receiver bank interface. ex: "Payment for service 1234"
//	- Fee [int]: Fee charged by this deposit. ex: 50 (= R$ 0.50)
//	- TransactionIds [slice of strings]: Ledger transaction ids linked to this deposit (if there are more than one, all but the first are reversals or failed reversal chargebacks). ex: []string{"19827356981273"}
//	- Created [time.Time]: Creation datetime for the Deposit. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the Deposit. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Deposit struct {
	Id             		string     `json:",omitempty"`
	Name           		string     `json:",omitempty"`
	TaxId          		string     `json:",omitempty"`
	BankCode       		string     `json:",omitempty"`
	BranchCode     		string     `json:",omitempty"`
	AccountNumber  		string     `json:",omitempty"`
	AccountType    		string     `json:",omitempty"`
	Amount         		int        `json:",omitempty"`
	Type           		string     `json:",omitempty"`
	Status         		string     `json:",omitempty"`
	Tags           		[]string   `json:",omitempty"`
	DisplayDescription  string	   `json:",omitempty"`
	Fee            		int        `json:",omitempty"`
	TransactionIds 		[]string   `json:",omitempty"`
	Created        		*time.Time `json:",omitempty"`
	Updated        		*time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "Deposit"}

func Get(id string, user user.User) (Deposit, Error.StarkErrors) {
	//	Retrieve a specific Deposit by its id
	//
	//	Receive a single Deposit struct from the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Deposit struct that corresponds to the given id
	var deposit Deposit
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &deposit)
	if unmarshalError != nil {
		return deposit, err
	}
	return deposit, err
}

func Query(params map[string]interface{}, user user.User) (chan Deposit, chan Error.StarkErrors) {
	//	Retrieve Deposit structs
	//
	//	Receive a channel of Deposit structs from the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//		- sort [string, default "-created"]: Sort order considered in response. Valid options are "created" or "-created".
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of Deposit structs with updated attributes
	var deposit Deposit
	deposits := make(chan Deposit)
	depositsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &deposit)
			if err != nil {
				depositsError <- Error.UnknownError(err.Error())
				continue
			}
			deposits <- deposit
		}
		for err := range errorChannel {
			depositsError <- err
		}
		close(deposits)
		close(depositsError)
	}()
	return deposits, depositsError
}

func Page(params map[string]interface{}, user user.User) ([]Deposit, string, Error.StarkErrors) {
	//	Retrieve paged Deposit structs
	//
	//	Receive a slice of up to 100 Deposit structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//		- sort [string, default "-created"]: Sort order considered in response. Valid options are "created" or "-created".
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Deposit structs with updated attributes
	//	- Cursor to retrieve the next page of Deposit structs
	var deposit []Deposit
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &deposit)
	if unmarshalError != nil {
		return deposit, cursor, err
	}
	return deposit, cursor, err
}

func Update(id string, amount int, user user.User) (Deposit, Error.StarkErrors) {
	//	Update Deposit entity
	//
	//	Update a Deposit by passing its id to be partially or fully reversed.
	//
	//	Parameters (required):
	//	- amount [int]: the new amount of the Deposit. If the amount = 0 the Deposit will be fully reversed
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Target Deposit with updated attributes
	var deposit Deposit

	payload := map[string]interface{} {
		"amount": amount,
	}

	update, err := utils.Patch(resource, id, payload, user)
	unmarshalError := json.Unmarshal(update, &deposit)
	if unmarshalError != nil {
		return deposit, err
	}
	return deposit, err
}
