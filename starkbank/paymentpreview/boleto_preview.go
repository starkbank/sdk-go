package paymentpreview

import "time"

//	BoletoPreview struct
//
//	A BoletoPreview is used to get information from a Boleto payment you received before confirming the payment.
//
//	Attributes (return-only):
//	- Status [string]: Current boleto status. ex: "active", "expired" or "inactive"
//	- Amount [int]: Final amount to be paid. ex: 23456 (= R$ 234.56)
//	- DiscountAmount [int]: Discount amount to be paid. ex: 23456 (= R$ 234.56)
//	- FineAmount [int]: Fine amount to be paid. ex: 23456 (= R$ 234.56)
//	- InterestAmount [int]: Interest amount to be paid. ex: 23456 (= R$ 234.56)
//	- Due [time.Time]: Boleto due date. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- Expiration [time.Time]: Boleto expiration date. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//	- Name [string]: Beneficiary full name. ex: "Anthony Edward Stark"
//	- TaxId [string]: Beneficiary tax ID (CPF or CNPJ). ex: "20.018.183/0001-80"
//	- ReceiverName [string]: Receiver (Sacador Avalista) full name. ex: "Anthony Edward Stark"
//	- ReceiverTaxId [string]: Receiver (Sacador Avalista) tax ID (CPF or CNPJ). ex: "20.018.183/0001-80"
//	- PayerName [string]: Payer full name. ex: "Anthony Edward Stark"
//	- PayerTaxId [string]: Payer tax ID (CPF or CNPJ). ex: "20.018.183/0001-80"
//	- Line [string]: Number sequence that identifies the payment. ex: "34191.09008 63571.277308 71444.640008 5 81960000000062"
//	- BarCode [string]: Bar code number that identifies the payment. ex: "34195819600000000621090063571277307144464000"

type BoletoPreview struct {
	Status         string     `json:",omitempty"`
	Amount         int        `json:",omitempty"`
	DiscountAmount int        `json:",omitempty"`
	FineAmount     int        `json:",omitempty"`
	InterestAmount int        `json:",omitempty"`
	Due            *time.Time `json:",omitempty"`
	Expiration     *time.Time `json:",omitempty"`
	Name           string     `json:",omitempty"`
	TaxId          string     `json:",omitempty"`
	ReceiverName   string     `json:",omitempty"`
	ReceiverTaxId  string     `json:",omitempty"`
	PayerName      string     `json:",omitempty"`
	PayerTaxId     string     `json:",omitempty"`
	Line           string     `json:",omitempty"`
	BarCode        string     `json:",omitempty"`
}

var PreviewBoleto BoletoPreview
