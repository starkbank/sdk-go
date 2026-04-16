package splitprofile

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	SplitProfile struct
//
//	When you create a Split, the entity SplitProfile will be automatically created. If you haven't
//	created a Split yet, you can use the Put method to create your SplitProfile.
//
//	Parameters (optional):
//	- Interval [string]: Frequency of transfer, default "week". Options: "day", "week", "month"
//	- Delay [int]: How long the amount will stay at the workspace in milliseconds. ex: 604800
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"tony", "stark"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the SplitProfile is created. ex: "5656565656565656"
//	- Status [string]: Current SplitProfile status. ex: "created"
//	- Created [time.Time]: Creation datetime for the SplitProfile. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC)
//	- Updated [time.Time]: Latest update datetime for the SplitProfile. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC)

type SplitProfile struct {
	Interval string     `json:",omitempty"`
	Delay    int        `json:",omitempty"`
	Tags     []string   `json:",omitempty"`
	Id       string     `json:",omitempty"`
	Status   string     `json:",omitempty"`
	Created  *time.Time `json:",omitempty"`
	Updated  *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "SplitProfile"}

func Put(profiles []SplitProfile, user user.User) ([]SplitProfile, Error.StarkErrors) {
	//	Create or update SplitProfiles
	//
	//	Send a slice of SplitProfile structs for creation or update in the Stark Bank API.
	//	If a SplitProfile already exists, it will be updated.
	//
	//	Parameters (required):
	//	- profiles [slice of SplitProfile structs]: Slice of SplitProfile structs to be created or updated in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of SplitProfile structs with updated attributes
	body := map[string]interface{}{
		"profiles": profiles,
	}
	put, err := utils.PutRaw("split-profile/", body, user, nil, "Joker", false)
	var result []SplitProfile
	unmarshalError := json.Unmarshal(put.Content, &result)
	if unmarshalError != nil {
		return result, err
	}
	return result, err
}

func Get(id string, user user.User) (SplitProfile, Error.StarkErrors) {
	//	Retrieve a specific SplitProfile by its id
	//
	//	Receive a single SplitProfile struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- SplitProfile struct that corresponds to the given id
	var splitProfile SplitProfile
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &splitProfile)
	if unmarshalError != nil {
		return splitProfile, err
	}
	return splitProfile, err
}

func Query(params map[string]interface{}, user user.User) (chan SplitProfile, chan Error.StarkErrors) {
	//	Retrieve SplitProfile structs
	//
	//	Receive a channel of SplitProfile structs from the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Channel of SplitProfile structs with updated attributes
	var splitProfile SplitProfile
	splitProfiles := make(chan SplitProfile)
	splitProfilesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &splitProfile)
			if err != nil {
				splitProfilesError <- Error.UnknownError(err.Error())
				continue
			}
			splitProfiles <- splitProfile
		}
		for err := range errorChannel {
			splitProfilesError <- err
		}
		close(splitProfiles)
		close(splitProfilesError)
	}()
	return splitProfiles, splitProfilesError
}

func Page(params map[string]interface{}, user user.User) ([]SplitProfile, string, Error.StarkErrors) {
	//	Retrieve paged SplitProfile structs
	//
	//	Receive a slice of up to 100 SplitProfile structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: List of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "created"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of SplitProfile structs with updated attributes
	//	- Cursor to retrieve the next page of SplitProfile structs
	var splitProfiles []SplitProfile
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &splitProfiles)
	if unmarshalError != nil {
		return splitProfiles, cursor, err
	}
	return splitProfiles, cursor, err
}
