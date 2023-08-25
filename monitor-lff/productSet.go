package lff

// productSet stores unique products identified by their link
type productSet map[string]ProductWrapper

func newProductSet() productSet {
	return make(productSet)
}

// Add adds a new product
func (s *productSet) Add(p ProductWrapper) {
	(*s)[p.Link] = p
}

// Get returns product from productSet, returns false if not found
func (s *productSet) Get(link string) (p ProductWrapper, found bool) {
	p, found = (*s)[link]
	return p, found
}

// Slice returns all products stored in productSet
func (s *productSet) Slice() (products []ProductWrapper) {
	for _, p := range *s {
		products = append(products, p)
	}

	return products
}

// Unload returns all products from productSet end epties it
func (s *productSet) Unload() (products []ProductWrapper) {
	products = s.Slice()
	*s = newProductSet()

	return products
}
