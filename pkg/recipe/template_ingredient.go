package recipe

// TemplateIngredient represents an ingredient in a template recipe.
type TemplateIngredient struct {
	Name       string  // Name of the ingredient
	IsCacao    bool    // Indicates if the ingredient is cacao
	Percentage float64 // Percentage of the ingredient in the recipe
}
