package starkbank

import (
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/hosts"
)

var SdkVersion = "0.6.0"
var Timeout = 15
var ApiVersion = "v2"
var Host = hosts.Bank
var Language = "pt-BR"
var User user.User = nil
