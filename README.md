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
			Amount:   250.00,
			Currency: "USD",
			Email:    "customer@example.com",
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
        Amount:   100.00,
        Currency: "TRY",
        UserID:   "user-123",
    },
})
```

### With a Ruleset

```go
resp, err := client.Process.Execute(ctx, &aurora.ProcessRequest{
    RulesetID: "ruleset-id",
    Transaction: &aurora.Transaction{
        Amount:        500.00,
        Currency:      "EUR",
        PaymentMethod: "bank_deposit",
        SenderCountry: "DE",
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

The `Transaction` struct supports 50+ fields across these categories:

| Category | Fields |
|---|---|
| **Core** | `Amount`, `Currency`, `Category`, `Date`, `Description` |
| **Payment** | `PaymentID`, `PaymentMethod`, `ReferenceID`, `MerchantID`, `MCC` |
| **Card** | `BinNumber`, `CardHolder`, `CardIssuer`, `CardLastFour`, `CardNetwork`, `CardToken`, `CardType` |
| **User** | `UserID`, `Email`, `Phone`, `IPAddress` |
| **Location** | `Country`, `Latitude`, `Longitude` |
| **Device** | `BrowserAgent`, `ConnectionType` |
| **Login** | `LastLoginIP`, `LastLoginTime` |
| **Profile** | `RegisteredIncome`, `IncomeMultiplier` |
| **Transfer** | `TransferType`, `TransferPurpose`, `SourceOfFunds`, `ExchangeRate`, `TargetCurrency`, `Relationship`, `IsFirstTransfer`, `TotalAmount24h`, `TransferCount24h` |
| **Sender** | `SenderName`, `SenderSurname`, `SenderCountry`, `SenderAddress`, `SenderIBAN`, `SenderIdentityPassport`, `SenderWalletID` |
| **Receiver** | `ReceiverName`, `ReceiverSurname`, `ReceiverCountry`, `ReceiverAddress`, `ReceiverIBAN`, `ReceiverIdentityPassport`, `ReceiverWalletID` |

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

## License

MIT
