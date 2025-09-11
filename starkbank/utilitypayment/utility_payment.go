package utilitypayment

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	UtilityPayment struct
//
//	When you initialize a UtilityPayment, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the slice of created structs.
//
//	Parameters (required):
//	- Description [string]: Text to be displayed in your statement (min. 10 characters). ex: "payment ABC"
//
//	Parameters (conditionally required):
//	- Line [string, default nil]: Number sequence that describes the payment. Either 'line' or 'barCode' parameters are required. If both are sent, they must match. ex: "34191.09008 63571.277308 71444.640008 5 81960000000062"
//	- BarCode [string, default nil]: Bar code number that describes the payment. Either 'line' or 'barCode' parameters are required. If both are sent, they must match. ex: "34195819600000000621090063571277307144464000"
//
//	Parameters (optional):
//	- Scheduled [time.Time, default today]: payment scheduled date. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- Tags [slice of strings, default nil]: slice of strings for tagging. ex: []string{"John", "Paul"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when payment is created. ex: "5656565656565656"
//	- Status [string]: current payment status. ex: "success" or "failed"
//	- Amount [int]: amount automatically calculated from line or barCode. ex: 23456 (= R$ 234.56)
//	- Fee [int]: fee charged when utility payment is created. ex: 200 (= R$ 2.00)
//  - Type [string]: payment type. ex: "utility"
//	- TransactionIds [slice of strings]: ledger transaction ids linked to this UtilityPayment. ex: []string{"19827356981273"}
//	- Created [time.Time]: creation datetime for the payment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: latest update datetime for the payment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type UtilityPayment struct {
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
	Created        *time.Time `json:",omitempty"`
	Updated        *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "UtilityPayment"}

func Create(payments []UtilityPayment, user user.User) ([]UtilityPayment, Error.StarkErrors) {
	//	Create UtilityPayments
	//
	//	Send a slice of UtilityPayment structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- payments [slice of UtilityPayment structs]: slice of UtilityPayment structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of UtilityPayment structs with updated attributes
	create, err := utils.Multi(resource, payments, nil, user)
	unmarshalError := json.Unmarshal(create, &payments)
	if unmarshalError != nil {
		return payments, err
	}
	return payments, err
}

func Get(id string, user user.User) (UtilityPayment, Error.StarkErrors) {
	//	Retrieve a specific UtilityPayment by its id
	//
	//	Receive a single UtilityPayment struct previously created by the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- UtilityPayment struct with updated attributes
	var utilityPayment UtilityPayment
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &utilityPayment)
	if unmarshalError != nil {
		return utilityPayment, err
	}
	return utilityPayment, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific UtilityPayment pdf file
	//
	//	Receive a single UtilityPayment pdf file generated in the Stark Bank API by its id.
	//	Only valid for utility payments with "success" status.
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- UtilityPayment pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}

func Query(params map[string]interface{}, user user.User) (chan UtilityPayment, chan Error.StarkErrors) {
	//	Retrieve UtilityPayments
	//
	//	Receive a channel of UtilityPayment structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of UtilityPayment structs with updated attributes
	var utilityPayment UtilityPayment
	payments := make(chan UtilityPayment)
	paymentsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &utilityPayment)
			if err != nil {
				paymentsError <- Error.UnknownError(err.Error())
				continue
			}
			payments <- utilityPayment
		}
		for err := range errorChannel {
			paymentsError <- err
		}
		close(payments)
		close(paymentsError)
	}()
	return payments, paymentsError
}

func Page(params map[string]interface{}, user user.User) ([]UtilityPayment, string, Error.StarkErrors) {
	//	Retrieve paged UtilityPayments
	//
	//	Receive a slice of up to 100 UtilityPayment structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of UtilityPayment structs with updated attributes
	//	- cursor to retrieve the next page of UtilityPayment structs
	var utilityPayments []UtilityPayment
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &utilityPayments)
	if unmarshalError != nil {
		return utilityPayments, cursor, err
	}
	return utilityPayments, cursor, err
}

func Delete(id string, user user.User) (UtilityPayment, Error.StarkErrors) {
	//	Delete a UtilityPayment entity
	//
	//	Delete a UtilityPayment entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: UtilityPayment unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Deleted UtilityPayment struct
	var utilityPayment UtilityPayment
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &utilityPayment)
	if unmarshalError != nil {
		return utilityPayment, err
	}
	return utilityPayment, err
}
