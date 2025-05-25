// @title Recipe API
// @version 1.0
// @description Recipe management API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /recipe
// @schemes http

// RecipeRequest represents the request body for creating/updating a recipe
// @name RecipeRequest
// @description Request body for creating or updating a recipe
// @param name string true "Recipe name"
// @param description string false "Recipe description"
// @param ingredients []Ingredient true "List of ingredients"
// @param instructions string true "Recipe instructions"
// @example name "Chocolate Cake"
// @example description "A delicious chocolate cake recipe"
// @example instructions "1. Preheat oven to 350°F\n2. Mix ingredients\n3. Bake for 30 minutes"
// @example ingredients [{"name": "flour", "quantity": {"amount": 2, "unit": "cup"}}, {"name": "sugar", "quantity": {"amount": 1.5, "unit": "cup"}}]
package main

// @Summary Create a new Recipe
// @Description Create a new recipe with the provided details
// @Tags recipes
// @Accept json
// @Produce json
// @Param recipe body RecipeRequest true "Recipe details"
// @Success 201 {object} recipe.Recipe
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router / [post]
func createRecipe() {}

// @Summary Get a Recipe by ID
// @Description Get a recipe by its unique ID
// @Tags recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Success 200 {object} recipe.Recipe
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /{id} [get]
func getRecipeByID() {}

// @Summary Update a Recipe
// @Description Update an existing recipe
// @Tags recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Param recipe body RecipeRequest true "Recipe details"
// @Success 204
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /{id} [put]
func updateRecipe() {}

// @Summary Delete a Recipe
// @Description Delete a recipe by its ID
// @Tags recipes
// @Accept json
// @Produce json
// @Param id path string true "Recipe ID"
// @Success 204
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /{id} [delete]
func deleteRecipe() {}

// @Summary List Recipes
// @Description Get a list of all recipes with pagination
// @Tags recipes
// @Accept json
// @Produce json
// @Param limit query int false "Number of items per page"
// @Param offset query int false "Offset for pagination"
// @Success 200 {array} recipe.Recipe
// @Failure 500 {object} string
// @Router / [get]
func listRecipes() {}

// @Summary Count Recipes
// @Description Get the total number of recipes
// @Tags recipes
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /count [get]
func countRecipes() {}
