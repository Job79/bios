package model

// Model that should be implemented by every struct
type Model interface {
	GetUID() string // Get the uid of the model, used to create maps and identify the struct
	SetID(int)      // Set the id, used by the attach function
	Verify() error  // Verify whether the model has a valid state
}
