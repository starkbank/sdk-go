package paymentpreview

import (
	"encoding/json"
	"github.com/starkbank/sdk-go/starkbank/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"time"
)

//	PaymentPreview struct
//
//	A PaymentPreview is used to get information from a payment code before confirming the payment.
//	This resource can be used to preview BR Codes and bar codes of boleto, tax and utility payments
//
//	Parameters (required):
//	- Id [string]: Main identification of the payment. This should be the BR Code for Pix payments and lines or bar codes for payment slips. ex: "34191.09008 63571.277308 71444.640008 5 81960000000062", "00020126580014br.gov.bcb.pix0136a629532e-7693-4846-852d-1bbff817b5a8520400005303986540510.005802BR5908T'Challa6009Sao Paulo62090505123456304B14A"
//
//	Parameters (optional):
//	- Scheduled [time.time, default today]: intended payment date. Right now, this parameter only has effect on BrcodePreviews. ex: time.Date(2020, 3, 10, 0, 0, 10, 0, time.UTC),
//
//	Attributes (return-only):
//	- Type [string]: Payment type. ex: "brcode-payment", "boleto-payment", "utility-payment" or "tax-payment"
//	- Payment [BrcodePreview struct, BoletoPreview struct, UtilityPreview or TaxPreview struct]: Information preview of the informed payment.

type PaymentPreview struct {
	Id        string      `json:",omitempty"`
	Payment   interface{} `json:",omitempty"`
	Type      string      `json:",omitempty"`
	Scheduled interface{} `json:",omitempty"`
}

var subresource = map[string]string{"name": "PaymentPreview"}

func Create(previews []PaymentPreview, user user.User) ([]PaymentPreview, Error.StarkErrors) {
	//	Create PaymentPreviews
	//
	//	Send a slice of PaymentPreviews structs for processing in the Stark Bank API
	//
	//	Parameters (required):
	//	- previews [slice of PaymentPreview structs]: slice of PaymentPreview structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkbank.User was set before function call
	//
	//	Return:
	//	- Slice of PaymentPreview structs with updated attributes
	create, err := utils.Multi(subresource, previews, nil, user)
	if err.Errors != nil {
		return nil, err
	}

	unmarshalError := json.Unmarshal(create, &previews)
	if unmarshalError != nil {
		return nil, Error.UnknownError(unmarshalError.Error())
	}

	parsedPreviews, parseErr := ParsePreviews(previews)
	if parseErr.Errors != nil {
		return nil, parseErr
	}
	return parsedPreviews, Error.StarkErrors{}
}

func (e PaymentPreview) ParsePreview() (PaymentPreview, Error.StarkErrors) {
	if e.Type == "tax-payment" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &PreviewTax)
		if unmarshalError != nil {
			return PaymentPreview{}, Error.UnknownError(unmarshalError.Error())
		}
		scheduled, _ := time.Parse("2006-01-02", e.Scheduled.(string))
		e.Scheduled = scheduled
		e.Payment = PreviewTax
		return e, Error.StarkErrors{}
	}
	if e.Type == "brcode-payment" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &PreviewBrcode)
		if unmarshalError != nil {
			return PaymentPreview{}, Error.UnknownError(unmarshalError.Error())
		}
		scheduled, _ := time.Parse("2006-01-02", e.Scheduled.(string))
		e.Scheduled = scheduled
		e.Payment = PreviewBrcode
		return e, Error.StarkErrors{}
	}
	if e.Type == "boleto-payment" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &PreviewBoleto)
		if unmarshalError != nil {
			return PaymentPreview{}, Error.UnknownError(unmarshalError.Error())
		}
		scheduled, _ := time.Parse("2006-01-02", e.Scheduled.(string))
		e.Scheduled = scheduled
		e.Payment = PreviewBoleto
		return e, Error.StarkErrors{}
	}
	if e.Type == "utility-payment" {
		marshal, _ := json.Marshal(e.Payment)
		unmarshalError := json.Unmarshal(marshal, &PreviewUtility)
		if unmarshalError != nil {
			return PaymentPreview{}, Error.UnknownError(unmarshalError.Error())
		}
		scheduled, _ := time.Parse("2006-01-02", e.Scheduled.(string))
		e.Scheduled = scheduled
		e.Payment = PreviewUtility
		return e, Error.StarkErrors{}
	}
	return e, Error.StarkErrors{}
}

func ParsePreviews(previews []PaymentPreview) ([]PaymentPreview, Error.StarkErrors) {
	var err Error.StarkErrors
	for i := 0; i < len(previews); i++ {
		previews[i], err = previews[i].ParsePreview()
		if err.Errors != nil {
			return nil, err
		}
	}
	return previews, Error.StarkErrors{}
}
