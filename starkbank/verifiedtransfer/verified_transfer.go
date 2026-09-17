package verifiedtransfer

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/transfer/rule"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	VerifiedTransfer struct
//
//	When you initialize a VerifiedTransfer, the entity will not be automatically
//	created in the Stark Bank API. The 'create' function sends a slice of VerifiedTransfer
//	structs to the Stark Bank API and returns the slice of created structs.
//
//	Parameters (required):
//	- Amount [int]: transfer value in cents. ex: 1234 (= R$ 12.34)
//	- AccountId [string]: receiver's VerifiedAccount id. ex: "5656565656565656"
//
//	Parameters (optional):
//	- AccountType [string, default "checking"]: receiver bank account type. This parameter only has effect on Pix Transfers. ex: "checking", "savings", "salary" or "payment"
//	- ExternalId [string, default nil]: url safe string that must be unique among all your transfers. Duplicated external_ids will cause failures. By default, this parameter will block any transfer that repeats amount and receiver information on the same date. ex: "my-internal-id-123456"
//	- Scheduled [time.Time, default now]: date or datetime when the transfer will be processed. May be pushed to next business day if necessary. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC)
//	- Description [string, default nil]: optional description to override default description to be shown in the bank statement. ex: "Payment for service #1234"
//	- DisplayDescription [string, default nil]: optional description to be shown in the receiver bank interface. ex: "Payment for service #1234"
//	- Tags [slice of strings, default nil]: slice of strings for reference when searching for verified transfers. ex: []string{"employees", "monthly"}
//	- Rules [slice of Transfer.Rule structs, default nil]: slice of Transfer.Rule structs for modifying transfer behavior. ex: []rule.Rule{{Key: "resendingLimit", Value: 5}}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the VerifiedTransfer is created. ex: "5656565656565656"
//	- Fee [int]: fee charged when the transfer is created. ex: 200 (= R$ 2.00)
//	- Status [string]: current verified transfer status. ex: "created", "processing", "success" or "failed"
//	- TransactionIds [slice of strings]: ledger transaction ids linked to this transfer (if there are two, the second is the chargeback). ex: []string{"19827356981273"}
//	- Metadata [map[string]interface{}]: object used to store additional information about the VerifiedTransfer struct.
//	- Created [time.Time]: creation datetime for the VerifiedTransfer. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Updated [time.Time]: latest update datetime for the VerifiedTransfer. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type VerifiedTransfer struct {
	Id                 string                 `json:",omitempty"`
	Amount             int                    `json:",omitempty"`
	AccountId          string                 `json:",omitempty"`
	AccountType        string                 `json:",omitempty"`
	ExternalId         string                 `json:",omitempty"`
	Scheduled          *time.Time             `json:",omitempty"`
	Description        string                 `json:",omitempty"`
	DisplayDescription string                 `json:",omitempty"`
	Tags               []string               `json:",omitempty"`
	Rules              []rule.Rule            `json:",omitempty"`
	Fee                int                    `json:",omitempty"`
	Status             string                 `json:",omitempty"`
	TransactionIds     []string               `json:",omitempty"`
	Metadata           map[string]interface{} `json:",omitempty"`
	Created            *time.Time             `json:",omitempty"`
	Updated            *time.Time             `json:",omitempty"`
}

var resource = map[string]string{"name": "VerifiedTransfer"}

func Create(verifiedTransfers []VerifiedTransfer, user user.User) ([]VerifiedTransfer, Error.StarkErrors) {
	//	Create VerifiedTransfers
	//
	//	Send a slice of VerifiedTransfer structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- verifiedTransfers [slice of VerifiedTransfer structs]: slice of VerifiedTransfer structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of VerifiedTransfer structs with updated attributes
	create, err := utils.Multi(resource, verifiedTransfers, nil, user)
	unmarshalError := json.Unmarshal(create, &verifiedTransfers)
	if unmarshalError != nil {
		return verifiedTransfers, err
	}
	return verifiedTransfers, err
}
