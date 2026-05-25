package parser

import (
	"encoding/csv"
	"log"
	"os"
)

// WriteCSV writes the provided orders to a CSV file at the given path.
// The output contains a header row: Date, Restaurant Name, Amount.
func WriteCSV(orders []Order, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println("Error closing file:", err)
		}
	}(file)

	w := csv.NewWriter(file)
	defer w.Flush()

	if err := w.Write([]string{"Date", "Restaurant Name", "Amount"}); err != nil {
		return err
	}

	for _, o := range orders {
		if err := w.Write([]string{o.Date, o.Restaurant, o.Amount}); err != nil {
			return err
		}
	}

	return nil
}
