package invoicepullsubscription

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	InvoicePullSubscription struct
//
//	When you initialize an Invoice, the entity will not be automatically
//	sent to the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the slice of created structs.
//
//	Parameters (required):
//	- Start [time.Time]: subscription start date. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)
//	- Interval [string]: subscription installment interval. Options: "week", "month", "quarter", "semester", "year"
//	- PullMode [string]: subscription pull mode. Options: "manual", "automatic". Automatic mode will create the Invoice Pull Requests automatically
//	- PullRetryLimit [int]: subscription pull retry limit. Options: 0, 3
//	- Type [string]: subscription type. Options: "push", "qrcode", "qrcodeAndPayment", "paymentAndOrQrcode"
//
//	Parameters (conditionally required):
//	- Amount [int, default 0]: subscription amount in cents. Required if an amount_min_limit is not informed. Minimum = 1 (R$ 0.01). ex: 100 (= R$ 1.00)
//	- AmountMinLimit [int, 0 nil]: subscription minimum amount in cents. Required if an amount is not informed. Minimum = 1 (R$ 0.01). ex: 100 (= R$ 1.00)
//
//	Parameters (optional):
//	- DisplayDescription [string, default nil]: Invoice description to be shown to the payer. ex: "Subscription payment"
//	- Due [time.Time, default nil]: subscription invoice due offset. Available only for type "push". ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)
//	- ExternalId [string, default nil]: string that must be unique among all your InvoicePullSubscriptions. Duplicated external_ids will cause failures. ex: "my-external-id"
// 	- ReferenceCode [string, default nil]: reference code for reconciliation. ex: "REF123456"
//	- End [time.Time, default nil]: subscription end date. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)
//	- Data [map[string]interface{}, default nil]: additional data for the subscription based on type
//	- Name [string, default nil]: subscription debtor name. ex: "Iron Bank S.A."
//	- TaxId [string, default nil]: subscription debtor tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- Tags [slice of strings, default nil]: slice of strings for tagging. ex: []string{"John", "Paul"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when InvoicePullSubscription is created. ex: "5656565656565656"
//	- Status [string]: current InvoicePullSubscription status. ex: "active", "canceled"
//	- BacenId [string]: unique authentication id at the Central Bank. ex: "RR2001818320250616dtsPkBVaBYs"
//	- Brcode [string]: Brcode string for the InvoicePullSubscription. ex: "00020101021126580014br.gov.bcb.pix0114+5599999999990210starkbank.com.br520400005303986540410000000000005802BR5913Stark Bank S.A.6009SAO PAULO62070503***6304D2B1"
//	- Created [time.Time]: creation datetime for the InvoicePullSubscription. time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)
//	- Updated [time.Time]: latest update datetime for the InvoicePullSubscription. time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)

type InvoicePullSubscription struct {
	Id                 string                 `json:",omitempty"`
	Start              *time.Time             `json:",omitempty"`
	Interval           string                 `json:",omitempty"`
	PullMode           string                 `json:",omitempty"`
	PullRetryLimit     int                    `json:",omitempty"`
	Type               string                 `json:",omitempty"`
	Amount             int                    `json:",omitempty"`
	AmountMinLimit     int                    `json:",omitempty"`
	DisplayDescription string                 `json:",omitempty"`
	Due                *time.Time             `json:",omitempty"`
	ExternalId         string                 `json:",omitempty"`
	ReferenceCode      string                 `json:",omitempty"`
	End                *time.Time             `json:",omitempty"`
	Data               map[string]interface{} `json:",omitempty"`
	Name               string                 `json:",omitempty"`
	TaxId              string                 `json:",omitempty"`
	Tags               []string               `json:",omitempty"`
	Status             string                 `json:",omitempty"`
	BacenId            string                 `json:",omitempty"`
	Brcode             string                 `json:",omitempty"`
	Created            *time.Time             `json:",omitempty"`
	Updated            *time.Time             `json:",omitempty"`
}

var resource = map[string]string{"name": "InvoicePullSubscription"}

func Create(subscriptions []InvoicePullSubscription, user user.User) ([]InvoicePullSubscription, Error.StarkErrors) {
	//	Create InvoicePullSubscription
	//
	//	Send a slice of InvoicePullSubscription structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- subscriptions [slice of InvoicePullSubscription structs]: slice of InvoicePullSubscription structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of InvoicePullSubscription structs with updated attributes
	create, err := utils.Multi(resource, subscriptions, nil, user)
	unmarshalError := json.Unmarshal(create, &subscriptions)
	if unmarshalError != nil {
		return subscriptions, err
	}
	return subscriptions, err
}

func Get(id string, user user.User) (InvoicePullSubscription, Error.StarkErrors) {
	//	Get a specific InvoicePullSubscription by its id
	//
	// 	Receive a single InvoicePullSubscription struct previously created in the Stark Bank API by its id
	//
	// 	Parameters (required):
	// 	- id [string]: unique id of the InvoicePullSubscription. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- InvoicePullSubscription struct that corresponds to the given id.
	var invoicePullSubscription InvoicePullSubscription
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &invoicePullSubscription)
	if unmarshalError != nil {
		return invoicePullSubscription, err
	}
	return invoicePullSubscription, err
}

func Query(params map[string]interface{}, user user.User) (chan InvoicePullSubscription, chan Error.StarkErrors) {
	//	Retrieve InvoicePullSubscription structs
	//
	//	Receive a channel of InvoicePullSubscription structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of InvoicePullSubscription structs with updated attributes
	var invoicePullSubscription InvoicePullSubscription
	invoicePullSubscriptions := make(chan InvoicePullSubscription)
	invoicePullSubscriptionsErrors := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &invoicePullSubscription)
			if err != nil {
				invoicePullSubscriptionsErrors <- Error.UnknownError(err.Error())
				continue
			}
			invoicePullSubscriptions <- invoicePullSubscription
		}
		for err := range errorChannel {
			invoicePullSubscriptionsErrors <- err
		}
		close(invoicePullSubscriptions)
		close(invoicePullSubscriptionsErrors)
	}()
	return invoicePullSubscriptions, invoicePullSubscriptionsErrors
}

func Page(params map[string]interface{}, user user.User) ([]InvoicePullSubscription, string, Error.StarkErrors) {
	//	Retrieve InvoicePullSubscription structs
	//
	//	Receive a slice of up to 100 InvoicePullSubscription structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: []string{"paid", "registered"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of InvoicePullSubscription structs with updated attributes
	//	- Cursor to retrieve the next page of InvoicePullSubscription structs
	var invoicePullSubscriptions []InvoicePullSubscription
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &invoicePullSubscriptions)
	if unmarshalError != nil {
		return invoicePullSubscriptions, cursor, err
	}
	return invoicePullSubscriptions, cursor, err
}

func Cancel(id string, user user.User) (InvoicePullSubscription, Error.StarkErrors) {
	//	Cancel a InvoicePullSubscription entity
	//
	//	Cancel a InvoicePullSubscription entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: InvoicePullSubscription unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- canceled InvoicePullSubscription struct
	var invoicePullSubscription InvoicePullSubscription
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &invoicePullSubscription)
	if unmarshalError != nil {
		return invoicePullSubscription, err
	}
	return invoicePullSubscription, err
}