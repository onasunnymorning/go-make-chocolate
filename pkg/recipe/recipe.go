package recipe

import "time"

// Error constants for recipe validation
var (
	ErrNameRequired         = &Error{"name_required", "Recipe name is required"}
	ErrIngredientsRequired  = &Error{"ingredients_required", "At least one ingredient is required"}
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
	ID              string
	Name            string
	Description     string
	Ingredients     []Ingredient
	Instructions    string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       string
	UpdatedBy       string
	CacaoPercentage float64 // Cacao percentage of the recipe, calculated from ingredients
}

// NewRecipe creates a new Recipe instance with the provided name, description, and ingredients. Cacao percentage is calculated automatically.
func NewRecipe(name, description string, ingredients []Ingredient) (*Recipe, error) {
	if name == "" {
		return nil, ErrNameRequired
	}
	if len(ingredients) == 0 {
		return nil, ErrIngredientsRequired
	}

	rcp := &Recipe{
		Name:        name,
		Description: description,
		Ingredients: ingredients,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rcp.CacaoPercentage = rcp.cacaoPercentage()

	return rcp, nil
}

// calculateCacaoPercentage calculates the cacao percentage of the recipe based on its ingredients.
func (r *Recipe) cacaoPercentage() float64 {
	var totalQuantity, cacaoQuantity float64
	for _, ingredient := range r.Ingredients {
		totalQuantity += ingredient.Quantity.Amount
		if ingredient.IsCacao {
			cacaoQuantity += ingredient.Quantity.Amount
		}
	}
	if totalQuantity == 0 {
		return 0 // Avoid division by zero
	}
	return (cacaoQuantity / totalQuantity) * 100
}
