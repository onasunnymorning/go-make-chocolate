package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/onasunnymorning/go-make-chocolate/internal/infra/db/mongo"
	"github.com/onasunnymorning/go-make-chocolate/internal/interface/rest"
	"github.com/onasunnymorning/go-make-chocolate/internal/service"
	"github.com/onasunnymorning/go-make-chocolate/pkg/recipe"

	_ "github.com/onasunnymorning/go-make-chocolate/cmd/recipe_api/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type RecipeRequest struct {
	Name         string              `json:"name" binding:"required"`
	Description  string              `json:"description"`
	Ingredients  []recipe.Ingredient `json:"ingredients" binding:"required,dive"`
	Instructions string              `json:"instructions" binding:"required"`
}

// @title           Recipee API
// @version         0.1
// @description     Manage Recipes.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /recipe

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Create a new Gin router
	r := gin.Default()

	mongoClient, err := mongo.NewClient(os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(context.Background())

	// Initialize recipe store and service
	db := mongoClient.Database("recipe_db")
	recipeStore := mongo.NewMongoDBRecipeStore(db)
	recipeService := service.NewRecipeService(recipeStore)
	recipeController := rest.NewRecipeController(recipeService)

	// Add a health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Recipe API is running",
		})
	})

	// Serve the swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DocExpansion("none"))) // collapse all endpoints by default

	// Recipe endpoints
	recipeGroup := r.Group("/recipe")
	{
		// Create a new recipe
		recipeGroup.POST("", recipeController.CreateRecipe)
		// Get recipe by ID
		recipeGroup.GET(":id", recipeController.GetRecipeByID)
		// Update recipe
		recipeGroup.PUT(":id", recipeController.UpdateRecipe)
		// Delete recipe
		recipeGroup.DELETE(":id", recipeController.DeleteRecipe)
		// List recipes
		recipeGroup.GET("", recipeController.ListRecipes)
		// Get recipe count
		recipeGroup.GET("/count", recipeController.CountRecipes)
	}

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
