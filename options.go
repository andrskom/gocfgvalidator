package gocfgvalidtor

// MustWithDeepOfRecursion set deep of cfg recursion
// Rise panic if have problem
func MustWithDeepOfRecursion(deep int) optionFunc { // nolint , we need not exported result type
	if deep <= 0 {
		panic("deep must be grate than 0")
	}
	return func(c *Component) {
		c.deepOfRecursion = deep
	}
}

//WithStrictMode set strict mode for component
func WithStrictMode(enable bool) optionFunc { // nolint , we need not exported result type
	return func(c *Component) {
		c.strictMode = enable
	}
}
