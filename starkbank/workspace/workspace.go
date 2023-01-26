package workspace

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

//	Workspace struct
//
//	Workspaces are bank accounts. They have independent balances, statements, operations and permissions.
//	The only property that is shared between your workspaces is that they are linked to your organization,
//	which carries your basic informations, such as tax ID, name, etc..
//
//	Parameters (required):
//	- Username [string]: Simplified name to define the workspace URL. This name must be unique across all Stark Bank Workspaces. Ex: "starkbankworkspace"
//	- Name [string]: Full name that identifies the Workspace. This name will appear when people access the Workspace on our platform, for example. Ex: "Stark Bank Workspace"
//
//	Parameters (optional):
//	- AllowedTaxIds [slice of strings, default nil]: slice of tax IDs that will be allowed to send Deposits to this Workspace. If empty, all are allowed. []string{"012.345.678-90", "20.018.183/0001-80"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the workspace is created. ex: "5656565656565656"

type Workspace struct {
	Username      string   `json:",omitempty"`
	Name          string   `json:",omitempty"`
	AllowedTaxIds []string `json:",omitempty"`
	Id            string   `json:",omitempty"`
}

var object Workspace
var objects []Workspace
var resource = map[string]string{"name": "Workspace"}

func Create(workspace Workspace, user user.User) (Workspace, Error.StarkErrors) {
	//	Create Workspace
	//
	//	Send a Workspace for creation in the Stark Bank API
	//
	//	Parameters (required):
	//	- user [Organization struct, default nil]: Organization struct. Not necessary if starkbank.user was set before function call
	//	- workspace [Workspace struct]: workspaceData parameters for the creation of the workspace
	//		Parameters (required):
	//		- username [string]: Simplified name to define the workspace URL. This name must be unique across all Stark Bank Workspaces. Ex: "starkbankworkspace"
	//			- name [string]: Full name that identifies the Workspace. This name will appear when people access the Workspace on our platform, for example. Ex: "Stark Bank Workspace"
	//		Parameters (optional):
	//			- allowedTaxIds [slice of strings, default nil]: slice of tax IDs that will be allowed to send Deposits to this Workspace. If empty, all are allowed. ex: []string{"012.345.678-90", "20.018.183/0001-80"}
	//
	//	Return:
	//	- Workspace struct with updated attributes
	create, err := utils.Single(resource, workspace, user)
	unmarshalError := json.Unmarshal(create, &workspace)
	if unmarshalError != nil {
		return workspace, err
	}
	return workspace, err
}

func Get(id string, user user.User) (Workspace, Error.StarkErrors) {
	//	Retrieve a specific Workspace by its id
	//
	//	Receive a single Workspace struct previously created in the Stark Bank API by passing its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- Workspace struct with updated attributes
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan Workspace {
	//	Retrieve Workspaces
	//
	//	Receive a generator of Workspace structs previously created in the Stark Bank API.
	//	If no filters are passed and the user is an Organization, all of the Organization Workspaces
	//	will be retrieved.
	//
	//	Parameters (required):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Parameters (optional):
	//	- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- username [string, default nil]: query by the simplified name that defines the workspace URL. This name is always unique across all Stark Bank Workspaces. Ex: "starkbankworkspace"
	//	- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Return:
	//	- generator of Workspace structs with updated attributes
	workspaces := make(chan Workspace)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				panic(err)
			}
			workspaces <- object
		}
		close(workspaces)
	}()
	return workspaces
}

func Update(id string, patchData map[string]interface{}, user user.User) (Workspace, Error.StarkErrors) {
	//	Update Workspace entity
	//
	//	Update a Workspace by passing its ID.
	//
	//	Parameters (required):
	//	- id [string]: Workspace ID. ex: '5656565656565656'
	//	- patchData [map[string]interface{}]: map containing the attributes to be updated. Allowed parameters: "name", "username and "allowedTaxIds". ex: map[string]interface{}{"name": "So Far Away"}
	//		Parameters (conditionally required):
	//		- pictureType [string]: picture MIME type. This parameter will be required if the picture parameter is informed ex: "image/png" or "image/jpeg"
	//		Parameters (optional):
	//		- username [string, default nil]: Simplified name to define the workspace URL. This name must be unique across all Stark Bank Workspaces. ex: "starkbank-workspace"
	//		- name [string, default nil]: Full name that identifies the Workspace. This name will appear when people access the Workspace on our platform, for example. ex: "Stark Bank Workspace"
	//		- allowedTaxIds [slice of strings, default nil]: slice of tax IDs that will be allowed to send Deposits to this Workspace. If empty, all are allowed. ex: ["012.345.678-90", "20.018.183/0001-80"]
	//		- picture [bytes, default nil]: Binary buffer of the picture. ex: ioutil.ReadFile("file.png")
	//	- user [Organization/Project struct, default nil]: Organization or Project object. Not necessary if starkbank.user was set before function call
	//
	//	Return:
	//	- target Workspace with updated attributes
	if patchData["picture"] != nil {
		patchData["picture"] = fmt.Sprintf("data:%v;base64,%v", patchData["pictureType"], base64.StdEncoding.EncodeToString(patchData["picture"].([]byte)))
		delete(patchData, "pictureType")
	}
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Page(params map[string]interface{}, user user.User) ([]Workspace, string, Error.StarkErrors) {
	//	Retrieve paged Workspaces
	//
	//	Receive a slice of up to 100 Workspace structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (required):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.user was set before function call
	//
	//	Parameters (optional):
	//	- cursor [string, default nil]: cursor returned on the previous page function call
	//	- limit [int, default 100]: maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//	- username [string, default nil]: query by the simplified name that defines the workspace URL. This name is always unique across all Stark Bank Workspaces. Ex: "starkbankworkspace"
	//	- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Return:
	//	- slice of Workspace structs with updated attributes
	//	- cursor to retrieve the next page of Workspace structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
