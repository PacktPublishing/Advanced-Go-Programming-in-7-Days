package set

type HashSet struct {
	items map[interface{}]status
}

type status bool
const statusExists status = true


// New instantiates a new empty set
func New() *HashSet {
	return &HashSet{items: make(map[interface{}]status)}
}

// Implements the Set.Add method
func (set *HashSet) Add(item interface{}) bool {
	if _, exists:= set.items[item]; exists {
		return false
	}
	//if set.IsElementOf(item) {
	//	return false
	//}
	set.items[item] = statusExists
	return true
}

// Implements the Set.Remove method
func (set *HashSet) Remove(item interface{}) bool {
	if _, exists:= set.items[item]; !exists {
		return false
	}
	delete(set.items, item)
	return true
}

// Implements the Set.Size method
func (set *HashSet) Size() int {
	return len(set.items)
}

// Implements the Set.IsElementOf method
func (set *HashSet) IsElementOf(item interface{}) bool {
	if _, exists := set.items[item]; !exists {
		return false
	} else {
		return true
	}
}

// Implements the Clearer.Empty method
func (set *HashSet) Empty() {
	set.items = make(map[interface{}]status)
}

// Values returns all items in the set as a slice
func (set *HashSet) Values() []interface{} {
	setValues := make([]interface{}, len(set.items), len(set.items))

	count := 0
	for item := range set.items {
		setValues[count] = item
		count++
	}

	return setValues
}