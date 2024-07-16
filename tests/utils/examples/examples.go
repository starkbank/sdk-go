package examples

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank/boleto"
	"github.com/starkbank/sdk-go/starkbank/boletoholmes"
	"github.com/starkbank/sdk-go/starkbank/boletopayment"
	"github.com/starkbank/sdk-go/starkbank/brcodepayment"
	"github.com/starkbank/sdk-go/starkbank/corporatecard"
	"github.com/starkbank/sdk-go/starkbank/corporateholder"
	"github.com/starkbank/sdk-go/starkbank/corporateholder/permission"
	"github.com/starkbank/sdk-go/starkbank/corporateinvoice"
	"github.com/starkbank/sdk-go/starkbank/corporatewithdrawal"
	"github.com/starkbank/sdk-go/starkbank/darfpayment"
	"github.com/starkbank/sdk-go/starkbank/dynamicbrcode"
	"github.com/starkbank/sdk-go/starkbank/invoice"
	Rule "github.com/starkbank/sdk-go/starkbank/invoice/rule"
	"github.com/starkbank/sdk-go/starkbank/paymentpreview"
	"github.com/starkbank/sdk-go/starkbank/paymentrequest"
	"github.com/starkbank/sdk-go/starkbank/taxpayment"
	"github.com/starkbank/sdk-go/starkbank/transaction"
	"github.com/starkbank/sdk-go/starkbank/transfer"
	"github.com/starkbank/sdk-go/starkbank/transfer/rule"
	"github.com/starkbank/sdk-go/starkbank/utilitypayment"
	"github.com/starkbank/sdk-go/starkbank/webhook"
	"github.com/starkbank/sdk-go/starkbank/workspace"
	Utils "github.com/starkbank/sdk-go/tests/utils"
	"math/rand"
	"os"
	"time"
)

func CorporateCard() corporatecard.CorporateCard {

	var holderIds []corporateholder.CorporateHolder
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)
	params["status"] = "active"

	holders := corporateholder.Query(params, Utils.ExampleProject)
	for holder := range holders {
		holderIds = append(holderIds, holder)
	}
	card := corporatecard.CorporateCard{
		HolderId: holderIds[0].Id,
	}
	return card
}

func CorporateHolder() []corporateholder.CorporateHolder {

	holders := []corporateholder.CorporateHolder{
		{
			Name:        "Iron Bank S.A.10",
			Tags:        []string{"Traveler Employee"},
			Permissions: []permission.Permission{{OwnerId: os.Getenv("HOLDER_ID"), OwnerType: "project"}},
		},
	}
	return holders
}

func CorporateInvoice() corporateinvoice.CorporateInvoice {

	invoice := corporateinvoice.CorporateInvoice{
		Amount: 1000,
	}
	return invoice
}

func CorporateWithdrawal() corporatewithdrawal.CorporateWithdrawal {

	withdrawal := corporatewithdrawal.CorporateWithdrawal{
		Amount:     1000,
		ExternalId: "123456789",
	}
	return withdrawal
}

func Holmes(id string) []boletoholmes.BoletoHolmes {

	holmes := []boletoholmes.BoletoHolmes{
		{
			BoletoId: id,
		},
	}
	return holmes
}

func BoletoPayment() boletopayment.BoletoPayment {

	payments := boletopayment.BoletoPayment{
		TaxId:       "20.018.183/0001-80",
		Description: "SDK-Go-Boleto-Payment-Test",
		Line:        "34191.09008 76038.597308 71444.640008 4 92150000028000",
	}

	return payments
}

func BoletosPayment() []boletopayment.BoletoPayment {

	var boletoList []boleto.Boleto
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)
	params["status"] = "registered"

	boletos := boleto.Query(params, Utils.ExampleProject)
	for boleto := range boletos {
		boletoList = append(boletoList, boleto)
	}

	payments := []boletopayment.BoletoPayment{
		{
			TaxId:       boletoList[rand.Intn(len(boletoList))].TaxId,
			Description: "SDK-Go-Boleto-Payment-Test",
			Line:        boletoList[rand.Intn(len(boletoList))].Line,
		},
	}

	return payments
}

func Boleto() []boleto.Boleto {

	due := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	boletos := []boleto.Boleto{
		{
			Amount:      400000,
			Name:        "Iron Bank S.A.",
			TaxId:       "20.018.183/0001-80",
			StreetLine1: "Av. Faria Lima, 1844",
			StreetLine2: "CJ 13",
			District:    "Itaim Bibi",
			City:        "São Paulo",
			StateCode:   "SP",
			ZipCode:     "01500-000",
			Due:         &due,
		},
		{
			Amount:      800000,
			Name:        "Iron Bank S.A.",
			TaxId:       "38.446.231/0001-04",
			StreetLine1: "Kubasch Street, 900",
			StreetLine2: "wefwe",
			District:    "Ronny",
			City:        "Emmet City",
			StateCode:   "SP",
			ZipCode:     "01420-020",
		},
	}
	return boletos
}

func BrcodePayment() []brcodepayment.BrcodePayment {

	invoices, errCreate := invoice.Create(Invoice(), Utils.ExampleProject)
	if errCreate.Errors != nil {
		for _, e := range errCreate.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	payments := []brcodepayment.BrcodePayment{
		{
			Brcode:      invoices[0].Brcode,
			TaxId:       invoices[0].TaxId,
			Description: "this will be fast",
		},
	}
	return payments
}

func Darf() []darfpayment.DarfPayment {

	competence := time.Date(2022, 10, 28, 0, 0, 0, 0, time.UTC)
	due := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	payments := []darfpayment.DarfPayment{
		{
			Description:    "92886864d3322df8b76a14ca4f0903c2c4eae69f6dc180501a440ee8d7febeff",
			RevenueCode:    "4333",
			TaxId:          "16.281.034/0001-31",
			Competence:     &competence,
			NominalAmount:  961,
			FineAmount:     90,
			InterestAmount: 18,
			Due:            &due,
		},
	}
	return payments
}

func Invoice() []invoice.Invoice {

	invoices := []invoice.Invoice{
		{
			Amount: 996699999,
			Name:   "Tony Stark",
			TaxId:  "38.446.231/0001-04",
			Rules: []Rule.Rule{
				{
					"allowedTaxIds",
					[]string{"45.059.493/0001-73"},
				},
			},
		},
	}
	return invoices
}

func DynamicBrcode() []dynamicbrcode.DynamicBrcode {

	brcodes := []dynamicbrcode.DynamicBrcode{
		{
			Amount:     rand.Intn(400000),
			Expiration: rand.Intn(3600),
			Tags:       []string{"SDK-Golang-Test"},
		},
	}
	return brcodes
}

func PaymentPreviewBoleto() []paymentpreview.PaymentPreview {

	previews := []paymentpreview.PaymentPreview{
		{
			Id: "34197923400000040001091042751897307144464000",
		},
	}
	return previews
}

func PaymentPreviewBrcode() []paymentpreview.PaymentPreview {

	var invoiceList []invoice.Invoice
	var params = map[string]interface{}{}
	params["limit"] = rand.Intn(100)

	invoices := invoice.Query(params, Utils.ExampleProject)
	for invoice := range invoices {
		invoiceList = append(invoiceList, invoice)
	}

	previews := []paymentpreview.PaymentPreview{
		{
			Id:        invoiceList[rand.Intn(len(invoiceList))].Brcode,
			Scheduled: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC),
		},
	}
	return previews
}

func PaymentPreviewTaxPreview() []paymentpreview.PaymentPreview {

	previews := []paymentpreview.PaymentPreview{
		{
			Id: "81660000005003657010074119002551100010601813",
		},
	}
	return previews
}

func PaymentPreviewUtility() []paymentpreview.PaymentPreview {

	utilities := []paymentpreview.PaymentPreview{
		{
			Id: "83660000001984801380074119002551100010601813",
		},
	}
	return utilities
}

func PaymentRequest() []paymentrequest.PaymentRequest {

	requests := []paymentrequest.PaymentRequest{
		{
			CenterId: "5763106043068416",
			Payment:  BoletoPayment(),
			Type:     "boleto-payment",
		},
	}
	return requests
}

func TaxPayment() []taxpayment.TaxPayment {

	requests := []taxpayment.TaxPayment{
		{

			BarCode:     "83660000001084301380074119002551100010601813",
			Description: "just a random description",
		},
	}
	return requests
}

func Transaction() []transaction.Transaction {

	transactions := []transaction.Transaction{
		{
			Amount:      10000,
			ReceiverId:  "5768064935133184",
			Description: "Paying my debts",
			ExternalId:  fmt.Sprintf("external_id%v,%v", time.Now().Day(), time.Now().Nanosecond()),
		},
	}
	return transactions
}

func Transfer() []transfer.Transfer {

	transfers := []transfer.Transfer{
		{
			Amount:        10000,
			Name:          "Steve Rogers",
			TaxId:         "330.731.970-10",
			BankCode:      "001",
			BranchCode:    "1234",
			AccountNumber: "123456-0",
			DisplayDescription: "Payment for service 1234",
			Rules:         []rule.Rule{{Key: "resendingLimit", Value: 0}},
		},
	}
	return transfers
}

func Utility() []utilitypayment.UtilityPayment {

	utilities := []utilitypayment.UtilityPayment{
		{
			Line:        "83660000001084301380074119002551100010601813",
			Description: "just a random description",
		},
	}
	return utilities
}

func Webhook() webhook.Webhook {

	webhookExample := webhook.Webhook{
		Url:           fmt.Sprintf("https://webhook.site/%v", rand.Intn(20-11)),
		Subscriptions: []string{"boleto"},
	}
	return webhookExample
}

func Workspace() workspace.Workspace {

	workspaceExample := workspace.Workspace{
		Username: "testGolang2",
		Name:     "TesteGolang",
	}
	return workspaceExample
}
