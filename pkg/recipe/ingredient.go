package recipe

// Ingredient represents a single ingredient and its quantity.
type Ingredient struct {
	Name     string
	IsCacao  bool // Indicates if the ingredient is cacao, user for determining cacao percentage
	Quantity Quantity
}
