package boleto

//	Boleto object
//	When you initialize a Boleto, the entity will not be automatically
//	sent to the Stark Bank API. The 'create' function sends the objects
//	to the Stark Bank API and returns the list of created objects.
//
//	Parameters (required):
//	- amount [integer]: Boleto value in cents. Minimum = 200 (R$2,00). ex: 1234 (= R$ 12.34)
//	- name [string]: payer full name. ex: "Anthony Edward Stark"
//	- tax_id [string]: payer tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- street_line_1 [string]: payer main address. ex: Av. Paulista, 200
//	- street_line_2 [string]: payer address complement. ex: Apto. 123
//	- district [string]: payer address district / neighbourhood. ex: Bela Vista
//	- city [string]: payer address city. ex: Rio de Janeiro
//	- state_code [string]: payer address state. ex: GO
//	- zip_code [string]: payer address zip code. ex: 01311-200
//
//	Parameters (optional):
//	- due [datetime.date or string, default today + 2 days]: Boleto due date in ISO format. ex: 2020-04-30
//	- fine [float, default 0.0]: Boleto fine for overdue payment in %. ex: 2.5
//	- interest [float, default 0.0]: Boleto monthly interest for overdue payment in %. ex: 5.2
//	- overdue_limit [integer, default 59]: limit in days for payment after due date. ex: 7 (max: 59)
//	- descriptions [list of dictionaries, default None]: list of dictionaries with "text":string and (optional) "amount":int pairs
//	- discounts [list of dictionaries, default None]: list of dictionaries with "percentage":float and "date":datetime.datetime or string pairs
//	- tags [list of strings]: list of strings for tagging
//	- receiver_name [string]: receiver (Sacador Avalista) full name. ex: "Anthony Edward Stark"
//	- receiver_tax_id [string]: receiver (Sacador Avalista) tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//
//	Attributes (return-only):
//	- id [string, default None]: unique id returned when Boleto is created. ex: "5656565656565656"
//	- fee [integer, default None]: fee charged when Boleto is paid. ex: 200 (= R$ 2.00)
//	- line [string, default None]: generated Boleto line for payment. ex: "34191.09008 63571.277308 71444.640008 5 81960000000062"
//	- bar_code [string, default None]: generated Boleto bar-code for payment. ex: "34195819600000000621090063571277307144464000"
//	- status [string, default None]: current Boleto status. ex: "registered" or "paid"
//	- transaction_ids [list of strings]: ledger transaction ids linked to this boleto. ex: ["19827356981273"]
//	- created [datetime.datetime, default None]: creation datetime for the Boleto. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- our_number [string, default None]: Reference number registered at the settlement bank. ex:"10131474"

type Boleto struct {
	Amount        int           `json:"amount"`
	Name          string        `json:"name"`
	TaxId         string        `json:"taxId"`
	StreetLine1   string        `json:"streetLine1"`
	StreetLine2   string        `json:"streetLine2"`
	District      string        `json:"district"`
	City          string        `json:"city"`
	StateCode     string        `json:"stateCode"`
	ZipCode       string        `json:"zipCode"`
	Due           string        `json:"due"`
	Fine          float32       `json:"fine"`
	Interest      float32       `json:"interest"`
	OverdueLimit  int           `json:"overdueLimit"`
	Descriptions  []Description `json:"descriptions"`
	Discounts     []Discount    `json:"discounts"`
	Tags          []string      `json:"tags"`
	ReceiverName  string        `json:"receiverName"`
	ReceiverTaxId string        `json:"receiverTaxId"`
	Id            string        `json:"id"`
	Fee           int           `json:"fee"`
	Line          string        `json:"line"`
	BarCode       string        `json:"barCode"`
	Transactions  []string      `json:"transactions"`
	Created       string        `json:"created"`
	OurNumber     string        `json:"ourNumber"`
}

var resource = map[string]any{"class": Boleto{}, "name": "Boleto"}

func Create() {
	//	Create Boletos
	//	Send a list of Boleto objects for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- boletos [list of Boleto objects]: list of Boleto objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- list of Boleto objects with updated attributes
}

func Get() {
	//	Retrieve a specific Boleto
	//	Receive a single Boleto object previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- Boleto object with updated attributes
}

func Pdf() {
	//	Retrieve a specific Boleto pdf file
	//	Receive a single Boleto pdf file generated in the Stark Bank API by its id.
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- layout [string]: Layout specification. Available options are "default" and "booklet"
	//	- hidden_fields [list of strings, default None]: List of string fields to be hidden in Boleto pdf. ex: ["customerAddress"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- Boleto pdf file
}

func Query() {
	//	Retrieve Boletos
	//	Receive a generator of Boleto objects previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None] date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None] date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- status [string, default None]: filter for status of retrieved objects. ex: "paid" or "registered"
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- generator of Boleto objects with updated attributes
}

func Page() {
	//	Retrieve paged Boletos
	//	Receive a list of up to 100 Boleto objects previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. It must be an integer between 1 and 100. ex: 50
	//	- after [datetime.date or string, default None] date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None] date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- status [string, default None]: filter for status of retrieved objects. ex: "paid" or "registered"
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- list of Boleto objects with updated attributes
	//	- cursor to retrieve the next page of Boleto objects
}

func Delete() {
	//	Delete a Boleto entity
	//	Delete a Boleto entity previously created in the Stark Bank API
	//
	//	Parameters (required):
	//	- id [string]: Boleto unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- deleted Boleto object
}
