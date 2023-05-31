package utils

import (
	"github.com/starkbank/sdk-go/starkbank"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/parse"
)

func ParseAndVerify(content string, signature string, key string, user user.User) string {
	if user == nil {
		return parse.ParseAndVerify(content, signature, starkbank.SdkVersion, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.Host, starkbank.User, key).(string)
	}
	return parse.ParseAndVerify(content, signature, starkbank.SdkVersion, starkbank.ApiVersion, starkbank.Language, starkbank.Timeout, starkbank.Host, user, key).(string)
}
