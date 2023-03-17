package taxpayment

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	TaxPayment struct
//
//	When you initialize a TaxPayment, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the slice of created structs.
//
//	Parameters (required):
//	- Description [string]: Text to be displayed in your statement (min. 10 characters). ex: "payment ABC"
//
//	Parameters (conditionally required):
//	- Line [string, default nil]: Number sequence that describes the payment. Either 'line' or 'barCode' parameters are required. If both are sent, they must match. ex: "85800000003 0 28960328203 1 56072020190 5 22109674804 0"
//	- BarCode [string, default nil]: Bar code number that describes the payment. Either 'line' or 'barCode' parameters are required. If both are sent, they must match. ex: "83660000001084301380074119002551100010601813"
//
//	Parameters (optional):
//	- Scheduled [time.Time, default today]: payment scheduled date. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- Tags [slice of strings]: slice of strings for tagging. ex: []string{"John", "Paul"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when payment is created. ex: "5656565656565656"
//	- Type [string]: tax type. ex: "das"
//	- Status [string]: current payment status. ex: "success" or "failed"
//	- Amount [int]: amount automatically calculated from line or bar_code. ex: 23456 (= R$ 234.56)
//	- Fee [int]: fee charged when tax payment is created. ex: 200 (= R$ 2.00)
//	- TransactionIds [slice of strings]: ledger transaction ids linked to this TaxPayment. []string{"19827356981273"}
//	- Updated [time.Time]: latest update datetime for the payment. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC),
//	- Created [time.Time]: creation datetime for the payment. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC),

type TaxPayment struct {
	Id             string     `json:",omitempty"`
	Line           string     `json:",omitempty"`
	BarCode        string     `json:",omitempty"`
	Description    string     `json:",omitempty"`
	Scheduled      *time.Time `json:",omitempty"`
	Tags           []string   `json:",omitempty"`
	Status         string     `json:",omitempty"`
	Amount         int        `json:",omitempty"`
	Fee            int        `json:",omitempty"`
	Type           string     `json:",omitempty"`
	TransactionIds []string   `json:",omitempty"`
	Updated        *time.Time `json:",omitempty"`
	Created        *time.Time `json:",omitempty"`
}

var Object TaxPayment
var objects []TaxPayment
var resource = map[string]string{"name": "TaxPayment"}

func Create(payments []TaxPayment, user user.User) ([]TaxPayment, Error.StarkErrors) {
	//	Create TaxPayments
	//
	//	Send a slice of TaxPayment structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- payments [slice of TaxPayment structs]: slice of TaxPayment structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of TaxPayment structs with updated attributes
	create, err := utils.Multi(resource, payments, nil, user)
	unmarshalError := json.Unmarshal(create, &payments)
	if unmarshalError != nil {
		return payments, err
	}
	return payments, err
}

func Get(id string, user user.User) (TaxPayment, Error.StarkErrors) {
	//	Retrieve a specific TaxPayment
	//
	//	Receive a single TaxPayment struct previously created by the Stark Bank API by passing its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- TaxPayment struct with updated attributes
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &Object)
	if unmarshalError != nil {
		return Object, err
	}
	return Object, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific TaxPayment .pdf file
	//
	//	Receive a single TaxPayment pdf file generated in the Stark Bank API by passing its id.
	//	Only valid for tax payments with "success" status.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- TaxPayment .pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}

func Query(params map[string]interface{}, user user.User) chan TaxPayment {
	//	Retrieve TaxPayment structs
	//
	//	Receive a channel of TaxPayment structs previously created by this user in the Stark Bank API
	//
	//	Parameters (required):
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: date filter for structs created only after specified date.
	//		- before [string, default nil]: date filter for structs created only before specified date.
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	 - Channel of TaxPayment structs with updated attributes\
	payments := make(chan TaxPayment)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &Object)
			if err != nil {
				panic(err)
			}
			payments <- Object
		}
		close(payments)
	}()
	return payments
}

func Page(params map[string]interface{}, user user.User) ([]TaxPayment, string, Error.StarkErrors) {
	//	Retrieve paged TaxPayment structs
	//
	//	Receive a slice of up to 100 TaxPayment structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	// 	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: date filter for structs created only after specified date.
	//		- before [string, default nil]: date filter for structs created only before specified date.
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of TaxPayment structs with updated attributes
	//	- cursor to retrieve the next page of TaxPayment structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}

func Delete(id string, user user.User) (TaxPayment, Error.StarkErrors) {
	//	Delete a TaxPayment struct
	//
	//	Delete a TaxPayment struct previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: TaxPayment unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- deleted TaxPayment struct
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &Object)
	if unmarshalError != nil {
		return Object, err
	}
	return Object, err
}
