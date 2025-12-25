# Sunbay Nexus SDK for Go

Official Go SDK for Sunbay Nexus Payment Platform

## Features

- ✅ Simple and intuitive API
- ✅ Config-based client initialization (Go idiomatic)
- ✅ Support Go 1.18+
- ✅ Automatic authentication
- ✅ Automatic retry for GET requests
- ✅ Comprehensive error handling
- ✅ Minimal dependencies
- ✅ Flexible logging support

## Installation

```bash
# Install latest version
go get github.com/sunbay-developer/sunbay-nexus-sdk-go

# Install specific version
go get github.com/sunbay-developer/sunbay-nexus-sdk-go@v1.0.0
```

## Quick Start

### 1. Initialize Client

The `NexusClient` is thread-safe and can be reused across multiple goroutines.

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    nexus "github.com/sunbay-developer/sunbay-nexus-sdk-go"
    "github.com/sunbay-developer/sunbay-nexus-sdk-go/errors"
    "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/request"
    "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"
)

func main() {
    // Create NexusClient with Config
    config := &nexus.Config{
        APIKey: "your-api-key",
        BaseURL: "https://open.sunbay.us",
    }
    client, err := nexus.NewNexusClient(config)
    if err != nil {
        log.Fatal(err)
    }
    
    // Build request
    amount := 100.00
    req := &request.SaleRequest{
        AppID:               "app_123456",
        MerchantID:          "mch_789012",
        ReferenceOrderID:    "ORDER20231119001",
        TransactionRequestID: "PAY_REQ_1234567890",
        Amount: &common.SaleAmount{
            OrderAmount:     &amount,
            PricingCurrency: "USD",
        },
        Description: "Product purchase",
        TerminalSN:  "T1234567890",
    }
    
    // Execute transaction
    ctx := context.Background()
    resp, err := client.Sale(ctx, req)
    if err != nil {
        // Error handling
        if bizErr, ok := err.(*errors.BusinessError); ok {
            log.Printf("Business error: code=%s, msg=%s, traceID=%s",
                bizErr.Code(), bizErr.Message(), bizErr.TraceID())
        } else if netErr, ok := err.(*errors.NetworkError); ok {
            log.Printf("Network error: %v (retryable: %v)", netErr, netErr.IsRetryable())
        }
        return
    }
    
    // Handle response (if no error, response is successful)
    fmt.Printf("Transaction ID: %s\n", resp.TransactionID)
}
```

### 2. Advanced Configuration

```go
config := &nexus.Config{
    APIKey:         "your-api-key",
    BaseURL:        "https://open.sunbay.us",
    ConnectTimeout: 30 * time.Second,
    ReadTimeout:    60 * time.Second,
    MaxRetries:     3,
    MaxTotal:       200,
    MaxPerRoute:    20,
}
client, err := nexus.NewNexusClient(config)
```

### 3. Custom Logger

The SDK supports custom loggers. If no logger is provided, it uses a default console logger.

```go
// Use default logger
config := &nexus.Config{
    APIKey: "your-api-key",
}
client, err := nexus.NewNexusClient(config)

// Use custom logger (implement the Logger interface)
type myLogger struct {
    // your logger implementation
}

func (l *myLogger) Debug(args ...interface{}) { /* ... */ }
func (l *myLogger) Debugf(format string, args ...interface{}) { /* ... */ }
func (l *myLogger) Info(args ...interface{})  { /* ... */ }
func (l *myLogger) Infof(format string, args ...interface{}) { /* ... */ }
func (l *myLogger) Warn(args ...interface{})  { /* ... */ }
func (l *myLogger) Warnf(format string, args ...interface{}) { /* ... */ }
func (l *myLogger) Error(args ...interface{}) { /* ... */ }
func (l *myLogger) Errorf(format string, args ...interface{}) { /* ... */ }

config := &nexus.Config{
    APIKey: "your-api-key",
    Logger: &myLogger{},
}
client, err := nexus.NewNexusClient(config)
```

## API Methods

### Transaction APIs

- `Sale(ctx, req)` - Sale transaction
- More APIs coming soon...

### Query APIs

- Coming soon...

## Error Handling

The SDK returns two types of errors:

- **BusinessError**: Business logic errors (parameter validation, API business errors, etc.)
- **NetworkError**: Network-related errors (connection timeout, network error, etc.)

Always check error type:

```go
resp, err := client.Sale(ctx, req)
if err != nil {
    if bizErr, ok := err.(*errors.BusinessError); ok {
        // Handle business error
        log.Printf("Business error: code=%s, msg=%s", bizErr.Code(), bizErr.Message())
    } else if netErr, ok := err.(*errors.NetworkError); ok {
        // Handle network error
        if netErr.IsRetryable() {
            // Can retry
        }
    }
}
```

## Requirements

- Go 1.18 or higher

## License

MIT License

