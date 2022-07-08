package log

type Log struct {
	Id      string        `json:"id"`
	Payment BrcodePayment `json:"payment"`
	Type    string        `json:"type"`
	Created string        `json:"created"`
}

var resource = map[string]any{"class": Log{}, "name": "BrcodePaymentLog"}
