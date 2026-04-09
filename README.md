# Aurora Go SDK

Official Go SDK for the [Aurora API](https://kints.dev) — a rule-based transaction processing engine.

## Installation

```bash
go get github.com/kintsdev/aurora-go-sdk
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	aurora "github.com/kintsdev/aurora-go-sdk"
)

func main() {
	client := aurora.NewClient("your-api-key",
		aurora.WithBaseURL("https://api.example.com"),
	)

	resp, err := client.Process.Execute(context.Background(), &aurora.ProcessRequest{
		RuleID: "your-rule-id",
		Transaction: &aurora.Transaction{
			Common: &aurora.Common{
				Amount:   250.00,
				Currency: "USD",
				Email:    "customer@example.com",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Allowed: %v | Rejected: %v | Score: %d\n", resp.Allow, resp.Rejected, resp.Score)
}
```

## Configuration

### Client Options

| Option | Description |
|---|---|
| `WithBaseURL(url)` | Set the API host (e.g. `https://api.example.com`). The path `/api/v1` is appended automatically. |
| `WithHTTPClient(client)` | Provide a custom `*http.Client`. |
| `WithTimeout(duration)` | Set HTTP client timeout (default: 30s). |

### Authentication

The SDK sends the API key as a `Bearer` token in the `Authorization` header.

```go
client := aurora.NewClient("your-api-key")
```

## Process

Execute a rule against a transaction using `client.Process.Execute()`.

### With a Single Rule

```go
resp, err := client.Process.Execute(ctx, &aurora.ProcessRequest{
    RuleID: "rule-id",
    Transaction: &aurora.Transaction{
        Common: &aurora.Common{
            Amount:   100.00,
            Currency: "TRY",
            UserID:   "user-123",
        },
    },
})
```

### With a Ruleset

```go
resp, err := client.Process.Execute(ctx, &aurora.ProcessRequest{
    RulesetID: "ruleset-id",
    Transaction: &aurora.Transaction{
        Common: &aurora.Common{
            Amount:   500.00,
            Currency: "EUR",
        },
        Transfer: &aurora.Transfer{
            PaymentMethod: "bank_deposit",
            SenderCountry: "DE",
        },
    },
})
```

### Full Example

```go
resp, err := client.Process.Execute(ctx, &aurora.ProcessRequest{
    RuleID:        "1e7b3b7b-0b3b-4b7b-8b3b-0b3b7b0b3b7b",
    RulesetID:     "2a8c4d9e-1f2a-4b3c-9d8e-7f6a5b4c3d2e",
    TransactionID: "txn_20260326_0001",
    Transaction: &aurora.Transaction{
        Type:            "card",
        TransactionType: "purchase",
        AccountAgeDays:  "180",
        IsNewDevice:     "true",
        IsNewIP:         "false",
        IsUnusualLocation: "true",
        Common: &aurora.Common{
            Amount:      99.99,
            Currency:    "USD",
            Category:    "ecommerce",
            Country:     "US",
            Email:       "john@doe.com",
            IPAddress:   "127.0.0.1",
            UserID:      "user_1234567890",
            MerchantID:  "merch_1234567890",
            PaymentID:   "pay_1234567890",
            ReferenceID: "order_1234567890",
        },
        Card: &aurora.Card{
            BinNumber: "123456",
            Holder:    "John Doe",
            Issuer:    "Bank of America",
            LastFour:  "1234",
            MCC:       "5411",
            Network:   "Visa",
            Token:     "card_tok_1234567890",
            Type:      "credit",
        },
        Transfer: &aurora.Transfer{
            TransferType:    "wallet_to_wallet",
            TransferPurpose: "family_support",
            SenderName:      "John Doe",
            SenderCountry:   "US",
            ReceiverName:    "Jane Smith",
            ReceiverCountry: "GB",
        },
    },
})
```

### Response

```go
type ProcessResponse struct {
    Allow          bool          // transaction is allowed
    AllowMessage   string        // message when allowed
    Rejected       bool          // transaction is rejected
    RejectMessage  string        // reason for rejection
    NeedInspect    bool          // transaction needs manual review
    InspectMessage string        // reason for inspection
    Score          int           // risk score
    Error          bool          // processing error occurred
    ErrorMessage   string        // error details
    ExecutionTime  time.Duration // rule execution duration
}
```

### Transaction Fields

The `Transaction` struct uses nested sub-structs to organize fields by category. Only populate the sections relevant to your use case — all sub-struct pointers and fields use `omitempty`.

#### Top-Level Scalar Fields

| Field | JSON Key | Description |
|---|---|---|
| `Type` | `type` | Transaction type (e.g. `card`) |
| `TransactionType` | `transaction_type` | Transaction action (e.g. `purchase`) |
| `AccountAgeDays` | `account_age_days` | Account age in days |
| `DeclinedCount` | `declined_count` | Number of declined transactions |
| `IncomeMultiplier` | `income_multiplier` | Ratio of amount to registered income |
| `IsFirstTransfer` | `is_first_transfer` | Whether this is the first transfer |
| `IsNewDevice` | `is_new_device` | Whether a new device is used |
| `IsNewIP` | `is_new_ip` | Whether a new IP is used |
| `IsUnusualLocation` | `is_unusual_location` | Whether location is unusual |
| `PasswordChangedRecently` | `password_changed_recently` | Recent password change flag |
| `ProfileCompletion` | `profile_completion` | Profile completion percentage |
| `RefundCount` | `refund_count` | Number of refunds |
| `RefundRatio` | `refund_ratio` | Refund to transaction ratio |
| `RegisteredIncome` | `registered_income` | User's registered income |
| `TotalAmount24h` | `total_amount_24h` | Total transaction amount in 24h |
| `TotalAmount7d` | `total_amount_7d` | Total transaction amount in 7 days |
| `TransactionHour` | `transaction_hour` | Hour of the transaction |
| `TransferCount24h` | `transfer_count_24h` | Number of transfers in 24h |
| `UniqueRecipients` | `unique_recipients` | Number of unique recipients |

#### Nested Sub-Structs

| Sub-Struct | JSON Key | Fields |
|---|---|---|
| `Common` | `common` | `Amount`, `BrowserAgent`, `Category`, `ConnectionType`, `Country`, `Currency`, `Date`, `Description`, `DeviceFingerprint`, `DeviceID`, `Email`, `IPAddress`, `LastLoginIP`, `LastLoginTime`, `Latitude`, `Longitude`, `MerchantID`, `PaymentID`, `Phone`, `ReferenceID`, `UserID` |
| `Card` | `card` | `BinNumber`, `Holder`, `Issuer`, `LastFour`, `MCC`, `Network`, `Token`, `Type` |
| `Transfer` | `transfer` | `ExchangeRate`, `PaymentMethod`, `ReceiverAddress`, `ReceiverCountry`, `ReceiverIBAN`, `ReceiverIdentityPassport`, `ReceiverName`, `ReceiverSurname`, `ReceiverWalletID`, `Relationship`, `SenderAddress`, `SenderCountry`, `SenderIBAN`, `SenderIdentityPassport`, `SenderName`, `SenderSurname`, `SenderWalletID`, `SourceOfFunds`, `TargetCurrency`, `TransferPurpose`, `TransferType` |
| `AccountChange` | `account_change` | `ChangeType`, `PreviousValueHash`, `TimeSinceLastChange`, `VerificationMethod` |
| `AccountLogin` | `account_login` | `FailedAttempts`, `LoginMethod`, `LoginStatus`, `MFAMethod`, `MFAUsed`, `SessionID` |
| `AccountOpening` | `account_opening` | `DocumentType`, `IdentityVerificationStatus`, `RegistrationMethod` |
| `BNPL` | `bnpl` | `BNPLProvider`, `InstallmentCount`, `OutstandingBNPLAmount` |
| `Crypto` | `crypto` | `CryptoAmount`, `CryptoCurrency`, `ExchangeName`, `IsSmartContract`, `WalletAddress` |
| `Deposit` | `deposit` | `CheckNumber`, `DepositMethod`, `DepositSource`, `IsRemoteDeposit` |
| `Invoice` | `invoice` | `InvoiceDueDate`, `InvoiceNumber`, `IsRecurringVendor`, `PurchaseOrderID`, `VendorID`, `VendorName` |
| `Loan` | `loan` | `CreditScore`, `DebtToIncomeRatio`, `EmploymentStatus`, `LoanAmount`, `LoanTerm`, `LoanType` |
| `P2P` | `p2p` | `PaymentNote`, `Platform`, `RecipientAccountAge` |
| `Refund` | `refund` | `DaysSincePurchase`, `OriginalTransactionID`, `RefundMethod`, `RefundReason` |
| `Withdrawal` | `withdrawal` | `AccountType`, `ATMID`, `DailyWithdrawalCount`, `WithdrawalMethod` |

## Error Handling

API errors are returned as `*aurora.APIError`:

```go
resp, err := client.Process.Execute(ctx, req)
if err != nil {
    var apiErr *aurora.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API error %d: %s\n", apiErr.StatusCode, apiErr.Message)
    }
    log.Fatal(err)
}
```

## Callback

Send a callback transaction using `client.Callback.Transaction()`.

```go
resp, err := client.Callback.Transaction(ctx, &aurora.CallbackTransactionRequest{
    Message:   "Payment completed",
    PaymentID: "pay_1234567890",
    Status:    "success",
})
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Success: %v | Message: %s\n", resp.Success, resp.Message)
```

### Request

| Field | JSON Key | Description |
|---|---|---|
| `Message` | `message` | Callback message |
| `PaymentID` | `payment_id` | Payment identifier |
| `Status` | `status` | Transaction status |

### Response

```go
type CallbackTransactionResponse struct {
    Success bool   // operation succeeded
    Message string // response message
}
```

## License

MIT
