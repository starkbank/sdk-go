package corporatebalance

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	CorporateBalance struct
//
//	The CorporateBalance object displays the current corporate balance of the Workspace,
//	which is the result of the sum of all transactions within this
//	Workspace. The balance is never generated by the user, but it
//	can be retrieved to see the available information.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when CorporateBalance is created. ex: "5656565656565656"
//	- Amount [int]: current corporate balance amount of the Workspace in cents. ex: 200 (= R$ 2.00)
//	- Limit [int]: The maximum negative balance allowed by the user. ex: 10000 (= R$ 100.00)
//	- MaxLimit [int]: The maximum negative balance allowed by StarkBank. ex: 100000 (= R$ 1000.00)
//	- Currency [string]: Currency of the current Workspace. Expect others to be added eventually. ex: "BRL"
//	- Updated [string]: Latest update datetime for the CorporateBalance. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type CorporateBalance struct {
	Id       string     `json:",omitempty"`
	Amount   int        `json:",omitempty"`
	Limit    int        `json:",omitempty"`
	MaxLimit int        `json:",omitempty"`
	Currency string     `json:",omitempty"`
	Updated  *time.Time `json:",omitempty"`
}

var object CorporateBalance
var resource = map[string]string{"name": "CorporateBalance"}

func Get(user user.User) CorporateBalance {
	//	Retrieve the CorporateBalance struct
	//
	//	Receive the CorporateBalance struct linked to your Workspace in the Stark Bank API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- CorporateBalance struct with updated attributes
	balance := make(chan CorporateBalance)
	query := utils.Query(resource, nil, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err.Error())
			}
			balance <- object
		}
		close(balance)
	}()
	return <-balance
}
