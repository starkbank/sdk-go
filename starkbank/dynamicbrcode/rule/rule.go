package rule

//	DynamicBrCode.Rule struct
//
//	The DynamicBrCode.Rule struct modifies the behavior of DynamicBrCode structs when passed as an argument upon their creation.
//
//	Attributes (return-only):
//	- Key [string]: Rule to be customized, describes what DynamicBrCode behavior will be altered. ex: "allowedTaxIds"
//	- Value [list of string]: Value of the rule. ex: ["012.345.678-90", "45.059.493/0001-73"]

type Rule struct {
	Key   string   `json:",omitempty"`
	Value []string `json:",omitempty"`
}
