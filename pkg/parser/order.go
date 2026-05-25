package parser

// Order represents a single Swiggy order extracted from a PDF.
type Order struct {
	Date       string
	Restaurant string
	Amount     string
}
