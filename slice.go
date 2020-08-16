package alphanum

// StringSlice attaches a sort.Interface implementation to []string. It's similar to sort.StringSlice.
type StringSlice []string

// Len returns the length of the StringSlice.
func (p StringSlice) Len() int {
	return len(p)
}

// Less checks if the element at index i is less than the element at index j.
func (p StringSlice) Less(i, j int) bool {
	return Less(p[i], p[j])
}

// Swap swaps the elements at index i and index j.
func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
