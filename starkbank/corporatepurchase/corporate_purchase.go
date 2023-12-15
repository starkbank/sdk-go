package corporatepurchase

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	CorporatePurchase struct
//
//	Displays the CorporatePurchase structs created in your Workspace.
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when CorporatePurchase is created. ex: "5656565656565656"
//	- HolderID [string]: card holder unique id. ex: "5656565656565656"
//	- HolderName [string]: card holder name. ex: "Tony Stark"
//	- CenterID [string]: target cost center ID. ex: "5656565656565656"
//	- CardID [string]: unique id returned when CorporateCard is created. ex: "5656565656565656"
//	- CardEnding [string]: last 4 digits of the card number. ex: "1234"
//	- Description [string]: purchase descriptions. ex: "my_description"
//	- Amount [int]: CorporatePurchase value in cents. Minimum = 0. ex: 1234 (= R$ 12.34)
//	- Tax [int]: IOF amount taxed for international purchases. ex: 1234 (= R$ 12.34)
//	- IssuerAmount [int]: issuer amount. ex: 1234 (= R$ 12.34)
//	- IssuerCurrencyCode [string]: issuer currency code. ex: "USD"
//	- IssuerCurrencySymbol [string]: issuer currency symbol. ex: "$"
//	- MerchantAmount [int]: merchant amount. ex: 1234 (= R$ 12.34)
//	- MerchantCurrencyCode [string]: merchant currency code. ex: "USD"
//	- MerchantCurrencySymbol [string]: merchant currency symbol. ex: "$"
//	- MerchantCategoryCode [string]: merchant category code. ex: "fastFoodRestaurants"
//	- MerchantCategoryType [string]: merchant category type. ex: "health"
//	- MerchantCountryCode [string]: merchant country code. ex: "USA"
//	- MerchantName [string]: merchant name. ex: "Google Cloud Platform"
//	- MerchantDisplayName [string]: merchant name. ex: "Google Cloud Platform"
//	- MerchantDisplayUrl [string]: public merchant icon (png image). ex: "https://sandbox.api.starkbank.com/v2/corporate-icon/merchant/ifood.png"
//	- MerchantFee [int]: fee charged by the merchant to cover specific costs, such as ATM withdrawal logistics, etc. ex: 200 (= R$ 2.00)
//	- MethodCode [string]: method code. Options: "chip", "token", "server", "manual", "magstripe" or "contactless"
//	- Tags [slice of strings]: list of strings for tagging returned by the sub-issuer during the authorization. ex: new List<string>{ "travel", "food" }
//	- CorporateTransactionIds [slice of strings]: ledger transaction ids linked to this Purchase
//	- Status [string]: current CorporateCard status. Options: "approved", "canceled", "denied", "confirmed" or "voided"
//	- Updated [DateTime]: latest update DateTime for the CorporatePurchase. ex: DateTime(2020, 3, 10, 10, 30, 0, 0)
//	- Created [DateTime]: creation DateTime for the CorporatePurchase. ex: DateTime(2020, 3, 10, 10, 30, 0, 0)

type CorporatePurchase struct {
	Id                      string     `json:",omitempty"`
	HolderID                string     `json:",omitempty"`
	HolderName              string     `json:",omitempty"`
	CenterID                string     `json:",omitempty"`
	CardID                  string     `json:",omitempty"`
	CardEnding              string     `json:",omitempty"`
	Description             string     `json:",omitempty"`
	Amount                  int        `json:",omitempty"`
	Tax                     int        `json:",omitempty"`
	IssuerAmount            int        `json:",omitempty"`
	IssuerCurrencyCode      string     `json:",omitempty"`
	IssuerCurrencySymbol    string     `json:",omitempty"`
	MerchantAmount          int        `json:",omitempty"`
	MerchantCurrencyCode    string     `json:",omitempty"`
	MerchantCurrencySymbol  string     `json:",omitempty"`
	MerchantCategoryCode    string     `json:",omitempty"`
	MerchantCategoryType    string     `json:",omitempty"`
	MerchantCountryCode     string     `json:",omitempty"`
	MerchantName            string     `json:",omitempty"`
	MerchantDisplayName     string     `json:",omitempty"`
	MerchantDisplayUrl      string     `json:",omitempty"`
	MerchantFee             int        `json:",omitempty"`
	MethodCode              string     `json:",omitempty"`
	Tags                    []string   `json:",omitempty"`
	CorporateTransactionIds []string   `json:",omitempty"`
	Status                  string     `json:",omitempty"`
	Updated                 *time.Time `json:",omitempty"`
	Created                 *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "CorporatePurchase"}

func Get(id string, user user.User) (CorporatePurchase, Error.StarkErrors) {
	//	Retrieve a specific CorporatePurchase by its id
	//
	//	Receive a single CorporatePurchase struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- corporatePurchase struct that corresponds to the given id.
	var corporatePurchase CorporatePurchase
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &corporatePurchase)
	if unmarshalError != nil {
		return corporatePurchase, err
	}
	return corporatePurchase, err
}

func Query(params map[string]interface{}, user user.User) chan CorporatePurchase {
	//	Retrieve CorporatePurchase structs
	//
	//	Receive a channel of CorporatePurchase structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- merchantCategoryTypes [slice of strings, default nil]: merchant category type. ex: []string]{"health"}
	//		- holderIds [slice of strings, default nil]: Card holder IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- cardIds [slice of strings, default nil]: Card  IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"approved", "canceled", "denied", "confirmed", "voided"}
	//		- ids [slice of strings, default nil, default nil]: Purchase IDs
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- channel of CorporatePurchase structs with updated attributes
	var corporatePurchase CorporatePurchase
	purchases := make(chan CorporatePurchase)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &corporatePurchase)
			if err != nil {
				print(err.Error())
			}
			purchases <- corporatePurchase
		}
		close(purchases)
	}()
	return purchases
}

func Page(params map[string]interface{}, user user.User) ([]CorporatePurchase, string, Error.StarkErrors) {
	//	Retrieve paged CorporatePurchase structs
	//
	//	Receive a slice of up to 100 CorporatePurchase structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- merchantCategoryTypes [slice of strings, default nil]: merchant category type. ex: []string]{"health"}
	//		- holderIds [slice of strings, default nil]: Card holder IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- cardIds [slice of strings, default nil]: Card  IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"approved", "canceled", "denied", "confirmed", "voided"}
	//		- ids [slice of strings, default nil, default nil]: Purchase IDs
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of CorporatePurchase structs with updated attributes
	//	- cursor to retrieve the next page of CorporatePurchase structs
	var corporatePurchases []CorporatePurchase
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &corporatePurchases)
	if unmarshalError != nil {
		return corporatePurchases, cursor, err
	}
	return corporatePurchases, cursor, err
}

func Parse(content string, signature string, user user.User) CorporatePurchase {
	//	Create single verified CorporatePurchase authorization request from a content string
	//
	//	Use this method to parse and verify the authenticity of the authorization request received at the informed endpoint.
	//	Authorization requests are posted to your registered endpoint whenever CorporatePurchases are received.
	//	They present CorporatePurchase data that must be analyzed and answered with approval or declination.
	//	If the provided digital signature does not check out with the StarkInfra public key, a stark.exception.InvalidSignatureException will be raised.
	//	If the authorization request is not answered within 2 seconds or is not answered with an HTTP status code 200 the CorporatePurchase will go through the pre-configured stand-in validation.
	//
	//	Parameters (required):
	//	- content [string]: Response content from request received at user endpoint (not parsed)
	//	- signature [string]: Base-64 digital signature received at response header "Digital-Signature"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- parsed CorporatePurchase struct
	var corporatePurchase CorporatePurchase
	unmarshalError := json.Unmarshal([]byte(utils.ParseAndVerify(content, signature, "", user)), &corporatePurchase)
	if unmarshalError != nil {
		return corporatePurchase
	}
	return corporatePurchase
}

func Response(authorization map[string]interface{}) string {
	//	Helps you respond CorporatePurchase requests
	//
	//	Parameters (required):
	//	- status [string]: Sub-issuer response to the authorization. ex: "approved" or "denied"
	//
	//	Parameters (conditionally required):
	//	- reason [string]: Denial reason. Options: "other", "blocked", "lostCard", "stolenCard", "invalidPin", "invalidCard", "cardExpired", "issuerError", "concurrency", "standInDenial", "subIssuerError", "invalidPurpose", "invalidZipCode", "invalidWalletId", "inconsistentCard", "settlementFailed", "cardRuleMismatch", "invalidExpiration", "prepaidInstallment", "holderRuleMismatch", "insufficientBalance", "tooManyTransactions", "invalidSecurityCode", "invalidPaymentMethod", "confirmationDeadline", "withdrawalAmountLimit", "insufficientCardLimit", "insufficientHolderLimit"
	//
	//	Parameters (optional):
	//	- amount [int, default nil]: Amount in cents that was authorized. ex: 1234 (= R$ 12.34)
	//	- tags [slice of strings, default nil]: Tags to filter retrieved struct. ex: []string{"tony", "stark"}
	//
	//	Return:
	//	- dumped JSON string that must be returned to us on the CorporatePurchase request
	params := map[string]map[string]interface{}{
		"authorization": authorization,
	}
	response, _ := json.MarshalIndent(params, "", "  ")
	return string(response)
}
