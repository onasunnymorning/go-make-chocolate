package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/onasunnymorning/go-make-chocolate/pkg/recipe"
)

// RecipeStore defines the interface for recipe database operations
type RecipeStore interface {
	Create(ctx context.Context, recipe *recipe.Recipe) (*recipe.Recipe, error)
	GetByID(ctx context.Context, id string) (*recipe.Recipe, error)
	Update(ctx context.Context, recipe *recipe.Recipe) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int64) ([]*recipe.Recipe, error)
	Count(ctx context.Context) (int64, error)
}

// MongoDBRecipeStore implements the RecipeStore interface using MongoDB
type MongoDBRecipeStore struct {
	collection *mongo.Collection
}

// NewMongoDBRecipeStore creates a new MongoDBRecipeStore
func NewMongoDBRecipeStore(db *mongo.Database) *MongoDBRecipeStore {
	return &MongoDBRecipeStore{
		collection: db.Collection("recipes"),
	}
}

// Create inserts a new recipe into the database
func (s *MongoDBRecipeStore) Create(ctx context.Context, recipe *recipe.Recipe) (*recipe.Recipe, error) {
	if recipe.ID == "" {
		recipe.ID = primitive.NewObjectID().Hex()
	}
	recipe.CreatedAt = time.Now()
	recipe.UpdatedAt = time.Now()

	doc := ToMongo(recipe)
	_, err := s.collection.InsertOne(ctx, doc)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

// GetByID retrieves a recipe by its ID
func (s *MongoDBRecipeStore) GetByID(ctx context.Context, id string) (*recipe.Recipe, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var doc RecipeDoc
	err = s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return doc.ToDomain(), nil
}

// Update updates an existing recipe
func (s *MongoDBRecipeStore) Update(ctx context.Context, recipe *recipe.Recipe) error {
	oid, err := primitive.ObjectIDFromHex(recipe.ID)
	if err != nil {
		return err
	}

	recipe.UpdatedAt = time.Now()
	doc := ToMongo(recipe)

	_, err = s.collection.ReplaceOne(ctx, bson.M{"_id": oid}, doc)
	return err
}

// Delete removes a recipe by its ID
func (s *MongoDBRecipeStore) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

// List retrieves recipes with pagination
func (s *MongoDBRecipeStore) List(ctx context.Context, limit, offset int64) ([]*recipe.Recipe, error) {
	cursor, err := s.collection.Find(ctx, bson.M{},
		options.Find().SetLimit(limit).SetSkip(offset))
	if err != nil {
		return nil, err
	}
	docs := make([]*RecipeDoc, 0)
	if err = cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	recipes := make([]*recipe.Recipe, len(docs))
	for i, doc := range docs {
		recipes[i] = doc.ToDomain()
	}

	return recipes, nil
}

// Count returns the total number of recipes
func (s *MongoDBRecipeStore) Count(ctx context.Context) (int64, error) {
	return s.collection.CountDocuments(ctx, bson.M{})
}
