package recipe

import "strconv"

// TemplateRecipe is derrived from a recipe.Recipe and holds a list of ingredients and the percentage they each contribute to the recipe.
// It is used to reproduce a recipe while allowing the user to adjust the desired batch size.
type TemplateRecipe struct {
	RecipeID        string
	Name            string // Name of the recipe
	Description     string // Description of the recipe
	Ingredients     []TemplateIngredient
	Instructions    string  // Instructions for the recipe
	CacaoPercentage float64 // Cacao percentage of the recipe
}

// ToRecipe converts a TemplateRecipe to a Recipe with recalculated ingredient quantities based on the desired yield.
// The yield is specified in grams, and the ingredient quantities are adjusted accordingly.
// The recipe ID is preserved, and the name and description are modified to reflect the new yield.
// If the yield is zero or negative, it returns nil.
func (tr *TemplateRecipe) ToRecipe(yield float64) *Recipe {
	ingredients := make([]Ingredient, len(tr.Ingredients))
	for i, ing := range tr.Ingredients {
		quantity := ing.Percentage * yield / 100
		ingredients[i] = Ingredient{
			Name:     ing.Name,
			IsCacao:  ing.IsCacao,
			Quantity: Quantity{Unit: "grams", Amount: quantity},
		}
	}
	return &Recipe{
		ID:              tr.RecipeID,
		Name:            tr.Name + "-toYield" + strconv.FormatFloat(yield, 'f', 0, 64),
		Description:     tr.Description + " - with quantities recalculated for a batch size of " + strconv.FormatFloat(yield, 'f', 0, 64) + " grams",
		Ingredients:     ingredients,
		Instructions:    tr.Instructions,
		CacaoPercentage: tr.CacaoPercentage,
		Yield:           Quantity{Unit: "grams", Amount: yield},
	}
}
