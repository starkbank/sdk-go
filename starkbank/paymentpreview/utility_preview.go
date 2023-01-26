package paymentpreview

//	UtilityPreview struct
//
//	A UtilityPreview is used to get information from a Utility Payment you received before confirming the payment.
//
//	Attributes (return-only):
//	- Amount [int]: final amount to be paid. ex: 23456 (= R$ 234.56)
//	- Name [string]: beneficiary full name. ex: "Iron Throne"
//	- Description [string]: utility payment description. ex: "Utility Payment - Light Company"
//	- Line [string]: Number sequence that identifies the payment. ex: "85660000006 6 67940064007 5 41190025511 7 00010601813 8"
//	- BarCode [string]: Bar code number that identifies the payment. ex: "85660000006679400640074119002551100010601813"

type UtilityPreview struct {
	Amount      int    `json:",omitempty"`
	Name        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Line        string `json:",omitempty"`
	BarCode     string `json:",omitempty"`
}

var PreviewUtility UtilityPreview
