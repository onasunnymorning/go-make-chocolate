package rest

import (
	"strconv"

	gin "github.com/gin-gonic/gin"
	service "github.com/onasunnymorning/go-make-chocolate/internal/service"
	recipe "github.com/onasunnymorning/go-make-chocolate/pkg/recipe"
)

// RecipeController handles HTTP requests related to recipes
type RecipeController struct {
	recipeService service.RecipeService
}

// NewRecipeController creates a new instance of RecipeController
func NewRecipeController(recipeService service.RecipeService) *RecipeController {
	return &RecipeController{
		recipeService: recipeService,
	}
}

// RecipeRequest represents the request body for creating/updating a recipe
type RecipeRequest struct {
	Name         string              `json:"name" binding:"required"`
	Description  string              `json:"description"`
	Ingredients  []recipe.Ingredient `json:"ingredients" binding:"required,dive"`
	Instructions string              `json:"instructions" binding:"required"`
}

// GetRecipyByID godoc
// @Summary Get a Recipe by ID
// @Description Get a Recipe by ID
// @Tags recipes
// @Produce json
// @Param id path string true "Recipe ID"
// @Success 200 {object} recipe.Recipe
// @Failure 404
// @Failure 500
func (rc *RecipeController) GetRecipeByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	recipe, err := rc.recipeService.GetByID(ctx, id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Recipe not found"})
		return
	}

	ctx.JSON(200, recipe)
}

// GetRecipeTemplate godoc
// @Summary Get a Recipe and return a template
// @Description Get a Recipe and return a template
// @Tags recipes
// @Produce json
// @Param id path string true "Recipe ID"
// @Success 200 {object} recipe.TemplateRecipe
// @Failure 404
// @Failure 500
// @Router /{id}/template [get]
func (rc *RecipeController) GetRecipeTemplate(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	template, err := rc.recipeService.GetTemplateByID(ctx, id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Recipe not found"})
		return
	}

	ctx.JSON(200, template)
}

// CreateRecipe godoc
// @Summary Create a new Recipe
// @Description Create a new Recipe
// @Tags recipes
// @Accept json
// @Produce json
// @Param recipe body RecipeRequest true "Recipe Request"
// @Success 201 {object} recipe.Recipe
// @Failure 400
// @Failure 500
func (rc *RecipeController) CreateRecipe(ctx *gin.Context) {
	var req RecipeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	recipe := &recipe.Recipe{
		Name:         req.Name,
		Description:  req.Description,
		Ingredients:  req.Ingredients,
		Instructions: req.Instructions,
	}

	createdRecipe, err := rc.recipeService.Create(ctx, recipe)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, createdRecipe)
}

// UpdateRecipe godoc
// @Summary Update a Recipe
// @Description Update a Recipe
// @Tags recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Param recipe body RecipeRequest true "Recipe Request"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
func (rc *RecipeController) UpdateRecipe(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	var req RecipeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	recipe := &recipe.Recipe{
		ID:           id,
		Name:         req.Name,
		Description:  req.Description,
		Ingredients:  req.Ingredients,
		Instructions: req.Instructions,
	}

	if err := rc.recipeService.Update(ctx, recipe); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(204)
}

// DeleteRecipe godoc
// @Summary Delete a Recipe
// @Description Delete a Recipe
// @Tags recipes
// @Param id path string true "Recipe ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
func (rc *RecipeController) DeleteRecipe(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{"error": "ID is required"})
		return
	}

	if err := rc.recipeService.Delete(ctx, id); err != nil {
		ctx.JSON(404, gin.H{"error": "Recipe not found"})
		return
	}

	ctx.Status(204)
}

// ListRecipes godoc
// @Summary List Recipes
// @Description List Recipes with pagination
// @Tags recipes
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} recipe.Recipe
// @Failure 500
func (rc *RecipeController) ListRecipes(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 0
	}

	recipes, err := rc.recipeService.List(ctx, limit, offset)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, recipes)
}

// CountRecipes godoc
// @Summary CountRecipes Recipes
// @Description CountRecipes total number of Recipes
// @Tags recipes
// @Produce json
// @Success 200 {object} map[string]int64
// @Failure 500
func (rc *RecipeController) CountRecipes(ctx *gin.Context) {
	count, err := rc.recipeService.Count(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"count": count})
}
