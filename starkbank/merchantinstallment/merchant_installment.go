package merchantinstallment

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

// MerchantInstallment struct
//
// The MerchantInstallment resource stores information about each installment of a MerchantPurchase made with InstallmentCount greater than 1, so the individual installments can be tracked and reconciled as they settle.
//
// Attributes (return-only):
// - Id [string]: unique id returned when the installment is created. ex: "5656565656565656"
// - Amount [int]: installment value in cents. ex: 100 (= R$1.00)
// - Network [string]: card network.
// - FundingType [string]: funding type used for the purchase. Options: "credit", "debit"
// - PurchaseId [string]: unique id of the MerchantPurchase that originated this installment. ex: "5656565656565656"
// - Status [string]: current installment status.
// - TransactionIds [slice of strings]: ledger transaction ids linked to this installment.
// - Tags [slice of strings]: tags associated with the purchase that originated this installment.
// - Due [time.Time]: installment due date.
// - Fee [int]: fee charged over this installment. ex: 200 (= R$ 2.00)
// - Created [time.Time]: creation datetime for the installment.
// - Updated [time.Time]: latest update datetime for the installment.

type MerchantInstallment struct {
	Id             string     `json:",omitempty"`
	Amount         int        `json:",omitempty"`
	Network        string     `json:",omitempty"`
	FundingType    string     `json:",omitempty"`
	PurchaseId     string     `json:",omitempty"`
	Status         string     `json:",omitempty"`
	TransactionIds []string   `json:",omitempty"`
	Tags           []string   `json:",omitempty"`
	Created        *time.Time `json:",omitempty"`
	Updated        *time.Time `json:",omitempty"`
	Due            *time.Time `json:",omitempty"`
	Fee            int        `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantInstallment"}

func Get(id string, user user.User) (MerchantInstallment, Error.StarkErrors){
	// Retrieve a specific MerchantInstallment by its id
	//
	// Receive a single MerchantInstallment struct previously created in the Stark Bank API by its id
	//
	// Parameters (required):
	// - id [string]: MerchantInstallment unique id. ex: "5656565656565656"
	//
	// Parameters (optional):
	// - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - MerchantInstallment struct that corresponds to the given id.
	var merchantInstallment MerchantInstallment
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &merchantInstallment)
	if unmarshalError != nil {
		return merchantInstallment, err
	}
	return merchantInstallment, err
}

func Query(params map[string]interface{}, user user.User) (chan MerchantInstallment, chan Error.StarkErrors) {
	// Retrieve MerchantInstallment structs
	//
	// Receive a channel of MerchantInstallment structs previously created in the Stark Bank API
	//
	// Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs.
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Channel of MerchantInstallment structs with updated attributes
	merchantInstallments := make(chan MerchantInstallment)
	merchantInstallmentsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var merchantInstallment MerchantInstallment
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &merchantInstallment)
			if err != nil {
				merchantInstallmentsError <- Error.UnknownError(err.Error())
				continue
			}
			merchantInstallments <- merchantInstallment
		}
		for err := range errorChannel {
			merchantInstallmentsError <- err
		}
		close(merchantInstallments)
		close(merchantInstallmentsError)
	}()
	return merchantInstallments, merchantInstallmentsError
}

func Page(params map[string]interface{}, user user.User) ([]MerchantInstallment, string, Error.StarkErrors) {
	// Retrieve paged MerchantInstallment structs
	//
	// Receive a slice of up to 100 MerchantInstallment structs previously created in the Stark Bank API and the cursor to the next page.
	// Use this function instead of query if you want to manually page your requests.
	//
	// Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs.
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"John", "Paul"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	// Return:
	// - Slice of MerchantInstallment structs with updated attributes
	// - Cursor to retrieve the next page of MerchantInstallment structs
	var merchantInstallments []MerchantInstallment
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &merchantInstallments)
	if unmarshalError != nil {
		return merchantInstallments, cursor, err
	}
	return merchantInstallments, cursor, err
}
