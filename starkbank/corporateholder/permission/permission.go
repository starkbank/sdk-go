package permission

import "time"

//	CorporateHolder.Permission struct
//
//	Permission object represents access granted to an user for a particular cardholder
//
//  Parameters (optional):
//  - owner_id [string, default nil]: owner unique id. ex: "5656565656565656"
//  - owner_type [string, default nil]: owner type. ex: "project"
//
//  Attributes (return only):
//  - owner_email [string]: email address of the owner. ex: "tony@starkbank.com
//  - owner_name [string]: name of the owner. ex: "Tony Stark"
//  - owner_picture_url [string]: Profile picture Url of the owner. ex: "https://storage.googleapis.com/api-ms-workspace-dev.appspot.com/pictures/workspace/5647143184367616.png?20230528223305"
//  - owner_status [string]: current owner status. ex: "active", "blocked", "canceled"
//	- Created [time.Time]: Creation datetime for the CorporateHolder.Permission. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Permission struct {
	OwnerId         string     `json:",omitempty"`
	OwnerType       string     `json:",omitempty"`
	OwnerEmail      string     `json:",omitempty"`
	OwnerName       string     `json:",omitempty"`
	OwnerPictureUrl string     `json:",omitempty"`
	OwnerStatus     string     `json:",omitempty"`
	Created         *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "Permission"}
