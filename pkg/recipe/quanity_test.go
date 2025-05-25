package recipe

import (
	"fmt"
	"testing"
)

func TestQuantityString(t *testing.T) {
	q := Quantity{
		Amount: 1.5,
		Unit:   Gram,
	}
	want := fmt.Sprintf("%g %s", q.Amount, q.Unit)
	got := q.String()
	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestParseQuantity(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Quantity
		wantErr bool
	}{
		{
			name:  "Valid quantity",
			input: "1.5 kg",
			want:  Quantity{Amount: 1.5, Unit: Unit("kg")},
		},
		{
			name:    "Invalid format (missing space)",
			input:   "1.5kg",
			wantErr: true,
		},
		{
			name:    "Invalid amount",
			input:   "abc kg",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseQuantity(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseQuantity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if got.Amount != tt.want.Amount || got.Unit != tt.want.Unit {
					t.Errorf("ParseQuantity() = %+v, want %+v", got, tt.want)
				}
			}
		})
	}
}

func TestSupportedUnits(t *testing.T) {
	units := SupportedUnits()
	expected := []Unit{
		Gram,
	}
	if len(units) != len(expected) {
		t.Errorf("SupportedUnits() returned %d units, want %d", len(units), len(expected))
	}
	for i, u := range expected {
		if units[i] != u {
			t.Errorf("SupportedUnits()[%d] = %q, want %q", i, units[i], u)
		}
	}
}
