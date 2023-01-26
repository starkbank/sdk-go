package sdk

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Workspace "github.com/starkbank/sdk-go/starkbank/workspace"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	Example "github.com/starkbank/sdk-go/tests/utils/examples"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestWorkspacePost(t *testing.T) {

	workspace, err := Workspace.Create(Example.Workspace(), Utils.ExampleOrganization)
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	assert.NotNil(t, workspace.Id)
}

func TestWorkspaceGet(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var workspaceList []Workspace.Workspace
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	workspaces := Workspace.Query(params, nil)
	for workspace := range workspaces {
		workspaceList = append(workspaceList, workspace)
	}

	workspace, err := Workspace.Get(workspaceList[rand.Intn(len(workspaceList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, workspace.Id)
}

func TestWorkspaceQuery(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var i int
	var params = map[string]interface{}{}
	params["limit"] = 1

	workspaces := Workspace.Query(params, nil)

	for workspace := range workspaces {
		assert.NotNil(t, workspace.Id)
		i++
	}
	assert.Equal(t, 1, i)
}

func TestWorkspacePage(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 1

	workspaces, cursor, err := Workspace.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, workspace := range workspaces {
		ids = append(ids, workspace.Id)
		assert.NotNil(t, workspace.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 1)
}

func TestWorkspaceReplace(t *testing.T) {

	var ids []string
	var params = map[string]interface{}{}
	params["limit"] = 1

	workspaces, cursor, err := Workspace.Page(params, Utils.ExampleOrganization.Replace("1234567890098"))
	if err.Errors != nil {
		for _, erro := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}
	for _, workspace := range workspaces {
		ids = append(ids, workspace.Id)
		assert.NotNil(t, workspace.Id)
		assert.NotNil(t, cursor)
	}
	assert.Len(t, ids, 1)
}

func TestWorkspaceUpdate(t *testing.T) {

	starkbank.User = Utils.ExampleProject

	bytes, _ := ioutil.ReadFile("file.png")

	var patchData = map[string]interface{}{}
	patchData["picture"] = bytes
	patchData["pictureType"] = "image/png"

	workspace, err := Workspace.Update("5647143184367616", patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, workspace.Id)
}
