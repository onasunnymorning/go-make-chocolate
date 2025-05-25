package recipe

import (
	"fmt"
	"strings"
)

// Quantity defines a numeric amount and a unit.
type Quantity struct {
	Amount float64
	Unit   Unit
}

// String returns a human-readable representation of the quantity.
func (q Quantity) String() string {
	return fmt.Sprintf("%g %s", q.Amount, q.Unit)
}

// ParseQuantity creates a Quantity from a string like "1.5 kg".
func ParseQuantity(s string) (Quantity, error) {
	parts := strings.Fields(s)
	if len(parts) != 2 {
		return Quantity{}, fmt.Errorf("invalid format, expected '<amount> <unit>'")
	}
	var q Quantity
	_, err := fmt.Sscanf(parts[0], "%f", &q.Amount)
	if err != nil {
		return Quantity{}, fmt.Errorf("invalid amount: %v", err)
	}
	q.Unit = Unit(parts[1])
	return q, nil
}

// SupportedUnits returns the list of all defined units.
func SupportedUnits() []Unit {
	return []Unit{
		Gram,
	}
}
