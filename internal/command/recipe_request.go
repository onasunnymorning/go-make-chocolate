package command

import "github.com/onasunnymorning/go-make-chocolate/pkg/recipe"

// RecipeRequest represents the request body for creating/updating a recipe
type RecipeRequest struct {
	Name         string              `json:"name" binding:"required"`
	Description  string              `json:"description"`
	Ingredients  []recipe.Ingredient `json:"ingredients" binding:"required,dive"`
	Instructions string              `json:"instructions" binding:"required"`
}
