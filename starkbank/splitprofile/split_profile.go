package splitprofile

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"time"
)

//	SplitProfile struct
//
//	When you create a Split, the SplitProfile entity is automatically created. If you haven't
//	created a Split yet, you can use the 'Put' function to create your SplitProfile.
//
//	Parameters (optional):
//	- Interval [string, default "week"]: frequency of transfer. Options: "day", "week", "month"
//	- Delay [int, default nil]: how long the amount will stay at the workspace, in milliseconds. ex: 604800
//	- Tags [slice of strings, default nil]: slice of strings for tagging
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the SplitProfile is created. ex: "5656565656565656"
//	- Status [string]: current SplitProfile status. ex: "created"
//	- Created [time.Time]: creation datetime for the SplitProfile. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)
//	- Updated [time.Time]: latest update datetime for the SplitProfile. ex: time.Date(2020, 3, 10, 30, 30, 0, 0, time.UTC)

type SplitProfile struct {
	Id       string     `json:",omitempty"`
	Interval string     `json:",omitempty"`
	Delay    *int       `json:",omitempty"`
	Tags     []string   `json:",omitempty"`
	Status   string     `json:",omitempty"`
	Created  *time.Time `json:",omitempty"`
	Updated  *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "SplitProfile"}

func Put(profiles []SplitProfile, user user.User) ([]SplitProfile, Error.StarkErrors) {
	//	Create SplitProfiles or update them if they already exist
	//
	//	Send a slice of SplitProfile structs for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- profiles [slice of SplitProfile structs]: slice of SplitProfile structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of SplitProfile structs with updated attributes
	data := map[string][]map[string]interface{}{}
	payload := api.ApiJson(profiles, resource)
	raw, err := utils.PutRaw(api.Endpoint(resource), payload, user, nil, "", true)
	if err.Errors != nil {
		return profiles, err
	}
	unmarshalErrorRaw := json.Unmarshal(raw.Content, &data)
	if unmarshalErrorRaw != nil {
		return profiles, err
	}
	jsonBytes, _ := json.Marshal(data[api.LastNamePlural(resource)])
	unmarshalError := json.Unmarshal(jsonBytes, &profiles)
	if unmarshalError != nil {
		return profiles, err
	}
	return profiles, err
}

func Get(id string, user user.User) (SplitProfile, Error.StarkErrors) {
	//	Retrieve a specific SplitProfile by its id
	//
	//	Receive a single SplitProfile struct previously created in the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- SplitProfile struct that corresponds to the given id.
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
	//	Receive a channel of SplitProfile structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: date filter for objects created or updated only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: date filter for objects created or updated only before specified date. ex: "2020-03-10"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- channel of SplitProfile structs with updated attributes
	splitProfiles := make(chan SplitProfile)
	splitProfilesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var splitProfile SplitProfile
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
		close(splitProfilesError)
		close(splitProfiles)
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
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- after [string, default nil]: date filter for objects created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: date filter for objects created only before specified date. ex: "2020-03-10"
	//		- tags [slice of strings, default nil]: tags to filter retrieved objects. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: filter for status of retrieved objects. ex: "success"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- slice of SplitProfile structs with updated attributes
	//	- cursor to retrieve the next page of SplitProfile structs
	var splitProfiles []SplitProfile
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &splitProfiles)
	if unmarshalError != nil {
		return splitProfiles, cursor, err
	}
	return splitProfiles, cursor, err
}
