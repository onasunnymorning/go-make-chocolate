package recipe

import "time"

// Error constants for recipe validation
var (
	ErrNameRequired        = &Error{"name_required", "Recipe name is required"}
	ErrIngredientsRequired = &Error{"ingredients_required", "At least one ingredient is required"}
	ErrInstructionsRequired = &Error{"instructions_required", "Recipe instructions are required"}
)

// Error represents a recipe-specific error
type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

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
