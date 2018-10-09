package set

// Package Set Interface
type Interface interface {
	// Adds the specified element to the set and returns true if the operation was successful or false
	// or false if the element already exists
	Add(v interface{}) bool

	// Removes the specified element from this set if it is present
	Remove(v interface{}) bool

	// Checks weather the element exists in the Set
	IsElementOf(v interface{}) bool

	// Get the size of the Set
	Size() int
}

// Optional types
type Emptier interface {
	// Removes all of the elements from a collection.
	// The collection will be empty after this call returns
	Empty()
}