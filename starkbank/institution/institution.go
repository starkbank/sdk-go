package institution

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

//	Institution struct
//
//	This resource is used to get information on the institutions that are recognized by the Brazilian Central Bank.
//	Besides the display name and full name, they also include the STR code (used for TEDs) and the SPI Code
//	(used for Pix) for the institutions. Either of these codes may be empty if the institution is not registered on
//	that Central Bank service.
//
//	Attributes (return-only):
//	- DisplayName [string]: Short version of the institution name that should be displayed to end users. ex: "Stark Bank"
//	- Name [string]: Full version of the institution name. ex: "Stark Bank S.A."
//	- SpiCode [string]: Spi code used to identify the institution on Pix transactions. ex: "20018183"
//	- StrCode [string]: Str code used to identify the institution on TED transactions. ex: "123"

type Institution struct {
	DisplayName string `json:",omitempty"`
	Name        string `json:",omitempty"`
	SpiCode     string `json:",omitempty"`
	StrCode     string `json:",omitempty"`
}

var object Institution
var resource = map[string]string{"name": "Institution"}

func Query(params map[string]interface{}, user user.User) chan Institution {
	//	Retrieve Bacen Institutions
	//
	//	Receive a slice of Institution structs that are recognized by the Brazilian Central bank for Pix and TED transactions
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- search [string, default nil]: Part of the institution name to be searched. ex: "stark"
	//		- spiCodes [slice of strings, default nil]: List of SPI (Pix) codes to be searched. ex: []string{"20018183"}
	//		- strCodes [slice of strings, default nil]: List of STR (TED) codes to be searched. ex: []string{"260"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Institution structs with updated attributes
	var object Institution
	institutions := make(chan Institution)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			institutions <- object
		}
		close(institutions)
	}()
	return institutions
}
