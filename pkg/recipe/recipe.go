package recipe

import "time"

// Recipe represents a complete recipe.
type Recipe struct {
	ID           string
	Name         string
	Description  string
	Ingredients  []Ingredient
	Instructions string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreatedBy    string
	UpdatedBy    string
}
