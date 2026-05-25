package main

import (
	"flag"
	"log"

	"swiggy-order-history-parser/pkg/parser"
)

func main() {
	outputPath := flag.String("o", "orders.csv", "output CSV file path")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		log.Fatalln("Usage: swiggy-order-history-parser [-o output.csv] <input.pdf>")
	}

	inputPath := args[0]

	orders, err := parser.ParsePDF(inputPath)
	if err != nil {
		log.Fatalf("Error parsing PDF: %v", err)
	}

	if err := parser.WriteCSV(orders, *outputPath); err != nil {
		log.Fatalf("Error writing CSV: %v", err)
	}

	log.Printf("Wrote %d orders to %s\n", len(orders), *outputPath)
}
