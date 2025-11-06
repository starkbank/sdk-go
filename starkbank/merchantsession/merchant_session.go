package merchantsession

import (
	"encoding/json"
	"time"
	AllowedInstallment "github.com/starkbank/sdk-go/starkbank/merchantsession/allowedinstallment"
	"github.com/starkbank/sdk-go/starkbank/utils"
	"github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

// Check out our API Documentation at https://starkbank.com/docs/api#merchant-session

type MerchantSession struct {
	Id                  string                                  `json:",omitempty"`
	AllowedFundingTypes []string                                `json:",omitempty"`
	AllowedInstallments []AllowedInstallment.AllowedInstallment `json:",omitempty"`
	AllowedIps          []string                                `json:",omitempty"`
	ChallengeMode       string                                  `json:",omitempty"`
	Created             *time.Time                              `json:",omitempty"`
	Expiration          int                                     `json:",omitempty"`
	Status              string                                  `json:",omitempty"`
	Tags                []string                                `json:",omitempty"`
	Updated             *time.Time                              `json:",omitempty"`
	Uuid                string                                  `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantSession"}

func Create(merchantSession MerchantSession, user user.User) (MerchantSession, error.StarkErrors) {
	create, err := utils.Single(resource, merchantSession, user)
	unmarshalError := json.Unmarshal(create, &merchantSession)
	if unmarshalError != nil {
		return merchantSession, err
	}
	return merchantSession, err
}

func Get(id string, user user.User) (MerchantSession, error.StarkErrors) {
	var merchantSession MerchantSession
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &merchantSession)
	if unmarshalError != nil {
		return merchantSession, err
	}
	return merchantSession, err
}

func Query(params map[string]interface{}, user user.User) (chan MerchantSession, chan error.StarkErrors) {
	var merchantSession MerchantSession
	merchantSessions := make(chan MerchantSession)
	merchantSessionsError := make(chan error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &merchantSession)
			if err != nil {
				merchantSessionsError <- error.UnknownError(err.Error())
				continue
			}
			merchantSessions <- merchantSession
		}
		for err := range errorChannel {
			merchantSessionsError <- err
		}
		close(merchantSessions)
		close(merchantSessionsError)
	}()
	return merchantSessions, merchantSessionsError
}

func Page(params map[string]interface{}, user user.User) ([]MerchantSession, string, error.StarkErrors) {
	var merchantSessions []MerchantSession
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &merchantSessions)
	if unmarshalError != nil {
		return merchantSessions, cursor, err
	}
	return merchantSessions, cursor, err
}

func PostPurchase(uuid string, payload Purchase, user user.User) (Purchase, error.StarkErrors) {
	post, err := utils.PostSubResource(resource, payload, uuid, user, SubResourcePurchase)
	unmarshalError := json.Unmarshal(post, &purchase)
	if unmarshalError != nil {
		return purchase, err
	}
	return purchase, err
}
