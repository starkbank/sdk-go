package merchantcard

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

// Check out our API Documentation at https://starkbank.com/docs/api#merchant-card

type MerchantCard struct {
	Id 			   	string                 	`json:",omitempty"`
	Ending 	   		string                 	`json:",omitempty"`
	FundingType		string                 	`json:",omitempty"`
	HolderName 		string                 	`json:",omitempty"`
	Network 	   	string                 	`json:",omitempty"`
	Status 		   	string                 	`json:",omitempty"`
	Tags 		   	[]string               	`json:",omitempty"`
	Expiration 	   	string             		`json:",omitempty"`
	Created 	   	*time.Time             	`json:",omitempty"`
	Updated 	   	*time.Time             	`json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantCard"}


func Get(id string, user user.User) (MerchantCard, Error.StarkErrors){
	var merchantCard MerchantCard
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &merchantCard)
	if unmarshalError != nil {
		return merchantCard, err
	}
	return merchantCard, err
}


func Query(params map[string]interface{}, user user.User) chan MerchantCard {
	var merchantCard MerchantCard
	merchantCards := make(chan MerchantCard)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &merchantCard)
			if err != nil {
				panic(err)
			}
			merchantCards <- merchantCard
		}
		close(merchantCards)
	}()
	return merchantCards
}


func Page(params map[string]interface{}, user user.User) ([]MerchantCard, string, Error.StarkErrors) {
	var merchantCards []MerchantCard
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &merchantCards)
	if unmarshalError != nil {
		return merchantCards, cursor, err
	}
	return merchantCards, cursor, err
}
