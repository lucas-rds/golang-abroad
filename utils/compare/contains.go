package compare

// StringsContains checks if a list of string contains a particular string
func StringsContains(values []string, query string) bool {
	for _, value := range values {
		if value == query {
			return true
		}
	}
	return false
}
