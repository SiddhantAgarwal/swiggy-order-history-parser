package parser

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteCSV(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		orders      []Order
		wantRecords [][]string
		wantErr     bool
	}{
		{
			name: "multiple orders",
			orders: []Order{
				{Date: "2024-01-01", Restaurant: "Foo Bar", Amount: "123.45"},
				{Date: "2024-01-02", Restaurant: "Baz Qux", Amount: "67.89"},
			},
			wantRecords: [][]string{
				{"Date", "Restaurant Name", "Amount"},
				{"2024-01-01", "Foo Bar", "123.45"},
				{"2024-01-02", "Baz Qux", "67.89"},
			},
		},
		{
			name:        "empty orders",
			orders:      []Order{},
			wantRecords: [][]string{{"Date", "Restaurant Name", "Amount"}},
		},
		{
			name:    "invalid path",
			orders:  []Order{{Date: "2024-01-01", Restaurant: "Foo", Amount: "10.00"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var path string
			if tt.wantErr {
				path = filepath.Join(t.TempDir(), "nonexistent", "out.csv")
			} else {
				path = filepath.Join(t.TempDir(), "out.csv")
			}

			err := WriteCSV(tt.orders, path)
			if (err != nil) != tt.wantErr {
				t.Fatalf("WriteCSV() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			f, err := os.Open(path)
			if err != nil {
				t.Fatalf("failed to open output file: %v", err)
			}
			defer func(f *os.File) {
				_ = f.Close()
			}(f)

			got, err := csv.NewReader(f).ReadAll()
			if err != nil {
				t.Fatalf("failed to read output CSV: %v", err)
			}

			if len(got) != len(tt.wantRecords) {
				t.Fatalf("record count mismatch: got %d, want %d", len(got), len(tt.wantRecords))
			}

			for i := range got {
				if len(got[i]) != len(tt.wantRecords[i]) {
					t.Fatalf("field count mismatch at row %d: got %d, want %d", i, len(got[i]), len(tt.wantRecords[i]))
				}

				for j := range got[i] {
					if got[i][j] != tt.wantRecords[i][j] {
						t.Errorf("row %d field %d mismatch: got %q, want %q", i, j, got[i][j], tt.wantRecords[i][j])
					}
				}
			}
		})
	}
}
