package dynamicbrcode

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/dynamicbrcode/rule"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	DynamicBrcode struct
//
//	When you initialize an DynamicBrcode, the entity will not be automatically
//	sent to the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the slice of created structs.
//
//	DynamicBrcodes are conciliated BR Codes that can be used to receive Pix transactions in a convenient way.
//	When a DynamicBrcode is paid, a Deposit is created with the tags parameter containing the character “dynamic-brcode/” followed by the DynamicBrcode’s uuid "dynamic-brcode/{uuid}" for conciliation.
//	Additionally, all tags passed on the DynamicBrcode will be transferred to the respective Deposit resource.
//
//	Parameters (required):
//	- amount [int]: DynamicBrcode value in cents. Minimum = 0 (any value will be accepted). ex: 1234 (= R$ 12.34
//
//	Parameters (optional):
//	- Expiration [int, default 3600 (1 hour)]: time interval in seconds between due date and expiration date. ex: 123456789
//	- Tags [slice of strings, default []]: list of strings for tagging, these will be passed to the respective Deposit resource when paid
//
//	Attributes (return-only):
//	- Id [string]: id returned on creation, this is the BR code. ex: "00020126360014br.gov.bcb.pix0114+552840092118152040000530398654040.095802BR5915Jamie Lannister6009Sao Paulo620705038566304FC6C"
//	- Uuid [string]: unique uuid returned when the DynamicBrcode is created. ex: "4e2eab725ddd495f9c98ffd97440702d"
// 	- PictureUrl [string]: public Dynamic Brcode picture url. ex: "https://development.api.starkbank.com/v2/dynamic-brcode/d3ebb1bd92024df1ab6e5a353ee799a4.png",
//	- Created [time.Time]: creation datetime for the DynamicBrcode. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: latest update datetime for the DynamicBrcode. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type DynamicBrcode struct {
	Id                        string           `json:",omitempty"`
	Amount                    int              `json:",omitempty"`
	Expiration                int              `json:",omitempty"`
	Tags                      []string         `json:",omitempty"`
	Uuid                      string           `json:",omitempty"`
	PictureUrl                string           `json:",omitempty"`
	Rules                     []rule.Rule      `json:",omitempty"`
	DisplayDescription        string           `json:",omitempty"`
	Created                   *time.Time       `json:",omitempty"`
	Updated                   *time.Time       `json:",omitempty"`
}

var resource = map[string]string{"name": "DynamicBrcode"}

func Create(brcodes []DynamicBrcode, user user.User) ([]DynamicBrcode, Error.StarkErrors) {
	//	Create DynamicBrcodes
	//
	//	Send a slice of DynamicBrcode structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- brcodes [slice of DynamicBrcode structs]: slice of DynamicBrcode structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of DynamicBrcode structs with updated attributes
	create, err := utils.Multi(resource, brcodes, nil, user)
	unmarshalError := json.Unmarshal(create, &brcodes)
	if unmarshalError != nil {
		return brcodes, err
	}
	return brcodes, err
}

func Get(uuid string, user user.User) (DynamicBrcode, Error.StarkErrors) {
	//	Retrieve a specific DynamicBrcode by its uuid
	//
	//	Receive a single DynamicBrcode struct previously created in the Stark Bank API by its uuid
	//
	//	Parameters (required):
	//	- uuid [string]: Struct unique uuid. ex: "901e71f2447c43c886f58366a5432c4b"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- DynamicBrcode struct that corresponds to the given uuid.
	var dynamicBrcode DynamicBrcode
	get, err := utils.Get(resource, uuid, nil, user)
	unmarshalError := json.Unmarshal(get, &dynamicBrcode)
	if unmarshalError != nil {
		return dynamicBrcode, err
	}
	return dynamicBrcode, err
}

func Query(params map[string]interface{}, user user.User) (chan DynamicBrcode, chan Error.StarkErrors) {
	//	Retrieve DynamicBrcode structs
	//
	//	Receive a channel of DynamicBrcode structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//		- uuids [slice of strings, default nil]: list of uuids to filter retrieved objects. ex: ["901e71f2447c43c886f58366a5432c4b", "4e2eab725ddd495f9c98ffd97440702d"]
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of DynamicBrcode structs with updated attributes
	var dynamicBrcode DynamicBrcode
	brcodes := make(chan DynamicBrcode)
	brcodesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &dynamicBrcode)
			if err != nil {
				brcodesError <- Error.UnknownError(err.Error())
				continue
			}
			brcodes <- dynamicBrcode
		}
		for err := range errorChannel {
			brcodesError <- err
		}
		close(brcodes)
		close(brcodesError)
	}()
	return brcodes, brcodesError
}

func Page(params map[string]interface{}, user user.User) ([]DynamicBrcode, string, Error.StarkErrors) {
	//	Retrieve paged DynamicBrcode structs
	//
	//	Receive a slice of up to 100 DynamicBrcode structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- uuids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of DynamicBrcode structs with updated attributes
	//	- Cursor to retrieve the next page of DynamicBrcode structs
	var dynamicBrcodes []DynamicBrcode
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &dynamicBrcodes)
	if unmarshalError != nil {
		return dynamicBrcodes, cursor, err
	}
	return dynamicBrcodes, cursor, err
}
