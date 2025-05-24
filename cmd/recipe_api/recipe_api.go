package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onasunnymorning/go-make-chocolate/internal/infra/db/mongo"
	"github.com/onasunnymorning/go-make-chocolate/internal/service"
	"github.com/onasunnymorning/go-make-chocolate/pkg/recipe"

	docs "github.com/onasunnymorning/go-make-chocolate/cmd/recipe_api/docs" // Import docs pkg to be able to access docs.json https://github.com/swaggo/swag/issues/830#issuecomment-725587162
	swaggerFiles "github.com/swaggo/files"                                  // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"                              // gin-swagger middleware
)

// RecipeRequest represents the request body for creating/updating a recipe
type RecipeRequest struct {
	Name         string              `json:"name" binding:"required"`
	Description  string              `json:"description"`
	Ingredients  []recipe.Ingredient `json:"ingredients" binding:"required,dive"`
	Instructions string              `json:"instructions" binding:"required"`
}

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

	// Add a health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Recipe API is running",
		})
	})

	// Recipe endpoints
	recipeGroup := r.Group("/recipe")
	{
		// Create a new recipe
		recipeGroup.POST("", func(c *gin.Context) {
			var req RecipeRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			recipe := &recipe.Recipe{
				Name:         req.Name,
				Description:  req.Description,
				Ingredients:  req.Ingredients,
				Instructions: req.Instructions,
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			createdRecipe, err := recipeService.Create(ctx, recipe)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			c.JSON(201, createdRecipe)
		})

		// Get recipe by ID
		recipeGroup.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				c.JSON(400, gin.H{"error": "ID is required"})
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			recipe, err := recipeService.GetByID(ctx, id)
			if err != nil {
				c.JSON(404, gin.H{"error": "Recipe not found"})
				return
			}

			c.JSON(200, recipe)
		})

		// Update recipe
		recipeGroup.PUT("/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				c.JSON(400, gin.H{"error": "ID is required"})
				return
			}

			var req RecipeRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			recipe := &recipe.Recipe{
				ID:           id,
				Name:         req.Name,
				Description:  req.Description,
				Ingredients:  req.Ingredients,
				Instructions: req.Instructions,
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := recipeService.Update(ctx, recipe); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			c.Status(204)
		})

		// Delete recipe
		recipeGroup.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				c.JSON(400, gin.H{"error": "ID is required"})
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := recipeService.Delete(ctx, id); err != nil {
				c.JSON(404, gin.H{"error": "Recipe not found"})
				return
			}

			c.Status(204)
		})

		// List recipes
		recipeGroup.GET("", func(c *gin.Context) {
			limitStr := c.DefaultQuery("limit", "10")
			offsetStr := c.DefaultQuery("offset", "0")

			limit, err := strconv.ParseInt(limitStr, 10, 64)
			if err != nil {
				limit = 10
			}
			offset, err := strconv.ParseInt(offsetStr, 10, 64)
			if err != nil {
				offset = 0
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			recipes, err := recipeService.List(ctx, limit, offset)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, recipes)
		})

		// Get recipe count
		recipeGroup.GET("/count", func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			count, err := recipeService.Count(ctx)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			c.JSON(200, gin.H{"count": count})
		})
	}

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
