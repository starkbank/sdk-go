package boletopayment

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	BoletoPayment struct
//
//	When you initialize a BoletoPayment, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the list of created structs.
//
//	Parameters (required):
//	- TaxId [string]: Receiver tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- Description [string]: Text to be displayed in your statement (min. 10 characters). ex: "payment ABC"
//
//	Parameters (conditionally required):
//	- Line [string, default nil]: Number sequence that describes the payment. Either 'line' or 'barCode' parameters are required. If both are sent, they must match. ex: "34191.09008 63571.277308 71444.640008 5 81960000000062"
//	- BarCode [string, default nil]: Bar code number that describes the payment. Either 'line' or 'barCode' parameters are required. If both are sent, they must match. ex: "34195819600000000621090063571277307144464000"
//
//	Parameters (optional):
//	- Amount [int, default nil]: Amount to be paid. If nil is informed, the current boleto value will be used. ex: 23456 (= R$ 234.56)
//	- Scheduled [time.Time, default today]: Payment scheduled date. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"John", "Paul"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when payment is created. ex: "5656565656565656"
//	- Status [string]: Current payment status. ex: "success" or "failed"
//	- Fee [int]: Fee charged when the Boleto payment is created. ex: 200 (= R$ 2.00)
//	- TransactionIds [slice of strings]: Ledger transaction ids linked to this BoletoPayment. ex: []string{"19827356981273"}
//	- Created [time.Time]: Creation datetime for the payment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type BoletoPayment struct {
	Id             string     `json:",omitempty"`
	Line           string     `json:",omitempty"`
	BarCode        string     `json:",omitempty"`
	TaxId          string     `json:",omitempty"`
	Description    string     `json:",omitempty"`
	Amount         int        `json:",omitempty"`
	Scheduled      string     `json:",omitempty"`
	Tags           []string   `json:",omitempty"`
	Status         string     `json:",omitempty"`
	Fee            int        `json:",omitempty"`
	TransactionIds []string   `json:",omitempty"`
	Created        *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "BoletoPayment"}

func Create(payments []BoletoPayment, user user.User) ([]BoletoPayment, Error.StarkErrors) {
	//	Create BoletoPayments
	//
	//	Send a list of BoletoPayment structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- payments [slice of BoletoPayment structs]: list of BoletoPayment structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of BoletoPayment structs with updated attributes
	create, err := utils.Multi(resource, payments, nil, user)
	unmarshalError := json.Unmarshal(create, &payments)
	if unmarshalError != nil {
		return payments, err
	}
	return payments, err
}

func Get(id string, user user.User) (BoletoPayment, Error.StarkErrors) {
	//	Retrieve a specific BoletoPayment by its id
	//
	//	Receive a single BoletoPayment struct previously created by the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- BoletoPayment struct with updated attributes
	var boletoPayment BoletoPayment
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &boletoPayment)
	if unmarshalError != nil {
		return boletoPayment, err
	}
	return boletoPayment, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific BoletoPayment .pdf file
	//
	//	Receive a single BoletoPayment pdf file generated in the Stark Bank API by its id.
	//	Only valid for boleto payments with "success" status.
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- BoletoPayment .pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}

func Query(params map[string]interface{}, user user.User) (chan BoletoPayment, chan Error.StarkErrors) {
	//	Retrieve BoletoPayment structs
	//
	//	Receive a channel of BoletoPayment structs previously created in the Stark Bank API
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
	//	- Channel of BoletoPayment structs with updated attributes
	var boletoPayment BoletoPayment
	payments := make(chan BoletoPayment)
	paymentsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &boletoPayment)
			if err != nil {
				paymentsError <- Error.UnknownError(err.Error())
				continue
			}
			payments <- boletoPayment
		}
		for err := range errorChannel {
			paymentsError <- err
		}
		close(payments)
		close(paymentsError)
	}()
	return payments, paymentsError
}

func Page(params map[string]interface{}, user user.User) ([]BoletoPayment, string, Error.StarkErrors) {
	//	Retrieve paged BoletoPayment structs
	//
	//	Receive a slice of up to 100 BoletoPayment structs previously created in the Stark Bank API and the cursor to the next page.
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
	//	- Slice of BoletoPayment structs with updated attributes
	//	- Cursor to retrieve the next page of BoletoPayment structs
	var boletoPayment []BoletoPayment
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &boletoPayment)
	if unmarshalError != nil {
		return boletoPayment, cursor, err
	}
	return boletoPayment, cursor, err
}

func Delete(id string, user user.User) (BoletoPayment, Error.StarkErrors) {
	//	Delete a BoletoPayment entity
	//
	//	Delete a BoletoPayment entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: BoletoPayment unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Deleted BoletoPayment struct
	var boletoPayment BoletoPayment
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &boletoPayment)
	if unmarshalError != nil {
		return boletoPayment, err
	}
	return boletoPayment, err
}
