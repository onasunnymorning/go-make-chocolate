package recipe

import (
	"fmt"
	"testing"
)

func TestQuantityString(t *testing.T) {
	q := Quantity{
		Amount: 1.5,
		Unit:   Kilogram,
	}
	want := fmt.Sprintf("%g %s", q.Amount, q.Unit)
	got := q.String()
	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestConvertTo(t *testing.T) {
	tests := []struct {
		name    string
		q       Quantity
		target  Unit
		want    Quantity
		wantErr bool
	}{
		{
			name:   "Gram to Kilogram",
			q:      Quantity{Amount: 1500, Unit: Gram},
			target: Kilogram,
			want:   Quantity{Amount: 1.5, Unit: Kilogram},
		},
		{
			name:   "Kilogram to Gram",
			q:      Quantity{Amount: 2, Unit: Kilogram},
			target: Gram,
			want:   Quantity{Amount: 2000, Unit: Gram},
		},
		{
			name:    "Unsupported conversion",
			q:       Quantity{Amount: 1, Unit: "blbal"},
			target:  Gram,
			want:    Quantity{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.ConvertTo(tt.target)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ConvertTo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if got.Amount != tt.want.Amount || got.Unit != tt.want.Unit {
					t.Errorf("ConvertTo() = %+v, want %+v", got, tt.want)
				}
			}
		})
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
		Gram, Kilogram,
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
