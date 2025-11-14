# EcoCash Client Go

A Go client library for integrating with the EcoCash payment API in Zimbabwe. This library provides a simple and clean interface for processing payments, refunds, and querying transaction status.

## Features

- 🚀 Simple and intuitive API
- 💳 Process payments (charge subscribers)
- 💰 Handle refunds
- 🔍 Query transaction status
- 🛡️ Type-safe request/response structures
- 📝 Comprehensive error handling
- 🔧 Configurable HTTP client with debug support

## Installation

```bash
go get github.com/scape-labs/ecocash-client-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "github.com/scape-labs/ecocash-client-go"
)

func main() {
    // Configure the client
    config := ecocash.ClientConfig{
        BaseURL:             "https://api.ecocash.co.zw",
        Username:            "your-username",
        Password:            "your-password",
        MerchantCode:        "your-merchant-code",
        MerchantPin:         "your-merchant-pin",
        MerchantNumber:      "your-merchant-number",
        MerchantName:        "Your Business Name",
        SuperMerchantName:   "Your Super Merchant Name",
        TerminalID:          "your-terminal-id",
        Location:            "Harare",
        CountryCode:         "ZW",
        NotifyURL:           "https://your-domain.com/webhook",
        TraceRequests:       true,
        DisableHttpWarnings: false,
    }

    // Create a new client
    client := ecocash.NewClient(config)

    // Process a payment
    chargeReq := ecocash.ChargeSubscriberRequest{
        ReferenceCode: "REF-123456",
        PhoneNumber:   "263712345678",
        Amount:        10.50,
        Currency:      "USD",
        Description:   "Payment for goods",
    }

    response, err := client.Charge(chargeReq)
    if err != nil {
        log.Fatalf("Payment failed: %v", err)
    }

    fmt.Printf("Payment successful! Reference: %s\n", response.ServerReferenceCode)
}
```

## Configuration

The `ClientConfig` struct requires the following fields:

| Field | Type | Description |
|-------|------|-------------|
| `BaseURL` | string | EcoCash API base URL |
| `Username` | string | API username |
| `Password` | string | API password |
| `MerchantCode` | string | Your merchant code |
| `MerchantPin` | string | Your merchant PIN |
| `MerchantNumber` | string | Your merchant number |
| `MerchantName` | string | Your business name |
| `SuperMerchantName` | string | Super merchant name |
| `TerminalID` | string | Terminal ID |
| `Location` | string | Your location |
| `CountryCode` | string | Country code (usually "ZW") |
| `NotifyURL` | string | Webhook URL for notifications |
| `TraceRequests` | bool | Enable HTTP request debugging |
| `DisableHttpWarnings` | bool | Disable HTTP client warnings |

## API Methods

### Charge a Subscriber

Process a payment from a customer's EcoCash account.

```go
chargeReq := ecocash.ChargeSubscriberRequest{
    ReferenceCode: "UNIQUE-REF-123",
    PhoneNumber:   "263712345678",
    Amount:        25.00,
    Currency:      "USD", // or "ZWG"
    Description:   "Product purchase",
}

response, err := client.Charge(chargeReq)
```

### Process a Refund

Refund a previous transaction.

```go
refundReq := ecocash.SimpleRefund{
    ReferenceCode:            "REFUND-REF-123",
    PhoneNumber:              "263712345678",
    Amount:                   25.00,
    Currency:                 "USD",
    OriginalEcocashReference: "ORIGINAL-TRANSACTION-REF",
    Description:              "Customer refund",
}

response, err := client.Refund(refundReq)
```

### Query Transaction Status

Check the status of a transaction.

```go
response, err := client.QueryTransaction("263712345678", "REFERENCE-CODE")
```

## Response Structure

All API methods return a `TransactionResponse` struct:

```go
type TransactionResponse struct {
    ClientCorrelator           string        `json:"clientCorrelator"`
    EndTime                    int64         `json:"endTime"`
    StartTime                  int64         `json:"startTime"`
    ServerReferenceCode        string        `json:"serverReferenceCode"`
    TransactionOperationStatus string        `json:"transactionOperationStatus"`
    ResponseCode               string        `json:"responseCode"`
    EcocashReference           string        `json:"ecocashReference"`
    PaymentAmount              paymentAmount `json:"paymentAmount"`
}
```

## Error Handling

The client returns errors for various scenarios:

- Network connectivity issues
- Invalid API credentials
- Malformed requests
- API server errors

Always check for errors when calling API methods:

```go
response, err := client.Charge(chargeReq)
if err != nil {
    // Handle error appropriately
    log.Printf("Payment failed: %v", err)
    return
}
```

## Supported Currencies

- `USD` - US Dollars
- `ZWG` - Zimbabwe Gold (ZiG)

## Phone Number Format

Phone numbers should be in international format without the `+` sign:
- Correct: `263712345678`
- Incorrect: `+263712345678` or `0712345678`

## Development

### Prerequisites

- Go 1.23.0 or later

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support and questions:
- Create an issue in this repository
- Contact the maintainers

## Disclaimer

This is an unofficial client library for the EcoCash API. Please ensure you have proper authorization and agreements with EcoCash before using this library in production.

The authors are not responsible for any financial losses or issues that may arise from the use of this software. Always test thoroughly in a sandbox environment before deploying to production.