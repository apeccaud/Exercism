package twofer

// ShareWith returns the following string : "One for X, one for me."
// where X is the value of the "name" parameter or "you" if no name 
// is provided
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
