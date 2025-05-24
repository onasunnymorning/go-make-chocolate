package mongo

import (
	"time"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/onasunnymorning/go-make-chocolate/pkg/recipe"
)

// RecipeDoc represents a recipe document in MongoDB
type RecipeDoc struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	Ingredients  []IngredientDoc    `bson:"ingredients"`
	Instructions string             `bson:"instructions"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	CreatedBy    string             `bson:"created_by"`
	UpdatedBy    string             `bson:"updated_by"`
}

// IngredientDoc represents an ingredient document in MongoDB
type IngredientDoc struct {
	Name     string      `bson:"name"`
	Quantity QuantityDoc `bson:"quantity"`
}

// QuantityDoc represents a quantity document in MongoDB
type QuantityDoc struct {
	Amount float64 `bson:"amount"`
	Unit   string  `bson:"unit"`
}

// ToDomain converts a MongoDB document to a domain model
func (r *RecipeDoc) ToDomain() *recipe.Recipe {
	return &recipe.Recipe{
		ID:           r.ID.Hex(),
		Name:         r.Name,
		Description:  r.Description,
		Ingredients:  toDomainIngredients(r.Ingredients),
		Instructions: r.Instructions,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
		CreatedBy:    r.CreatedBy,
		UpdatedBy:    r.UpdatedBy,
	}
}

// ToMongo converts a domain model to a MongoDB document
func ToMongo(r *recipe.Recipe) *RecipeDoc {
	id, _ := primitive.ObjectIDFromHex(r.ID)
	return &RecipeDoc{
		ID:           id,
		Name:         r.Name,
		Description:  r.Description,
		Ingredients:  toMongoIngredients(r.Ingredients),
		Instructions: r.Instructions,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
		CreatedBy:    r.CreatedBy,
		UpdatedBy:    r.UpdatedBy,
	}
}

func toDomainIngredients(docs []IngredientDoc) []recipe.Ingredient {
	ingredients := make([]recipe.Ingredient, len(docs))
	for i, doc := range docs {
		ingredients[i] = recipe.Ingredient{
			Name:     doc.Name,
			Quantity: toDomainQuantity(doc.Quantity),
		}
	}
	return ingredients
}

func toMongoIngredients(ingredients []recipe.Ingredient) []IngredientDoc {
	docs := make([]IngredientDoc, len(ingredients))
	for i, ing := range ingredients {
		docs[i] = IngredientDoc{
			Name:     ing.Name,
			Quantity: toMongoQuantity(ing.Quantity),
		}
	}
	return docs
}

func toDomainQuantity(doc QuantityDoc) recipe.Quantity {
	return recipe.Quantity{
		Amount: doc.Amount,
		Unit:   recipe.Unit(doc.Unit),
	}
}

func toMongoQuantity(q recipe.Quantity) QuantityDoc {
	return QuantityDoc{
		Amount: q.Amount,
		Unit:   string(q.Unit),
	}
}
