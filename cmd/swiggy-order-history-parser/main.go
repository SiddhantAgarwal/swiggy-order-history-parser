package main

import (
	"log"
	"os"

	"swiggy-order-history-parser/pkg/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: swiggy-order-history-parser <input.pdf>")
	}

	inputPath := os.Args[1]

	orders, err := parser.ParsePDF(inputPath)
	if err != nil {
		log.Fatalf("Error parsing PDF: %v", err)
	}

	outputPath := "orders.csv"
	if err := parser.WriteCSV(orders, outputPath); err != nil {
		log.Fatalf("Error writing CSV: %v", err)
	}

	log.Printf("Wrote %d orders to %s\n", len(orders), outputPath)
}
