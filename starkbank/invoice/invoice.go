package invoice

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/invoice/rule"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	Invoice struct
//
//	When you initialize an Invoice, the entity will not be automatically
//	sent to the Stark Bank API. The 'create' function sends the structs
//	to the Stark Bank API and returns the slice of created structs.
//
//	To create scheduled Invoices, which will display the discount, interest, etc. on the final users banking interface,
//	use dates instead of datetimes on the "due" and "discounts" fields.
//
//	Parameters (required):
//	- Amount [int]: Invoice value in cents. Minimum = 0 (any value will be accepted). ex: 1234 (= R$ 12.34)
//	- TaxId [string]: payer tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- Name [string]: payer name. ex: "Iron Bank S.A."
//
//	Parameters (optional):
//	- Due [time.Time, default now + 2 days]: Invoice due date in UTC ISO format. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC), for immediate invoices and time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC) for scheduled invoices
//	- Expiration [int, default 5097600 (59 days)]: time interval in seconds between due date and expiration date. ex: 123456789
//	- Fine [float64, default 2.0]: Invoice fine for overdue payment in %. ex: 2.5
//	- Interest [float64, default 1.0]: Invoice monthly interest for overdue payment in %. ex: 5.2
//	- Discounts [slice of maps, default nil]: slice of maps with "percentage":float64 and "due":time.Time or string pairs
//	- Tags [slice of strings, default nil]: slice of strings for tagging. ex: []string{"John", "Paul"}
//	- Rules [slice of Invoice.Rule structs, default nil]: slice of Invoice.Rule structs for modifying transfer behavior. ex: []rule.Rule{{Key: "allowedTaxIds", Value: []string{"012.345.678-90", "45.059.493/0001-73"}}},
//	- Descriptions [slice of maps, default nil]: slice of maps with "key":string and (optional) "value":string pairs
//
//	Attributes (return-only):
//	- Pdf [string]: public Invoice PDF URL. ex: "https://invoice.starkbank.com/pdf/d454fa4e524441c1b0c1a729457ed9d8"
//	- Link [string]: public Invoice webpage URL. ex: "https://my-workspace.sandbox.starkbank.com/invoicelink/d454fa4e524441c1b0c1a729457ed9d8"
//	- NominalAmount [int]: Invoice emission value in cents (will change if invoice is updated, but not if it's paid). ex: 400000
//	- FineAmount [int]: Invoice fine value calculated over nominalAmount. ex: 20000
//	- InterestAmount [int]: Invoice interest value calculated over nominalAmount. ex: 10000
//	- DiscountAmount [int]: Invoice discount value calculated over nominalAmount. ex: 3000
//	- Id [string]: unique id returned when Invoice is created. ex: "5656565656565656"
//	- Brcode [string]: BR Code for the Invoice payment. ex: "00020101021226800014br.gov.bcb.pix2558invoice.starkbank.com/f5333103-3279-4db2-8389-5efe335ba93d5204000053039865802BR5913Arya Stark6009Sao Paulo6220051656565656565656566304A9A0"
//	- Status [string]: current Invoice status. ex: "registered" or "paid"
//	- Fee [int]: fee charged by this Invoice. ex: 200 (= R$ 2.00)
//	- TransactionIds [slice of strings]: ledger transaction ids linked to this Invoice (if there are more than one, all but the first are reversals or failed reversal chargebacks). ex: []string{"19827356981273"}
//	- Created [time.Time]: creation datetime for the Invoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: latest update datetime for the Invoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Invoice struct {
	Id             string                   `json:",omitempty"`
	Amount         int                      `json:",omitempty"`
	Name           string                   `json:",omitempty"`
	TaxId          string                   `json:",omitempty"`
	Due            *time.Time               `json:",omitempty"`
	Expiration     int                      `json:",omitempty"`
	Fine           float64                  `json:",omitempty"`
	Interest       float64                  `json:",omitempty"`
	Discounts      []map[string]interface{} `json:",omitempty"`
	Tags           []string                 `json:",omitempty"`
	Rules          []rule.Rule              `json:",omitempty"`
	Descriptions   []map[string]interface{} `json:",omitempty"`
	Pdf            string                   `json:",omitempty"`
	Link           string                   `json:",omitempty"`
	NominalAmount  int                      `json:",omitempty"`
	FineAmount     int                      `json:",omitempty"`
	InterestAmount int                      `json:",omitempty"`
	DiscountAmount int                      `json:",omitempty"`
	Brcode         string                   `json:",omitempty"`
	Status         string                   `json:",omitempty"`
	Fee            int                      `json:",omitempty"`
	TransactionIds []string                 `json:",omitempty"`
	Created        *time.Time               `json:",omitempty"`
	Updated        *time.Time               `json:",omitempty"`
}

var object Invoice
var objects []Invoice
var resource = map[string]string{"name": "Invoice"}

func Create(invoices []Invoice, user user.User) ([]Invoice, Error.StarkErrors) {
	//	Create Invoices
	//
	//	Send a slice of Invoice structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- invoices [slice of Invoice structs]: slice of Invoice structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of Invoice structs with updated attributes
	create, err := utils.Multi(resource, invoices, nil, user)
	unmarshalError := json.Unmarshal(create, &invoices)
	if unmarshalError != nil {
		return invoices, err
	}
	return invoices, err
}

func Get(id string, user user.User) (Invoice, Error.StarkErrors) {
	//	Retrieve a specific Invoice by its id
	//
	//	Receive a single Invoice struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Invoice struct that corresponds to the given id.
	var object Invoice
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan Invoice {
	//	Retrieve Invoice structs
	//
	//	Receive a channel of Invoice structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "paid" or "registered"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of Invoice structs with updated attributes
	var object Invoice
	invoices := make(chan Invoice)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			invoices <- object
		}
		close(invoices)
	}()
	return invoices
}

func Page(params map[string]interface{}, user user.User) ([]Invoice, string, Error.StarkErrors) {
	//	Retrieve paged Invoice structs
	//
	//	Receive a slice of up to 100 Invoice structs previously created in the Stark Bank API and the cursor to the next page.
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
	//	- Slice of Invoice structs with updated attributes
	//	- Cursor to retrieve the next page of Invoice structs
	var objects []Invoice
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (Invoice, Error.StarkErrors) {
	//	Update Invoice entity
	//
	//	Update an Invoice by passing id, if it hasn't been paid yet.
	//
	//	Parameters (required):
	//	- patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"amount": 9090}
	//		Parameters (optional):
	//		- status [string]: You may cancel the invoice by passing 'canceled' in the status
	//		- amount [string]: Nominal amount charged by the invoice. ex: 100 (R$1.00)
	//		- due [string, default now + 2 days]: Invoice due date in UTC ISO format. ex: "2020-10-28"
	//		- expiration [int, default nil]: time interval in seconds between the due date and the expiration date. ex: 123456789
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Target Invoice with updated attributes
	var object Invoice
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Qrcode(id string, params map[string]interface{}, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific Invoice QR Code png
	//
	//	Receive a single Invoice QR Code in png format generated in the Stark Bank API by the invoice ID.
	//
	//	Parameters (required):
	//	- id [string]: Invoice unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- size [int, default 7]: number of pixels in each "box" of the QR code. Minimum = 1, maximum = 50. ex: 12
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Invoice .png blob
	return utils.GetContent(resource, id, params, user, "qrcode")
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific Invoice .pdf file
	//
	//	Receive a single Invoice pdf receipt file generated in the Stark Bank API by its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Invoice .pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}

func GetPayment(id string, user user.User) (Payment, Error.StarkErrors) {
	//	Retrieve a specific Invoice payment information
	//
	//	Receive the Invoice.Payment sub-resource associated with a paid Invoice.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Invoice.Payment sub-resource
	get, err := utils.SubResource(resource, id, user, SubResourcePayment)
	unmarshalError := json.Unmarshal(get, &payment)
	if unmarshalError != nil {
		return payment, err
	}
	return payment, err
}
