package mongo

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/onasunnymorning/go-make-chocolate/pkg/recipe"
)

func TestConversions(t *testing.T) {
	// Setup test data
	id := primitive.NewObjectID()
	testTime := time.Date(2025, 5, 24, 13, 40, 17, 0, time.UTC)
		ingredients := []recipe.Ingredient{
		{
			Name: "Sugar",
			Quantity: recipe.Quantity{
				Amount: 100,
				Unit:   recipe.Gram,
			},
		},
		{
			Name: "Water",
			Quantity: recipe.Quantity{
				Amount: 500,
				Unit:   recipe.Milliliter,
			},
		},
	}

	// Create domain recipe
	domainRecipe := &recipe.Recipe{
		ID:           id.Hex(),
		Name:         "Test Recipe",
		Description:  "A test recipe",
		Ingredients:  ingredients,
		Instructions: "Mix ingredients",
		CreatedAt:    testTime,
		UpdatedAt:    testTime,
		CreatedBy:    "test_user",
		UpdatedBy:    "test_user",
	}

	// Convert to MongoDB doc
	doc := ToMongo(domainRecipe)

	// Verify MongoDB doc
	if doc.ID.Hex() != id.Hex() {
		t.Errorf("Expected ID %s, got %s", id.Hex(), doc.ID.Hex())
	}
	if doc.Name != domainRecipe.Name {
		t.Errorf("Expected Name %s, got %s", domainRecipe.Name, doc.Name)
	}
	if doc.Description != domainRecipe.Description {
		t.Errorf("Expected Description %s, got %s", domainRecipe.Description, doc.Description)
	}
	if len(doc.Ingredients) != len(domainRecipe.Ingredients) {
		t.Errorf("Expected %d ingredients, got %d", len(domainRecipe.Ingredients), len(doc.Ingredients))
	}
	if doc.Instructions != domainRecipe.Instructions {
		t.Errorf("Expected Instructions %s, got %s", domainRecipe.Instructions, doc.Instructions)
	}
	if doc.CreatedAt != domainRecipe.CreatedAt {
		t.Errorf("Expected CreatedAt %v, got %v", domainRecipe.CreatedAt, doc.CreatedAt)
	}
	if doc.UpdatedAt != domainRecipe.UpdatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", domainRecipe.UpdatedAt, doc.UpdatedAt)
	}
	if doc.CreatedBy != domainRecipe.CreatedBy {
		t.Errorf("Expected CreatedBy %s, got %s", domainRecipe.CreatedBy, doc.CreatedBy)
	}
	if doc.UpdatedBy != domainRecipe.UpdatedBy {
		t.Errorf("Expected UpdatedBy %s, got %s", domainRecipe.UpdatedBy, doc.UpdatedBy)
	}

	// Convert back to domain
	convertedRecipe := doc.ToDomain()

	// Verify domain conversion
	if convertedRecipe.ID != domainRecipe.ID {
		t.Errorf("Expected ID %s, got %s", domainRecipe.ID, convertedRecipe.ID)
	}
	if convertedRecipe.Name != domainRecipe.Name {
		t.Errorf("Expected Name %s, got %s", domainRecipe.Name, convertedRecipe.Name)
	}
	if convertedRecipe.Description != domainRecipe.Description {
		t.Errorf("Expected Description %s, got %s", domainRecipe.Description, convertedRecipe.Description)
	}
	if len(convertedRecipe.Ingredients) != len(domainRecipe.Ingredients) {
		t.Errorf("Expected %d ingredients, got %d", len(domainRecipe.Ingredients), len(convertedRecipe.Ingredients))
	}
	if convertedRecipe.Instructions != domainRecipe.Instructions {
		t.Errorf("Expected Instructions %s, got %s", domainRecipe.Instructions, convertedRecipe.Instructions)
	}
	if !convertedRecipe.CreatedAt.Equal(domainRecipe.CreatedAt) {
		t.Errorf("Expected CreatedAt %v, got %v", domainRecipe.CreatedAt, convertedRecipe.CreatedAt)
	}
	if !convertedRecipe.UpdatedAt.Equal(domainRecipe.UpdatedAt) {
		t.Errorf("Expected UpdatedAt %v, got %v", domainRecipe.UpdatedAt, convertedRecipe.UpdatedAt)
	}
	if convertedRecipe.CreatedBy != domainRecipe.CreatedBy {
		t.Errorf("Expected CreatedBy %s, got %s", domainRecipe.CreatedBy, convertedRecipe.CreatedBy)
	}
	if convertedRecipe.UpdatedBy != domainRecipe.UpdatedBy {
		t.Errorf("Expected UpdatedBy %s, got %s", domainRecipe.UpdatedBy, convertedRecipe.UpdatedBy)
	}

	// Test ingredient conversion
	ingredientDoc := IngredientDoc{
		Name: "Test Ingredient",
		Quantity: QuantityDoc{
			Amount: 100,
			Unit:   "g",
		},
	}

	convertedIngredient := toDomainIngredients([]IngredientDoc{ingredientDoc})[0]
	if convertedIngredient.Name != ingredientDoc.Name {
		t.Errorf("Expected ingredient name %s, got %s", ingredientDoc.Name, convertedIngredient.Name)
	}
	if convertedIngredient.Quantity.Amount != ingredientDoc.Quantity.Amount {
		t.Errorf("Expected quantity amount %f, got %f", ingredientDoc.Quantity.Amount, convertedIngredient.Quantity.Amount)
	}
	if convertedIngredient.Quantity.Unit != recipe.Gram {
		t.Errorf("Expected quantity unit %s, got %s", recipe.Gram, convertedIngredient.Quantity.Unit)
	}

	// Test quantity conversion
	quantityDoc := QuantityDoc{
		Amount: 200,
		Unit:   "ml",
	}

	convertedQuantity := toDomainQuantity(quantityDoc)
	if convertedQuantity.Amount != quantityDoc.Amount {
		t.Errorf("Expected quantity amount %f, got %f", quantityDoc.Amount, convertedQuantity.Amount)
	}
	if convertedQuantity.Unit != recipe.Milliliter {
		t.Errorf("Expected quantity unit %s, got %s", recipe.Milliliter, convertedQuantity.Unit)
	}

	// Test all unit conversions
	testUnits := []recipe.Unit{
		recipe.Gram,
		recipe.Kilogram,
		recipe.Milliliter,
		recipe.Liter,
		recipe.Cup,
		recipe.Tablespoon,
		recipe.Teaspoon,
		recipe.Piece,
	}

	for _, unit := range testUnits {
		quantity := recipe.Quantity{
			Amount: 100,
			Unit:   unit,
		}
		converted := toMongoQuantity(quantity)
		if converted.Amount != quantity.Amount {
			t.Errorf("Unit %s: Expected amount %f, got %f", unit, quantity.Amount, converted.Amount)
		}
		if converted.Unit != string(unit) {
			t.Errorf("Unit %s: Expected unit %s, got %s", unit, unit, converted.Unit)
		}
	}
}
