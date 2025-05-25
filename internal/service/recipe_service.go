package service

import (
	"context"
	"time"

	"github.com/onasunnymorning/go-make-chocolate/internal/infra/db/mongo"
	"github.com/onasunnymorning/go-make-chocolate/pkg/recipe"
)

// RecipeService defines the contract for recipe operations
type RecipeService interface {
	Create(ctx context.Context, recipe *recipe.Recipe) (*recipe.Recipe, error)
	GetByID(ctx context.Context, id string) (*recipe.Recipe, error)
	GetTemplateByID(ctx context.Context, id string) (*recipe.TemplateRecipe, error)
	Update(ctx context.Context, recipe *recipe.Recipe) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int64) ([]*recipe.Recipe, error)
	Count(ctx context.Context) (int64, error)
}

// recipeService implements the RecipeService interface
type recipeService struct {
	store mongo.RecipeStore
}

// NewRecipeService creates a new RecipeService
func NewRecipeService(store mongo.RecipeStore) *recipeService {
	return &recipeService{
		store: store,
	}
}

// Create creates a new recipe
func (s *recipeService) Create(ctx context.Context, rcp *recipe.Recipe) (*recipe.Recipe, error) {
	if rcp.Name == "" {
		return nil, recipe.ErrNameRequired
	}
	if len(rcp.Ingredients) == 0 {
		return nil, recipe.ErrIngredientsRequired
	}
	if rcp.Instructions == "" {
		return nil, recipe.ErrInstructionsRequired
	}

	rcp.CreatedAt = time.Now()
	rcp.UpdatedAt = time.Now()

	newRecipe, err := recipe.NewRecipe(rcp.Name, rcp.Description, rcp.Ingredients)
	if err != nil {
		return nil, err
	}

	return s.store.Create(ctx, newRecipe)
}

// GetByID retrieves a recipe by its ID
func (s *recipeService) GetByID(ctx context.Context, id string) (*recipe.Recipe, error) {
	return s.store.GetByID(ctx, id)
}

// GetTemplate retrieves a recipe and returns it as a template
func (s *recipeService) GetTemplateByID(ctx context.Context, id string) (*recipe.TemplateRecipe, error) {
	rcp, err := s.store.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	template := rcp.ToTemplate()

	return template, nil
}

// Update updates an existing recipe
func (s *recipeService) Update(ctx context.Context, rcp *recipe.Recipe) error {
	if rcp.Name == "" {
		return recipe.ErrNameRequired
	}
	if len(rcp.Ingredients) == 0 {
		return recipe.ErrIngredientsRequired
	}
	if rcp.Instructions == "" {
		return recipe.ErrInstructionsRequired
	}

	rcp.UpdatedAt = time.Now()

	return s.store.Update(ctx, rcp)
}

// Delete removes a recipe by its ID
func (s *recipeService) Delete(ctx context.Context, id string) error {
	return s.store.Delete(ctx, id)
}

// List retrieves recipes with pagination
func (s *recipeService) List(ctx context.Context, limit, offset int64) ([]*recipe.Recipe, error) {
	return s.store.List(ctx, limit, offset)
}

// Count returns the total number of recipes
func (s *recipeService) Count(ctx context.Context) (int64, error) {
	return s.store.Count(ctx)
}
