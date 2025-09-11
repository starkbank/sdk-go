package corporateinvoice

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	CorporateInvoice struct
//
//	The CorporateInvoice structs created in your Workspace load your Corporate balance when paid.
//
//	When you initialize a CorporateInvoice, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the objects
//	to the Stark Bank API and returns the created object.
//
//	Parameters (required):
//	- Amount [int]: CorporateInvoice value in cents. ex: 1234 (= R$ 12.34)
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: slice of strings for tagging. ex: []string{"travel", "food"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when CorporateInvoice is created. ex: "5656565656565656"
//	- Name [string]: payer name. ex: "Iron Bank S.A."
//	- TaxId [string]: payer tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//  - Brcode [string]: BR Code for the Invoice payment. ex: "00020101021226930014br.gov.bcb.pix2571brcode-h.development.starkinfra.com/v2/d7f6546e194d4c64a153e8f79f1c41ac5204000053039865802BR5925Stark Bank S.A. - Institu6009Sao Paulo62070503***63042109"
//  - Due [time.Time]: Invoice due and expiration date in UTC ISO format. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//  - Link [string]: public Invoice webpage URL. ex: "https://starkbank-card-issuer.development.starkbank.com/invoicelink/d7f6546e194d4c64a153e8f79f1c41ac"
//	- Status [string]: current CorporateInvoice status. ex: "created", "expired", "overdue", "paid"
//	- CorporateTransactionId [string]: ledger transaction ids linked to this CorporateInvoice. ex: "corporate-invoice/5656565656565656"
//	- Updated [time.Time]: latest update datetime for the CorporateInvoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: creation datetime for the CorporateInvoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type CorporateInvoice struct {
	Id                     string     `json:",omitempty"`
	Amount                 int        `json:",omitempty"`
	TaxId                  string     `json:",omitempty"`
	Name                   string     `json:",omitempty"`
	Tags                   []string   `json:",omitempty"`
	Brcode                 string     `json:",omitempty"`
	Due                    *time.Time `json:",omitempty"`
	Link                   string     `json:",omitempty"`
	Status                 string     `json:",omitempty"`
	CorporateTransactionId string     `json:",omitempty"`
	Updated                *time.Time `json:",omitempty"`
	Created                *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "CorporateInvoice"}

func Create(invoice CorporateInvoice, user user.User) (CorporateInvoice, Error.StarkErrors) {
	//	Create a CorporateInvoice
	//
	//	Send a CorporateInvoice struct for creation at the Stark Bank API
	//
	//	Parameters (required):
	//	- invoice [CorporateInvoice struct]: CorporateInvoice struct to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- corporateInvoice struct with updated attributes
	create, err := utils.Single(resource, invoice, user)
	unmarshalError := json.Unmarshal(create, &invoice)
	if unmarshalError != nil {
		return invoice, err
	}
	return invoice, err
}

func Query(params map[string]interface{}, user user.User) (chan CorporateInvoice, chan Error.StarkErrors) {
	//	Retrieve CorporateInvoice
	//
	//	Receive a channel of CorporateInvoices structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"created", "expired", "overdue", "paid"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- channel of CorporateInvoices structs with updated attributes
	var corporateInvoice CorporateInvoice
	invoices := make(chan CorporateInvoice)
	invoicesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &corporateInvoice)
			if err != nil {
				invoicesError <- Error.UnknownError(err.Error())
				continue
			}
			invoices <- corporateInvoice
		}
		for err := range errorChannel {
			invoicesError <- err
		}
		close(invoices)
		close(invoicesError)
	}()
	return invoices, invoicesError
}

func Page(params map[string]interface{}, user user.User) ([]CorporateInvoice, string, Error.StarkErrors) {
	//	Retrieve CorporateInvoices
	//
	//	Receive a slice of up to 100 CorporateInvoice structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"created", "expired", "overdue", "paid"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of CorporateInvoices structs with updated attributes
	//	- cursor to retrieve the next page of CorporateInvoices structs
	var corporateInvoices []CorporateInvoice
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &corporateInvoices)
	if unmarshalError != nil {
		return corporateInvoices, cursor, err
	}
	return corporateInvoices, cursor, err
}
