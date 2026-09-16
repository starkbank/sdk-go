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
//	created in the Stark Bank API. The 'Create' function sends the objects
//	to the Stark Bank API and returns the slice of created objects.
//
//	Parameters (required):
//	- Name [string]: receiver full name. ex: "Anthony Edward Stark"
//	- TaxId [string]: receiver account tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- BankCode [string]: code of the receiver bank institution in Brazil. If an ISPB (8 digits) is informed, a PIX SplitReceiver will be created, else a TED will be issued. ex: "20018183" or "341"
//	- BranchCode [string]: receiver bank account branch. Use '-' in case there is a verifier digit. ex: "1357-9"
//	- AccountNumber [string]: receiver bank account number. Use '-' before the verifier digit. ex: "876543-2"
//	- AccountType [string]: receiver bank account type. This parameter only has effect on Pix SplitReceivers. ex: "checking", "savings", "salary" or "payment"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: slice of strings for reference when searching for receivers. ex: []string{"seller/123456"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the SplitReceiver is created. ex: "5656565656565656"
//	- Status [string]: current SplitReceiver status. ex: "success" or "failed"
//	- Created [time.Time]: creation datetime for the SplitReceiver. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)
//	- Updated [time.Time]: latest update datetime for the SplitReceiver. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)

type SplitReceiver struct {
	Id            string     `json:",omitempty"`
	Name          string     `json:",omitempty"`
	TaxId         string     `json:",omitempty"`
	BankCode      string     `json:",omitempty"`
	BranchCode    string     `json:",omitempty"`
	AccountNumber string     `json:",omitempty"`
	AccountType   string     `json:",omitempty"`
	Tags          []string   `json:",omitempty"`
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
	//	- receivers [slice of SplitReceiver structs]: slice of SplitReceiver structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of SplitReceiver structs with updated attributes
	create, err := utils.Multi(resource, receivers, nil, user)
	unmarshalError := json.Unmarshal(create, &receivers)
	if unmarshalError != nil {
		return receivers, err
	}
	return receivers, err
}

func Get(id string, user user.User) (SplitReceiver, Error.StarkErrors) {
	//	Retrieve a specific SplitReceiver by its id
	//
	//	Receive a single SplitReceiver struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- SplitReceiver struct that corresponds to the given id.
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
	//	Receive a channel of SplitReceiver structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: date filter for objects created or updated only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: date filter for objects created or updated only before specified date. ex: "2020-03-10"
	//		- transactionIds [slice of strings, default nil]: slice of transaction IDs linked to the desired SplitReceivers. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved objects. ex: "success" or "failed"
	//		- taxId [string, default nil]: filter for SplitReceivers sent to the specified tax ID. ex: "012.345.678-90"
	//		- sort [string, default "-created"]: sort order considered in response. Valid options are "created", "-created", "updated" or "-updated".
	//		- tags [slice of strings, default nil]: tags to filter retrieved objects. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- channel of SplitReceiver structs with updated attributes
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
		close(splitReceiversError)
		close(splitReceivers)
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
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: date filter for objects created or updated only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: date filter for objects created or updated only before specified date. ex: "2020-03-10"
	//		- status [string, default nil]: filter for status of retrieved objects. ex: "success" or "failed"
	//		- sort [string, default "-created"]: sort order considered in response. Valid options are "created", "-created", "updated" or "-updated".
	//		- tags [slice of strings, default nil]: tags to filter retrieved objects. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of SplitReceiver structs with updated attributes
	//	- cursor to retrieve the next page of SplitReceiver structs
	var splitReceivers []SplitReceiver
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &splitReceivers)
	if unmarshalError != nil {
		return splitReceivers, cursor, err
	}
	return splitReceivers, cursor, err
}
