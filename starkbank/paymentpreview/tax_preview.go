package paymentpreview

//	TaxPreview struct
//
//	A TaxPreview is used to get information from a Tax Payment you received before confirming the payment.
//
//	Attributes (return-only):
//	- Amount [int]: final amount to be paid. ex: 23456 (= R$ 234.56)
//	- Name [string]: beneficiary full name. ex: "Iron Throne"
//	- Description [string]: tax payment description. ex: "ISS Payment - Iron Throne"
//	- Line [string]: Number sequence that identifies the payment. ex: "85660000006 6 67940064007 5 41190025511 7 00010601813 8"
//	- BarCode [string]: Bar code number that identifies the payment. ex: "85660000006679400640074119002551100010601813"

type TaxPreview struct {
	Amount      int    `json:",omitempty"`
	Name        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Line        string `json:",omitempty"`
	BarCode     string `json:",omitempty"`
}

var PreviewTax TaxPreview
