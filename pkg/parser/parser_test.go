package parser

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParsePDF(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		setup   func(t *testing.T) string
		wantErr bool
	}{
		{
			name: "non-existent path",
			setup: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "does-not-exist.pdf")
			},
			wantErr: true,
		},
		{
			name: "non-pdf file",
			setup: func(t *testing.T) string {
				p := filepath.Join(t.TempDir(), "not-a-pdf.txt")
				if err := os.WriteFile(p, []byte("hello world"), 0644); err != nil {
					t.Fatalf("failed to write temp file: %v", err)
				}

				return p
			},
			wantErr: true,
		},
		{
			name: "empty file",
			setup: func(t *testing.T) string {
				p := filepath.Join(t.TempDir(), "empty.pdf")
				if err := os.WriteFile(p, []byte{}, 0644); err != nil {
					t.Fatalf("failed to write temp file: %v", err)
				}

				return p
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			path := tt.setup(t)

			_, err := ParsePDF(path)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParsePDF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParseOrderText(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		text string
		want []Order
	}{
		{
			name: "multiple orders",
			text: "01-05-20261234567890123451947 Restaurant         ₹359.00\n" +
				"03-05-2026123456789012346Wholesome Bowlsome by Potful₹270.00\n" +
				"03-05-20261234567890123471947 Restaurant         ₹516.00",
			want: []Order{
				{Date: "01-05-2026", Restaurant: "1947 Restaurant", Amount: "₹359.00"},
				{Date: "03-05-2026", Restaurant: "Wholesome Bowlsome by Potful", Amount: "₹270.00"},
				{Date: "03-05-2026", Restaurant: "1947 Restaurant", Amount: "₹516.00"},
			},
		},
		{
			name: "single order",
			text: "22-05-2026123456789012348Growgreenherbs Modern RetailsPvt Ltd - Bengaluru Badootaa₹336.00",
			want: []Order{
				{Date: "22-05-2026", Restaurant: "Growgreenherbs Modern RetailsPvt Ltd - Bengaluru Badootaa", Amount: "₹336.00"},
			},
		},
		{
			name: "no matches",
			text: "This is just some random text without any order data.",
			want: nil,
		},
		{
			name: "whitespace trimming",
			text: "01-05-2026123456789012349   Some Restaurant Name   ₹100.00",
			want: []Order{
				{Date: "01-05-2026", Restaurant: "Some Restaurant Name", Amount: "₹100.00"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := parseOrderText(tt.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOrderText() = %v, want %v", got, tt.want)
			}
		})
	}
}
