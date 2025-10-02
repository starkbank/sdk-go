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
//	- Payment [Transfer struct, BoletoPayment struct, UtilityPayment struct, BrcodePayment struct or Transaction]: payment entity that should be approved and executed.
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
//	- Description [string]: payment request description. ex: "Tony Stark's Suit"
//	- Status [string]: current PaymentRequest status. ex: "pending" or "approved"
//	- Actions [slice of maps]: slice of actions that are affecting this PaymentRequest. ex: [{"type": "member", "id": "56565656565656, "action": "requested"}]
//	- Updated [time.Time]: latest update datetime for the PaymentRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: creation datetime for the PaymentRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PaymentRequest struct {
	CenterId    string                   `json:",omitempty"`
	Payment     interface{}              `json:",omitempty"`
	Type        string                   `json:",omitempty"`
	Due         *time.Time               `json:",omitempty"`
	Tags        []string                 `json:",omitempty"`
	Amount      int                      `json:",omitempty"`
	Description string                   `json:",omitempty"`
	Status      string                   `json:",omitempty"`
	Actions     []map[string]interface{} `json:",omitempty"`
	Updated     *time.Time               `json:",omitempty"`
	Created     *time.Time               `json:",omitempty"`
}

var resource = map[string]string{"name": "PaymentRequest"}

func Create(requests []PaymentRequest, user user.User) ([]PaymentRequest, Error.StarkErrors) {
	//	Create PaymentRequests
	//
	//	Send a slice of PaymentRequest structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- requests [slice of PaymentRequest structs]: slice of PaymentRequest structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of PaymentRequest structs with updated attributes
	create, err := utils.Multi(resource, requests, nil, user)
	if err.Errors != nil {
		return nil, err
	}

	unmarshalError := json.Unmarshal(create, &requests)
	if unmarshalError != nil {
		return nil, Error.UnknownError(unmarshalError.Error())
	}

	parsedRequests, err := ParseRequests(requests)
	if err.Errors != nil {
		return nil, err
	}
	return parsedRequests, Error.StarkErrors{}
}

func Query(centerId string, params map[string]interface{}, user user.User) (chan PaymentRequest, chan Error.StarkErrors) {
	//	Retrieve PaymentRequest structs
	//
	//	Receive a channel of PaymentRequest structs previously created by this user in the Stark Bank API
	//
	//	Parameters (required):
	//	- centerId [string]: target cost center ID. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: date filter for structs created or updated only after specified date.
	//		- before [string, default nil]: date filter for structs created or updated only before specified date.
	//		- sort [string, default "-created"]: sort order considered in response. Valid options are "-created" or "-due".
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success" or "failed"
	//		- type [string, default nil]: payment type, inferred from the payment parameter if it is not a dictionary. ex: "transfer", "boleto-payment"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of PaymentRequest structs with updated attributes
	var param = map[string]interface{}{}
	var paymentRequest PaymentRequest
	for k, v := range params {
		param[k] = v
	}
	param["centerId"] = centerId
	requests := make(chan PaymentRequest)
	requestsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, param, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &paymentRequest)
			if err != nil {
				requestsError <- Error.UnknownError(err.Error())
				continue
			}
			parsedRequest, parseErr := paymentRequest.ParseRequest()
			if parseErr.Errors != nil {
				requestsError <- parseErr
			}
			requests <- parsedRequest
		}
		for err := range errorChannel {
			requestsError <- err
		}
		close(requests)
		close(requestsError)
	}()
	return requests, requestsError
}

func Page(centerId string, params map[string]interface{}, user user.User) ([]PaymentRequest, string, Error.StarkErrors) {
	//	Retrieve paged PaymentRequest structs
	//
	//	Receive a slice of up to 100 PaymentRequest structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (required):
	//	- centerId [string]: target cost center ID. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: date filter for structs created or updated only after specified date.
	//		- before [string, default nil]: date filter for structs created or updated only before specified date.
	//		- sort [string, default "-created"]: sort order considered in response. Valid options are "-created" or "-due".
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success" or "failed"
	//		- type [string, default nil]: payment type, inferred from the payment parameter if it is not a dictionary. ex: "transfer", "boleto-payment"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of PaymentRequest structs with updated attributes
	//	- Cursor to retrieve the next page of PaymentRequest structs
	var param = map[string]interface{}{}
	var paymentRequests []PaymentRequest
	for k, v := range params {
		param[k] = v
	}
	param["centerId"] = centerId
	page, cursor, err := utils.Page(resource, param, user)
	if err.Errors != nil {
		return nil, "", err
	}

	unmarshalError := json.Unmarshal(page, &paymentRequests)
	if unmarshalError != nil {
		return nil, "", Error.UnknownError(unmarshalError.Error())
	}
	
	parsedRequests, err := ParseRequests(paymentRequests)
	if err.Errors != nil {
		return nil, "", err
	}
	return parsedRequests, cursor, Error.StarkErrors{}
}

func (e PaymentRequest) ParseRequest() (PaymentRequest, Error.StarkErrors) {
	if e.Type == "transfer" {
		var transfer Transfer.Transfer
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &transfer)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Payment = transfer
		return e, Error.StarkErrors{}
	}
	if e.Type == "transaction" {
		var transaction Transaction.Transaction
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &transaction)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Payment = transaction
		return e, Error.StarkErrors{}
	}
	if e.Type == "tax-payment" {
		var taxPayment TaxPayment.TaxPayment
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &taxPayment)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Payment = taxPayment
		return e, Error.StarkErrors{}
	}
	if e.Type == "brcode-payment" {
		var brcodePayment BrcodePayment.BrcodePayment
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &brcodePayment)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Payment = brcodePayment
		return e, Error.StarkErrors{}
	}
	if e.Type == "boleto-payment" {
		var boletoPayment BoletoPayment.BoletoPayment
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &boletoPayment)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Payment = boletoPayment
		return e, Error.StarkErrors{}
	}
	if e.Type == "utility-payment" {
		var utilityPayment UtilityPayment.UtilityPayment
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &utilityPayment)
		if unmarshalError != nil {
			return e, Error.UnknownError(unmarshalError.Error())
		}
		e.Payment = utilityPayment
		return e, Error.StarkErrors{}
	}
	return e, Error.StarkErrors{}
}

func ParseRequests(previews []PaymentRequest) ([]PaymentRequest, Error.StarkErrors) {
	var err Error.StarkErrors
	for i := 0; i < len(previews); i++ {
		previews[i], err = previews[i].ParseRequest()
		if err.Errors != nil {
			return nil, err
		}
	}
	return previews, Error.StarkErrors{}
}
