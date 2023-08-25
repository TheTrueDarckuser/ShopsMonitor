package parser

import (
	"strings"
)

var (
// Checkers contains all the checkers
/*Checkers = map[string]AvailabilityChecker{
	"https://www.solebox.com":          CheckSolebox,
	"https://www.nakedcph.com":         CheckNaked,
	"https://www.nittygrittystore.com": CheckNittygritty,
	"https://thewebster.us":            CheckThewebster,
	"https://www.woodwood.com":         CheckWoodwood,
	"https://www.luisaviaroma.com":     CheckLuisaviaroma,
	"https://www.slamjamsocialism.com": CheckSlamjamsocialism,
	"https://www.sneakers76.com":       CheckSneakers76,
	"https://www.43einhalb.com":        Check43einhlab,
	"https://www.size.co.uk":           CheckSize,
	"https://hbx.com":                  CheckHbx,
	"https://caliroots.com":            CheckCaliroots,
	"https://www.finishline.com":       CheckFinishline,
	"https://www.allikestore.com":      CheckAllikestore,
	"https://wellgosh.com":             CheckWellgosh,
	"https://www.thenextdoor.fr":       CheckThenextdoor,
	"https://www.cornerstreet.fr":      CheckCornerstreet,
	"https://www.endclothing.com":      CheckEnd,
	"https://asphaltgold.de":           CheckAsphaltgold,
	"https://www.footshop.eu":          CheckFootshop,
	"https://www.dtlr.com":             CheckDtlr,
	"http://www.consortium.co.uk":      CheckConsortium,
	"https://chmielna20.pl":            CheckChmielna20,
	"https://goodhoodstore.com":        CheckGoodhoodstore,
	"https://www.patta.nl":             CheckPatta,
	"https://www.ruvilla.com":          CheckRuvilla,
	"https://www.shinzo.paris":         CheckShinzo,
	"https://www.sneakersnstuff.com":   CheckSneakersnstuff,
	"https://sneakerstudio.com":        CheckSneakerstudio,
	"https://en.titolo.ch":             CheckTitolo,
	"https://tres-bien.com":            CheckTresbien,
	"https://www.urbanjunglestore.com": CheckUrbanjungle,
	"https://www.vooberlin.com":        CheckVooberlin,
	"https://www.1st-og.ch":            Check1stog,
	"https://www.ssense.com":           CheckSsense,
	"https://www.mrporter.com":         CheckMrporter,
	"https://www.norsestore.com":       CheckNorsestore,
	"http://www.shiekhshoes.com":       CheckShiek,
	"https://www.citygear.com":         CheckCitygear,
	"https://www.barneys.com":          CheckBarneys,
	"https://www.footlocker.co.uk":     CheckFootlocker,
	"https://ist.290sqm.com":           Check290sqm,
	"https://www.footpatrol.com":       CheckFootpatrol,
	"https://www.bstn.com":             CheckBstn,
	"https://www.ymeuniverse.com":      CheckYmeuniverse,
	"http://www.jimmyjazz.com":         CheckJimmyjazz,
	"https://shop.nordstrom.com":       CheckNordstorm,
	"https://www.net-a-porter.com":     CheckNetaporter,
	"https://www.starcowparis.com":     CheckStarcowparis,
	"https://footdistrict.com":         CheckFootdistrict,
	"https://www.antonioli.eu":         CheckAntonioli,
	"https://en.aw-lab.com":            CheckAwlab,
	"https://bouncewear.com":           CheckBouncewear,
	"https://www.excelsiormilano.com":  CheckExcelsiormilano,
	"https://www.footish.se":           CheckFootish,
}*/
)

// CheckSivasdelcalzo parses new arrivals
func CheckSivasdelcalzo(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add-to-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSolebox parses new arrivals try
func CheckSolebox(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#toBasket").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckNaked parses new arrivals try
func CheckNaked(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-form").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckNittyGritty parses new arrivals try
func CheckNittyGritty(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.button--accent").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckThewebster parses new arrivals try
func CheckThewebster(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#addToCartButton").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckWoodwood parses new arrivals try
func CheckWoodwood(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("li.form-itemselector-li-submit").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckLuisaviaroma parses new arrivals try
func CheckLuisaviaroma(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#a-bnt-add-bag").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSlamjamsocialism parses new arrivals try
func CheckSlamjamsocialism(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add_to_cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSneakers76 parses new arrivals try
func CheckSneakers76(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add_to_cart").Length() > 0 {
		available = true
	}
	return available, err
}

// Check43einhlab parses new arrivals try
func Check43einhlab(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#submitAddToCart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSize parses new arrivals try
func CheckSize(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#addToBasket").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckHbx parses new arrivals try
func CheckHbx(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.add-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckCaliroots parses new arrivals try
func CheckCaliroots(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("div.add-to-cart-form.unavailable").Length() == 0 {
		available = true
	}
	return available, err
}

// CheckFinishline parses new arrivals try
func CheckFinishline(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#buttonAddToCart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckAllikestore parses new arrivals try
func CheckAllikestore(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("div.add-to-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckWellgosh parses new arrivals try
func CheckWellgosh(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("div.add-to-cart-buttons").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckThenextdoor parses new arrivals try
func CheckThenextdoor(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("span.st-label-success").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckCornerstreet parses new arrivals try
func CheckCornerstreet(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.btn-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckEnd parses new arrivals try
func CheckEnd(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-addtocart-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckAsphaltgold parses new arrivals try
func CheckAsphaltgold(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.btn-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckFootshop parses new arrivals try
func CheckFootshop(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add-to-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckConsortium parses new arrivals try
func CheckConsortium(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-addtocart-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckChmielna20 parses new arrivals try
func CheckChmielna20(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("input.btn.sd-basket-add").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckGoodhoodstore parses new arrivals try
func CheckGoodhoodstore(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("input.Btn-Action.fat").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckPatta parses new arrivals try
func CheckPatta(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.action.primary.tocart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckRuvilla parses new arrivals try
func CheckRuvilla(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-addtocart-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckShinzo parses new arrivals try
func CheckShinzo(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.button-add-cart.nosizeselected").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSneakersnstuff parses new arrivals try
func CheckSneakersnstuff(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add-to-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSneakerstudio parses new arrivals try
func CheckSneakerstudio(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#projector_button_basket").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckTitolo parses new arrivals try
func CheckTitolo(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckTresbien parses new arrivals try
func CheckTresbien(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-addtocart-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckUrbanjungle parses new arrivals try
func CheckUrbanjungle(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-addtocart-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckVooberlin parses new arrivals try
func CheckVooberlin(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.buybox--button.block.btn.is--primary.is--icon-right.is--center.is--large").Length() > 0 {
		available = true
	}
	return available, err
}

// Check1stog parses new arrivals try
func Check1stog(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSsense parses new arrivals try
func CheckSsense(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.full-width.btn-add-to-bag.button.full-width.no-border.button-primary").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckMrporter parses new arrivals try
func CheckMrporter(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.btn.product-button.add-to-bag-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckNorsestore parses new arrivals try
func CheckNorsestore(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	_, t := doc.Find("#commodity-show-addcart-submit").Attr("value")
	if t {
		available = true
	}
	return available, err
}

// CheckShiek parses new arrivals try
func CheckShiek(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.action.tocart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckCitygear parses new arrivals try
func CheckCitygear(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.btn-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckBarneys parses new arrivals try
func CheckBarneys(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	_, t := doc.Find("#atg_behavior_addItemToCart").Attr("value")
	if t {
		available = true
	}
	return available, err
}

// Check290sqm parses new arrivals try
func Check290sqm(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	t := doc.Find("table.table").Text()
	if strings.Contains(t, "Stokta var") {
		available = true
	}
	return available, err
}

// CheckFootpatrol parses new arrivals try
func CheckFootpatrol(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#addToBasket").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckBstn parses new arrivals try
func CheckBstn(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#submitAddToCart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckYmeuniverse www.ymeuniverse.com parses new arrivals try
func CheckYmeuniverse(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.product-form-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckJimmyjazz parses new arrivals try
func CheckJimmyjazz(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add_to_bag_btn").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckNordstorm parses new arrivals try
func CheckNordstorm(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckStarcowparis parses new arrivals try
func CheckStarcowparis(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add_to_cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckNetaporter parses new arrivals try
func CheckNetaporter(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add-to-bag").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckFootdistrict parses new arrivals try
func CheckFootdistrict(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.btn-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckAntonioli parses new arrivals try
func CheckAntonioli(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-variants > p").Text() == "Available Now" {
		available = true
	}
	return available, err
}

// CheckAwlab parses new arrivals try
func CheckAwlab(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("div.add-to-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckBouncewear parses new arrivals try
func CheckBouncewear(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add_to_cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckExcelsiormilano parses new arrivals try
func CheckExcelsiormilano(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add_to_cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckFootish parses new arrivals try
func CheckFootish(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("a.buy-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSsenseGB parses new arrivals try
func CheckSsenseGB(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.full-width.btn-add-to-bag.button.full-width.no-border.button-primary").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSsenseCA parses new arrivals try
func CheckSsenseCA(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.full-width.btn-add-to-bag.button.full-width.no-border.button-primary").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckSpace23 parses new arrivals try
func CheckSpace23(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.btn.btn-roll-text.btn-lg.btn-primary.btn-group.buy-overlay-menu__buy-btn").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckCruvoir parses new arrivals try
func CheckCruvoir(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if t, _ := doc.Find("div > link").Attr("href"); strings.Contains(t, "InStock") {
		available = true
	}
	return available, err
}

// CheckBrandShop parses new arrivals try
func CheckBrandShop(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.btn.btn-cart").Text() == "Добавить в корзину" {
		available = true
	}
	return available, err
}

// CheckBaitMe parses new arrivals try
func CheckBaitMe(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.button.btn-cart > span > span").Text() == "Add to Cart" {
		available = true
	}
	return available, err
}

// CheckDoubleClutch parses new arrivals try
func CheckDoubleClutch(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-addtocart-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckUrbanStaroma parses new arrivals try
func CheckUrbanStaroma(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.btn.btn-large.btn_primary.exclusive > span").Text() == "Add to cart" {
		available = true
	}
	return available, err
}

// CheckLeray parses new arrivals try
func CheckLeray(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("div.ajout").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckFeatureSneakerBoutique parses new arrivals try
func CheckFeatureSneakerBoutique(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#AddToCartText").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckKickz parses new arrivals try
func CheckKickz(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#addToBasket").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckStorm parses new arrivals try
func CheckStorm(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("a.buy.size-variants").Text() == "Buy" {
		available = true
	}
	return available, err
}

// CheckBasket4ballers parses new arrivals try
func CheckBasket4ballers(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add_to_cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckTheGoodLifeSpace parses new arrivals try
func CheckTheGoodLifeSpace(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	stock, _ := doc.Find("button.action_button.add_to_cart").Attr("data-label")
	if stock == "Add to Cart" {
		available = true
	}
	return available, err
}

// CheckTike parses new arrivals try
func CheckTike(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.shop-button-qnt.btn.btn-success.btn-icon-right.icon-cart-f.ease").Text() == "Dodaj u korpu" {
		available = true
	}
	return available, err
}

// CheckSupplyStore parses new arrivals try
func CheckSupplyStore(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	stock, _ := doc.Find("#Form_Form_action_doform").Attr("title")
	if stock == "Add To Cart" {
		available = true
	}
	return available, err
}

// CheckSelfridges parses new arrivals try
func CheckSelfridges(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("div.actions.actionsGlobal > button > span").Text() == "Add to bag" {
		available = true
	}
	return available, err
}

// CheckBrownFashion parses new arrivals try
func CheckBrownFashion(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add-to-bag-button > button > span").Text() == "Add To Bag" {
		available = true
	}
	return available, err
}

// CheckDistance parses new arrivals try
func CheckDistance(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#AddToCartText").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckKicksStore parses new arrivals try
func CheckKicksStore(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	stock, _ := doc.Find("#add-basket").Attr("value")
	if stock == "ADD TO BASKET" {
		available = true
	}
	return available, err
}

// CheckKm20 parses new arrivals try
func CheckKm20(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if t := doc.Find("div.prod_tocart  > a > span").Text(); strings.Contains(t, "В корзину") {
		available = true
	}
	return available, err
}

// CheckSummerstore parses new arrivals try
func CheckSummerstore(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#add_to_cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckWiseboutique parses new arrivals try
func CheckWiseboutique(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#addshopping2").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckWalkinmycloset parses new arrivals try
func CheckWalkinmycloset(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if t := doc.Find("button.btn").Text(); strings.Contains(t, "Add to Cart") {
		available = true
	}
	return available, err
}

// CheckOffspring parses new arrivals try
func CheckOffspring(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#ajaxAdd").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckOqium parses new arrivals try
func CheckOqium(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.shopify-payment-button__button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckNosstore parses new arrivals try
func CheckNosstore(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#cphContent_btnAddToCart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckTitan22 parses new arrivals try
func CheckTitan22(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-addtocart-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckThebrokenarm parses new arrivals try
func CheckThebrokenarm(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.exclusive").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckFootlocker parses new arrivals try
func CheckFootlocker(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("span.fl-product-details--stock hide").Text() == "In Stock" {
		available = true
	}
	return available, err
}

// CheckProdirectsoccer parses new arrivals try
func CheckProdirectsoccer(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#size").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckShoezGalleryEn parses new arrivals try
func CheckShoezGalleryEn(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("button.btn.btn-block.btn-shoez.btn-lg.text-uppercase").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckNighshop parses new arrivals try
func CheckNighshop(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-addtocart-button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckNtrStore parses new arrivals try
func CheckNtrStore(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("i.icon-cart").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckOkiNi parses new arrivals try
func CheckOkiNi(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#product-addtocart-button > span").Text() == "add to bag" {
		available = true
	}
	return available, err
}

// CheckOverKill parses new arrivals try
func CheckOverKill(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("div.span8.qty-add").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckTriads parses new arrivals try
func CheckTriads(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#js-product-in-stock-default").Text() == "Item in Stock" {
		available = true
	}
	return available, err
}

// CheckUebervart parses new arrivals try
func CheckUebervart(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("div.variations_button").Length() > 0 {
		available = true
	}
	return available, err
}

// CheckNowhere parses new arrivals try
func CheckNowhere(link string) (bool, error) {
	available := false
	doc, err := getDoc(link, false)
	if err != nil {
		return available, err
	}
	if doc.Find("#AddToCart").Length() > 0 {
		available = true
	}
	return available, err
}
