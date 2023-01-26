package invoice

//	Invoice.Payment struct
//
//	When an Invoice is paid, its Payment sub-resource will become available.
//	It carries all the available information about the invoice payment.
//
//	Attributes (return-only):
//	- Amount [int]: Amount in cents that was paid. ex: 1234 (= R$ 12.34)
//	- Name [string]: Payer full name. ex: "Anthony Edward Stark"
//	- TaxId [string]: Payer tax ID (CPF or CNPJ). ex: "20.018.183/0001-80"
//	- BankCode [string]: Code of the payer bank institution in Brazil. ex: "20018183"
//	- BranchCode [string]: Payer bank account branch. ex: "1357-9"
//	- AccountNumber [string]: Payer bank account number. ex: "876543-2"
//	- AccountType [string]: Payer bank account type. ex: "checking", "savings", "salary" or "payment"
//	- EndToEndId [string]: Central bank's unique transaction ID. ex: "E79457883202101262140HHX553UPqeq"
//	- Method [string]: Payment method that was used. ex: "pix"

type Payment struct {
	Amount        int    `json:",omitempty"`
	Name          string `json:",omitempty"`
	TaxId         string `json:",omitempty"`
	BankCode      string `json:",omitempty"`
	BranchCode    string `json:",omitempty"`
	AccountNumber string `json:",omitempty"`
	AccountType   string `json:",omitempty"`
	EndToEndId    string `json:",omitempty"`
	Method        string `json:",omitempty"`
}

var payment Payment
var SubResourcePayment = map[string]string{"name": "Payment"}
