package dictkey

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

//	DictKey struct
//
//	DictKey represents a Pix key registered in Bacen's DICT system.
//
//	Parameters (optional):
//	- Id [string, default nil]: DictKey struct unique id. ex: "tony@starkbank.com", "722.461.430-04", "20.018.183/0001-80", "+5511988887777", "b6295ee1-f054-47d1-9e90-ee57b74f60d9"
//
//	Attributes (return-only):
//	- Type [string]: Dict key type. ex: "email", "cpf", "cnpj", "phone" or "evp"
//	- Name [string]: Key owner full name. ex: "Tony Stark"
//	- TaxId [string]: Key owner tax ID (CNPJ or masked CPF). ex: "***.345.678-**" or "20.018.183/0001-80"
//	- OwnerType [string]: Dict key owner type. ex: "naturalPerson" or "legalPerson"
//	- BankName [string]: Bank name associated with the DICT key. ex: "Stark Bank"
//	- Ispb [string]: Bank ISPB associated with the DICT key. ex: "20018183"
//	- BranchCode [string]: Encrypted bank account branch code associated with the DICT key. ex: "ZW5jcnlwdGVkLWJyYW5jaC1jb2Rl"
//	- AccountNumber [string]: Encrypted bank account number associated with the DICT key. ex: "ZW5jcnlwdGVkLWFjY291bnQtbnVtYmVy"
//	- AccountType [string]: Bank account type associated with the DICT key. ex: "checking", "savings", "salary" or "payment"
//	- Status [string]: Current DICT key status. ex: "created", "registered", "canceled" or "failed" 

type DictKey struct {
	Id             string     `json:",omitempty"`
	Type           string     `json:",omitempty"`
	Name           string     `json:",omitempty"`
	TaxId          string     `json:",omitempty"`
	OwnerType      string     `json:",omitempty"`
	BankName       string     `json:",omitempty"`
	Ispb           string     `json:",omitempty"`
	BranchCode     string     `json:",omitempty"`
	AccountNumber  string     `json:",omitempty"`
	AccountType    string     `json:",omitempty"`
	Status         string     `json:",omitempty"`
}

var object DictKey
var objects []DictKey
var resource = map[string]string{"name": "DictKey"}

func Get(id string, user user.User) (DictKey, Error.StarkErrors) {
	//	Retrieve a specific DictKey by its id
	//
	//	Receive a single DictKey struct by its id
	//
	//	Parameters (required):
	//	- id [string]: DictKey struct unique id and Pix key itself. ex: "tony@starkbank.com", "722.461.430-04", "20.018.183/0001-80", "+5511988887777", "b6295ee1-f054-47d1-9e90-ee57b74f60d9"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- DictKey struct that corresponds to the given id
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan DictKey {
	//	Retrieve DictKey structs
	//
	//	Receive a channel of DictKey structs associated with your Stark Bank Workspace
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- type [string, default nil]: DictKey type. ex: "cpf", "cnpj", "phone", "email" or "evp"
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of DictKey structs with updated attributes
	keys := make(chan DictKey)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			keys <- object
		}
		close(keys)
	}()
	return keys
}

func Page(params map[string]interface{}, user user.User) ([]DictKey, string, Error.StarkErrors) {
	//	Retrieve paged DictKey structs
	//
	//	Receive a slice of up to 100 DictKey structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- type [string, default nil]: DictKey type. ex: "cpf", "cnpj", "phone", "email" or "evp"
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of DictKey structs with updated attributes
	//	- Cursor to retrieve the next page of DictKey structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
