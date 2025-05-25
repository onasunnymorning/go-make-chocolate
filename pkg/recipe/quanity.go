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

// ConvertTo converts the quantity to a different unit, if conversion is supported.
func (q Quantity) ConvertTo(target Unit) (Quantity, error) {
	switch {
	case q.Unit == Gram && target == Kilogram:
		return Quantity{Amount: q.Amount / 1000, Unit: target}, nil
	case q.Unit == Kilogram && target == Gram:
		return Quantity{Amount: q.Amount * 1000, Unit: target}, nil
	default:
		return Quantity{}, fmt.Errorf("conversion from %s to %s not supported", q.Unit, target)
	}
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
		Gram, Kilogram,
	}
}
