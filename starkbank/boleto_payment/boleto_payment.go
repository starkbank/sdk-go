package boleto_payment

//	BoletoPayment object
//	When you initialize a BoletoPayment, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the objects
//	to the Stark Bank API and returns the list of created objects.
//
//	Parameters (conditionally required):
//	- line [string, default None]: Number sequence that describes the payment. Either 'line' or 'bar_code' parameters are required. If both are sent, they must match. ex: "34191.09008 63571.277308 71444.640008 5 81960000000062"
//	- bar_code [string, default None]: Bar code number that describes the payment. Either 'line' or 'barCode' parameters are required. If both are sent, they must match. ex: "34195819600000000621090063571277307144464000"
//
//	Parameters (required):
//	- tax_id [string]: receiver tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- description [string]: Text to be displayed in your statement (min. 10 characters). ex: "payment ABC"
//
//	Parameters (optional):
//	- amount [int, default None]: amount to be paid. If none is informed, the current boleto value will be used. ex: 23456 (= R$ 234.56)
//	- scheduled [datetime.date or string, default today]: payment scheduled date. ex: datetime.date(2020, 3, 10)
//	- tags [list of strings]: list of strings for tagging
//
//	Attributes (return-only):
//	- id [string, default None]: unique id returned when payment is created. ex: "5656565656565656"
//	- status [string, default None]: current payment status. ex: "success" or "failed"
//	- fee [integer, default None]: fee charged when the boleto payment is created. ex: 200 (= R$ 2.00)
//	- transaction_ids [list of strings]: ledger transaction ids linked to this BoletoPayment. ex: ["19827356981273"]
//	- created [datetime.datetime, default None]: creation datetime for the payment. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type BoletoPayment struct {
	Line           int      `json:"line"`
	BarCode        []string `json:"barCode"`
	TaxId          string   `json:"taxId"`
	Description    string   `json:"description"`
	Amount         int      `json:"amount"`
	Scheduled      string   `json:"scheduled"`
	Tags           []string `json:"tags"`
	Id             string   `json:"id"`
	Status         string   `json:"status"`
	Fee            int      `json:"fee"`
	TransactionIds []string `json:"transactionIds"`
	Created        string   `json:"created"`
}

var resource = map[string]any{"class": BoletoPayment{}, "name": "BoletoPayment"}
