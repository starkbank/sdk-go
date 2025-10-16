package invoicepullrequest

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	InvoicePullRequest struct
//
//	When you initialize an InvoicePullRequest, the entity will not be automatically
//	sent to the Stark Bank API. The 'create' function sends the objects
//	to the Stark Bank API and returns the list of created objects.
// 
//	Parameters (required):
//	- SubscriptionId [string]: Unique of the InvoicePullSubscription related to the invoice. ex: "5656565656565656"
//	- InvoiceId [string]: Id of the invoice previously created to be sent for payment. ex: "5656565656565656"
//	- Due [time.Time]: payment scheduled date in UTC ISO format. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)
//
//	Parameters (optional):
//	- AttemptType [string, default "default"]: attempt type for the payment. options: "default", "retry".
//	- Tags [slice of strings, default nil]: list of strings for tagging. ex: []string{"John", "Paul"}
//	- ExternalId [string, default nil]: a string that must be unique among all your InvoicePullRequests. Duplicated external_ids will cause failures. ex: "my-external-id"
//	- DisplayDescription [string, default nil]: Description to be shown to the payer. ex: "Payment for services"
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when InvoicePullRequest is created. ex: "5656565656565656"
//	- Status [string]: current InvoicePullRequest status. ex: "pending", "scheduled", "success", "failed", "canceled"
//	- InstallmentId [string]: unique id of the installment related to this request. ex: "5656565656565656"
//	- Created [time.Time]: creation datetime for the InvoicePullRequest. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)
//	- Updated [time.Time]: latest update datetime for the InvoicePullRequest. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)

type InvoicePullRequest struct {
	Id                 string     `json:",omitempty"`
	SubscriptionId     string     `json:",omitempty"`
	InvoiceId          string     `json:",omitempty"`
	Due                *time.Time `json:",omitempty"`
	AttemptType        string     `json:",omitempty"`
	Tags               []string   `json:",omitempty"`
	ExternalId         string     `json:",omitempty"`
	DisplayDescription string     `json:",omitempty"`
	Status             string     `json:",omitempty"`
	InstallmentId      string     `json:",omitempty"`
	Created            *time.Time `json:",omitempty"`
	Updated            *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "InvoicePullRequest"}

func Create(requests []InvoicePullRequest, user user.User) ([]InvoicePullRequest, Error.StarkErrors) {
	//	Create InvoicePullRequest
	//
	//	Send a slice of InvoicePullRequest structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- requests [slice of InvoicePullRequest structs]: list of InvoicePullRequest structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of InvoicePullRequest structs with updated attributes
	create, err := utils.Multi(resource, requests, nil, user)
	unmarshalError := json.Unmarshal(create, &requests)
	if unmarshalError != nil {
		return requests, err
	}
	return requests, err
}

func Get(id string, user user.User) (InvoicePullRequest, Error.StarkErrors) {
	//	Get a specific InvoicePullRequest by its id
	//
	// 	Receive a single Invoice struct previously created in the Stark Bank API by its id
	//
	// 	Parameters (required):
	// 	- id [string]: unique id of the InvoicePullRequest. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- InvoicePullRequest struct that corresponds to the given id.
	var invoicePullRequest InvoicePullRequest
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &invoicePullRequest)
	if unmarshalError != nil {
		return invoicePullRequest, err
	}
	return invoicePullRequest, err
}

func Query(params map[string]interface{}, user user.User) (chan InvoicePullRequest, chan Error.StarkErrors) {
	//	Retrieve InvoicePullRequest structs
	//
	//	Receive a channel of InvoicePullRequest structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	var invoicePullRequest InvoicePullRequest
	invoicePullRequests := make(chan InvoicePullRequest)
	invoicePullRequestsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &invoicePullRequest)
			if err != nil {
				invoicePullRequestsError <- Error.UnknownError(err.Error())
				continue
			}
			invoicePullRequests <- invoicePullRequest
		}
		for err := range errorChannel {
			invoicePullRequestsError <- err
		}
		close(invoicePullRequestsError)
		close(invoicePullRequests)
	}()
	return invoicePullRequests, invoicePullRequestsError
}

func Page(params map[string]interface{}, user user.User) ([]InvoicePullRequest, string, Error.StarkErrors) {
	//	Retrieve paged InvoicePullRequest structs
	//
	//	Receive a slice of up to 100 InvoicePullRequest structs previously created in the Stark Bank API and the cursor to the next page.
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
	//	- Slice of InvoicePullRequest structs with updated attributes
	//	- Cursor to retrieve the next page of InvoicePullRequest structs
	var invoicePullRequests []InvoicePullRequest
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &invoicePullRequests)
	if unmarshalError != nil {
		return invoicePullRequests, cursor, err
	}
	return invoicePullRequests, cursor, err
}

func Cancel(id string, user user.User) (InvoicePullRequest, Error.StarkErrors) {
	//	Cancel a InvoicePullRequest entity
	//
	//	Cancel a InvoicePullRequest entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: InvoicePullRequest unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- canceled InvoicePullRequest struct
	var invoicePullRequest InvoicePullRequest
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &invoicePullRequest)
	if unmarshalError != nil {
		return invoicePullRequest, err
	}
	return invoicePullRequest, err
}