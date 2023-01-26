package paymentrequest

import (
	"encoding/json"
	BoletoPayment "github.com/starkbank/sdk-go/starkbank/boletopayment"
	BrcodePayment "github.com/starkbank/sdk-go/starkbank/brcodepayment"
	TaxPayment "github.com/starkbank/sdk-go/starkbank/taxpayment"
	Transaction "github.com/starkbank/sdk-go/starkbank/transaction"
	Transfer "github.com/starkbank/sdk-go/starkbank/transfer"
	UtilityPayment "github.com/starkbank/sdk-go/starkbank/utilitypayment"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	PaymentRequest struct
//
//	A PaymentRequest is an indirect request to access a specific cash-out service
//	(such as Transfers, BoletoPayments, etc.) which goes through the cost center
//	approval flow on our website. To emit a PaymentRequest, you must direct it to
//	a specific cost center by its ID, which can be retrieved on our website at the
//	cost center page.
//
//	Parameters (required):
//	- CenterId [string]: target cost center ID. ex: "5656565656565656"
//	- Payment [Transfer struct, BoletoPayment struct, UtilityPayment struct, BrcodePayment struct and Transaction]: payment entity that should be approved and executed.
//
//	Parameters (conditionally required):
//	- Type [string]: payment type, inferred from the payment parameter if it is not a dictionary. ex: "transfer", "boleto-payment"
//
//	Parameters (optional):
//	- Due [time.Time, default today]: Payment target date in ISO format. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- Tags [slice of strings, default nil]: slice of strings for tagging. ex: []string{"John", "Paul"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when a PaymentRequest is created. ex: "5656565656565656"
//	- Amount [int]: PaymentRequest amount. ex: 100000 = R$1.000,00
//	- Status [string]: current PaymentRequest status. ex: "pending" or "approved"
//	- Actions [slice of maps]: slice of actions that are affecting this PaymentRequest. ex: [{"type": "member", "id": "56565656565656, "action": "requested"}]
//	- Updated [time.Time]: latest update datetime for the PaymentRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: creation datetime for the PaymentRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PaymentRequest struct {
	CenterId string                   `json:",omitempty"`
	Payment  interface{}              `json:",omitempty"`
	Type     string                   `json:",omitempty"`
	Due      *time.Time               `json:",omitempty"`
	Tags     []string                 `json:",omitempty"`
	Amount   int                      `json:",omitempty"`
	Status   string                   `json:",omitempty"`
	Actions  []map[string]interface{} `json:",omitempty"`
	Updated  *time.Time               `json:",omitempty"`
	Created  *time.Time               `json:",omitempty"`
}

var object PaymentRequest
var objects []PaymentRequest
var resource = map[string]string{"name": "PaymentRequest"}

func Create(requests []PaymentRequest, user user.User) ([]PaymentRequest, Error.StarkErrors) {
	//	Create PaymentRequests
	//
	//	Send a slice of PaymentRequest structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- requests [slice of PaymentRequest structs]: slice of PaymentRequest structs to be created in the API
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- slice of PaymentRequest structs with updated attributes
	create, err := utils.Multi(resource, requests, nil, user)
	unmarshalError := json.Unmarshal(create, &requests)
	if unmarshalError != nil {
		return ParseRequests(requests), err
	}
	return ParseRequests(requests), err
}

func Query(centerId string, params map[string]interface{}, user user.User) chan PaymentRequest {
	//	Retrieve PaymentRequest structs
	//
	//	Receive a generator of PaymentRequest structs previously created by this user in the Stark Bank API
	//
	//	Parameters (required):
	//	- centerId [string]: target cost center ID. ex: "5656565656565656"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Parameters (optional):
	//	- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- after [string, default nil]: date filter for structs created or updated only after specified date.
	//	- before [string, default nil]: date filter for structs created or updated only before specified date.
	//	- sort [string, default "-created"]: sort order considered in response. Valid options are "-created" or "-due".
	//	- status [string, default nil]: filter for status of retrieved structs. ex: "success" or "failed"
	//	- type [string, default nil]: payment type, inferred from the payment parameter if it is not a dictionary. ex: "transfer", "boleto-payment"
	//	- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//	- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Return:
	//	- Generator of PaymentRequest structs with updated attributes
	var param = map[string]interface{}{}
	for k, v := range params {
		param[k] = v
	}
	param["centerId"] = centerId
	requests := make(chan PaymentRequest)
	query := utils.Query(resource, param, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			requests <- object.ParseRequest()
		}
		close(requests)
	}()
	return requests
}

func Page(centerId string, params map[string]interface{}, user user.User) ([]PaymentRequest, string, Error.StarkErrors) {
	//	Retrieve paged PaymentRequest structs
	//
	//	Receive a slice of up to 100 PaymentRequest structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (required):
	//	- centerId [string]: target cost center ID. ex: "5656565656565656"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Parameters (optional):
	//	- cursor [string, default nil]: cursor returned on the previous page function call
	//	- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//	- after [string, default nil]: date filter for structs created or updated only after specified date.
	//	- before [string, default nil]: date filter for structs created or updated only before specified date.
	//	- sort [string, default "-created"]: sort order considered in response. Valid options are "-created" or "-due".
	//	- status [string, default nil]: filter for status of retrieved structs. ex: "success" or "failed"
	//	- type [string, default nil]: payment type, inferred from the payment parameter if it is not a dictionary. ex: "transfer", "boleto-payment"
	//	- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//	- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Return:
	//	- slice of PaymentRequest structs with updated attributes
	//	- Cursor to retrieve the next page of PaymentRequest structs
	var param = map[string]interface{}{}
	for k, v := range params {
		param[k] = v
	}
	param["centerId"] = centerId
	page, cursor, err := utils.Page(resource, param, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return ParseRequests(objects), cursor, err
	}
	return ParseRequests(objects), cursor, err
}

func (e PaymentRequest) ParseRequest() PaymentRequest {
	if e.Type == "transfer" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &Transfer.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Payment = Transfer.Object
		return e
	}
	if e.Type == "transaction" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &Transaction.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Payment = Transaction.Object
		return e
	}
	if e.Type == "tax-payment" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &TaxPayment.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Payment = TaxPayment.Object
		return e
	}
	if e.Type == "brcode-payment" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &BrcodePayment.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Payment = BrcodePayment.Object
		return e
	}
	if e.Type == "boleto-payment" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &BoletoPayment.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Payment = BoletoPayment.Object
		return e
	}
	if e.Type == "utility-payment" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &UtilityPayment.Object)
		if unmarshalError != nil {
			panic(unmarshalError)
		}
		e.Payment = UtilityPayment.Object
		return e
	}
	return e
}

func ParseRequests(previews []PaymentRequest) []PaymentRequest {
	for i := 0; i < len(previews); i++ {
		previews[i] = previews[i].ParseRequest()
	}
	return previews
}
