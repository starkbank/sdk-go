package rule

//	Transfer.Rule struct
//
//	The Transfer.Rule struct modifies the behavior of Transfer structs when passed as an argument upon their creation.
//
//	Attributes (return-only):
//	- Key [string]: Rule to be customized, describes what Transfer behavior will be altered. ex: "resendingLimit"
//	- Value [int]: Value of the rule. ex: 5

type Rule struct {
	Key   string `json:",omitempty"`
	Value int
}
