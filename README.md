# Stark Bank Golang SDK

Welcome to the Stark Bank Golang SDK! This tool is made for Golang
developers who want to easily integrate with our API.
This SDK version is compatible with the Stark Bank API v2.

If you have no idea what Stark Bank is, check out our [website](https://www.starkbank.com/)
and discover a world where receiving or making payments
is as easy as sending a text message to your client!

# Introduction

## Index

- [Introduction](#introduction)
    - [Supported Golang versions](#supported-golang-versions)
    - [API documentation](#stark-bank-api-documentation)
    - [Versioning](#versioning)
- [Setup](#setup)
    - [Install our SDK](#1-install-our-sdk)
    - [Create your Private and Public Keys](#2-create-your-private-and-public-keys)
    - [Register your user credentials](#3-register-your-user-credentials)
    - [Setting up the user](#4-setting-up-the-user)
    - [Setting up the error language](#5-setting-up-the-error-language)
- [Resource listing and manual pagination](#resource-listing-and-manual-pagination)
- [Testing in Sandbox](#testing-in-sandbox)
- [Usage](#usage)
    - [Transactions](#create-transactions): Account statement entries
    - [Balance](#get-balance): Account balance
    - [Transfers](#create-transfers): Wire transfers (TED and manual Pix)
    - [DictKeys](#get-dict-key): Pix Key queries to use with Transfers
    - [Institutions](#query-bacen-institutions): Institutions recognized by the Central Bank
    - [Invoices](#create-invoices): Reconciled receivables (dynamic Pix QR Codes)
    - [DynamicBrcode](#create-dynamicbrcodes): Simplified reconciled receivables (dynamic Pix QR Codes)
    - [Deposits](#query-deposits): Other cash-ins (static Pix QR Codes, manual Pix, etc)
    - [Boletos](#create-boletos): Boleto receivables
    - [BoletoHolmes](#investigate-a-boleto): Boleto receivables investigator
    - [BrcodePayments](#pay-a-br-code): Pay Pix QR Codes
    - [BoletoPayments](#pay-a-boleto): Pay Boletos
    - [UtilityPayments](#create-utility-payments): Pay Utility bills (water, light, etc.)
    - [TaxPayments](#create-tax-payment): Pay taxes
    - [DarfPayments](#create-darf-payment): Pay DARFs
    - [PaymentPreviews](#preview-payment-information-before-executing-the-payment): Preview all sorts of payments
    - [PaymentRequest](#create-payment-requests-to-be-approved-by-authorized-people-in-a-cost-center): Request a payment
      approval to a cost center
    - [CorporateHolders](#create-corporateholders): Manage cardholders
    - [CorporateCards](#create-corporatecards): Create virtual and/or physical cards
    - [CorporateInvoices](#create-corporateinvoices): Add money to your corporate balance
    - [CorporateWithdrawals](#create-corporatewithdrawals): Send money back to your Workspace from your corporate balance
    - [CorporateBalance](#get-your-corporatebalance): View your corporate balance
    - [CorporateTransactions](#query-corporatetransactions): View the transactions that have affected your corporate balance
    - [CorporateEnums](#corporate-enums): Query enums related to the corporate purchases, such as merchant categories, countries and card purchase methods
    - [MerchantCard](#query-merchantcards): Stores information about approved purchase cards for reuse.
    - [MerchantSession](#create-a-merchantsession): Manages a session to create a purchase with a new card.
    - [MerchantPurchase](#create-a-merchantpurchase): Allows a merchant to charge their customers using debit or credit cards
    - [MerchantInstallment](#query-merchantinstallments): Tracks the lifecycle of purchase installments
    - [Webhooks](#create-a-webhook-subscription): Configure your webhook endpoints and subscriptions
    - [WebhookEvents](#process-webhook-events): Manage webhook events
    - [WebhookEventAttempts](#query-failed-webhook-event-delivery-attempts-information): Query failed webhook event
      deliveries
    - [Workspaces](#create-a-new-workspace): Manage your accounts
- [Handling errors](#handling-errors)
- [Help and Feedback](#help-and-feedback)

## Supported Golang Versions

This library supports the following Golang versions:

* Golang 1.17 or later

## Stark Bank API documentation

Feel free to take a look at our [API docs](https://www.starkbank.com/docs/api).

## Versioning

This project adheres to the following versioning pattern:

Given a version number MAJOR.MINOR.PATCH, increment:

- MAJOR version when the **API** version is incremented. This may include backwards incompatible changes;
- MINOR version when **breaking changes** are introduced OR **new functionalities** are added in a backwards compatible
  manner;
- PATCH version when backwards compatible bug **fixes** are implemented.

# Setup

## 1. Install our SDK

1.1 In go.mod file, add the path in the required packages

```golang
github.com/starkbank/sdk-go v1.2.0

```

1.2 You can also explicitly go get the package into a project:

```sh
go get -u github.com/starkbank/sdk-go
```

## 2. Create your Private and Public Keys

We use ECDSA. That means you need to generate a secp256k1 private
key to sign your requests to our API, and register your public key
with us, so we can validate those requests.

You can use one of following methods:

2.1. Check out the options in our [tutorial](https://starkbank.com/faq/how-to-create-ecdsa-keys).

2.2. Use our SDK:

```golang
package main

import (
  "github.com/starkinfra/core-go/starkcore/key"
)

func main() {

  privateKey, publicKey := key.Create("")

  // or, to also save .pem files in a specific path
  privateKey, publicKey := key.Create("files/keys/")
}

```

**NOTE**: When you are creating new credentials, it is recommended that you create the
keys inside the infrastructure that will use it, in order to avoid risky internet
transmissions of your **private-key**. Then you can export the **public-key** alone to the
computer where it will be used in the new Project creation.

## 3. Register your user credentials

You can interact directly with our API using two types of users: Projects and Organizations.

- **Projects** are workspace-specific users, that is, they are bound to the workspaces they are created in.
  One workspace can have multiple Projects.
- **Organizations** are general users that control your entire organization.
  They can control all your Workspaces and even create new ones. The Organization is bound to your company's tax ID
  only.
  Since this user is unique in your entire organization, only one credential can be linked to it.

3.1. To create a Project in Sandbox:

3.1.1. Log into [Starkbank Sandbox](https://web.sandbox.starkbank.com)

3.1.2. Go to Menu > Integrations

3.1.3. Click on the "New Project" button

3.1.4. Create a Project: Give it a name and upload the public key you created in section 2

3.1.5. After creating the Project, get its Project ID

3.1.6. Use the Project ID and private key to create the struct below:

```golang
package main

import (
  "github.com/starkinfra/core-go/starkcore/user/project"
  "github.com/starkinfra/core-go/starkcore/utils/checks"
)

// Get your private key from an environment variable or an encrypted database.
// This is only an example of a private key content. You should use your own key.

var privateKeyContent = "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEILChZrjrrtFnyCLhcxm/hp+9ljWSmG7Wv9HRugf+FnhkoAcGBSuBBAAK\noUQDQgAEpIAM/tMqXEfLeR93rRHiFcpDB9I18MrnCJyTVk0MdD1J9wgEbRfvAZEL\nYcEGhTFYp2X3B7K7c4gDDCr0Pu1L3A==\n-----END EC PRIVATE KEY-----\n"

var project = project.Project{
  Id:          "5656565656565656",
  PrivateKey:  checks.CheckPrivateKey(privateKeyContent),
  Environment: checks.CheckEnvironment("sandbox"),
}

```

3.2. To create Organization credentials in Sandbox:

3.2.1. Log into [Starkbank Sandbox](https://web.sandbox.starkbank.com)

3.2.2. Go to Menu > Integrations

3.2.3. Click on the "Organization public key" button

3.2.4. Upload the public key you created in section 2 (only a legal representative of the organization can upload the
public key)

3.2.5. Click on your profile picture and then on the "Organization" menu to get the Organization ID

3.2.6. Use the Organization ID and private key to create the struct below:

```golang
package main

import (
  "fmt"
  Balance "github.com/starkbank/sdk-go/starkbank/balance"
  "github.com/starkinfra/core-go/starkcore/user/organization"
  "github.com/starkinfra/core-go/starkcore/utils/checks"
)

// Get your private key from an environment variable or an encrypted database.
// This is only an example of a private key content. You should use your own key.

var privateKeyContent = "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEILChZrjrrtFnyCLhcxm/hp+9ljWSmG7Wv9HRugf+FnhkoAcGBSuBBAAK\noUQDQgAEpIAM/tMqXEfLeR93rRHiFcpDB9I18MrnCJyTVk0MdD1J9wgEbRfvAZEL\nYcEGhTFYp2X3B7K7c4gDDCr0Pu1L3A==\n-----END EC PRIVATE KEY-----\n"

var organization = organization.Organization{
  Id:          "5656565656565656",
  PrivateKey:  checks.CheckPrivateKey(privateKeyContent),
  Environment: checks.CheckEnvironment("sandbox"),
}

// To dynamically use your organization credentials in a specific workspaceId,
// you can use the organization.Replace() function:

func main() {
	
  balance := Balance.Get(organization.Replace("4848484848484848"))
  fmt.Println(balance)
}

```

NOTE 1: Never hard-code your private key. Get it from an environment variable or an encrypted database.

NOTE 2: We support `'sandbox'` and `'production'` as environments.

NOTE 3: The credentials you registered in `sandbox` do not exist in `production` and vice versa.

## 4. Setting up the user

There are three kinds of users that can access our API: **Organization**, **Project** and **Member**.

- `Project` and `Organization` are designed for integrations and are the ones meant for our SDKs.
- `Member` is the one you use when you log into our webpage with your e-mail.

There are two ways to inform the user to the SDK:

4.1 Passing the user as argument in all functions:

```golang
package main

import (
  "fmt"
  Balance "github.com/starkbank/sdk-go/starkbank/balance"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  balance := Balance.Get(utils.ExampleProject) // or organization
  fmt.Println(balance)
}

```

4.2 Set it as a default user in the SDK:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Balance "github.com/starkbank/sdk-go/starkbank/balance"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  balance := Balance.Get(nil) // or organization
  fmt.Println(balance)
}

```

## 5. Setting up the error language

The error language can also be set in the same way as the default user:

```golang
package main

import (
	"github.com/starkbank/sdk-go/starkbank"
)

func main() {
	starkbank.Language = "pt-BR"
}

```

Language options are "en-US" for english and "pt-BR" for brazilian portuguese. English is default.

# Resource listing and manual pagination

Almost all SDK resources provide a `query` and a `page` function.

- The `query` function provides a straight forward way, through a `channel`, to efficiently iterate through all results
  that match the filters you inform, seamlessly retrieving the next batch of elements from the API only when you reach
  the end of the current batch.
  If you are not worried about data volume or processing time, this is the way to go.

- In this function, in particular, we return a second `channel` for error handling. This error channel will receive any
  errors that occur during the query process, including API errors, network issues, or data parsing problems. You should
  monitor this channel alongside the main data channel to handle errors appropriately.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transaction "github.com/starkbank/sdk-go/starkbank/transaction"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 200

  transactions, errorChannel := Transaction.Query(params, nil)
  loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					fmt.Printf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transaction, ok := <-transactions:
			if !ok {
				break loop
			}
			fmt.Println(transaction)
		}
	}
}

```

- The `page` function gives you full control over the API pagination. With each function call, you receive up to
  100 results and the cursor to retrieve the next batch of elements. This allows you to stop your queries and
  pick up from where you left off whenever it is convenient. When there are no more elements to be retrieved, the
  returned cursor will be `nil`.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transaction "github.com/starkbank/sdk-go/starkbank/transaction"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 50

  for true {
    transactions, cursor, err := Transaction.Page(params, nil)
    if err.Errors != nil {
      for _, e := range err.Errors {
        fmt.Printf("code: %s, message: %s", e.Code, e.Message)
      }
    }
	
    for _, transaction := range transactions {
      fmt.Println(transaction)
    }
	
    if cursor == "" {
      break
    }
  }
}

```

To simplify the following SDK examples, we will only use the `query` function, but feel free to use `page` instead.

# Testing in Sandbox

Your initial balance is zero. For many operations in Stark Bank, you'll need funds
in your account, which can be added to your balance by creating an Invoice or a Boleto.

In the Sandbox environment, most of the created Invoices and Boletos will be automatically paid,
so there's nothing else you need to do to add funds to your account. Just create
a few Invoices and wait around a bit.

In Production, you (or one of your clients) will need to actually pay this Invoice or Boleto
for the value to be credited to your account.

# Usage

Here are a few examples on how to use the SDK. If you have any doubts, check out the function or class docstring to get
more info or go straight to our [API docs].

## Create transactions

To send money between Stark Bank accounts, you can create transactions:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transaction "github.com/starkbank/sdk-go/starkbank/transaction"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  transactions, err := Transaction.Create(
    []Transaction.Transaction{
      {
        Amount:      10000,
        ReceiverId:  "5768064935133184",
        Description: "Paying my debts",
        ExternalId:  "my_external_id",
      },
    }, utils.Project)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, transaction := range transactions {
    fmt.Println(transaction)
  }
}

```

**Note**: Instead of using Transaction structs, you can also pass each transaction element in map format

## Query transactions

To understand your balance changes (bank statement), you can query
transactions. Note that our system creates transactions for you when
you receive boleto payments, pay a bill or make transfers, for example.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transaction "github.com/starkbank/sdk-go/starkbank/transaction"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 200

  transactions, errorChannel := Transaction.Query(params, nil)
  loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					fmt.Printf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transaction, ok := <-transactions:
			if !ok {
				break loop
			}
			fmt.Println(transaction)
		}
	}
}

```

## Get a transaction

You can get a specific transaction by its id:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transaction "github.com/starkbank/sdk-go/starkbank/transaction"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  transaction, err := Transaction.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(transaction)
}

```

## Get balance

To know how much money you have in your workspace, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Balance "github.com/starkbank/sdk-go/starkbank/balance"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  balance := Balance.Get(nil)
  fmt.Println(balance)
}

```

## Create transfers

You can also create transfers in the SDK (TED/Pix).

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transfer "github.com/starkbank/sdk-go/starkbank/transfer"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  transfers, err := Transfer.Create(
    []Transfer.Transfer{
      {
        Amount:        100,
        Name:          "Tony Stark",
        TaxId:         "012.345.678-90",
        BankCode:      "033", // TED
        BranchCode:    "0001",
        AccountNumber: "10000-0",
      },
      {
        Amount:        200,
        Name:          "Jon Snow",
        TaxId:         "012.345.678-90",
        BankCode:      "20018183", //Pix
        BranchCode:    "1234",
        AccountNumber: "123456-7",
        AccountType:   "salary",
        ExternalId:    "my-internal-id-12345",
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, transfer := range transfers {
    fmt.Println(transfer)
  }
}

```

**Note**: Instead of using Transfer structs, you can also pass each transfer element in map format

## Query transfers

You can query multiple transfers according to filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transfer "github.com/starkbank/sdk-go/starkbank/transfer"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["after"] = "2020-01-01"
  params["before"] = "2020-04-01"

  transfers, errorChannel := Transfer.Query(params, nil)
  loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					fmt.Printf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transfer, ok := <-transfers:
			if !ok {
				break loop
			}
			fmt.Println(transfer)
		}
	}
}

```

## Cancel a scheduled transfer

To cancel a single scheduled transfer by its id, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transfer "github.com/starkbank/sdk-go/starkbank/transfer"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  transfer, err := Transfer.Delete("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(transfer)
}

```

## Get a transfer

To get a single transfer by its id, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transfer "github.com/starkbank/sdk-go/starkbank/transfer"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  transfer, err := Transfer.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(transfer)
}

```

## Get a transfer PDF

A transfer PDF may also be retrieved by its id.
This operation is only valid if the transfer status is "processing" or "success".

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transfer "github.com/starkbank/sdk-go/starkbank/transfer"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject

  pdf, err := Transfer.Pdf("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.pdf", "transfer", "5155165527080960")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw pdf content,
as it may yield abnormal results in the final file, such as missing images
and strange characters.

## Query transfer logs

You can query transfer logs to better understand transfer life cycles.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/transfer/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 50

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a transfer log

You can also get a specific log by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/transfer/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(log)
}

```

## Get DICT key

You can get the Pix key's parameters by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  DictKey "github.com/starkbank/sdk-go/starkbank/dictkey"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  key, err := DictKey.Get("tony@starkbank.com", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(key)
}

```

## Query your DICT keys

To take a look at the Pix keys linked to your workspace, just run the following:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  DictKey "github.com/starkbank/sdk-go/starkbank/dictkey"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["status"] = "registered"

  keys, errorChannel := DictKey.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case key, ok := <-keys:
      if !ok {
        break loop
      }
      fmt.Println(key)
    }
  }
}

```

## Query Bacen institutions

You can query institutions registered by the Brazilian Central Bank for Pix and TED transactions.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Institution "github.com/starkbank/sdk-go/starkbank/institution"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["search"] = "stark"

  institutions, errorChannel := Institution.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case institution, ok := <-institutions:
      if !ok {
        break loop
      }
      fmt.Println(institution)
    }
  }
}

```

## Create invoices

You can create dynamic QR Code invoices to charge customers or to receive money from accounts you have in other banks.

Since the banking system only understands value modifiers (discounts, fines and interest) when dealing with **dates** (
instead of **datetimes**), these values will only show up in the end user banking interface if you use **dates** in
the "due" and "discounts" fields.

If you use **datetimes** instead, our system will apply the value modifiers in the same manner, but the end user will
only see the final value to be paid on his interface.

Also, other banks will most likely only allow payment scheduling on invoices defined with **dates** instead of *
*datetimes**.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
  "github.com/starkbank/sdk-go/tests/utils"
  "time"
)

func main() {

  starkbank.User = utils.ExampleProject

  due := time.Now().Add(time.Hour * 1)
  due2 := time.Date(2022, 03, 20, 0, 0, 0, 0, time.UTC)

  invoices, err := Invoice.Create(
    []Invoice.Invoice{
      {
        Amount:   23571, // R$ 235,71
        Name:     "Buzz Aldrin",
        TaxId:    "012.345.678-90",
        Due:      &due,
        Fine:     5,   // 5%
        Interest: 2.5, // 2.5% per month
        Tags:     []string{"imediate"},
      },
      {
        Amount:   923571, // R$ 235,71
        Name:     "Buzz Aldrin",
        TaxId:    "012.345.678-90",
        Due:      &due2,
        Fine:     5,   // 5%
        Interest: 2.5, // 2.5% per month
        Tags:     []string{"scheduled"},
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, invoice := range invoices {
    fmt.Println(invoice)
  }
}

```

**Note**: Instead of using Invoice structs, you can also pass each invoice element in map format

## Get an invoice

After its creation, information on an invoice may be retrieved by its id.
Its status indicates whether it's been paid.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  invoice, err := Invoice.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(invoice)
}

```

## Get an invoice PDF

After its creation, an invoice PDF may be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject

  pdf, err := Invoice.Pdf("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.pdf", "invoice", "5155165527080960")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw pdf content,
as it may yield abnormal results in the final file, such as missing images
and strange characters.

## Get an invoice QR Code

After its creation, an Invoice QR Code may be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["size"] = 10

  qrcode, err := Invoice.Qrcode("5155165527080960", params, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.png", "invoice", "5155165527080960")
  errFile := ioutil.WriteFile(filename, qrcode, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw png content,
as it may corrupt the file.

## Cancel an invoice

You can also cancel an invoice by its id.
Note that this is not possible if it has been paid already.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var patchData = map[string]interface{}{}
  patchData["status"] = "canceled"

  invoice, err := Invoice.Update("5155165527080960", patchData, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(invoice)
}

```

## Update an invoice

You can update an invoice's amount, due date and expiration by its id.
If the invoice has already been paid, only the amount can be
decreased, which will result in a payment reversal. To fully reverse
the invoice, pass amount = 0.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
  "github.com/starkbank/sdk-go/tests/utils"
  "time"
)

func main() {

  starkbank.User = utils.ExampleProject

  var patchData = map[string]interface{}{}
  patchData["amount"] = 100
  patchData["expiration"] = 0
  patchData["due"] = time.Now().Add(time.Hour * 1)

  invoice, err := Invoice.Update("5155165527080960", patchData, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(invoice)
}

```

## Query invoices

You can get a list of created invoices given some filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["after"] = "2020-01-01"
  params["before"] = "2020-03-01"

  invoices, errorChannel := Invoice.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case invoice, ok := <-invoices:
      if !ok {
        break loop
      }
      fmt.Println(invoice)
    }
  }
}

```

## Query invoice logs

Logs are pretty important to understand the life cycle of an invoice.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/invoice/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 150

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get an invoice log

You can get a single log by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/invoice/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(log)
}

```

## Get a reversed invoice log PDF

Whenever an Invoice is successfully reversed, a reversed log will be created.
To retrieve a specific reversal receipt, you can request the corresponding log PDF:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/invoice/log"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject
  
  pdf, err := Log.Pdf("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.pdf", "invoice-log", "5155165527080960")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw pdf content,
as it may yield abnormal results in the final file, such as missing images
and strange characters.

## Get an invoice payment information

Once an invoice has been paid, you can get the payment information using the Invoice.Payment sub-resource:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Invoice "github.com/starkbank/sdk-go/starkbank/invoice"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := Invoice.GetPayment("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Create DynamicBrcodes

You can create simplified dynamic QR Codes to receive money using Pix transactions.
When a DynamicBrcode is paid, a Deposit is created with the tags parameter containing the character “dynamic-brcode/”
followed by the DynamicBrcode’s uuid "dynamic-brcode/{uuid}" for conciliation.

The differences between an Invoice and the DynamicBrcode are the following:

|                       | Invoice | DynamicBrcode |
|-----------------------|:-------:|:-------------:|
| Expiration            |    ✓    |       ✓       | 
| Due, fine and fee     |    ✓    |       X       | 
| Discount              |    ✓    |       X       | 
| Description           |    ✓    |       X       |
| Can be updated        |    ✓    |       X       |
| Can only be paid once |    ✓    |       ✓       |

**Note:** In order to check if a BR code has expired, you must first calculate its expiration date (add the expiration
to the creation date).
**Note:** To know if the BR code has been paid, you need to query your Deposits by the tag "dynamic-brcode/{uuid}" to
check if it has been paid.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/dynamicbrcode"
  "github.com/starkbank/sdk-go/tests/utils"
  "math/rand"
)

func main() {

  starkbank.User = utils.ExampleProject

  brcodes, err := dynamicbrcode.Create(
    []dynamicbrcode.DynamicBrcode{
      {
        Amount:     23571,
        Expiration: rand.Intn(3600),
      }, {
        Amount:     23571, // R$ 235,71
        Expiration: rand.Intn(3600),
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for brcode := range brcodes {
    fmt.Println(brcode)
  }
}

```

**Note**: Instead of using DynamicBrcode objects, you can also pass each brcode element in dictionary format

## Get a DynamicBrcode

After its creation, information on a DynamicBrcode may be retrieved by its uuid.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/dynamicbrcode"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  brcode, err := dynamicbrcode.Get("bb9cd43ea6f4403391bf7ef6aa876600", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(brcode)
}

```

## Query DynamicBrcodes

You can get a list of created DynamicBrcodes given some filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/dynamicbrcode"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 4

  brcodes, errorChannel := dynamicbrcode.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case brcode, ok := <-brcodes:
      if !ok {
        break loop
      }
      fmt.Println(brcode)
    }
  }
}

```

## Query deposits

You can get a list of created deposits given some filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Deposit "github.com/starkbank/sdk-go/starkbank/deposit"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["after"] = "2020-01-01"
  params["before"] = "2020-03-01"

  deposits, errorChannel := Deposit.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case deposit, ok := <-deposits:
      if !ok {
        break loop
      }
      fmt.Println(deposit)
    }
  }
}

```

## Get a deposit

After its creation, information on a deposit may be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Deposit "github.com/starkbank/sdk-go/starkbank/deposit"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  deposit, err := Deposit.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(deposit)
}

```

## Query deposit logs

Logs are pretty important to understand the life cycle of a deposit.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/deposit/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 150

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a deposit log

You can get a single log by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/deposit/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(log)
}

```

## Create boletos

You can create boletos to charge customers or to receive money from accounts
you have in other banks.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Boleto "github.com/starkbank/sdk-go/starkbank/boleto"
  "github.com/starkbank/sdk-go/tests/utils"
  "time"
)

func main() {

  starkbank.User = utils.ExampleProject

  due := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

  boletos, err := Boleto.Create(
    []Boleto.Boleto{
      {
        Amount:      23571, //R$235,71
        Name:        "Buzz Aldrin",
        TaxId:       "012.345.678-90",
        StreetLine1: "Av. Paulista, 200",
        StreetLine2: "10 Andar",
        District:    "Bela Vista",
        City:        "São Paulo",
        StateCode:   "SP",
        ZipCode:     "01420-020",
        Due:         &due,
        Fine:        5,   //5%
        Interest:    2.5, //2.5% per month
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, boleto := range boletos {
    fmt.Println(boleto)
  }
}

```

**Note**: Instead of using Boleto structs, you can also pass each boleto element in map format

## Get a boleto

After its creation, information on a boleto may be retrieved by its id.
Its status indicates whether it's been paid.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Boleto "github.com/starkbank/sdk-go/starkbank/boleto"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  boleto, err := Boleto.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(boleto)
}

```

## Get a boleto PDF

After its creation, a boleto PDF may be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Boleto "github.com/starkbank/sdk-go/starkbank/boleto"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["layout"] = "booklet"

  pdf, err := Boleto.Pdf("5155165527080960", params, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.pdf", "boleto", "5155165527080960")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw pdf content,
as it may yield abnormal results in the final file, such as missing images
and strange characters.

## Delete a boleto

You can also cancel a boleto by its id.
Note that this is not possible if it has been processed already.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Boleto "github.com/starkbank/sdk-go/starkbank/boleto"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  boleto, err := Boleto.Delete("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(boleto)
}

```

## Query boletos

You can get a list of created boletos given some filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Boleto "github.com/starkbank/sdk-go/starkbank/boleto"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["after"] = "2020-01-01"
  params["before"] = "2020-03-01"

  boletos, errorChannel := Boleto.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case boleto, ok := <-boletos:
      if !ok {
        break loop
      }
      fmt.Println(boleto)
    }
  }
}

```

## Query boleto logs

Logs are pretty important to understand the life cycle of a boleto.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/boleto/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 150

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a boleto log

You can get a single log by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/boleto/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(log)
}

```

## Investigate a boleto

You can discover if a StarkBank boleto has been recently paid before we receive the response on the next day.
This can be done by creating a BoletoHolmes struct, which fetches the updated status of the corresponding
Boleto struct according to CIP to check, for example, whether it is still payable or not. The investigation
happens asynchronously and the most common way to retrieve the results is to register a "boleto-holmes" webhook
subscription, although polling is also possible.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Holmes "github.com/starkbank/sdk-go/starkbank/boletoholmes"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  holmes, err := Holmes.Create(
    []Holmes.BoletoHolmes{
      {
        BoletoId: "5656565656565656",
      },
      {
        BoletoId: "4848484848484848",
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, holmes := range holmes {
    fmt.Println(holmes)
  }
}

```

**Note**: Instead of using BoletoHolmes structs, you can also pass each payment element in map format

## Get a boleto holmes

To get a single Holmes by its id, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Holmes "github.com/starkbank/sdk-go/starkbank/boletoholmes"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  holmes, err := Holmes.Get("19278361897236187236", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(holmes)
}

```

## Query boleto holmes

You can search for boleto Holmes using filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Holmes "github.com/starkbank/sdk-go/starkbank/boletoholmes"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["holmesIds"] = []string{"customer_1", "customer_2"}

  holmes, errorChannel := Holmes.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case holmes, ok := <-holmes:
      if !ok {
        break loop
      }
      fmt.Println(holmes)
    }
  }
}

```

## Query boleto holmes logs

Searches are also possible with boleto holmes logs:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/boletoholmes/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["holmesIds"] = []string{"5155165527080960", "76551659167801921"}

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a boleto holmes log

You can also get a boleto holmes log by specifying its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/boletoholmes/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("19278361897236187236", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(log)
}

```

## Pay a BR Code

Paying a BR Code is also simple. After extracting the BRCode encoded in the Pix QR Code, you can do the following:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BrcodePayment "github.com/starkbank/sdk-go/starkbank/brcodepayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payments, err := BrcodePayment.Create(
    []BrcodePayment.BrcodePayment{
      {
        Brcode:      "00020101021226930014br.gov.bcb.pix2571brcode-h.sandbox.starkinfra.com/v2/09a7970542fe4399ab2af079982bb1005204000053039865802BR5925Stark Bank S.A. - Institu6009Sao Paulo62070503***63044AC2",
        TaxId:       "38.446.231/0001-04",
        Description: "this will be fast",
        Tags:        []string{"pix", "qrcode"},
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, payment := range payments {
    fmt.Println(payment)
  }
}

```

**Note**: Instead of using BrcodePayment structs, you can also pass each payment element in map format

## Get a BR Code payment

To get a single BR Code payment by its id, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BrcodePayment "github.com/starkbank/sdk-go/starkbank/brcodepayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := BrcodePayment.Get("19278361897236187236", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Get a BR Code payment PDF

After its creation, a BR Code payment PDF may be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BrcodePayment "github.com/starkbank/sdk-go/starkbank/brcodepayment"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject

  pdf, err := BrcodePayment.Pdf("19278361897236187236", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.pdf", "brcode-payment", "19278361897236187236")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw pdf content,
as it may yield abnormal results in the final file, such as missing images
and strange characters.

## Cancel a BR Code payment

You can cancel a BR Code payment by changing its status to "canceled".
Note that this is not possible if it has been processed already.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BrcodePayment "github.com/starkbank/sdk-go/starkbank/brcodepayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var patchData = map[string]interface{}{}
  patchData["status"] = "canceled"

  payment, err := BrcodePayment.Update("19278361897236187236", patchData, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Query BR Code payments

You can search for brcode payments using filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BrcodePayment "github.com/starkbank/sdk-go/starkbank/brcodepayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["tags"] = []string{"company_1", "company_2"}

  payments, errorChannel := BrcodePayment.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case payment, ok := <-payments:
      if !ok {
        break loop
      }
      fmt.Println(payment)
    }
  }
}

```

## Query BR Code payment logs

Searches are also possible with BR Code payment logs:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/brcodepayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["paymentIds"] = []string{"5155165527080960", "76551659167801921"}

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a BR Code payment log

You can also get a BR Code payment log by specifying its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/brcodepayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := Log.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Pay a boleto

Paying a boleto is also simple.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BoletoPayment "github.com/starkbank/sdk-go/starkbank/boletopayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payments, err := BoletoPayment.Create(
    []BoletoPayment.BoletoPayment{
      {
        TaxId:       "20.018.183/0001-80",
        Description: "take my money",
        Line:        "34191.09008 76038.597308 71444.640008 4 92150000028000",
      },
      {
        TaxId:       "20.018.183/0001-80",
        Description: "take my money one more time",
        BarCode:     "34197819200000000011090063609567307144464000",
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, payment := range payments {
    fmt.Println(payment)
  }
}

```

**Note**: Instead of using BoletoPayment structs, you can also pass each payment element in map format

## Get a boleto payment

To get a single boleto payment by its id, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BoletoPayment "github.com/starkbank/sdk-go/starkbank/boletopayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := BoletoPayment.Get("19278361897236187236", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Get a boleto payment PDF

After its creation, a boleto payment PDF may be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BoletoPayment "github.com/starkbank/sdk-go/starkbank/boletopayment"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject

  pdf, err := BoletoPayment.Pdf("19278361897236187236", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.pdf", "boleto-payment", "19278361897236187236")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw pdf content,
as it may yield abnormal results in the final file, such as missing images
and strange characters.

## Delete a boleto payment

You can also cancel a boleto payment by its id.
Note that this is not possible if it has been processed already.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BoletoPayment "github.com/starkbank/sdk-go/starkbank/boletopayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := BoletoPayment.Delete("19278361897236187236", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Query boleto payments

You can search for boleto payments using filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  BoletoPayment "github.com/starkbank/sdk-go/starkbank/boletopayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["tags"] = []string{"company_1", "company_2"}

  payments, errorChannel := BoletoPayment.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case payment, ok := <-payments:
      if !ok {
        break loop
      }
      fmt.Println(payment)
    }
  }
}

```

## Query boleto payment logs

Searches are also possible with boleto payment logs:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/boletopayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["paymentIds"] = []string{"5155165527080960", "76551659167801921"}

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a boleto payment log

You can also get a boleto payment log by specifying its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/boletopayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(log)
}

```

## Create utility payments

It's also simple to pay utility bills (such as electricity and water bills) in the SDK.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  UtilityPayment "github.com/starkbank/sdk-go/starkbank/utilitypayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payments, err := UtilityPayment.Create(
    []UtilityPayment.UtilityPayment{
      {
        Line:        "3419109107 41224987309 71444640008 9 91800999999999",
        Description: "take my money",
        Tags:        []string{"take", "my", "money"},
      },
      {
        BarCode:     "34194918109999999991091041242887307144464000",
        Description: "take my money one more time",
        Tags:        []string{"again"},
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, payment := range payments {
    fmt.Println(payment)
  }
}

```

**Note**: Instead of using UtilityPayment structs, you can also pass each payment element in map format

## Query utility payments

To search for utility payments using filters, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  UtilityPayment "github.com/starkbank/sdk-go/starkbank/utilitypayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["tags"] = []string{"eletricity", "gas"}

  payments, errorChannel := UtilityPayment.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case payment, ok := <-payments:
      if !ok {
        break loop
      }
      fmt.Println(payment)
    }
  }
}

```

## Get a utility payment

You can get a specific bill by its id:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  UtilityPayment "github.com/starkbank/sdk-go/starkbank/utilitypayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := UtilityPayment.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Get a utility payment PDF

After its creation, a utility payment PDF may also be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  UtilityPayment "github.com/starkbank/sdk-go/starkbank/utilitypayment"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject

  pdf, err := UtilityPayment.Pdf("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.pdf", "utility-payment", "5155165527080960")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw pdf content,
as it may yield abnormal results in the final file, such as missing images
and strange characters.

## Delete a utility payment

You can also cancel a utility payment by its id.
Note that this is not possible if it has been processed already.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  UtilityPayment "github.com/starkbank/sdk-go/starkbank/utilitypayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := UtilityPayment.Delete("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Query utility payment logs

You can search for payments by specifying filters. Use this to understand the
bills life cycles.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/utilitypayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["paymentIds"] = []string{"102893710982379182", "92837912873981273"}

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a utility payment log

If you want to get a specific payment log by its id, just run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/utilitypayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("1902837198237992", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(log)
}

```

## Create tax payment

It is also simple to pay taxes (such as ISS and DAS) using this SDK.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  TaxPayment "github.com/starkbank/sdk-go/starkbank/taxpayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payments, err := TaxPayment.Create(
    []TaxPayment.TaxPayment{
      {
        BarCode:     "83660000001084301380074119002551100010601813",
        Description: "fix the road",
        Tags:        []string{"take", "my", "money"},
      },
      {
        Line:        "85800000003 0 28960328203 1 56072020190 5 22109674804 0",
        Description: "build the hospital, hopefully",
        Tags:        []string{"expensive"},
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, payment := range payments {
    fmt.Println(payment)
  }
}

```

**Note**: Instead of using TaxPayment structs, you can also pass each payment element in map format

## Query tax payments

To search for tax payments using filters, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  TaxPayment "github.com/starkbank/sdk-go/starkbank/taxpayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["tags"] = []string{"das", "july"}

  payments, errorChannel := TaxPayment.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case payment, ok := <-payments:
      if !ok {
        break loop
      }
      fmt.Println(payment)
    }
  }
}

```

## Get tax payment

You can get a specific tax payment by its id:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  TaxPayment "github.com/starkbank/sdk-go/starkbank/taxpayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := TaxPayment.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Get tax payment PDF

After its creation, a tax payment PDF may also be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  TaxPayment "github.com/starkbank/sdk-go/starkbank/taxpayment"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject

  pdf, err := TaxPayment.Pdf("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.pdf", "tax-payment", "5155165527080960")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw pdf content,
as it may yield abnormal results in the final file, such as missing images
and strange characters.

## Delete tax payment

You can also cancel a tax payment by its id.
Note that this is not possible if it has been processed already.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  TaxPayment "github.com/starkbank/sdk-go/starkbank/taxpayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := TaxPayment.Delete("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Query tax payment logs

You can search for payment logs by specifying filters. Use this to understand each payment life cycle.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/taxpayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get tax payment log

If you want to get a specific payment log by its id, just run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/taxpayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("1902837198237992", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(log)
}

```

**Note**: Some taxes can't be paid with bar codes. Since they have specific parameters, each one of them has its own
resource and routes, which are all analogous to the TaxPayment resource. The ones we currently support are:

- DarfPayment, for DARFs

## Create DARF payment

If you want to manually pay DARFs without barcodes, you may create DarfPayments:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  DarfPayment "github.com/starkbank/sdk-go/starkbank/darfpayment"
  "github.com/starkbank/sdk-go/tests/utils"
  "time"
)

func main() {

  starkbank.User = utils.ExampleProject

  competence := time.Date(2022, 10, 28, 0, 0, 0, 0, time.UTC)
  due := time.Now().Add(time.Hour * 24 * 30)
  scheduled := time.Now().Add(time.Hour * 24 * 30)

  payments, err := DarfPayment.Create(
    []DarfPayment.DarfPayment{
      {
        Description:     "take my money",
        RevenueCode:     "1240",
        TaxId:           "012.345.678-90",
        ReferenceNumber: "2340978970",
        Competence:      &competence,
        NominalAmount:   1234,
        FineAmount:      12,
        InterestAmount:  34,
        Due:             &due,
        Scheduled:       &scheduled,
        Tags:            []string{"DARF", "making money"},
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, payment := range payments {
    fmt.Println(payment)
  }
}

```

**Note**: Instead of using DarfPayment structs, you can also pass each payment element in map format

## Query DARF payments

To search for DARF payments using filters, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  DarfPayment "github.com/starkbank/sdk-go/starkbank/darfpayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["tags"] = []string{"darf", "july"}

  payments, errorChannel := DarfPayment.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case payment, ok := <-payments:
      if !ok {
        break loop
      }
      fmt.Println(payment)
    }
  }
}

```

## Get DARF payment

You can get a specific DARF payment by its id:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  DarfPayment "github.com/starkbank/sdk-go/starkbank/darfpayment"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  payment, err := DarfPayment.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(payment)
}

```

## Get DARF payment PDF

After its creation, a DARF payment PDF may also be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  DarfPayment "github.com/starkbank/sdk-go/starkbank/darfpayment"
  "github.com/starkbank/sdk-go/tests/utils"
  "io/ioutil"
)

func main() {

  starkbank.User = utils.ExampleProject

  pdf, err := DarfPayment.Pdf("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  filename := fmt.Sprintf("%v%v.pdf", "darf-paymet", "5155165527080960")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}

```

Be careful not to accidentally enforce any encoding on the raw pdf content,
as it may yield abnormal results in the final file, such as missing images
and strange characters.

## Delete DARF payment

You can also cancel a DARF payment by its id.
Note that this is not possible if it has been processed already.

```golang
package main

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	DarfPayment "github.com/starkbank/sdk-go/starkbank/darfpayment"
	"github.com/starkbank/sdk-go/tests/utils"
)

func main() {

	starkbank.User = utils.ExampleProject

	payment, err := DarfPayment.Delete("5155165527080960", nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	
	fmt.Println(payment)
}

```

## Query DARF payment logs

You can search for payment logs by specifying filters. Use this to understand each payment life cycle.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/darfpayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get DARF payment log

If you want to get a specific payment log by its id, just run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/darfpayment/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("5155165527080960", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(log)
}

```

## Preview payment information before executing the payment

You can preview multiple types of payment to confirm any information before actually paying.
If the "scheduled" parameter is not informed, today will be assumed as the intended payment date.
Right now, the "scheduled" parameter only has effect on BrcodePreviews.
This resource is able to preview the following types of payment:
"brcode-payment", "boleto-payment", "utility-payment" and "tax-payment"

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  PaymentPreview "github.com/starkbank/sdk-go/starkbank/paymentpreview"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  previews, err := PaymentPreview.Create(
    []PaymentPreview.PaymentPreview{
      {
        Id: "00020101021226930014br.gov.bcb.pix2571brcode-h.sandbox.starkinfra.com/v2/09a7970542fe4399ab2af079982bb1005204000053039865802BR5925Stark Bank S.A. - Institu6009Sao Paulo62070503***63044AC2",
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, preview := range previews {
    fmt.Println(preview)
  }
}

```

**Note**: Instead of using PaymentPreview structs, you can also pass each request element in map format

## Create payment requests to be approved by authorized people in a cost center

You can also request payments that must pass through a specific cost center approval flow to be executed.
In certain structures, this allows double checks for cash-outs and also gives time to load your account
with the required amount before the payments take place.
The approvals can be granted at our website and must be performed according to the rules
specified in the cost center.

**Note**: The value of the centerId parameter can be consulted by logging into our website and going
to the desired cost center page.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  PaymentRequest "github.com/starkbank/sdk-go/starkbank/paymentrequest"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  requests, err := PaymentRequest.Create(
    []PaymentRequest.PaymentRequest{
      {
        Amount: 12345,
      },
      {
        Amount: 67890,
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, request := range requests {
    fmt.Println(request)
  }
}

```

**Note**: Instead of using PaymentRequest structs, you can also pass each request element in map format

## Query payment requests

To search for payment requests, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  PaymentRequest "github.com/starkbank/sdk-go/starkbank/paymentrequest"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["status"] = "approved"

  requests, errorChannel := PaymentRequest.Query("123456778890", params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case request, ok := <-requests:
      if !ok {
        break loop
      }
      fmt.Println(request)
    }
  }
}

```

## Create CorporateHolders

You can create card holders to which your cards will be bound.
They support spending rules that will apply to all underlying cards.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporateholder"
  "github.com/starkbank/sdk-go/starkbank/corporaterule"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  holders, err := corporateholder.Create(
    []corporateholder.CorporateHolder{
      {
        Name:  "Iron Bank S.A.",
        Tags:  []string{"Traveler Employee"},
        Rules: []corporaterule.CorporateRule{
			{
              Name: "General USD",
              Interval: "day",
              Amount: 100000,
              CurrencyCode: "USD",
            },
        },
      },
      {
        Name: "Iron Bank S.A.",
      },
    }, nil, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  for _, holder := range holders {
    fmt.Println(holder)
  }
}

```

**Note**: Instead of using CorporateHolder objects, you can also pass each element in dictionary format

## Query CorporateHolders

You can query multiple holders according to filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporateholder"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  holders, errorChannel := corporateholder.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case holder, ok := <-holders:
      if !ok {
        break loop
      }
      fmt.Println(holder)
    }
  }
}

```

## Cancel a CorporateHolder

To cancel a single Corporate Holder by its id, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporateholder"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  holder, err := corporateholder.Cancel("5353197895942144", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(holder)
}

```

## Get a CorporateHolder

To get a single Corporate Holder by its id, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporateholder"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  holder, err := corporateholder.Get("5353197895942144", nil, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(holder)
}

```

## Query CorporateHolder logs

You can query holder logs to better understand holder life cycles.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/corporateholder/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a CorporateHolder log

You can also get a specific log by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/corporateholder/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("5353197895942144", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(log)
}
```

## Create CorporateCard

You can issue cards with specific spending rules.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatecard"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  card, err := corporatecard.Create(
    corporatecard.CorporateCard{
      HolderId: "5155165527080960",
    }, nil, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(card)
}

```

## Query CorporateCards

You can get a list of created cards given some filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatecard"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  cards, errorChannel := corporatecard.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case card, ok := <-cards:
      if !ok {
        break loop
      }
      fmt.Println(card)
    }
  }
}

```

## Get a CorporateCard

After its creation, information on a card may be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatecard"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  card, err := corporatecard.Get("5353197895942144", nil, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(card)
}

```

## Update a CorporateCard

You can update a specific card by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatecard"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var patchData = map[string]interface{}{}
  patchData["status"] = "blocked"

  card, err := corporatecard.Update("5353197895942144", patchData, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(card)
}

```

## Cancel a CorporateCard

You can also cancel a card by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatecard"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  card, err := corporatecard.Cancel("5353197895942144", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(card)
}

```

## Query CorporateCard logs

Logs are pretty important to understand the life cycle of a card.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/corporatecard/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a CorporateCard log

You can get a single log by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/corporatecard/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("5353197895942144", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(log)
}

```

## Query CorporatePurchases

You can get a list of created purchases given some filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatepurchase"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  purchases, errorChannel := corporatepurchase.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case purchase, ok := <-purchases:
      if !ok {
        break loop
      }
      fmt.Println(purchase)
    }
  }
}

```

## Get a CorporatePurchase

After its creation, information on a purchase may be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatepurchase"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  purchase, err := corporatepurchase.Get("5353197895942144", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(purchase)
}

```

## Query CorporatePurchase logs

Logs are pretty important to understand the life cycle of a purchase.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/corporatepurchase/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  logs, errorChannel := Log.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case log, ok := <-logs:
      if !ok {
        break loop
      }
      fmt.Println(log)
    }
  }
}

```

## Get a CorporatePurchase log

You can get a single log by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Log "github.com/starkbank/sdk-go/starkbank/corporatepurchase/log"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  log, err := Log.Get("5353197895942144", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(log)
}

```

## Create CorporateInvoices

You can create Pix invoices to transfer money from accounts you have in any bank to your Corporate balance,
allowing you to run your corporate operation.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporateinvoice"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  invoice, err := corporateinvoice.Create(
    corporateinvoice.CorporateInvoice{
      Amount: 10000,
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(invoice)
}

```

**Note**: Instead of using CorporateInvoice objects, you can also pass each element in dictionary format

## Query CorporateInvoices

You can get a list of created invoices given some filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporateinvoice"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  invoices, errorChannel := corporateinvoice.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case invoice, ok := <-invoices:
      if !ok {
        break loop
      }
      fmt.Println(invoice)
    }
  }
}

```

## Create CorporateWithdrawals

You can create withdrawals to send cash back from your Corporate balance to your Banking balance
by using the Withdrawal resource.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatewithdrawal"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  withdrawal, err := corporatewithdrawal.Create(
    corporatewithdrawal.CorporateWithdrawal{
      Amount: 10000,
      ExternalId: "my-external-id",
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(withdrawal)
}

```

**Note**: Instead of using CorporateWithdrawal objects, you can also pass each element in dictionary format

## Get a CorporateWithdrawal

After its creation, information on a withdrawal may be retrieved by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatewithdrawal"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  withdrawal, err := corporatewithdrawal.Get("5353197895942144", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(withdrawal)
}

```

## Query CorporateWithdrawals

You can get a list of created withdrawals given some filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatewithdrawal"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  withdrawals, errorChannel := corporatewithdrawal.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case withdrawal, ok := <-withdrawals:
      if !ok {
        break loop
      }
      fmt.Println(withdrawal)
    }
  }
}

```

## Get your CorporateBalance

To know how much money you have available to run authorizations, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatebalance"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  balance := corporatebalance.Get(nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(balance)
}

```

## Query CorporateTransactions

To understand your balance changes (corporate statement), you can query
transactions. Note that our system creates transactions for you when
you make purchases, withdrawals, receive corporate invoice payments, for example.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatetransaction"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 10

  transactions, errorChannel := corporatetransaction.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case transaction, ok := <-transactions:
      if !ok {
        break loop
      }
      fmt.Println(transaction)
    }
  }
}

```

## Get a CorporateTransaction

You can get a specific transaction by its id:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/corporatetransaction"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  transaction, err := corporatetransaction.Get("5353197895942144", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Println(transaction)
}

```

## Corporate Enums

### Query MerchantCategories

You can query any merchant categories using this resource.
You may also use MerchantCategories to define specific category filters in CorporateRules.
Either codes (which represents specific MCCs) or types (code groups) will be accepted as filters.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/merchantcategory"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  categories, errorChannel := merchantcategory.Query(nil, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case category, ok := <-categories:
      if !ok {
        break loop
      }
      fmt.Println(category)
    }
  }
}
```

### Query MerchantCountries

You can query any merchant countries using this resource.
You may also use MerchantCountries to define specific country filters in CorporateRules.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/merchantcountry"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  countries, errorChannel := merchantcountry.Query(nil, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case country, ok := <-countries:
      if !ok {
        break loop
      }
      fmt.Println(country)
    }
  }
}

```

### Query CardMethods

You can query available card methods using this resource.
You may also use CardMethods to define specific purchase method filters in CorporateRules.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/cardmethod"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  methods, errorChannel := cardmethod.Query(nil, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case method, ok := <-methods:
      if !ok {
        break loop
      }
      fmt.Println(method)
    }
  }
}

```

## Query MerchantCards

The Merchant Card resource stores information about cards used in approved purchases.
These cards can be used in new purchases without the need to create a new session.

```go
import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/merchantcard"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  cards, errorChannel := merchantcard.Query(nil, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case card, ok := <-cards:
      if !ok {
        break loop
      }
      fmt.Println(card)
    }
  }
}
```

## Get a MerchantCard

Retrieve detailed information about a specific card by its id.

```go
import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  "github.com/starkbank/sdk-go/starkbank/merchantcard"
)

func main() {

  merchantCard, err := merchantcard.Get("5950134772826112", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Printf("%+v", merchantCard)
} 
```

## Create a MerchantSession

The Merchant Session allows you to create a session prior to a purchase.
Sessions are essential for defining the parameters of a purchase, including funding type, expiration, 3DS, and more.

```go
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	MerchantSession "github.com/starkbank/sdk-go/starkbank/merchantsession"
	AllowedInstallment "github.com/starkbank/sdk-go/starkbank/merchantsession/allowedinstallment"
)
func main() {

  merchantSession := MerchantSession.MerchantSession{
    AllowedFundingTypes: []string{"credit"},
    AllowedIps:          []string{"192.168.0.1"},
    AllowedInstallments: []AllowedInstallment.AllowedInstallment{
      {Count: 1, TotalAmount: 0},
      {Count: 2, TotalAmount: 120},
      {Count: 12, TotalAmount: 180},
    },
    Expiration:   		 60,
    ChallengeMode: 		 "disabled",
    Tags:          		 []string{"test"},
  }

  createdSession, err := MerchantSession.Create(merchantSession, nil)

  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  fmt.Println(createdSession)
}
```

You can create a MerchantPurchase through a MerchantSession by passing its UUID.
**Note**: This method must be implemented in your front-end to ensure that sensitive card data does not pass through the back-end of the integration.

## Create a MerchantSession Purchase

This route can be used to create a Merchant Purchase directly from the payer's client application.
The UUID of a Merchant Session that was previously created by the merchant is necessary to access this route.

```go
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
  Purchase "github.com/starkbank/sdk-go/starkbank/merchantsession"
)
func main() {

  	purchase := Purchase.Purchase{
		Amount:            180,
		InstallmentCount:  12,
		CardExpiration:    "2035-01",
		CardNumber:        "5102589999999913",
		CardSecurityCode:  "123",
		HolderName:        "Holder Name",
		HolderEmail:       "holdeName@email.com",
		HolderPhone:       "11111111111",
		FundingType:       "credit",
		BillingCountryCode: "BRA",
		BillingCity:       "São Paulo",
		BillingStateCode:  "SP",
		BillingStreetLine1: "Rua do Holder Name, 123",
		BillingStreetLine2: "casa",
		BillingZipCode:    "11111-111",
		Metadata: map[string]interface{}{
			"userAgent":      "Postman",
			"userIp":         "255.255.255.255",
			"language":       "pt-BR",
			"timezoneOffset": 3,
			"extraData":      "extraData",
		},
	}

	createdPurchase, err := MerchantSession.PostPurchase("0bb894a2697d41d99fe02cad2c00c9bc", purchase, nil)

	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}

  fmt.Println(createdPurchase)
}
```

## Query MerchantSessions

Get a list of merchant sessions in chunks of at most 100. If you need smaller chunks, use the limit parameter.

```go
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
  MerchantSession "github.com/starkbank/sdk-go/starkbank/merchantsession"
)

func main() {

  starkbank.User = utils.ExampleProject

  sessions, errorChannel := MerchantSession.Query(nil, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case session, ok := <-sessions:
      if !ok {
        break loop
      }
      fmt.Println(session)
    }
  }
}
```

## Get a MerchantSession

Retrieve detailed information about a specific session by its id.

```go
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
  MerchantSession "github.com/starkbank/sdk-go/starkbank/merchantsession"
)
func main() {
	sessions, err := MerchantSession.get("5950134772826112", nil)

	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}

  fmt.Println(sessions)
}
```

## Create a MerchantPurchase

The Merchant Purchase section allows users to retrieve detailed information of the purchases.

```go
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
  MerchantPurchase "github.com/starkbank/sdk-go/starkbank/merchantpurchase"
)

merchantPurchase := MerchantPurchase.MerchantPurchase{
  Amount:           	1000,
  FundingType: 		"credit",
  CardId: 		 	  "5920400605184000",
  ChallengeMode: 		"disabled",
}

createdMerchantPurchase, err := MerchantPurchase.Create(merchantPurchase, nil)

if err.Errors != nil {
  for _, e := range err.Errors {
    fmt.Printf("code: %s, message: %s", e.Code, e.Message)
  }
}
```

## Query MerchantPurchases

Get a list of merchant purchases in chunks of at most 100. If you need smaller chunks, use the limit parameter.

```go
import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  MerchantPurchase "github.com/starkbank/sdk-go/starkbank/merchantpurchase"
)

func main() {

  starkbank.User = utils.ExampleProject

  purchases, errorChannel := MerchantPurchase.Query(nil, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case purchase, ok := <-purchases:
      if !ok {
        break loop
      }
      fmt.Println(purchase)
    }
  }
}
```

## Get a MerchantPurchase

Retrieve detailed information about a specific purchase by its id.

```go
import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  MerchantPurchase "github.com/starkbank/sdk-go/starkbank/merchantpurchase"
)

func main() {

  merchantPurchase, err := MerchantPurchase.Get("5950134772826112", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Printf("%+v", merchantPurchase)
}
```

## Query MerchantInstallments

Merchant Installments are created for every installment in a purchase.
These resources will track its own due payment date and settlement lifecycle.

```go
import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  MerchantInstallment "github.com/starkbank/sdk-go/starkbank/merchantinstallment"
)

func main() {

  starkbank.User = utils.ExampleProject

  installments, errorChannel := MerchantInstallment.Query(nil, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case installment, ok := <-installments:
      if !ok {
        break loop
      }
      fmt.Println(installment)
    }
  }
}
```

## Get a MerchantInstallment

Retrieve detailed information about a specific installment by its id.

```go
import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  MerchantInstallment "github.com/starkbank/sdk-go/starkbank/merchantinstallment"
)

func main() {

  merchantinstallment, err := MerchantInstallment.Get("4848075206033408", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }

  fmt.Printf("%+v", merchantinstallment)
}  
```

## Create a webhook subscription

To create a webhook subscription and be notified whenever an event occurs, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Webhook "github.com/starkbank/sdk-go/starkbank/webhook"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  webhook, err := Webhook.Create(
    Webhook.Webhook{
      Url:           "https://webhook.site/dd784f26-1d6a-4ca6-81cb-fda0267761ec",
      Subscriptions: []string{"transfer", "boleto", "boleto-payment", "boleto-holmes", "brcode-payment", "utility-payment", "deposit", "invoice"},
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(webhook)
}

```

## Query webhooks

To search for registered webhooks, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Webhook "github.com/starkbank/sdk-go/starkbank/webhook"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  webhooks, errorChannel := Webhook.Query(nil, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case webhook, ok := <-webhooks:
      if !ok {
        break loop
      }
      fmt.Println(webhook)
    }
  }
}

```

## Get a webhook

You can get a specific webhook by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Webhook "github.com/starkbank/sdk-go/starkbank/webhook"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  webhook, err := Webhook.Get("10827361982368179", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(webhook)
}

```

## Delete a webhook

You can also delete a specific webhook by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Webhook "github.com/starkbank/sdk-go/starkbank/webhook"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  webhook, err := Webhook.Delete("10827361982368179", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(webhook)
}

```

## Process webhook events

It's easy to process events that arrived in your webhook. Remember to pass the
signature header so the SDK can make sure it's really StarkBank that sent you
the event.

```golang
package main

import (
	"fmt"
	Event "github.com/starkbank/sdk-go/starkbank/event"
	"github.com/starkbank/sdk-go/tests/utils"
)

func main() {

	request := listen() // this is the method you made to get the events posted to your webhook endpoint

	event := Event.Parse(
		request.Data,
		request.Headers["Digital-Signature"],
		"",
		utils.ExampleProject,
	)

	if event.Subscription == "transfer" {
		fmt.Println(event.Log)
	} else if event.Subscription == "boleto" {
		fmt.Println(event.Log)
	} else if event.Subscription == "boleto-payment" {
		fmt.Println(event.Log)
	} else if event.Subscription == "boleto-holmes" {
		fmt.Println(event.Log)
	} else if event.Subscription == "brcode-payment" {
		fmt.Println(event.Log)
	} else if event.Subscription == "utility-payment" {
		fmt.Println(event.Log)
	} else if event.Subscription == "deposit" {
		fmt.Println(event.Log)
	} else if event.Subscription == "invoice" {
		fmt.Println(event.Log)
	}
}

```

## Query webhook events

To search for webhooks events, run:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Event "github.com/starkbank/sdk-go/starkbank/event"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["isDelivered"] = "false"
  params["after"] = "2020-03-20"

  events, errorChannel := Event.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case event, ok := <-events:
      if !ok {
        break loop
      }
      fmt.Println(event)
    }
  }
}

```

## Get a webhook event

You can get a specific webhook event by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Event "github.com/starkbank/sdk-go/starkbank/event"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  event, err := Event.Get("10827361982368179", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(event)
}

```

## Delete a webhook event

You can also delete a specific webhook event by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Event "github.com/starkbank/sdk-go/starkbank/event"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  event, err := Event.Delete("10827361982368179", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(event)
}

```

## Set webhook events as delivered

This can be used in case you've lost events.
With this function, you can manually set events retrieved from the API as
"delivered" to help future event queries with `isDelivered=False`.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Event "github.com/starkbank/sdk-go/starkbank/event"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var patchData = map[string]interface{}{}
  patchData["isDelivered"] = true

  event, err := Event.Update("10827361982368179", patchData, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
	
  fmt.Println(event)
}

```

## Query failed webhook event delivery attempts information

You can also get information on failed webhook event delivery attempts.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Attempt "github.com/starkbank/sdk-go/starkbank/event/attempt"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["after"] = "2020-03-20"

  attempts, errorChannel := Attempt.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case attempt, ok := <-attempts:
      if !ok {
        break loop
      }
      fmt.Println(attempt)
    }
  }
}

```

## Get a failed webhook event delivery attempt information

To retrieve information on a single attempt, use the following function:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Attempt "github.com/starkbank/sdk-go/starkbank/event/attempt"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  attempt, err := Attempt.Get("1616161616161616", nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(attempt)
}

```

## Create a new Workspace

The Organization user allows you to create new Workspaces (bank accounts) under your organization.
Workspaces have independent balances, statements, operations and users.
The only link between your Workspaces is the Organization that controls them.

**Note**: This route will only work if the Organization user is used with `workspaceId=nil`.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Workspace "github.com/starkbank/sdk-go/starkbank/workspace"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  workspace, err := Workspace.Create(
    Workspace.Workspace{
      Username:      "iron-bank-workspace-1",
      Name:          "Iron Bank Workspace 1",
      AllowedTaxIds: []string{""},
    }, nil)

  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(workspace)
}

```

## List your Workspaces

This route lists Workspaces. If no parameter is passed, all the workspaces the user has access to will be listed, but
you can also find other Workspaces by searching for their usernames or IDs directly.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Workspace "github.com/starkbank/sdk-go/starkbank/workspace"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var params = map[string]interface{}{}
  params["limit"] = 30

  workspaces, errorChannel := Workspace.Query(params, nil)
  loop:
  for {
    select {
    case err := <-errorChannel:
      if err.Errors != nil {
        for _, e := range err.Errors {
          fmt.Printf("code: %s, message: %s", e.Code, e.Message)
        }
      }
    case workspace, ok := <-workspaces:
      if !ok {
        break loop
      }
      fmt.Println(workspace)
    }
  }
}

```

## Get a Workspace

You can get a specific Workspace by its id.

```golang
package main

import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Workspace "github.com/starkbank/sdk-go/starkbank/workspace"
	"github.com/starkbank/sdk-go/tests/utils"
)

func main() {

	starkbank.User = utils.ExampleProject

	workspace, err := Workspace.Get("10827361982368179", nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	
	fmt.Println(workspace)
}

```

## Update a Workspace

You can update a specific Workspace by its id.

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Workspace "github.com/starkbank/sdk-go/starkbank/workspace"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  var patchData = map[string]interface{}{}
  patchData["username"] = "new-username"
  patchData["name"] = "New Name"
  patchData["allowedTaxIds"] = []string{"012.345.678-90"}

  workspace, err := Workspace.Update("10827361982368179", patchData, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  fmt.Println(workspace)
}

```

**Note**: the Organization user can only update a workspace with the Workspace ID set.

# request

This resource allows you to send HTTP requests to StarkBank routes.

## GET

You can perform a GET request to any StarkBank route.

It's possible to get a single resource using its id in the path.

```golang
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
)
func main() {
	data := map[string]interface{}{}
	var path string

	path = "/invoice/5155165527080960"

	response, err := Request.Get(
		path,
		nil,
		nil,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		fmt.Printf(unmarshalError)
	}
  
  fmt.Println(data)
}
```

You can also get the specific resource log,

```golang
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
)
func main() {
	data := map[string]interface{}{}
	var path string

	path = "/invoice/log/5155165527080960"

	response, err := Request.Get(
		path,
		nil,
		nil,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		fmt.Printf(unmarshalError)
	}
  
  fmt.Println(data)
}
```

This same method will be used to list all created items for the requested resource.

```golang
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
)
func main() {
	data := map[string]interface{}{}
	var path string
	var query = map[string]interface{}{}

	path = "/invoice/"
	query["limit"] = 2
  query["status"] = "paid"

	response, err := Request.Get(
		path,
		query,
		nil,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		fmt.Printf(unmarshalError)
	}
  
  fmt.Println(data)
}
```

To list logs, you will use the same logic as for getting a single log.

```golang
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
)
func main() {
	data := map[string]interface{}{}
	var path string
	var query = map[string]interface{}{}

	path = "/invoice/log"
	query["limit"] = 2
  query["status"] = "paid"

	response, err := Request.Get(
		path,
		query,
		nil,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		fmt.Printf(unmarshalError)
	}
  
  fmt.Println(data)
}
```

You can get a resource file using this method.

```golang
import (
	"fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
)
func main() {
	var path string

	path = "/invoice/log/5155165527080960/pdf"

	response, err := Request.Get(
		path,
		nil,
		nil,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}
  filename := fmt.Sprintf("%v%v.pdf", "transfer", "5155165527080960")
  errFile := ioutil.WriteFile(filename, pdf, 0666)
  if errFile != nil {
    fmt.Print(errFile)
  }
}
```

## POST

You can perform a POST request to any StarkBank route.

This will create an object for each item sent in your request

**Note**: It's not possible to create multiple resources simultaneously. You need to send separate requests if you want to create multiple resources, such as invoices and boletos.

```golang
import (
  "fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
)
func main() {
	starkbank.User = Utils.ExampleProject
	data := map[string]interface{}{}
	var path string
	body := map[string][]map[string]interface{}{
        "invoices": {
            {
				"amount": 996699999,
				"name":   "Tony Stark",
				"taxId":  "38.446.231/0001-04",
			},
        },
    }
	path = "/invoice/"

	response, err := Request.Post(
		path,
		body,
		nil,
		Utils.ExampleProject,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		fmt.Printf(unmarshalError)
	}

  fmt.Println(data)
}
```

## PATCH

You can perform a PATCH request to any StarkBank route.

It's possible to update a single item of a StarkBank resource.
```golang
import (
  "fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
)
func main() {
  starkbank.User = Utils.ExampleProject

  body := map[string]interface{}{
    "amount" : 0,
  }
	path = "/invoice/log/5155165527080960"

  response, err := Request.Patch(
    path,
    body,
    nil,
    Utils.ExampleProject,
  )

  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  unmarshalError := json.Unmarshal(response.Content, &data)
  if unmarshalError != nil {
    fmt.Printf(unmarshalError)
  }
  fmt.Println(data)
}
```

## PUT

You can perform a PUT request to any StarkBank route.

It's possible to put a single item of a StarkBank resource.
```golang
import (
  "fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
)
func main() {
	starkbank.User = Utils.ExampleProject
	data := map[string]interface{}{}
	body := map[string][]map[string]interface{}{
        "profiles": {
            {
				"interval": "day",
				"delay": 0,
			},
        },
    }
	path := "split-profile/"
	response, err := Request.Put(
		path,
		body,
		nil,
		Utils.ExampleProject,
	)
	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		fmt.Printf(unmarshalError)
	}
  fmt.Println(data)
}
```

## DELETE

You can perform a DELETE request to any StarkBank route.

It's possible to delete a single item of a StarkBank resource.
```golang
import (
  "fmt"
	"github.com/starkbank/sdk-go/starkbank"
	Request "github.com/starkbank/sdk-go/starkbank/request"
)
func main() {
starkbank.User = Utils.ExampleProject
  path = "transfer/5155165527080960"
  response, err := Request.Delete(
    path,
    Utils.ExampleProject,
  )

  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  unmarshalError := json.Unmarshal(response.Content, &data)
  if unmarshalError != nil {
    fmt.Printf(unmarshalError)
  }
  fmt.Println(data)
}
```

# Handling errors

The SDK may return errors as the StarkErrors struct, which contains the "code" and "message" attributes.

It's highly recommended that you handle the errors returned from the functions used to get a feedback of the operation,
as the example below:

__InputErrors__ will be raised whenever the API detects an error in your request (status code 400).
If you catch such an error, you can get its elements to verify each of the
individual errors that were detected in your request by the API.

For example:

```golang
package main

import (
  "fmt"
  "github.com/starkbank/sdk-go/starkbank"
  Transaction "github.com/starkbank/sdk-go/starkbank/transaction"
  "github.com/starkbank/sdk-go/tests/utils"
)

func main() {

  starkbank.User = utils.ExampleProject

  transactions, err := Transaction.Create(
    []Transaction.Transaction{
      {
        Amount:      10000,
        ReceiverId:  "5768064935133184",
        Description: "Paying my debts",
        ExternalId:  "my_unique_external_id",
      },
    }, nil)
  if err.Errors != nil {
    for _, e := range err.Errors {
      fmt.Printf("code: %s, message: %s", e.Code, e.Message)
    }
  }
  
  for _, transaction := range transactions {
    fmt.Println(transaction)
  }
}

```

__InternalServerError__ will be raised if the API runs into an internal error.
If you ever stumble upon this one, rest assured that the development team
is already rushing in to fix the mistake and get you back up to speed.

__UnknownError__ will be raised if a request encounters an error that is
neither __InputErrors__ nor an __InternalServerError__, such as connectivity problems.

__InvalidSignatureError__ will be raised specifically by event.Parse()
when the provided content and signature do not check out with the Stark Bank public
key.

# Help and Feedback

If you have any questions about our SDK, just email us.
We will respond you quickly, pinky promise. We are here to help you integrate with us ASAP.
We also love feedback, so don't be shy about sharing your thoughts with us.

Email: help@starkbank.com
