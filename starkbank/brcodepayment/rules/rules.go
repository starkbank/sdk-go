package rules

//	BrcodePayment.Rule struct
//
//	The BrcodePayment.Rule struct modifies the behavior of BrcodePayment structs when passed as an argument upon their creation.
//
//	Attributes (return-only):
//	- Key [string]: Rule to be customized, describes what BrcodePayment behavior will be altered. ex: "resendingLimit"
//	- Value [int]: Value of the rule. ex: 5

type Rule struct {
	Key   string `json:",omitempty"`
	Value int    `json:",omitempty"`
}
