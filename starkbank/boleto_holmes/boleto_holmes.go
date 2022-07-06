package boleto_holmes

//	BoletoHolmes object
//	When you initialize a BoletoHolmes, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends the objects
//	to the Stark Bank API and returns the list of created objects.
//
//	Parameters (required):
//	- boleto_id [string]: investigated boleto entity ID. ex: "5656565656565656"
//
//	Parameters (optional):
//	- tags [list of strings]: list of strings for tagging
//
//	Attributes (return-only):
//	- id [string, default None]: unique id returned when holmes is created. ex: "5656565656565656"
//	- status [string, default None]: current holmes status. ex: "solving" or "solved"
//	- result [string, default None]: result of boleto status investigation. ex: "paid" or "cancelled"
//	- created [datetime.datetime, default None]: creation datetime for the holmes. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- updated [datetime.datetime, default None]: latest update datetime for the holmes. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type BoletoHolmes struct {
	BoletoId int      `json:"boletoId"`
	Tags     []string `json:"tags"`
	Id       string   `json:"id"`
	Status   int      `json:"status"`
	Result   string   `json:"result"`
	Created  string   `json:"created"`
	Updated  string   `json:"updated"`
}

var resource = map[string]any{"class": BoletoHolmes{}, "name": "BoletoHolmes"}

func Create() {
	//	Create Boletos Holmes
	//	Send a list of BoletoHolmes objects for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- holmes [list of BoletoHolmes objects]: list of BoletoHolmes objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- list of BoletoHolmes objects with updated attributes
}

func Get() {
	//	Retrieve a specific BoletoHolmes
	//	Receive a single BoletoHolmes object previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- BoletoHolmes object with updated attributes
}

func Query() {
	//	Retrieve BoletoHolmes
	//	Receive a generator of BoletoHolmes objects previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None] date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None] date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- status [string, default None]: filter for status of retrieved objects. ex: "paid" or "registered"
	//	- boleto_id [string, default None]: filter for holmes that investigate a specific boleto by its ID. ex: "5656565656565656"
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- generator of BoletoHolmes objects with updated attributes
}

func Page() {
	//	Retrieve paged BoletoHolmes
	//	Receive a list of up to 100 BoletoHolmes objects previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. It must be an integer between 1 and 100. ex: 50
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- list of BoletoHolmes objects with updated attributes
	//	- cursor to retrieve the next page of BoletoHolmes objects
}
