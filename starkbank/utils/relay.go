package utils

import (
	"github.com/starkbank/sdk-go/starkbank"
	Errors "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/rest"
)

func Page(resource map[string]string, params map[string]interface{}, user user.User) ([]byte, string, Errors.StarkErrors) {
	if user == nil {
		return rest.GetPage(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.User, resource, params)
	}
	return rest.GetPage(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, user, resource, params)
}

func Query(resource map[string]string, params map[string]interface{}, user user.User) chan map[string]interface{} {
	if user == nil {
		return rest.GetStream(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.User, resource, params)
	}
	return rest.GetStream(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, user, resource, params)
}

func Get(resource map[string]string, id string, query map[string]interface{}, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.GetId(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.User, resource, id, query)
	}
	return rest.GetId(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, user, resource, id, query)
}

func GetContent(resource map[string]string, id string, params map[string]interface{}, user user.User, content string) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.GetContent(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.User, resource, id, content, params)
	}
	return rest.GetContent(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, user, resource, id, content, params)
}

func SubResource(resource map[string]string, id string, user user.User, subResource map[string]string) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.GetSubResource(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.User, resource, id, subResource, nil)
	}
	return rest.GetSubResource(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, user, resource, id, subResource, nil)
}

func Multi(resource map[string]string, entities interface{}, query map[string]interface{}, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.PostMulti(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.User, resource, entities, query)
	}
	return rest.PostMulti(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, user, resource, entities, query)
}

func Single(resource map[string]string, entity interface{}, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.PostSingle(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.User, resource, entity, nil)
	}
	return rest.PostSingle(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, user, resource, entity, nil)
}

func Delete(resource map[string]string, id string, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.DeleteId(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.User, resource, id, nil)
	}
	return rest.DeleteId(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, user, resource, id, nil)
}

func Patch(resource map[string]string, id string, payload map[string]interface{}, user user.User) ([]byte, Errors.StarkErrors) {
	if user == nil {
		return rest.PatchId(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.User, resource, id, payload, nil)
	}
	return rest.PatchId(starkbank.SdkVersion, starkbank.Host, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, user, resource, id, payload, nil)
}
