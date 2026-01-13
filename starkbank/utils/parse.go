package utils

import (
	"strings"
	"github.com/starkbank/sdk-go/starkbank"
	Errors "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/parse"
)

func ParseAndVerify(content string, signature string, key string, user user.User) (string, Errors.StarkErrors) {
	if user == nil {
		response, err := parse.ParseAndVerify(content, signature, starkbank.SdkVersion, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.Host, starkbank.User, key)
		if err.Errors != nil {
			return "", err
		}
		return response.(string), err
	}
	response, err := parse.ParseAndVerify(content, signature, starkbank.SdkVersion, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.Host, user, key)
	if err.Errors != nil {
		return "", err
	}
	return response.(string), err
}

func ReplaceEmptyStringField(jsonStr, pattern, replacement string) string {
	return strings.ReplaceAll(jsonStr, pattern, replacement)
}