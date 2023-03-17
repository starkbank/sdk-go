package metadata

//	Transfer.Metadata struct
//
//	The Transfer.Metadata struct contains additional information about the Transfer struct.
//
//	Attributes (return-only):
//	- Authentication [string]: Central Bank’s unique ID for Pix transactions (EndToEndID). ex: “E200181832023031715008Scr7tD63TS”

type Metadata struct {
	Authentication string `json:",omitempty"`
}
