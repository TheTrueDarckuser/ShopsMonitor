package parser

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
// Updaters contains all the updaters
/*Updaters = []func() ([]*Product, error){
	ParseSivasdelcalzo, ParseSolebox, ParseNaked, ParseEnd, ParseNittyGritty, ParseWoodwood, ParseSlamjamsocialism,
	ParseSneakers76, Parse43einhalb, ParseAsphaltgold, ParseAwlab, ParseBouncewear, ParseChmielna20, ParseConsortium,
	ParseCruvoir, ParseDtlr, ParseFootshop, ParseGoodhoodstore, ParseRuvilla, ParseShinzo, ParseSneakerstuff,
	ParseSneakerstudio, ParseTitolo, ParseTresbien, ParseUrbanjungle, ParseVooberlin, ParseStarcowparis,
	Parse1stog, ParseSsense, ParseMrporter, ParseNetaporter, ParseNorsestore, ParseFootpatrol, ParseShiek,
	ParseCitygear, ParseBarneys, ParseAllikestore, ParseCaliroots, ParseWellgosh, ParseThenextdoor, ParseCornerstreet,
	ParseSneakersnstuff, ParseTriads, ParseUebervart, ParseUptherestore, ParseYcmc, ParseNowhere, ParseChicagocitysports,
	ParseBaitme, ParseDoubleclutch, ParseUrbanstaroma, ParseFeaturesneakerboutique, ParseStormfashion,
	ParseBasket4ballers, ParseThegoodlifespace, ParseJdsports, ParseTike, ParseSupplystore, ParseSelfridges,
	ParseBrownsfashion, ParseKicksstore, ParseSaksfifthavenue, ParseSil, ParseProdirectsoccer, ParseThebrokenarm,
	ParseTitan22, ParseNosstore, ParseOqium, ParseSize, ParseOffspring, ParseWalkinmycloset, ParseWiseboutique,
	ParseSummerstore, ParseHbx, ParseThewebster, ParsePatta, ParseAntonioli, ParseExcelsiormilano, ParseFootish,
	ParseNighshop, ParseNordstorm, ParseNtrstore, ParseOkini, ParseOverkillshop, ParseShoezgallery, ParseSpace23,
	ParseJimmyjazz, ParseKickzpremium, ParseYmeuniverse, ParseFootlocker, ParseFinishline, ParseFootdistrict,
}*/
)

// ParseSivasdelcalzo parses www.sivasdescalzo.com
func ParseSivasdelcalzo() (products []*Product, err error) {
	shopURL := "https://www.sivasdescalzo.com"
	doc, err := getDoc(shopURL+"/en/lifestyle/new-arrivals", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-list > li > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSolebox parses www.solebox.com
func ParseSolebox() (products []*Product, err error) {
	shopURL := "https://www.solebox.com"
	doc, err := getDoc(shopURL+"/en/New/", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.gridView > li > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div.gridPicture img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseNaked parses www.nakedcph.com
func ParseNaked() (products []*Product, err error) {
	shopURL := "https://www.nakedcph.com"
	doc, err := getDoc(shopURL+"/new-arrivals/s/6", true)
	if err != nil {
		return products, err
	}

	doc.Find("#products > div > a").Each(func(i int, s *goquery.Selection) {
		name := s.Find("div.card-block > h4 > span.product-name").Text()
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div.card-img-top > img").Attr("data-src")
		link = shopURL + link
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseEnd parses endclothing.com
func ParseEnd() (products []*Product, err error) {
	shopURL := "https://www.endclothing.com"
	doc, err := getDoc(shopURL+"/us/latest-products/new-this-week", true)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-items > div > div > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div > img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseNittyGritty parses nittygrittystore.com
func ParseNittyGritty() (products []*Product, err error) {
	shopURL := "https://www.nittygrittystore.com"
	doc, err := getDoc(shopURL+"/men/latest-men", false)
	if err != nil {
		return products, err
	}

	doc.Find("#filterItems > div > div > a.product-item__link").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("data-ecommerce-name")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseWoodwood parses woodwood.com
func ParseWoodwood() (products []*Product, err error) {
	shopURL := "https://www.woodwood.com"
	doc, err := getDoc(shopURL+"/men/new-arrivals", false)
	if err != nil {
		return products, err
	}

	doc.Find("#commodity-lister-list > li > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("span.list-commodity-image > span > img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSlamjamsocialism parses slamjamsocialism.com
func ParseSlamjamsocialism() (products []*Product, err error) {
	shopURL := "https://www.slamjamsocialism.com"
	doc, err := getDoc(shopURL+"/new-products", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.product_list > li > div.product-container").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("div.right-block > a").Attr("title")
		link, _ := s.Find("div.right-block > a").Attr("href")
		photoURL, _ := s.Find("div.left-block > div.product-image-container img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSneakers76 parses sneakers76.com
func ParseSneakers76() (products []*Product, err error) {
	shopURL := "https://www.sneakers76.com"
	doc, err := getDoc(shopURL+"/en/new-products", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product_list > div > div").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.right-block > div.product-meta > h5.name > a")
		name := t.Text()
		link, _ := t.Attr("href")
		photoURL, _ := s.Find("div.left-block > div.product-image-container > a.product_img_link > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// Parse43einhalb parses 43einhalb.com
func Parse43einhalb() (products []*Product, err error) {
	shopURL := "https://www.43einhalb.com"
	doc, err := getDoc(shopURL+"/en/new-arrivals", true)
	if err != nil {
		return products, err
	}

	doc.Find("ul.block-grid > li.item").Each(func(i int, s *goquery.Selection) {
		name := s.Find("div.pInfo > a.pName > span.productName").Text()
		link, _ := s.Find("div.pImageContainer > a").Attr("href")
		photoURL, _ := s.Find("div.pImageContainer > a > div.pImages > div > img").Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseAsphaltgold parses asphaltgold.de
func ParseAsphaltgold() (products []*Product, err error) {
	shopURL := "https://asphaltgold.de"
	doc, err := getDoc(shopURL+"/en/sneaker/new", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-grid > section.item > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("data-name")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseAwlab parses aw-lab.com
func ParseAwlab() (products []*Product, err error) {
	shopURL := "https://en.aw-lab.com"
	doc, err := getDoc(shopURL+"/shop/men/new-now/", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.products-grid > div").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.product-data-container > h2.product-name > a")
		name := t.Text()
		link, _ := t.Attr("href")
		photoURL, _ := s.Find("div.product-image-container > div.product-image > a > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseBouncewear parses bouncewear.com
func ParseBouncewear() (products []*Product, err error) {
	shopURL := "https://bouncewear.com"
	doc, err := getDoc(shopURL+"/category/schoenen", false)
	if err != nil {
		return products, err
	}

	doc.Find("#resultdiv > div > div").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.productImage > a")
		name := s.Find("div.productCaption > a > h5").Text()
		link, _ := t.Attr("href")
		photoURL, _ := t.Find("img").Attr("src")
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseChmielna20 parses chmielna20.pl
func ParseChmielna20() (products []*Product, err error) {
	shopURL := "https://chmielna20.pl"
	doc, err := getDoc(shopURL+"/en/menu/obuwie/meskie", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.row > div.products__item").Each(func(i int, s *goquery.Selection) {
		t := s.Find("a")
		name, _ := t.Attr("title")
		link, _ := t.Attr("href")
		photoURL, _ := t.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseConsortium parses consortium.co.uk
func ParseConsortium() (products []*Product, err error) {
	shopURL := "http://www.consortium.co.uk"
	doc, err := getDoc(shopURL+"/latest", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li").Each(func(i int, s *goquery.Selection) {
		t := s.Find("img")
		name, _ := t.Attr("title")
		link, _ := s.Find("div.product-titles > h2 > a").Attr("href")
		photoURL, _ := t.Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseCruvoir parses cruvoir.com
func ParseCruvoir() (products []*Product, err error) {
	shopURL := "https://cruvoir.com"
	doc, err := getDoc(shopURL+"/collections/mens-shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-grid > article.product").Each(func(i int, s *goquery.Selection) {
		name := s.Find("div.info > p.name").Text()
		link, _ := s.Find("a").Attr("href")
		photoURL, _ := s.Find("a > img").Attr("src")
		if !strings.Contains(link, "https://cruvoir.com") {
			link = shopURL + link
		}
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseDtlr parses dtlr.com
func ParseDtlr() (products []*Product, err error) {
	shopURL := "https://www.dtlr.com"
	doc, err := getDoc(shopURL+"/men/footwear/new.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("section.catalog > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("figure > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseFootshop parses footshop.com
func ParseFootshop() (products []*Product, err error) {
	shopURL := "https://www.footshop.eu"
	doc, err := getDoc(shopURL+"/en/5-mens-shoes/", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-list > div > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.product__img > img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("data-src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseGoodhoodstore parses goodhoodstore.com
func ParseGoodhoodstore() (products []*Product, err error) {
	shopURL := "https://goodhoodstore.com"
	doc, err := getDoc(shopURL+"/mens/all-mens-footwear", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.Products-List > li > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseRuvilla parses ruvilla.com
func ParseRuvilla() (products []*Product, err error) {
	shopURL := "https://www.ruvilla.com"
	doc, err := getDoc(shopURL+"/men/footwear/new.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("section.catalog > div.product > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("figure > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseShinzo parses shinzo.paris
func ParseShinzo() (products []*Product, err error) {
	shopURL := "https://www.shinzo.paris"
	doc, err := getDoc(shopURL+"/en/27-sportswear-shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-list-grid > div.product-item > div.product-inner > div.product-thumb > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSneakerstuff parses sneakersnstuff.com
func ParseSneakerstuff() (products []*Product, err error) {
	shopURL := "https://www.sneakersnstuff.com"
	doc, err := getDoc(shopURL+"/en/904/mens-sneakers", false)
	if err != nil {
		return products, err
	}

	doc.Find("#products > li").Each(func(i int, s *goquery.Selection) {
		name := s.Find("div.info > h4.name > a").Text()
		link, _ := s.Find("a").Attr("href")
		photoURL, _ := s.Find("a > img").Attr("src")
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSneakerstudio parses sneakersnstudio.com
func ParseSneakerstudio() (products []*Product, err error) {
	shopURL := "https://sneakerstudio.com"
	doc, err := getDoc(shopURL+"/eng_m_MEN-187.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("#search > div.product_wrapper > a").Each(func(i int, s *goquery.Selection) {
		name := ""
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("data-src")
		link = shopURL + link
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseTitolo parses en.titolo.ch
func ParseTitolo() (products []*Product, err error) {
	shopURL := "https://en.titolo.ch"
	doc, err := getDoc(shopURL+"/sneakers", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.columns > ul > li > div.list-inner-wrapper > div.img-wrapper > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseTresbien parses tres-bien.com
func ParseTresbien() (products []*Product, err error) {
	shopURL := "https://tres-bien.com"
	doc, err := getDoc(shopURL+"/new-footwear", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products > li > article > div.product-item-photo-wrapper > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseUrbanjungle parses urbanjunglestore.com
func ParseUrbanjungle() (products []*Product, err error) {
	shopURL := "https://www.urbanjunglestore.com"
	doc, err := getDoc(shopURL+"/it/latest-products.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > div.product-image-wrapper > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseVooberlin parses vooberlin.com
func ParseVooberlin() (products []*Product, err error) {
	shopURL := "https://www.vooberlin.com"
	doc, err := getDoc(shopURL+"/men/footwear/", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.listing > div > div > div.product--info > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("span > span > img").Attr("srcset")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseStarcowparis parses www.starcowparis.com
func ParseStarcowparis() (products []*Product, err error) {
	shopURL := "https://www.starcowparis.com"
	doc, err := getDoc(shopURL+"/40-footwear", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.productList > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("span.image > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// Parse1stog parses www.1st-og.ch
func Parse1stog() (products []*Product, err error) {
	shopURL := "https://www.1st-og.ch"
	doc, err := getDoc(shopURL+"/brands", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.columns > ul > li.item > div").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("a.product-name").Attr("title")
		link, _ := s.Find("a.product-name").Attr("href")
		photoURL, _ := s.Find("div > a > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSsense parses www.ssense.com
func ParseSsense() (products []*Product, err error) {
	shopURL := "https://www.ssense.com"
	doc, err := getDoc(shopURL+"/en-us/men/shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.browsing-product-list > figure > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("div > picture > img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div > picture > img").Attr("data-srcset")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseMrporter parses www.mrporter.com
func ParseMrporter() (products []*Product, err error) {
	shopURL := "https://www.mrporter.com"
	doc, err := getDoc(shopURL+"/en-us/mens/whats-new/shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-images > div.product-image > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseNetaporter parses www.net-a-porter.com
func ParseNetaporter() (products []*Product, err error) {
	shopURL := "https://www.net-a-porter.com"
	doc, err := getDoc(shopURL+"/us/en/d/Shop/Whats-New/Now/Shoes/All?cm_sp=topnav-_-shoes-_-whatsnew&pn=1&npp=180&image_view=product&dscroll=0", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products > li > div.product-image > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("data-src")
		link = shopURL + link
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseNorsestore parses www.norsestore.com
func ParseNorsestore() (products []*Product, err error) {
	shopURL := "https://www.norsestore.com"
	doc, err := getDoc(shopURL+"/section/mens-footwear", false)
	if err != nil {
		return products, err
	}

	doc.Find("#commodity-lister-list > li > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("span.list-commodity-image > span > img")
		name, _ := t.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseFootpatrol parses www.footpatrol.com
func ParseFootpatrol() (products []*Product, err error) {
	shopURL := "https://www.footpatrol.com"
	doc, err := getDoc(shopURL+"/footwear/latest/", false)
	if err != nil {
		return products, err
	}

	doc.Find("#productListMain > li > span > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("picture > img")
		name, _ := t.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("data-src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseShiek parses www.shiekhshoes.com
func ParseShiek() (products []*Product, err error) {
	shopURL := "http://www.shiekhshoes.com"
	doc, err := getDoc(shopURL+"/c-46-mens-boots.aspx", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.productWrapper > div > div > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.p-info")
		name := t.Find("div.b").Text() + " " + t.Find("div.n").Text()
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div.rollover-container > img.main-img").Attr("src")
		link = shopURL + "/" + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseCitygear parses www.citygear.com
func ParseCitygear() (products []*Product, err error) {
	shopURL := "https://www.citygear.com"
	doc, err := getDoc(shopURL+"/catalog/shoes/gender/men/sort-by/news_from_date/sort-direction/desc.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")

		photoURL, _ := s.Find("img").Attr("src")
		if strings.Contains(photoURL, "data:image") {
			photoURL, _ = s.Find("img").Attr("data-pagespeed-lsc-url")
		}

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseBarneys parses www.barneys.com
func ParseBarneys() (products []*Product, err error) {
	shopURL := "https://www.barneys.com"
	doc, err := getDoc(shopURL+"/category/men/shoes/sneakers/N-j68mm4", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.product-set > li > div > div.product-image > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseAllikestore parses www.allikestore.com
func ParseAllikestore() (products []*Product, err error) {
	shopURL := "https://www.allikestore.com"
	doc, err := getDoc(shopURL+"/default/sneakers/mens-sneakers.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseCaliroots parses caliroots.com
func ParseCaliroots() (products []*Product, err error) {
	shopURL := "https://caliroots.com"
	doc, err := getDoc(shopURL+"/mens-shoes/s/19", false)
	if err != nil {
		return products, err
	}

	doc.Find("#products > li > a").Each(func(i int, s *goquery.Selection) {
		name := s.Find("p.brand").Text() + " " + s.Find("p.name").Text()
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div.image > img").Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseWellgosh parses wellgosh.com
func ParseWellgosh() (products []*Product, err error) {
	shopURL := "https://wellgosh.com"
	doc, err := getDoc(shopURL+"/footwear/?dir=desc&order=created_at", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.row > article > figure > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseThenextdoor parses www.thenextdoor.fr
func ParseThenextdoor() (products []*Product, err error) {
	shopURL := "https://www.thenextdoor.fr"
	doc, err := getDoc(shopURL+"/en/34-shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.row > li > div > div > div.pro_first_box > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseCornerstreet parses www.cornerstreet.fr
func ParseCornerstreet() (products []*Product, err error) {
	shopURL := "https://www.cornerstreet.fr"
	doc, err := getDoc(shopURL+"/hommes.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSneakersnstuff parses www.sneakersnstuff.com
func ParseSneakersnstuff() (products []*Product, err error) {
	shopURL := "https://www.sneakersnstuff.com"
	doc, err := getDoc(shopURL+"/en/904/mens-sneakers", false)
	if err != nil {
		return products, err
	}

	doc.Find("#products > li").Each(func(i int, s *goquery.Selection) {
		name := s.Find("div.info > h4.name > a").Text()
		link, _ := s.Find("a").Attr("href")
		photoURL, _ := s.Find("a > img").Attr("src")
		link = shopURL + link
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseTriads www.triads.co.uk
func ParseTriads() (products []*Product, err error) {
	shopURL := "https://www.triads.co.uk"
	doc, err := getDoc(shopURL+"/triads-mens-c1/footwear-c24", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.row > li > div > div.product__image > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseUebervart www.uebervart-shop.de
func ParseUebervart() (products []*Product, err error) {
	shopURL := "https://www.uebervart-shop.de"
	doc, err := getDoc(shopURL+"/shoes/", false)
	if err != nil {
		return products, err
	}

	doc.Find("#shop-products > article > a").Each(func(i int, s *goquery.Selection) {
		name := s.Find("h3").Text()
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseUptherestore uptherestore.com
func ParseUptherestore() (products []*Product, err error) {
	shopURL := "https://uptherestore.com"
	doc, err := getDoc(shopURL+"/footwear", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > div > div.product-img-box > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseYcmc www.ycmc.com
func ParseYcmc() (products []*Product, err error) {
	shopURL := "https://www.ycmc.com"
	doc, err := getDoc(shopURL+"/men.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > div.product-image-wrapper > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseNowhere nowhere.ie
func ParseNowhere() (products []*Product, err error) {
	shopURL := "https://nowhere.ie"
	doc, err := getDoc(shopURL+"/collections/footwear#", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.grid > div > div.supports-js > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseChicagocitysports parses chicagocitysports.com
func ParseChicagocitysports() (products []*Product, err error) {
	shopURL := "https://chicagocitysports.com"
	doc, err := getDoc(shopURL+"/collections/shoes/MEN", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.collection-grid-item > article > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("figure > img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("src")
		link = shopURL + link
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseBaitme parses www.baitme.com
func ParseBaitme() (products []*Product, err error) {
	shopURL := "http://www.baitme.com"
	doc, err := getDoc(shopURL+"/footwear", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseDoubleclutch parses www.doubleclutch.it
func ParseDoubleclutch() (products []*Product, err error) {
	shopURL := "https://www.doubleclutch.it"
	doc, err := getDoc(shopURL+"/en/basketball-shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("ol.products > li > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("data-name")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div.img-wrapper > div.big > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseUrbanstaroma parses www.urbanstaroma.com
func ParseUrbanstaroma() (products []*Product, err error) {
	shopURL := "https://www.urbanstaroma.com"
	doc, err := getDoc(shopURL+"/en/5-men-s-sneakers", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.product_list > li > div > div > div.pro_first_box > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseFeaturesneakerboutique parses www.featuresneakerboutique.com
func ParseFeaturesneakerboutique() (products []*Product, err error) {
	shopURL := "https://www.featuresneakerboutique.com"
	doc, err := getDoc(shopURL+"/collections/whats-new", false)
	if err != nil {
		return products, err
	}

	doc.Find("#posts > a.entry").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("article > img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("article > img").Attr("src")
		link = shopURL + link
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseStormfashion parses stormfashion.dk
func ParseStormfashion() (products []*Product, err error) {
	shopURL := "https://stormfashion.dk"
	doc, err := getDoc(shopURL+"/category/shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("#products > ul > li > div.pro_img > ul > li > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(strings.Replace(name, "\n", " ", 1)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseBasket4ballers parses www.basket4ballers.com
func ParseBasket4ballers() (products []*Product, err error) {
	shopURL := "https://www.basket4ballers.com"
	doc, err := getDoc(shopURL+"/en/152-shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("#product_list > li > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("div.block-product__wrapper-img > img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div.block-product__wrapper-img > img").Attr("data-src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseThegoodlifespace parses www.thegoodlifespace.com
func ParseThegoodlifespace() (products []*Product, err error) {
	shopURL := "https://www.thegoodlifespace.com"
	doc, err := getDoc(shopURL+"/collections/latest-arrivals", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-list > div > div > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("div > img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div > img").Attr("data-src")
		link = shopURL + link
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseJdsports parses www.jdsports.co.uk
func ParseJdsports() (products []*Product, err error) {
	shopURL := "https://www.jdsports.co.uk"
	doc, err := getDoc(shopURL+"/men/mens-footwear/?sort=latest", false)
	if err != nil {
		return products, err
	}

	doc.Find("#productListMain > li > span > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("picture > img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("picture > img").Attr("data-src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseTike parses www.tike.rs
func ParseTike() (products []*Product, err error) {
	shopURL := "https://www.tike.rs"
	doc, err := getDoc(shopURL+"/patike", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.row > div > div > div > div.product-item-data-col").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.product-image-wrapper > img")
		name, _ := t.Attr("alt")
		link, _ := s.Find("div.product-info-wrapper > div.product-item-title > a").Attr("href")
		photoURL, _ := t.Attr("src")
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSupplystore parses www.supplystore.com.au
func ParseSupplystore() (products []*Product, err error) {
	shopURL := "https://www.supplystore.com.au"
	doc, err := getDoc(shopURL+"/brands/new-arrivals/c-28/c-150", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.collection-grid-content > li > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSelfridges parses www.selfridges.com
func ParseSelfridges() (products []*Product, err error) {
	shopURL := "http://www.selfridges.com"
	doc, err := getDoc(shopURL+"/US/en/cat/shoes/mens/?cm_sp=MegaMenu-_-Shoes-_-Mens&fh_sort_by=newness%20|ASC", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.products > div > div").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div > a.description")
		name := t.Text()
		link, _ := t.Attr("href")
		photoURL, _ := s.Find("a > img").Attr("data-mainsrc")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseBrownsfashion parses www.brownsfashion.com
func ParseBrownsfashion() (products []*Product, err error) {
	shopURL := "https://www.brownsfashion.com"
	doc, err := getDoc(shopURL+"/sets/man/mens-new-in/?categories=135968&sort=newItems&fsb=1", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-collection__container > article > figure > a.collection-item__link").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("data-src")
		link = shopURL + link
		photoURL = photoURL[strings.Index(photoURL, "lg:")+3:]

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseKicksstore parses kicksstore.eu
func ParseKicksstore() (products []*Product, err error) {
	shopURL := "https://kicksstore.eu"
	doc, err := getDoc(shopURL+"/menu/shoes/shoes/all", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product__container > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div.listing-item > div.listing-pic > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSaksfifthavenue parses www.saksfifthavenue.com
func ParseSaksfifthavenue() (products []*Product, err error) {
	shopURL := "https://www.saksfifthavenue.com"
	doc, err := getDoc(shopURL+"/Men/Shoes/shop/_/N-52flst/Ne-6lvnb5", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.pa-product-large").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.image-container-large > a.sfa-product-view")
		name := s.Find("div.product-text > a > p.product-description").Text()
		link, _ := t.Attr("href")
		photoURL, _ := t.Find("img.pa-product-large").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSil parses www.sil.lt
func ParseSil() (products []*Product, err error) {
	shopURL := "https://www.sil.lt"
	doc, err := getDoc(shopURL+"/shoes2.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.search_products > div > div.photoHolder > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		photoURL = shopURL + "/" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseProdirectsoccer parses www.prodirectsoccer.com
func ParseProdirectsoccer() (products []*Product, err error) {
	shopURL := "https://www.prodirectsoccer.com"
	doc, err := getDoc(shopURL+"/lists/all-mens-sneakers.aspx?listName=all-mens-sneakers&pp=32&o=latest", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.list > div").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.img-holder > a")
		name := s.Find("div.productTitlePriceWrap > a").Text()
		link, _ := t.Attr("href")
		photoURL, _ := t.Find("img").Attr("src")
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseThebrokenarm parses www.the-broken-arm.com
func ParseThebrokenarm() (products []*Product, err error) {
	shopURL := "https://www.the-broken-arm.com"
	doc, err := getDoc(shopURL+"/en/26-men", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.product_list > li > div > div.left-block > div	> a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img.ProdImg").Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img.ProdImg").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseTitan22 parses www.titan22.com
func ParseTitan22() (products []*Product, err error) {
	shopURL := "https://www.titan22.com"
	doc, err := getDoc(shopURL+"/new-arrivals.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("ol.products > li > div > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("span > span > img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseNosstore parses www.nosstore.eu
func ParseNosstore() (products []*Product, err error) {
	shopURL := "https://www.nosstore.eu"
	doc, err := getDoc(shopURL+"/en/filter/gender_4,1,3,6,2/type_44,48,52/", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.frame > article > div.post-holder").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.img-holder > div.box > div > a")
		name := s.Find("div.description > div.heading-box > strong").Text()
		link, _ := t.Attr("href")
		photoURL, _ := t.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseOqium parses www.oqium.com
func ParseOqium() (products []*Product, err error) {
	shopURL := "https://www.oqium.com"
	doc, err := getDoc(shopURL+"/sneakers/?vars=sort:snieuwste;doelgroep_84:true;", false)
	if err != nil {
		return products, err
	}

	doc.Find("#product-list > ul > li > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSize parses www.size.co.uk
func ParseSize() (products []*Product, err error) {
	shopURL := "https://www.size.co.uk"
	doc, err := getDoc(shopURL+"/mens/footwear/latest/", false)
	if err != nil {
		return products, err
	}

	doc.Find("#productListMain > li > span > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("picture > img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("data-src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseOffspring parses www.offspring.co.uk
func ParseOffspring() (products []*Product, err error) {
	shopURL := "https://www.offspring.co.uk"
	doc, err := getDoc(shopURL+"/view/category/offspring_catalog/NEWIN?pageSize=100&sort=releaseDate", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.productList > ul > li > div").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.productList_img > a")
		name := s.Find("div.productList_info > div > a").Text()
		link, _ := t.Attr("href")
		photoURL, _ := t.Find("img").Attr("data-media")
		link = shopURL + link
		photoURL = "https:" + photoURL[strings.LastIndex(photoURL, "//"):len(photoURL)-2]

		space := regexp.MustCompile(`\s+`)
		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(space.ReplaceAllString(name, " "))), link, shopURL, photoURL))
	})

	return products, err
}

// ParseWalkinmycloset parses www.walk-inmycloset.com
func ParseWalkinmycloset() (products []*Product, err error) {
	shopURL := "https://www.walk-inmycloset.com"
	doc, err := getDoc(shopURL+"/collections/new-arrivals-1", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.grid-uniform > div > div > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseWiseboutique parses www.wiseboutique.com
func ParseWiseboutique() (products []*Product, err error) {
	shopURL := "https://www.wiseboutique.com"
	doc, err := getDoc(shopURL+"/en/shoes-man?tp=news", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.colonna_shop > div.contfoto > div.cotienifoto > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSummerstore parses www.summer-store.com
func ParseSummerstore() (products []*Product, err error) {
	shopURL := "https://www.summer-store.com"
	doc, err := getDoc(shopURL+"/en/11-shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("section.product_list > article > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseHbx parses hbx.com
func ParseHbx() (products []*Product, err error) {
	shopURL := "https://hbx.com"
	doc, err := getDoc(shopURL+"/new-arrivals/sneakers", true)
	if err != nil {
		return products, err
	}

	doc.Find("div.picture-wrapper > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		photoURL = photoURL[:strings.Index(photoURL, "?")] + "?q=60&w=720&fit=max&v=1"

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseThewebster parses thewebster.us
func ParseThewebster() (products []*Product, err error) {
	shopURL := "https://thewebster.us"
	doc, err := getDoc(shopURL+"/shop/men/shoes.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > div > div.proImage > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("span.face > img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("span.face > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParsePatta parses www.patta.nl
func ParsePatta() (products []*Product, err error) {
	shopURL := "https://www.patta.nl"
	doc, err := getDoc(shopURL+"/footwear", false)
	if err != nil {
		return products, err
	}

	doc.Find("ol.products > li > div").Each(func(i int, s *goquery.Selection) {
		name := s.Find("div.product-item-details > strong > a").Text()
		link, _ := s.Find("a").Attr("href")
		photoURL, _ := s.Find("a > div > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseAntonioli parses www.antonioli.eu
func ParseAntonioli() (products []*Product, err error) {
	shopURL := "https://www.antonioli.eu"
	doc, err := getDoc(shopURL+"/en/US/men/section/new-arrivals", false)
	if err != nil {
		return products, err
	}

	doc.Find("section.products > article.product > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("span").Attr("content")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("figure > p > img.top").Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseExcelsiormilano parses www.excelsiormilano.com
func ParseExcelsiormilano() (products []*Product, err error) {
	shopURL := "https://www.excelsiormilano.com"
	doc, err := getDoc(shopURL+"/1258-what-s-new", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.product_list > li > div.product-container > div.left-block > div > a.product_img_link").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		name = strings.TrimSpace(s.Find("div.right-block > h5.newManu > a").Text()) + " " + name

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseFootish parses www.footish.se
func ParseFootish() (products []*Product, err error) {
	shopURL := "https://www.footish.se"
	doc, err := getDoc(shopURL+"/en/sneakers", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.product-list > li > div > article > div.product-image > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("data-original")
		link = shopURL + link
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// Parse290sqm parses ist.290sqm.com
func Parse290sqm() (products []*Product, err error) {
	shopURL := "https://ist.290sqm.com"
	doc, err := getDoc(shopURL+"/Men/Mens-Shoes", false)
	if err != nil {
		return products, err
	}

	doc.Find("#content > div > div > div > div.image > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseNighshop parses www.nighshop.com
func ParseNighshop() (products []*Product, err error) {
	shopURL := "https://www.nighshop.com"
	doc, err := getDoc(shopURL, false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.product-list > li > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div.image > img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseNordstorm parses shop.nordstrom.com
func ParseNordstorm() (products []*Product, err error) {
	shopURL := "https://shop.nordstrom.com"
	doc, err := getDoc(shopURL+"/c/mens-shoes?origin=topnav&cm_sp=Top%20Navigation-_-Men_-_-Shoes&top=72&offset=6&page=1&sort=Newest", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.row_Z2vyUxT > div > article").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.media_12txnm")
		name := s.Find("h3 > a > span > span").Text()
		link, _ := t.Find("a").Attr("href")
		photoURL, _ := t.Find("img").Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseNtrstore parses www.ntrstore.com"
func ParseNtrstore() (products []*Product, err error) {
	shopURL := "https://www.ntrstore.com"
	doc, err := getDoc(shopURL+"/sneakers", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > div > div.product-image-area > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseOkini parses oki-ni.com
func ParseOkini() (products []*Product, err error) {
	shopURL := "https://oki-ni.com"
	doc, err := getDoc(shopURL+"/latest", false)
	if err != nil {
		return products, err
	}

	doc.Find("ol.products > li > div > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("span > span > img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseOverkillshop parses www.overkillshop.com
func ParseOverkillshop() (products []*Product, err error) {
	shopURL := "https://www.overkillshop.com"
	doc, err := getDoc(shopURL+"/en/new-products.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseShoezgallery parses www.shoezgallery.com
func ParseShoezgallery() (products []*Product, err error) {
	shopURL := "https://www.shoezgallery.com"
	doc, err := getDoc(shopURL+"/en/32-latest", false)
	if err != nil {
		return products, err
	}

	doc.Find("#category_products > div > div.product-image > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("img.normal").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img.normal").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSpace23 parses www.space23.it
func ParseSpace23() (products []*Product, err error) {
	shopURL := "https://www.space23.it"
	doc, err := getDoc(shopURL+"/men.html", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.products-grid > li > div.product-media > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseJimmyjazz parses www.jimmyjazz.com
func ParseJimmyjazz() (products []*Product, err error) {
	shopURL := "http://www.jimmyjazz.com"
	doc, err := getDoc(shopURL+"/mens/specials/new-arrivals", false)
	if err != nil {
		return products, err
	}

	doc.Find("ul.product_grid > li > div.product_grid_image > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseKickzpremium parses www.kickzpremium.com
func ParseKickzpremium() (products []*Product, err error) {
	shopURL := "https://www.kickzpremium.com"
	doc, err := getDoc(shopURL+"/en/new", false)
	if err != nil {
		return products, err
	}

	doc.Find("#product_list_container > div > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div > div.categoryElementPicture > img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("link")
		photoURL, _ := t.Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseYmeuniverse parses www.ymeuniverse.com
func ParseYmeuniverse() (products []*Product, err error) {
	shopURL := "https://www.ymeuniverse.com"
	doc, err := getDoc(shopURL+"/en/216/new-arrivals", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-list > div").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div.card-image-container > a")
		name := s.Find("div.card-content > a > h4.card-title").Text()
		link, _ := t.Attr("href")
		photoURL, _ := t.Find("div > img").Attr("data-src")
		link = shopURL + link
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return products, err
}

// ParseFootlocker parses www.footlocker.co.uk
func ParseFootlocker() (products []*Product, err error) {
	shopURL := "https://www.footlocker.co.uk"
	doc, err := getDoc(shopURL+"/en/all/", true)
	if err != nil {
		return products, err
	}

	doc.Find("div.fl-category--productlist > div > div.fl-category--productlist--item").Each(func(i int, s *goquery.Selection) {
		name := s.Find("a").Text()
		link, _ := s.Find("a").Attr("href")
		photoURL := ""

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseFinishline parses www.finishline.com
func ParseFinishline() (products []*Product, err error) {
	shopURL := "https://www.finishline.com"
	doc, err := getDoc(shopURL+"/store/new-arrivals?mnid=newarrivals", false)
	if err != nil {
		return products, err
	}

	doc.Find("div.product-grid > div > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("div > div > img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("src")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseFootdistrict parses https://footdistrict.com
func ParseFootdistrict() (products []*Product, err error) {
	shopURL := "https://footdistrict.com"
	doc, err := getDoc(shopURL+"/en/sneakers/latest.html", false)
	if err != nil {
		return products, err
	}

	code, err := doc.Find("script").Html()
	if err == nil && code != "" {
		data := map[string]string{}
		translateJS(&code)
		code = strings.Replace(code, "e(r)", `data.answer = r;`, 1)
		code = strings.Replace(code, "while(l>=8){((a=(U>>>(l-=8))&amp;0xff)||(i<(L-2)))&amp;&amp;", "while(l>=8){((a=(U>>>(l-=8))&0xff)||(i<(L-2)))&0xff&0xff;", 1)
		code = `setEnrichFunction(function(data) {` + code + `});`
		runJS(code, &data)

		code = data["answer"]
		translateJS(&code)
		code = strings.Replace(code, " + ';path=/;max-age=86400'; location.reload();", "});", 1)
		code = strings.Replace(code, "document.cookie", "data.cookie", 1)
		code = `setEnrichFunction(function(data) {` + code
		code = code[:strings.Index(code, "});")+len("});")]
		runJS(code, &data)

		c := strings.Split(data["cookie"], "=")
		doc, err = getDoc(shopURL+"/en/sneakers/latest.html", false, cookie{c[0], c[1]})
		if err != nil {
			return products, err
		}
	}

	doc.Find("ul.products-grid > li > a.product-image").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseBstn parses www.bstn.com
func ParseBstn() (products []*Product, err error) {
	shopURL := "https://www.bstn.com"
	doc, err := getDoc(shopURL+"/new-arrivals", true)
	if err != nil {
		return products, err
	}

	doc.Find("ul.productlist > li > div > div.pImageContainer > a").Each(func(i int, s *goquery.Selection) {
		t := s.Find("img")
		name, _ := t.Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := t.Attr("src")
		link = shopURL + link
		photoURL = shopURL + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseSsenseGB parses www.ssense.com
func ParseSsenseGB() (products []*Product, err error) {
	shopURL := "https://www.ssense.com"
	doc, err := getDoc(shopURL+"/en-gb/men/shoes", false)
	if err != nil {
		return
	}

	doc.Find("div.browsing-product-list > figure > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("div > picture > img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div > picture > img").Attr("data-srcset")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return
}

// ParseSsenseCA parses www.ssense.com
func ParseSsenseCA() (products []*Product, err error) {
	shopURL := "https://www.ssense.com"
	doc, err := getDoc(shopURL+"/en-ca/men/shoes", false)
	if err != nil {
		return
	}

	doc.Find("div.browsing-product-list > figure > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Find("div > picture > img").Attr("alt")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("div > picture > img").Attr("data-srcset")
		link = shopURL + link

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return
}

// ParseKm20 parses https://www.km20.ru
func ParseKm20() (products []*Product, err error) {
	shopURL := "https://www.km20.ru"
	doc, err := getDoc(shopURL+"/catalog/men/new/", false)
	if err != nil {
		return
	}

	doc.Find("div.cat_wrapper > a").Each(func(i int, s *goquery.Selection) {
		name := s.Find("span.cat_item_name").Text()
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("span.cat_item_img > img").Attr("src")
		link = shopURL + link
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return
}

// ParseBrandShop parses
func ParseBrandShop() (products []*Product, err error) {
	shopURL := "https://brandshop.ru"
	doc, err := getDoc(shopURL+"/new/", false)
	if err != nil {
		return
	}

	doc.Find("div.product-container > div > a").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("title")
		link, _ := s.Attr("href")
		photoURL, _ := s.Find("img").Attr("src")

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return
}

// ParseKM20 parses
func ParseKM20() (products []*Product, err error) {
	shopURL := "https://www.km20.ru"
	doc, err := getDoc(shopURL+"/catalog/men/new/", false)
	if err != nil {
		return
	}

	doc.Find("div.cat_wrapper.js_parallax.clearfix > a").Each(func(i int, s *goquery.Selection) {
		name := s.Find("span.cat_item_name").Text()
		link, _ := s.Attr("href")
		link = shopURL + link
		photoURL, _ := s.Find("span.cat_item_img > img").Attr("src")
		photoURL = "https:" + photoURL

		products = append(products, NewProduct(strings.ToLower(strings.TrimSpace(name)), link, shopURL, photoURL))
	})

	return
}

/*
// ParseHolypopstore parses https://www.holypopstore.com
func ParseHolypopstore() (products []*Product, err error) {
	shopURL := "https://www.holypopstore.com"
	doc, err := getDoc(shopURL+"/en/latest", false)
	if err != nil {
		return products, err
	}

	log.Println(doc.Html())

	doc.Find("div.fl-category--productlist > div > div.fl-category--productlist--item").Each(func(i int, s *goquery.Selection) {
		name := s.Find("a").Text()
		link, _ := s.Find("a").Attr("href")
		photoURL := ""

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}

// ParseOneblockdown parses www.oneblockdown.it
func ParseOneblockdown() (products []*Product, err error) {
	shopURL := "https://www.oneblockdown.it"
	doc, err := getDoc(shopURL+"/en/latest", false)
	if err != nil {
		return products, err
	}

	js := doc.Find("script").Next().Next().Text()
	translateJS(&js)
	js = strings.Replace(js, `t = document.createElement('div');`, "", 1)
	js = strings.Replace(js, `t.innerHTML="<a href='/'>x</a>";`, "", 1)
	js = strings.Replace(js, `t = t.firstChild.href;r = t.match(/https?:\/\//)[0];`, "", 1)
	js = strings.Replace(js, `t = t.substr(r.length); t = t.substr(0,t.length-1);`, "", 1)
	js = strings.Replace(js, `a = document.getElementById('jschl-answer');`, "", 1)
	js = strings.Replace(js, `f = document.getElementById('challenge-form');`, "", 1)
	js = js[:strings.Index(js, "f.action += location.hash;")] + "});"
	js = js[strings.Index(js, "setTimeout(function(){"):]
	js = strings.Replace(js, "setTimeout(function(){", "setEnrichFunction(function(data) {", 1)
	js = strings.Replace(js, "a.value = ", "data.jschl_answer = (", 1)
	js = strings.Replace(js, " + t.length", ").toString()", 1)
	f, _ := os.Create("def.js")
	f.WriteString(js)

	n := doc.Find("#challenge-form > input").Nodes
	data := map[string]string{
		"action":   "/cdn-cgi/l/chk_jschl",
		"jschl_vc": n[0].Attr[2].Val,
		"pass":     n[1].Attr[2].Val,
	}

	runJS(js, &data)

	form := url.Values{}
	form.Add("jschl_vc", data["jschlvcValue"])
	form.Add("pass", data["pass"])
	form.Add("jschl_answer", data["jschl_answer"])

	response, err := client.PostForm(shopURL+data["action"]+"", form)
	if err != nil {
		log.Fatalf("FUCK ME HARDER %v\n", err)
	}

	/*req, err := http.NewRequest("GET", shopURL+data["action"], strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatalf("FUCK ME %v\n", err)
	}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("FUCK ME HARDER %v\n", err)
	}
	doc, err = goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalf("FUCK ME HARDERRRR %v\n", err)
	}

	log.Println(doc.Html())
	log.Println(data)

	doc.Find("div.fl-category--productlist > div > div.fl-category--productlist--item").Each(func(i int, s *goquery.Selection) {
		name := s.Find("a").Text()
		link, _ := s.Find("a").Attr("href")
		photoURL := ""

		products = append(products, NewProduct(strings.ToLower(name), link, shopURL, photoURL))
	})

	return products, err
}
*/
