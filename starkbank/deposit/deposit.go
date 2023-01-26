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
//	- Fee [int]: Fee charged by this deposit. ex: 50 (= R$ 0.50)
//	- TransactionIds [slice of strings]: Ledger transaction ids linked to this deposit (if there are more than one, all but the first are reversals or failed reversal chargebacks). ex: []string{"19827356981273"}
//	- Created [time.Time]: Creation datetime for the Deposit. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the Deposit. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Deposit struct {
	Id             string     `json:",omitempty"`
	Name           string     `json:",omitempty"`
	TaxId          string     `json:",omitempty"`
	BankCode       string     `json:",omitempty"`
	BranchCode     string     `json:",omitempty"`
	AccountNumber  string     `json:",omitempty"`
	AccountType    string     `json:",omitempty"`
	Amount         int        `json:",omitempty"`
	Type           string     `json:",omitempty"`
	Status         string     `json:",omitempty"`
	Tags           []string   `json:",omitempty"`
	Fee            int        `json:",omitempty"`
	TransactionIds []string   `json:",omitempty"`
	Created        *time.Time `json:",omitempty"`
	Updated        *time.Time `json:",omitempty"`
}

var object Deposit
var objects []Deposit
var resource = map[string]string{"name": "Deposit"}

func Get(id string, user user.User) (Deposit, Error.StarkErrors) {
	//	Retrieve a specific Deposit by its id
	//
	//	Receive a single Deposit struct from the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- Deposit struct that corresponds to the given id
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan Deposit {
	//	Retrieve Deposit structs
	//
	//	Receive a generator of Deposit structs from the Stark Bank API
	//
	//	Parameters (required):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Parameters (optional):
	//	- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//	- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//	- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//	- sort [string, default "-created"]: Sort order considered in response. Valid options are "created" or "-created".
	//	- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//	- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Return:
	//	- Generator of Deposit structs with updated attributes
	deposits := make(chan Deposit)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			deposits <- object
		}
		close(deposits)
	}()
	return deposits
}

func Page(params map[string]interface{}, user user.User) ([]Deposit, string, Error.StarkErrors) {
	//	Retrieve paged Deposit structs
	//
	//	Receive a list of up to 100 Deposit structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (required):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Parameters (optional):
	//	- cursor [string, default nil]: Cursor returned on the previous page function call
	//	- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//	- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//	- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//	- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//	- sort [string, default "-created"]: Sort order considered in response. Valid options are "created" or "-created".
	//	- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//	- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Return:
	//	- List of Deposit structs with updated attributes
	//	- Cursor to retrieve the next page of Deposit structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
