package discord

func help() string {
	return `
	I have following commands:
		**!addproduct** - add a product to monitor, takes:
			Name - whatever you want,
			Link - link to the product,
			Shop URL - URL of the shop(from https:// to first '/'),
			Photo URL - URL of the product's photo
	`
}

/*func addproduct(name string, link string, shopURL string, photoURL string) string {
	response := AC.AddProduct(*pp.NewProduct(name, link, shopURL, photoURL))
	return response
}*/
