package paymentpreview

//	BrcodePreview struct
//
//	A BrcodePreview is used to get information from a BR Code you received before confirming the payment.
//
//	Attributes (return-only):
//	- Status [string]: Payment status. ex: "active", "paid", "canceled" or "unknown"
//	- Name [string]: Payment receiver name. ex: "Tony Stark"
//	- TaxId [string]: Payment receiver tax ID. ex: "012.345.678-90"
//	- BankCode [string]: Payment receiver bank code. ex: "20018183"
//	- AccountType [string]: Payment receiver account type. ex: "checking"
//	- AllowChange [bool]: If True, the payment is able to receive amounts that are different from the nominal one. ex: True or False
//	- Amount [int]: Value in cents that this payment is expecting to receive. If 0, any value is accepted. ex: 123 (= R$1,23)
//	- NominalAmount [int]: Original value in cents that this payment was expecting to receive without the discounts, fines, etc.. If 0, any value is accepted. ex: 123 (= R$1,23)
//	- InterestAmount [int]: Current interest value in cents that this payment is charging. If 0, any value is accepted. ex: 123 (= R$1,23)
//	- FineAmount [int]: Current fine value in cents that this payment is charging. ex: 123 (= R$1,23)
//	- ReductionAmount [int]: Current value reduction value in cents that this payment is expecting. ex: 123 (= R$1,23)
//	- DiscountAmount [int]: Current discount value in cents that this payment is expecting. ex: 123 (= R$1,23)
//	- ReconciliationId [string]: Reconciliation ID linked to this payment. ex: "txId", "payment-123"

type BrcodePreview struct {
	Status           string `json:",omitempty"`
	Name             string `json:",omitempty"`
	TaxId            string `json:",omitempty"`
	BankCode         string `json:",omitempty"`
	AccountType      string `json:",omitempty"`
	AllowChange      bool   `json:",omitempty"`
	Amount           int    `json:",omitempty"`
	NominalAmount    int    `json:",omitempty"`
	InterestAmount   int    `json:",omitempty"`
	FineAmount       int    `json:",omitempty"`
	ReductionAmount  int    `json:",omitempty"`
	DiscountAmount   int    `json:",omitempty"`
	ReconciliationId string `json:",omitempty"`
}

var PreviewBrcode BrcodePreview
