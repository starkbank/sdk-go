package darfpayment

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	DarfPayment struct
//
//	When you initialize a DarfPayment, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the list of created structs.
//
//	Parameters (required):
//	- Description [string]: Text to be displayed in your statement (min. 10 characters). ex: "payment ABC"
//	- RevenueCode [string]: 4-digit tax code assigned by Federal Revenue. ex: "5948"
//	- TaxId [string]: tax id (formatted or unformatted) of the payer. ex: "12.345.678/0001-95"
//	- Competence [time.Time]: competence month of the service. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- NominalAmount [int]: amount due in cents without fee or interest. ex: 23456 (= R$ 234.56)
//	- FineAmount [int]: fixed amount due in cents for fines. ex: 234 (= R$ 2.34)
//	- InterestAmount [int]: amount due in cents for interest. ex: 456 (= R$ 4.56)
//	- Due [time.Time]: due date for payment. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//
//	Parameters (optional):
//	- ReferenceNumber [string, default nil]: number assigned to the region of the tax. ex: "08.1.17.00-4"
//	- Scheduled [time.Time, default today]: payment scheduled date. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- Tags [slice of strings, default nil]: slice of strings for tagging. ex: []string{"John", "Paul"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when payment is created. ex: "5656565656565656"
//	- Status [string]: current payment status. ex: "success" or "failed"
//	- Amount [int]: Total amount due calculated from other amounts. ex: 24146 (= R$ 241.46)
//	- Fee [int]: fee charged when the DarfPayment is processed. ex: 0 (= R$ 0.00)
//	- TransactionIds [slice of strings]: ledger transaction ids linked to this DarfPayment. ex: []string{"19827356981273"}
//	- Updated [time.Time]: Latest update datetime for the payment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: creation datetime for the payment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type DarfPayment struct {
	Id              string     `json:",omitempty"`
	Description     string     `json:",omitempty"`
	RevenueCode     string     `json:",omitempty"`
	TaxId           string     `json:",omitempty"`
	Competence      *time.Time `json:",omitempty"`
	NominalAmount   int        `json:",omitempty"`
	FineAmount      int        `json:",omitempty"`
	InterestAmount  int        `json:",omitempty"`
	Due             *time.Time `json:",omitempty"`
	ReferenceNumber string     `json:",omitempty"`
	Scheduled       *time.Time `json:",omitempty"`
	Tags            []string   `json:",omitempty"`
	Status          string     `json:",omitempty"`
	Amount          int        `json:",omitempty"`
	Fee             int        `json:",omitempty"`
	TransactionIds  []string   `json:",omitempty"`
	Updated         *time.Time `json:",omitempty"`
	Created         *time.Time `json:",omitempty"`
}

var object DarfPayment
var objects []DarfPayment
var resource = map[string]string{"name": "DarfPayment"}

func Create(payments []DarfPayment, user user.User) ([]DarfPayment, Error.StarkErrors) {
	//	Create DarfPayments
	//
	//	Send a list of DarfPayment structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- payments [slice of DarfPayment structs]: list of DarfPayment structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of DarfPayment structs with updated attributes
	create, err := utils.Multi(resource, payments, nil, user)
	unmarshalError := json.Unmarshal(create, &payments)
	if unmarshalError != nil {
		return payments, err
	}
	return payments, err
}

func Get(id string, user user.User) (DarfPayment, Error.StarkErrors) {
	//	Retrieve a specific DarfPayment by its id
	//
	//	Receive a single DarfPayment struct previously created by the Stark Bank API by passing its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- DarfPayment struct that corresponds to the given id
	var object DarfPayment
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific DarfPayment .pdf file
	//
	//	Receive a single DarfPayment pdf file generated in the Stark Bank API by passing its id.
	//	Only valid for darf payments with "success" status.
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- DarfPayment .pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}

func Query(params map[string]interface{}, user user.User) chan DarfPayment {
	//	Retrieve DarfPayment structs
	//
	//	Receive a channel of DarfPayment structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: list of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of DarfPayment structs with updated attributes
	var object DarfPayment
	payments := make(chan DarfPayment)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			payments <- object
		}
		close(payments)
	}()
	return payments
}

func Page(params map[string]interface{}, user user.User) ([]DarfPayment, string, Error.StarkErrors) {
	//	Retrieve paged DarfPayment structs
	//
	//	Receive a slice of up to 100 DarfPayment structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: list of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of DarfPayment structs with updated attributes
	//	- Cursor to retrieve the next page of DarfPayment structs
	var objects []DarfPayment
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}

func Delete(id string, user user.User) (DarfPayment, Error.StarkErrors) {
	//	Delete a DarfPayment entity
	//
	//	Delete a DarfPayment entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: DarfPayment unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Deleted DarfPayment struct
	var object DarfPayment
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}
