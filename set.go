package datastructures

type Set struct {
	values      map[interface{}]struct{}
	cardinality int
}

// Creates a new set.
func NewSet() *Set {
	return &Set{
		map[interface{}]struct{}{},
		0,
	}
}

// Add a new value to the set.
func (s *Set) Add(value interface{}) {
	_, exists := (*s).values[value]
	if !exists {
		(*s).values[value] = struct{}{}
		(*s).cardinality++
	}
}

// Check if the set contains a value.
func (s *Set) Contains(value interface{}) bool {
	_, exists := (*s).values[value]
	return exists
}

// Returns the current number of items in the set.
func (s *Set) Length() int {
	return (*s).cardinality
}

// Provides a channel for iterating over all values.
func (s *Set) Iterate() <-chan interface{} {
	ch := make(chan interface{}, (*s).cardinality)
	go func() {
		for k, _ := range (*s).values {
			ch <- k
		}
		close(ch)
	}()
	return ch
}

// Combines two sets into a new.
func (s *Set) Combine(o *Set) *Set {
	n := NewSet()

	// TODO: Avoid iteration by copy?
	// Benchmark: How is branching in Add
	// 			  affected when iterating?
	for v := range s.Iterate() {
		n.Add(v)
	}

	for v := range o.Iterate() {
		n.Add(v)
	}

	return n
}

func (s *Set) Array() []string {
	// TODO: Check if we can pass cardinality to capacity
	result := make([]string, 0)
	for v := range s.Iterate() {
		if str, ok := v.(string); ok {
			result = append(result, str)
		}
	}
	return result
}
