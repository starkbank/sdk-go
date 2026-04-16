package splitreceiver

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	SplitReceiver struct
//
//	When you initialize a SplitReceiver, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the objects
//	to the Stark Bank API and returns the list of created objects.
//
//	Parameters (required):
//	- Name [string]: Receiver full name. ex: "Anthony Edward Stark"
//	- TaxId [string]: Receiver account tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- BankCode [string]: Code of the receiver bank institution in Brazil. If an ISPB (8 digits) is informed, a PIX splitReceiver will be created, else a TED will be issued. ex: "20018183" or "341"
//	- BranchCode [string]: Receiver bank account branch. Use '-' in case there is a verifier digit. ex: "1357-9"
//	- AccountNumber [string]: Receiver bank account number. Use '-' before the verifier digit. ex: "876543-2"
//	- AccountType [string]: Receiver bank account type. This parameter only has effect on Pix SplitReceivers. ex: "checking", "savings", "salary" or "payment"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for receivers. ex: []string{"seller/123456"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the SplitReceiver is created. ex: "5656565656565656"
//	- Status [string]: Current SplitReceiver status. ex: "success" or "failed"
//	- Created [time.Time]: Creation datetime for the SplitReceiver. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC)
//	- Updated [time.Time]: Latest update datetime for the SplitReceiver. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC)

type SplitReceiver struct {
	Name          string     `json:",omitempty"`
	TaxId         string     `json:",omitempty"`
	BankCode      string     `json:",omitempty"`
	BranchCode    string     `json:",omitempty"`
	AccountNumber string     `json:",omitempty"`
	AccountType   string     `json:",omitempty"`
	Tags          []string   `json:",omitempty"`
	Id            string     `json:",omitempty"`
	Status        string     `json:",omitempty"`
	Created       *time.Time `json:",omitempty"`
	Updated       *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "SplitReceiver"}

func Create(receivers []SplitReceiver, user user.User) ([]SplitReceiver, Error.StarkErrors) {
	//	Create SplitReceivers
	//
	//	Send a slice of SplitReceiver structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- receivers [slice of SplitReceiver structs]: Slice of SplitReceiver structs to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of SplitReceiver structs with updated attributes
	create, err := utils.Multi(resource, receivers, nil, user)
	var splitReceivers []SplitReceiver
	unmarshalError := json.Unmarshal(create, &splitReceivers)
	if unmarshalError != nil {
		return splitReceivers, err
	}
	return splitReceivers, err
}

func Get(id string, user user.User) (SplitReceiver, Error.StarkErrors) {
	//	Retrieve a specific SplitReceiver by its id
	//
	//	Receive a single SplitReceiver struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- SplitReceiver struct that corresponds to the given id
	var splitReceiver SplitReceiver
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &splitReceiver)
	if unmarshalError != nil {
		return splitReceiver, err
	}
	return splitReceiver, err
}

func Query(params map[string]interface{}, user user.User) (chan SplitReceiver, chan Error.StarkErrors) {
	//	Retrieve SplitReceiver structs
	//
	//	Receive a channel of SplitReceiver structs from the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- transactionIds [slice of strings, default nil]: List of transaction IDs linked to the desired SplitReceivers. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "success" or "failed"
	//		- taxId [string, default nil]: Filter for SplitReceivers sent to the specified tax ID. ex: "012.345.678-90"
	//		- sort [string, default "-created"]: Sort order considered in response. Valid options are "created", "-created", "updated" or "-updated".
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of SplitReceiver structs with updated attributes
	var splitReceiver SplitReceiver
	splitReceivers := make(chan SplitReceiver)
	splitReceiversError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &splitReceiver)
			if err != nil {
				splitReceiversError <- Error.UnknownError(err.Error())
				continue
			}
			splitReceivers <- splitReceiver
		}
		for err := range errorChannel {
			splitReceiversError <- err
		}
		close(splitReceivers)
		close(splitReceiversError)
	}()
	return splitReceivers, splitReceiversError
}

func Page(params map[string]interface{}, user user.User) ([]SplitReceiver, string, Error.StarkErrors) {
	//	Retrieve paged SplitReceiver structs
	//
	//	Receive a slice of up to 100 SplitReceiver structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "success" or "failed"
	//		- sort [string, default "-created"]: Sort order considered in response. Valid options are "created", "-created", "updated" or "-updated".
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of SplitReceiver structs with updated attributes
	//	- Cursor to retrieve the next page of SplitReceiver structs
	var splitReceivers []SplitReceiver
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &splitReceivers)
	if unmarshalError != nil {
		return splitReceivers, cursor, err
	}
	return splitReceivers, cursor, err
}
