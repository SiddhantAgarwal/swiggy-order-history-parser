# Swiggy Order History Parser

[![Go Report Card](https://goreportcard.com/badge/github.com/SiddhantAgarwal/swiggy-order-history-parser)](https://goreportcard.com/report/github.com/SiddhantAgarwal/swiggy-order-history-parser)
[![Go Version](https://img.shields.io/github/go-mod/go-version/SiddhantAgarwal/swiggy-order-history-parser)](go.mod)
[![License](https://img.shields.io/github/license/SiddhantAgarwal/swiggy-order-history-parser)](LICENSE)

A small Go library and CLI tool to extract order history from Swiggy PDFs and write them to CSV.

## Installation

```bash
go get github.com/SiddhantAgarwal/swiggy-order-history-parser/pkg/parser
```

## Usage

### As a CLI

```bash
go run ./cmd/swiggy-order-history-parser <input.pdf>
```

This will generate an `orders.csv` file in the current directory.

### As a Library

```go
import "github.com/SiddhantAgarwal/swiggy-order-history-parser/pkg/parser"

orders, err := parser.ParsePDF("path/to/orders.pdf")
if err != nil {
    log.Fatal(err)
}

err = parser.WriteCSV(orders, "output.csv")
```

## Project Structure

```
.
├── cmd/swiggy-order-history-parser/   # CLI entry point
├── pkg/parser/                          # Reusable library code
│   ├── order.go                         # Order data model
│   ├── parser.go                        # PDF parsing logic
│   └── csv.go                           # CSV writer
├── go.mod
└── README.md
```

## License

MIT