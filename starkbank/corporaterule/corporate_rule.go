package corporaterule

import (
	CardMethod "github.com/starkbank/sdk-go/starkbank/cardmethod"
	MerchantCategory "github.com/starkbank/sdk-go/starkbank/merchantcategory"
	MerchantCountry "github.com/starkbank/sdk-go/starkbank/merchantcountry"
)

//	CorporateRule struct
//
//	The CorporateRule struct displays the spending rules of CorporateCards and CorporateHolders created in your Workspace.
//
//	Parameters (required):
//	- Name [string]: Rule name. ex: "Travel" or "Food"
//	- Amount [int]: Maximum amount that can be spent in the informed interval. ex: 200000 (= R$ 2000.00)
//
//	Parameters (optional):
//	- Id [string, default nil]: Unique id returned when a CorporateRule is created, used to update a specific CorporateRule. ex: "5656565656565656"
//	- Interval [string, default "lifetime"]: Interval after which the rule amount counter will be reset to 0. ex: "instant", "day", "week", "month", "year" or "lifetime"
//  - Schedule [string, default null]: schedule time for user to spend. ex: "every monday, wednesday from 00:00 to 23:59 in America/Sao_Paulo"
//  - Purposes [slice of strings default null]: list of strings representing the allowed purposes for card purchases, you can use this to restrict ATM withdrawals. ex: []string{"purchase", "withdrawal"}
//	- CurrencyCode [string, default "BRL"]: Code of the currency that the rule amount refers to. ex: "BRL" or "USD"
//	- Categories [slice of MerchantCategory structs, default nil]: Merchant categories accepted by the rule. ex: []string{MerchantCategory(code="fastFoodRestaurants")]
//  - Countries [slice of MerchantCountry structs, default nil]: Countries accepted by the rule. ex: []string{MerchantCountry(code="BRA")]
//  - Methods [slice of CardMethod structs, default nil]: Card purchase methods accepted by the rule. ex: []string{CardMethod(code="magstripe")]
//
//	Attributes (expanded return-only):
//	- CounterAmount [int]: Current rule spent amount. ex: 1000
//	- CurrencySymbol [string]: Currency symbol. ex: "R$"
//	- CurrencyName [string]: Currency name. ex: "Brazilian Real"

type CorporateRule struct {
	Name           string                              `json:",omitempty"`
	Amount         int                                 `json:",omitempty"`
	Id             string                              `json:",omitempty"`
	Interval       string                              `json:",omitempty"`
	Schedule       string                              `json:",omitempty"`
	Purposes       []string                            `json:",omitempty"`
	CurrencyCode   string                              `json:",omitempty"`
	Categories     []MerchantCategory.MerchantCategory `json:",omitempty"`
	Countries      []MerchantCountry.MerchantCountry   `json:",omitempty"`
	Methods        []CardMethod.CardMethod             `json:",omitempty"`
	CounterAmount  int                                 `json:",omitempty"`
	CurrencySymbol string                              `json:",omitempty"`
	CurrencyName   string                              `json:",omitempty"`
}
