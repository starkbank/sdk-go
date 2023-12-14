package boleto

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	Boleto struct
//
//	When you initialize a Boleto, the entity will not be automatically
//	sent to the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the list of created structs.
//
//	Parameters (required):
//	- Amount [int]: Boleto value in cents. Minimum = 200 (R$2,00). ex: 1234 (= R$ 12.34)
//	- Name [string]: Payer full name. ex: "Anthony Edward Stark"
//	- TaxId [string]: Payer tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- StreetLine1 [string]: Payer main address. ex: Av. Paulista, 200
//	- StreetLine2 [string]: Payer address complement. ex: Apto. 123
//	- District [string]: Payer address district/neighbourhood. ex: Bela Vista
//	- City [string]: Payer address city. ex: Rio de Janeiro
//	- StateCode [string]: Payer address state. ex: GO
//	- ZipCode [string]: Payer address zip code. ex: 01311-200
//
//	Parameters (optional):
//	- Due [time.Time, default today + 2 days]: Boleto due date in ISO format. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- Fine [float64, default 2.0]: Boleto fine for overdue payment in %. ex: 2.5
//	- Interest [float64, default 1.0]: Boleto monthly interest for overdue payment in %. ex: 5.2
//	- OverdueLimit [int, default 59]: Limit in days for payment after due date. ex: 7 (max: 59)
//	- Descriptions [slice of maps, default nil]: List of maps with "text":string and (optional) "amount":int pairs
//	- Discounts [slice of maps, default nil]: List of maps with "percentage":float64 and "date":time.Time or string pairs
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"John", "Paul"}
//	- ReceiverName [string, default nil]: Receiver (Sacador Avalista) full name. ex: "Anthony Edward Stark"
//	- ReceiverTaxId [string, default nil]: Receiver (Sacador Avalista) tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when Boleto is created. ex: "5656565656565656"
//	- Fee [int]: Fee charged when Boleto is paid. ex: 200 (= R$ 2.00)
//	- Line [string]: Generated Boleto line for payment. ex: "34191.09008 63571.277308 71444.640008 5 81960000000062"
//	- BarCode [string]: Generated Boleto bar-code for payment. ex: "34195819600000000621090063571277307144464000"
//	- Status [string]: Current Boleto status. ex: "registered" or "paid"
//	- TransactionIds [slice of strings]: Ledger transaction ids linked to this boleto. ex: []string{"19827356981273"}
//  - WorkspaceId [string]: ID of the Workspace that generated this Boleto. ex: "4545454545454545"
//	- Created [time.Time]: Creation datetime for the Boleto. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- OurNumber [string]: Reference number registered at the settlement bank. ex:"10131474"

type Boleto struct {
	Id            string                   `json:",omitempty"`
	Amount        int                      `json:",omitempty"`
	Name          string                   `json:",omitempty"`
	TaxId         string                   `json:",omitempty"`
	StreetLine1   string                   `json:",omitempty"`
	StreetLine2   string                   `json:",omitempty"`
	District      string                   `json:",omitempty"`
	City          string                   `json:",omitempty"`
	StateCode     string                   `json:",omitempty"`
	ZipCode       string                   `json:",omitempty"`
	Due           *time.Time               `json:",omitempty"`
	Fine          float64                  `json:",omitempty"`
	Interest      float64                  `json:",omitempty"`
	OverdueLimit  int                      `json:",omitempty"`
	Descriptions  []map[string]interface{} `json:",omitempty"`
	Discounts     []map[string]interface{} `json:",omitempty"`
	Tags          []string                 `json:",omitempty"`
	ReceiverName  string                   `json:",omitempty"`
	ReceiverTaxId string                   `json:",omitempty"`
	Fee           int                      `json:",omitempty"`
	Line          string                   `json:",omitempty"`
	BarCode       string                   `json:",omitempty"`
	Status        string                   `json:",omitempty"`
	Transactions  []string                 `json:",omitempty"`
	WorkspaceId   string                   `json:",omitempty"`
	Created       *time.Time               `json:",omitempty"`
	OurNumber     string                   `json:",omitempty"`
}

var object Boleto
var objects []Boleto
var resource = map[string]string{"name": "Boleto"}

func Create(boletos []Boleto, user user.User) ([]Boleto, Error.StarkErrors) {
	//	Create Boletos
	//
	//	Send a list of Boleto structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- boletos [slice of Boleto structs]: List of Boleto structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Boleto structs with updated attributes
	create, err := utils.Multi(resource, boletos, nil, user)
	unmarshalError := json.Unmarshal(create, &boletos)
	if unmarshalError != nil {
		return boletos, err
	}
	return boletos, err
}

func Get(id string, user user.User) (Boleto, Error.StarkErrors) {
	//	Retrieve a specific Boleto by its id
	//
	//	Receive a single Boleto struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Boleto struct that corresponds to the given id.
	var object Boleto
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Pdf(id string, params map[string]interface{}, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific Boleto .pdf file
	//
	//	Receive a single Boleto pdf file generated in the Stark Bank API by its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- layout [string]: Layout specification. Available options are "default" and "booklet"
	//	- hiddenFields [slice of strings, default nil]: List of string fields to be hidden in Boleto pdf. ex: []string{"customerAddress"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Boleto .pdf file
	return utils.GetContent(resource, id, params, user, "pdf")
}

func Query(params map[string]interface{}, user user.User) chan Boleto {
	//	Retrieve Boleto structs
	//
	//	Receive a channel of Boleto structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of Boleto structs with updated attributes
	var object Boleto
	boletos := make(chan Boleto)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			boletos <- object
		}
		close(boletos)
	}()
	return boletos
}

func Page(params map[string]interface{}, user user.User) ([]Boleto, string, Error.StarkErrors) {
	//	Retrieve paged Boleto structs
	//
	//	Receive a slice of up to 100 Boleto structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Boleto structs with updated attributes
	//	- Cursor to retrieve the next page of Boleto structs
	var objects []Boleto
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}

func Delete(id string, user user.User) (Boleto, Error.StarkErrors) {
	//	Delete a Boleto entity
	//
	//	Delete a Boleto entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: Boleto unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- deleted Boleto struct
	var object Boleto
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}
