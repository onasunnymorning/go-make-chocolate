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
	CacaoPercentage float64  // Cacao percentage of the recipe, calculated from ingredients
	Yield           Quantity // Batch size or yield of the recipe
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

	rcp.CacaoPercentage = rcp.CalculateCacaoPercentage()
	rcp.Yield = rcp.CalculateYield()

	return rcp, nil
}

// CalculateYield calculates the yield in grams of the recipe based on its ingredients.
func (r *Recipe) CalculateYield() Quantity {
	var totalQuantity float64
	for _, ingredient := range r.Ingredients {
		totalQuantity += ingredient.Quantity.Amount
	}
	return Quantity{
		Unit:   "grams",
		Amount: totalQuantity,
	}
}

// calculateCacaoPercentage calculates the cacao percentage of the recipe based on its ingredients.
func (r *Recipe) CalculateCacaoPercentage() float64 {
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

// ToTemplate converts the Recipe to a TemplateRecipe, which is used for creating new recipes based on templates.
// It calculates the percentages of each ingredient and returns a TemplateRecipe instance.
func (r *Recipe) ToTemplate() *TemplateRecipe {
	totalQuantity := 0.0
	for _, ingredient := range r.Ingredients {
		totalQuantity += ingredient.Quantity.Amount
	}
	ingredients := make([]TemplateIngredient, len(r.Ingredients))
	for i, ingredient := range r.Ingredients {
		percentage := 0.0
		if totalQuantity > 0 {
			percentage = (ingredient.Quantity.Amount / totalQuantity) * 100
		}
		ingredients[i] = TemplateIngredient{
			Name:       ingredient.Name,
			IsCacao:    ingredient.IsCacao,
			Percentage: percentage,
		}
	}
	return &TemplateRecipe{
		RecipeID:        r.ID,
		Ingredients:     ingredients,
		CacaoPercentage: r.CacaoPercentage,
	}

}
