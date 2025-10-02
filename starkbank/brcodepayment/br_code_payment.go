package brcodepayment

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/brcodepayment/rules"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	BrcodePayment struct
//
//	When you initialize a BrcodePayment, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the list of created structs.
//
//	Parameters (required):
//	- Brcode [string]: String loaded directly from the QR Code or copied from the invoice. ex: "00020126580014br.gov.bcb.pix0136a629532e-7693-4846-852d-1bbff817b5a8520400005303986540510.005802BR5908T'Challa6009Sao Paulo62090505123456304B14A"
//	- TaxId [string]: Receiver tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- Description [string]: Text to be displayed in your statement (min. 10 characters). ex: "payment ABC"
//
//	Parameters (conditionally required):
//	- Amount [int, default nil]: If the BR Code does not provide an amount, this parameter is mandatory, else it is optional. ex: 23456 (= R$ 234.56)
//
//	Parameters (optional):
//	- Scheduled [time.Time, default now]: Payment scheduled date or datetime. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Rules [slice of BrcodePayment.Rules, default nil]: slice of BrcodePayment.Rule structs for modifying transfer behavior. ex: []rule.Rule{{Key: "resendingLimit", Value: 5}},
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"John", "Paul"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when payment is created. ex: "5656565656565656"
//	- Name [string]: Receiver name. ex: "Jon Snow"
//	- Status [string]: Current payment status. ex: "success" or "failed"
//	- Type [string]: Brcode type. ex: "static" or "dynamic"
//	- TransactionIds [slice of strings]: Ledger transaction ids linked to this payment. ex: []string{"19827356981273"}
//	- Fee [int]: Fee charged by this brcode payment. ex: 50 (= R$ 0.50)
//	- Updated [time.Time]: Latest update datetime for the payment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the payment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type BrcodePayment struct {
	Id             string       `json:",omitempty"`
	Brcode         string       `json:",omitempty"`
	TaxId          string       `json:",omitempty"`
	Description    string       `json:",omitempty"`
	Amount         int          `json:",omitempty"`
	Scheduled      *time.Time   `json:",omitempty"`
	Rules          []rules.Rule `json:",omitempty"`
	Tags           []string     `json:",omitempty"`
	Name           string       `json:",omitempty"`
	Status         string       `json:",omitempty"`
	Type           string       `json:",omitempty"`
	TransactionIds []string     `json:",omitempty"`
	Fee            int          `json:",omitempty"`
	Updated        *time.Time   `json:",omitempty"`
	Created        *time.Time   `json:",omitempty"`
}

var resource = map[string]string{"name": "BrcodePayment"}

func Create(payments []BrcodePayment, user user.User) ([]BrcodePayment, Error.StarkErrors) {
	//	Create BrcodePayments
	//
	//	Send a list of BrcodePayment structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- payments [slice of BrcodePayment structs]: List of BrcodePayment structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of BrcodePayment structs with updated attributes
	create, err := utils.Multi(resource, payments, nil, user)
	unmarshalError := json.Unmarshal(create, &payments)
	if unmarshalError != nil {
		return payments, err
	}
	return payments, err
}

func Get(id string, user user.User) (BrcodePayment, Error.StarkErrors) {
	//	Retrieve a specific BrcodePayment by its id
	//
	//	Receive a single BrcodePayment struct previously created by the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- BrcodePayment struct that corresponds to the given id
	var brCodePayment BrcodePayment
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &brCodePayment)
	if unmarshalError != nil {
		return brCodePayment, err
	}
	return brCodePayment, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific BrcodePayment .pdf file
	//
	//	Receive a single BrcodePayment pdf receipt file generated in the Stark Bank API by its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- BrcodePayment .pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}

func Query(params map[string]interface{}, user user.User) (chan BrcodePayment, chan Error.StarkErrors) {
	//	Retrieve BrcodePayment structs
	//
	//	Receive a channel of BrcodePayment structs previously created in the Stark Bank API
	//
	//	Parameters (required):
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of BrcodePayment structs with updated attributes
	var brCodePayment BrcodePayment
	payments := make(chan BrcodePayment)
	paymentsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &brCodePayment)
			if err != nil {
				paymentsError <- Error.UnknownError(err.Error())
				continue
			}
			payments <- brCodePayment
		}
		for err := range errorChannel {
			paymentsError <- err
		}
		close(payments)
		close(paymentsError)
	}()
	return payments, paymentsError
}

func Page(params map[string]interface{}, user user.User) ([]BrcodePayment, string, Error.StarkErrors) {
	//	Retrieve paged BrcodePayment structs
	//
	//	Receive a slice of up to 100 BrcodePayment structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of BrcodePayment structs with updated attributes
	//	- Cursor to retrieve the next page of BrcodePayment structs
	var brCodePayments []BrcodePayment
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &brCodePayments)
	if unmarshalError != nil {
		return brCodePayments, cursor, err
	}
	return brCodePayments, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (BrcodePayment, Error.StarkErrors) {
	//	Update BrcodePayment entity
	//
	//	Update a BrcodePayment by passing its id, if it hasn't been paid yet.
	//
	//	Parameters (required):
	//	- id [string]: BrcodePayment id. ex: '5656565656565656'
	//	- patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"status": "canceled"}
	//		Parameters (required):
	//		- status [string]: If the BrcodePayment hasn't been paid yet, you may cancel it by passing "canceled" in the status
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Target BrcodePayment with updated attributes
	var brCodePayment BrcodePayment
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &brCodePayment)
	if unmarshalError != nil {
		return brCodePayment, err
	}
	return brCodePayment, err
}
