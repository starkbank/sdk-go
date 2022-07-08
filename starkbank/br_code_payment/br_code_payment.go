package br_code_payment

//	BrcodePayment object
//	When you initialize a BrcodePayment, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the objects
//	to the Stark Bank API and returns the list of created objects.
//
//	Parameters (required):
//	- brcode [string]: String loaded directly from the QR Code or copied from the invoice. ex: "00020126580014br.gov.bcb.pix0136a629532e-7693-4846-852d-1bbff817b5a8520400005303986540510.005802BR5908T'Challa6009Sao Paulo62090505123456304B14A"
//	- tax_id [string]: receiver tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- description [string]: Text to be displayed in your statement (min. 10 characters). ex: "payment ABC"
//
//	Parameters (conditionally required):
//	- amount [int, default None]: If the BRCode does not provide an amount, this parameter is mandatory, else it is optional. ex: 23456 (= R$ 234.56)
//
//	Parameters (optional):
//	- scheduled [datetime.date, datetime.datetime or string, default now]: payment scheduled date or datetime. ex: datetime.datetime(2020, 3, 10, 15, 17, 3)
//	- tags [list of strings, default None]: list of strings for tagging
//
//	Attributes (return-only):
//	- id [string, default None]: unique id returned when payment is created. ex: "5656565656565656"
//	- name [string]: receiver name. ex: "Jon Snow"
//	- status [string, default None]: current payment status. ex: "success" or "failed"
//	- type [string, default None]: brcode type. ex: "static" or "dynamic"
//	- transaction_ids [list of strings, default None]: ledger transaction ids linked to this payment. ex: ["19827356981273"]
//	- fee [integer, default None]: fee charged by this brcode payment. ex: 50 (= R$ 0.50)
//	- updated [datetime.datetime, default None]: latest update datetime for the payment. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- created [datetime.datetime, default None]: creation datetime for the payment. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type BrcodePayment struct {
	Brcode         string `json:"brcode"`
	TaxId          string `json:"taxId"`
	Description    string `json:"description"`
	Amount         string `json:"amount"`
	Scheduled      string `json:"scheduled"`
	Tags           string `json:"tags"`
	Id             string `json:"id"`
	Name           string `json:"name"`
	Status         string `json:"status"`
	Type           string `json:"type"`
	TransactionIds string `json:"transactionIds"`
	Fee            string `json:"fee"`
	Updated        string `json:"updated"`
	Created        string `json:"created"`
}

var resource = map[string]any{"class": BrcodePayment{}, "name": "BrcodePayment"}

func Create() {
	//	Create BrcodePayments
	//	Send a list of BrcodePayment objects for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- payments [list of BrcodePayment objects]: list of BrcodePayment objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- list of BrcodePayment objects with updated attributes

}

func Get() {
	//	Retrieve a specific BrcodePayment
	//	Receive a single BrcodePayment object previously created by the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- BrcodePayment object with updated attributes

}

func Pdf() {
	//	Retrieve a specific BrcodePayment pdf file
	//	Receive a single BrcodePayment pdf receipt file generated in the Stark Bank API by its id.
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- BrcodePayment pdf file

}

func Query() {
	//	Retrieve BrcodePayments
	//	Receive a generator of BrcodePayment objects previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None] date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None] date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- status [string, default None]: filter for status of retrieved objects. ex: "success"
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- generator of BrcodePayment objects with updated attributes

}

func Page() {
	//	Retrieve paged BrcodePayments
	//	Receive a list of up to 100 BrcodePayment objects previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. It must be an integer between 1 and 100. ex: 50
	//	- after [datetime.date or string, default None] date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None] date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- status [string, default None]: filter for status of retrieved objects. ex: "success"
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- list of BrcodePayment objects with updated attributes
	//	- cursor to retrieve the next page of BrcodePayment objects

}

func Update() {
	//	Update BrcodePayment entity
	//	Update a BrcodePayment by passing its Id, if it hasn't been paid yet.
	//
	//	Parameters (required):
	//	- id [string]: BrcodePayment id. ex: '5656565656565656'
	//	- status [string]: You may cancel the payment by passing 'canceled' in the status
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- target BrcodePayment with updated attributes

}
