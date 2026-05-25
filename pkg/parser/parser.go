package parser

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/ledongthuc/pdf"
)

func ParsePDF(path string) ([]Order, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println("Error closing file:", err)
		}
	}(f)

	var sb strings.Builder

	for i := 1; i <= r.NumPage(); i++ {
		p := r.Page(i)
		if p.V.IsNull() {
			continue
		}

		text, err := p.GetPlainText(nil)
		if err != nil {
			return nil, fmt.Errorf("page %d: %w", i, err)
		}

		sb.WriteString(text)
		sb.WriteString("\n")
	}

	text := sb.String()

	// Regex to capture each row: Date, OrderID, RestaurantName, Amount
	// Date: dd-mm-yyyy
	// OrderID: long digits
	// Amount: ₹digits.digits
	// Restaurant: everything in between
	rowRe := regexp.MustCompile(`(\d{2}-\d{2}-\d{4})(\d{15})(.+?)(₹[\d,]+\.\d{2})`)
	matches := rowRe.FindAllStringSubmatch(text, -1)

	var orders []Order

	for _, m := range matches {
		if len(m) >= 5 {
			orders = append(orders, Order{
				Date:       m[1],
				Restaurant: strings.TrimSpace(m[3]),
				Amount:     m[4],
			})
		}
	}

	return orders, nil
}
