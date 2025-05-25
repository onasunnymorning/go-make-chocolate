package recipe

// TemplateRecipe is derrived from a recipe.Recipe and holds a list of ingredients and the percentage they each contribute to the recipe.
// It is used to reproduce a recipe while allowing the user to adjust the desired batch size.
type TemplateRecipe struct {
	RecipeID    string
	Ingredients []TemplateIngredient
}
