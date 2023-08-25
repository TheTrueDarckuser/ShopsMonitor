package parser

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"github.com/valyala/fasthttp"
)

var (
	// Global stores all the monitor stuff
	Global = map[string]GlobalWrapper{
		"https://www.solebox.com":          {ParseSolebox, CheckSolebox, "solebox", "/warenkorb/"},
		"https://www.nakedcph.com":         {ParseNaked, CheckNaked, "naked", "/cart/view"},
		"https://www.nittygrittystore.com": {ParseNittyGritty, CheckNittyGritty, "nittygritty", "/selection"},
		"https://thewebster.us":            {ParseThewebster, CheckThewebster, "thewebster", "/shop/checkout/cart/"},
		"https://www.woodwood.com":         {ParseWoodwood, CheckWoodwood, "woodwood", "/cart"},
		//"https://www.luisaviaroma.com":     {ParseLuisaviaroma, CheckLuisaviaroma, "lui"},
		"https://www.slamjamsocialism.com": {ParseSlamjamsocialism, CheckSlamjamsocialism, "slamjam", "/quick-order"},
		"https://www.sneakers76.com":       {ParseSneakers76, CheckSneakers76, "sneakers76", "/it/order"},
		"https://www.43einhalb.com":        {Parse43einhalb, Check43einhlab, "43einhlab", "/warenkorb"},
		"https://www.size.co.uk":           {ParseSize, CheckSize, "size", "/cart/"},
		"https://hbx.com":                  {ParseHbx, CheckHbx, "hbx", "/cart"},
		"https://caliroots.com":            {ParseCaliroots, CheckCaliroots, "caliroots", "/cart/view"},
		"https://www.finishline.com":       {ParseFinishline, CheckFinishline, "finishline", "/store/cart/cart.jsp"},
		"https://www.allikestore.com":      {ParseAllikestore, CheckAllikestore, "allikestore", "/default/checkout/cart/"},
		"https://wellgosh.com":             {ParseWellgosh, CheckWellgosh, "wellgosh", "/checkout/cart"},
		"https://www.thenextdoor.fr":       {ParseThenextdoor, CheckThenextdoor, "thenextdoor", "/en/commande-rapide"},
		"https://www.cornerstreet.fr":      {ParseCornerstreet, CheckCornerstreet, "cornerstreet", "/onestepcheckout/"},
		"https://www.endclothing.com":      {ParseEnd, CheckEnd, "end", "/row/checkout/cart/"},
		"https://asphaltgold.de":           {ParseAsphaltgold, CheckAsphaltgold, "asphaltgold", "/de/checkout/cart/"},
		"https://www.footshop.eu":          {ParseFootshop, CheckFootshop, "footshop", "/en/index.php?controller=order"},
		"http://www.consortium.co.uk":      {ParseConsortium, CheckConsortium, "consortium", "/checkout/cart/"},
		"https://chmielna20.pl":            {ParseChmielna20, CheckChmielna20, "chmielna20", "/basket"},
		"https://goodhoodstore.com":        {ParseGoodhoodstore, CheckGoodhoodstore, "goodhoodstore", "/checkout"},
		"https://www.patta.nl":             {ParsePatta, CheckPatta, "patta", "/checkout/cart/"},
		"https://www.ruvilla.com":          {ParseRuvilla, CheckRuvilla, "ruvilla", "/checkout/cart/"},
		"https://www.shinzo.paris":         {ParseShinzo, CheckShinzo, "shinzo", "/fr/commande#/panier"},
		"https://www.sneakersnstuff.com":   {ParseSneakersnstuff, CheckSneakersnstuff, "sneakersnstuff", "/en/cart/view"},
		"https://sneakerstudio.com":        {ParseSneakerstudio, CheckSneakerstudio, "sneakerstudio", "/basketedit.php"},
		"https://en.titolo.ch":             {ParseTitolo, CheckTitolo, "titolo", "/checkout/cart/"},
		"https://tres-bien.com":            {ParseTresbien, CheckTresbien, "tresbien", "/checkout/"},
		"https://www.urbanjunglestore.com": {ParseUrbanjungle, CheckUrbanjungle, "urbanjungle", "/it/checkout/cart/"},
		"https://www.vooberlin.com":        {ParseVooberlin, CheckVooberlin, "vooberlin", "/checkout/"},
		"https://www.1st-og.ch":            {Parse1stog, Check1stog, "1stog", "/checkout/cart/"},
		"https://www.ssense.com/en-us":     {ParseSsense, CheckSsense, "ssenseUS", "/shopping-bag"},
		"https://www.ssense.com/en-gb":     {ParseSsenseGB, CheckSsenseGB, "ssenseGB", "/shopping-bag"},
		"https://www.ssense.com/en-ca":     {ParseSsenseCA, CheckSsenseCA, "ssenseCA", "/shopping-bag"},
		"https://www.mrporter.com":         {ParseMrporter, CheckMrporter, "mrporter", "/en-us/shoppingbag.mrp"},
		"https://www.norsestore.com":       {ParseNorsestore, CheckNorsestore, "norsestore", "/cart"},
		"http://www.shiekhshoes.com":       {ParseShiek, CheckShiek, "shiek", "/shoppingcart.aspx?"},
		"https://www.citygear.com":         {ParseCitygear, CheckCitygear, "citygear", "/checkout/cart/"},
		"https://www.barneys.com":          {ParseBarneys, CheckBarneys, "barneys", "/cart"},
		"https://www.footlocker.co.uk":     {ParseFootlocker, CheckFootlocker, "footlocker", "/en/cart"},
		"https://ist.290sqm.com":           {Parse290sqm, Check290sqm, "290sqm", "/index.php?route=checkout/cart"},
		"https://www.footpatrol.com":       {ParseFootpatrol, CheckFootpatrol, "footpatrol", "/cart/"},
		"https://www.bstn.com":             {ParseBstn, CheckBstn, "bstn", "/en/cart"},
		"https://www.ymeuniverse.com":      {ParseYmeuniverse, CheckYmeuniverse, "ymeuniverse", "/en/cart/view"},
		"http://www.jimmyjazz.com":         {ParseJimmyjazz, CheckJimmyjazz, "jimmyjazz", "/cart"},
		"https://shop.nordstrom.com":       {ParseNordstorm, CheckNordstorm, "nordstorm", "/ShoppingBag.aspx"},
		"https://www.net-a-porter.com":     {ParseNetaporter, CheckNetaporter, "netaporter", "/us/en/shoppingbag.nap"},
		"https://www.starcowparis.com":     {ParseStarcowparis, CheckStarcowparis, "starcowparis", "/commande"},
		"https://footdistrict.com":         {ParseFootdistrict, CheckFootdistrict, "footdistrict", "/en/checkout/cart/"},
		"https://www.antonioli.eu":         {ParseAntonioli, CheckAntonioli, "antonioli", "/en/US/cart"},
		"https://en.aw-lab.com":            {ParseAwlab, CheckAwlab, "awlab", "/shop/checkout/cart/"},
		"https://bouncewear.com":           {ParseBouncewear, CheckBouncewear, "bouncewear", "/shoppingcart"},
		"https://www.excelsiormilano.com":  {ParseExcelsiormilano, CheckExcelsiormilano, "excelsiormilano", "/order"},
		"https://www.footish.se":           {ParseFootish, CheckFootish, "footish", "/checkout"},
		"https://www.space23.it":           {ParseSpace23, CheckSpace23, "space23", "/checkout/cart/"},
		"https://cruvoir.com":              {ParseCruvoir, CheckCruvoir, "cruvoir", "/cart"},
		"https://www.km20.ru":              {ParseKm20, CheckKm20, "km20", "/cart"},
		"https://www.summer-store.com":     {ParseSummerstore, CheckSummerstore, "summerstore", "/en/commande"},
		"https://www.wiseboutique.com":     {ParseWiseboutique, CheckWiseboutique, "wiseboutique", "/en/shoppingbag.html"},
		"https://www.walk-inmycloset.com":  {ParseWalkinmycloset, CheckWalkinmycloset, "walkinmycloset", "/cart"},
		"https://www.offspring.co.uk":      {ParseOffspring, CheckOffspring, "offspring", "/view/content/basket"},
		"https://oqium.com":                {ParseOqium, CheckOqium, "oqium", "/cart"},
		"https://www.nosstore.eu":          {ParseNosstore, CheckNosstore, "nosstore", "/en/cart/"},
		"https://www.titan22.com":          {ParseTitan22, CheckTitan22, "titan22", "/checkout/cart/"},
		"https://www.the-broken-arm.com":   {ParseThebrokenarm, CheckThebrokenarm, "thebrokenarm", "/en/index.php?controller=order"},
		"https://www.prodirectsoccer.com":  {ParseProdirectsoccer, CheckProdirectsoccer, "prodirectsoccer", "/V3_1/V3_1_Basket.aspx"},
		"https://brandshop.ru":             {ParseBrandShop, CheckBrandShop, "brandshop", "/cart/"},
		"http://www.baitme.com":            {ParseBaitme, CheckBaitMe, "", "/checkout/cart"},
		"https://www.doubleclutch.it":      {ParseDoubleclutch, CheckDoubleClutch, "baitme", "/it/checkout/cart/"},
		"https://www.urbanstaroma.com":     {ParseUrbanstaroma, CheckUrbanStaroma, "urbanstaroma", "/en/ordine"},
		//"":			{, CheckLeray, "", ""},
		"https://www.featuresneakerboutique.com": {ParseFeaturesneakerboutique, CheckFeatureSneakerBoutique, "featuresneakerboutique", "/cart"},
		//"https://www.kickzpremium.com":			{ParseKickzpremium, , "", "/en/cart"},
		//"":			{, CheckKickz, "", "/checkout/cart"},
		//"https://stormfashion.dk":			{ParseStormfashion, , "", "None"},,
		"https://www.basket4ballers.com":   {ParseBasket4ballers, CheckBasket4ballers, "basket4ballers", "/en/commande"},
		"https://www.thegoodlifespace.com": {ParseThegoodlifespace, CheckTheGoodLifeSpace, "thegoodlifespace", "/cart"},
		"https://www.tike.rs":              {ParseTike, CheckTike, "tike", "/kupovina"},
		"https://www.supplystore.com.au":   {ParseSupplystore, CheckSupplyStore, "supplystore", "/shop/checkout/cart.aspx"},
		"http://www.selfridges.com":        {ParseSelfridges, CheckSelfridges, "selfridges", "/US/en/cat/Basket/"},
		"https://www.brownsfashion.com":    {ParseBrownsfashion, CheckBrownFashion, "brownsfashion", "/commerce/bag"},
		//"":			{, CheckDistance, "", ""},
		"https://kicksstore.eu": {ParseKicksstore, CheckKicksStore, "kicksstore", "/basket"},
		//"https://www.walk-inmycloset.com":			{ParseWalkinmycloset, CheckWalkinmycloset, "", ""},
		"https://www.ntrstore.com":      {ParseNtrstore, CheckNtrStore, "ntrstore", "/checkout/cart/"},
		"https://oki-ni.com":            {ParseOkini, CheckOkiNi, "oki-ni", "/checkout/cart/"},
		"https://www.overkillshop.com":  {ParseOverkillshop, CheckOverKill, "overkillshop", "/en/checkout/cart/"},
		"https://www.triads.co.uk":      {ParseTriads, CheckTriads, "triads", "/basket"},
		"https://www.uebervart-shop.de": {ParseUebervart, CheckUebervart, "uebervart-shop", "/cart/"},
		"https://nowhere.ie":            {ParseNowhere, CheckNowhere, "nowhere", "/cart"},
	}

	//AtcM contains all the atcs
	AtcM = map[string]ATC{
		"https://www.mrporter.com":         AtcMrporter,
		"https://goodhoodstore.com":        AtcGoodhoodstore,
		"https://www.nakedcph.com":         AtcNaked,
		"https://hbx.com":                  AtcHbx,
		"https://www.solebox.com":          AtcSolebox,
		"https://www.nittygrittystore.com": AtcNittygritty,
		"https://www.woodwood.com":         AtcWoodwood,
		"https://www.sneakers76.com":       AtcSneakers76,
		"https://www.43einhalb.com":        Atc43einhalb,
		"https://www.size.co.uk":           AtcSize,
		"https://en.titolo.ch":             AtcTitolo,
		"https://caliroots.com":            AtcCaliroots,
		"https://www.allikestore.com":      AtcAllikestore,
		"https://www.thenextdoor.fr":       AtcThenextdoor,
		"https://asphaltgold.de":           AtcAsphaltgold,
		"https://www.footshop.eu":          AtcFootshop,
		"https://chmielna20.pl":            AtcChmielna20,
		"https://www.shinzo.paris":         AtcShinzo,
		"https://www.sneakersnstuff.com":   AtcSneakersnstuff,
		"https://sneakerstudio.com":        AtcSneakerstudio,
		"https://tres-bien.com":            AtcTresbien,
		"https://www.urbanjunglestore.com": AtcUrbanjunglestore,
		"https://www.vooberlin.com":        AtcVooberlin,
		"https://www.ssense.com":           AtcSsense,
		"http://www.shiekhshoes.com":       AtcShiekhshoes,
		"https://www.barneys.com":          AtcBarneys,
		"https://ist.290sqm.com":           AtcIst290,
		"https://www.footpatrol.com":       AtcFootpatrol,
		"https://wellgosh.com":             AtcWellgosh,
		"https://www.km20.ru":              AtcKm20,
		"https://www.summer-store.com":     AtcSummerstore,
		"https://www.wiseboutique.com":     AtcWiseboutique,
		"https://www.walk-inmycloset.com":  AtcWalkinmycloset,
		"https://www.offspring.co.uk":      AtcOffspring,
		"https://oqium.com":                AtcOqium,
	}

	client = &fasthttp.Client{
		ReadTimeout:                   10 * time.Second,
		DisableHeaderNamesNormalizing: true,
		ReadBufferSize:                16384,
	}

	userAgentsIter = 0
)

// GlobalWrapper contains all the monitor stuff
type GlobalWrapper struct {
	Updater  ProductsUpdater
	Checker  AvailabilityChecker
	ShopName string
	Basket   string
}

// ATC implements atc
type ATC func(link string) map[string]string

// ProductsUpdater parses stuff
type ProductsUpdater func() ([]*Product, error)

// AvailabilityChecker shecks if product is restocked
type AvailabilityChecker func(string) (bool, error)

// Product is a schop product
type Product struct {
	Name     string `json:"name,required" bson:"name"`
	Link     string `json:"link,required" bson:"link"`
	ShopURL  string `json:"shop_url,required" bson:"shop_url"`
	PhotoURL string `json:"photo_url,omitempty" bson:"photo_url"`
	Type     string `json:"type" bson:"type"`
}

// Markdown returns markdown string containing product info
func (p Product) Markdown() string {
	return fmt.Sprintf(`
		**[%s](%s)**`,
		p.Name,
		p.Link,
	)
}

// Hash returns sha224 hash from product
func (p Product) Hash() [28]byte {
	return sha256.Sum224(
		[]byte(
			p.Name + p.Link,
		),
	)
}

type cookie struct {
	key   string
	value string
}

// NewProduct creates a new Product instance
func NewProduct(name, link, shopURL, photoURL string) *Product {
	product := new(Product)

	product.Name = name
	product.Link = link
	product.ShopURL = shopURL
	product.PhotoURL = photoURL
	product.Type = "new arrival"

	return product
}

func getDoc(shopURL string, sanitize bool, cookies ...cookie) (doc *goquery.Document, err error) {
	body, err := parseHTML(shopURL, cookies)
	if err != nil {
		return doc, err
	}

	if sanitize {
		hdata := strings.Replace(string(body), "<noscript>", "", -1)
		hdata = strings.Replace(hdata, "</noscript>", "", -1)
		doc, err = goquery.NewDocumentFromReader(strings.NewReader(hdata))
	} else {
		doc, err = goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	}

	return doc, err
}

func parseHTML(url string, cookies []cookie) ([]byte, error) {
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()
	cookie := fasthttp.AcquireCookie()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)
	defer fasthttp.ReleaseCookie(cookie)

	request.SetConnectionClose()
	request.SetRequestURI(url)
	request.Header.Set("User-Agent", getUserAgent())
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	request.Header.Set("X-API-Key", "5F9D749B65CD44479C1BA2AA21991925")
	request.Header.Set("Upgrade-Insecure-Requests", "1")
	request.Header.Set("Pragma", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Cache-Control", "no-cache")

	for _, c := range cookies {
		cookie.SetKey(c.key)
		cookie.SetValue(c.value)
	}

	request.Header.SetBytesV("Set-Cookie", cookie.Cookie())
	if err := client.Do(request, response); err != nil {
		return nil, err
	}

	return response.Body(), nil
}

func getUserAgent() string {
	userAgent := userAgents[userAgentsIter]
	userAgentsIter++
	if userAgentsIter == len(userAgents) {
		userAgentsIter = 0
	}
	//"FootPatrol/2.0 CFNetwork/808.3 Darwin/16.3.0 Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36"
	return "FootPatrol/2.0 CFNetwork/808.3 " + userAgent
}

func runJS(code string, data *map[string]string) {
	js := otto.New()

	var function otto.Value

	js.Set("setEnrichFunction", func(call otto.FunctionCall) otto.Value {
		function = call.Argument(0)
		if class := function.Class(); class != "Function" {
			log.Fatalf("setEnrichFunction: expected Function, got %s instead.", class)
		}
		return otto.UndefinedValue()
	})
	js.Run(code)

	arg, err := js.ToValue(*data)
	if err != nil {
		log.Fatalf("couldn't convert message to JS value")
	}

	_, err = function.Call(otto.NaNValue(), arg)
	if err != nil {
		log.Fatalf("calling enrich function failed: %v", err)
	}
}

func translateJS(code *string) {
	*code = strings.Replace(*code, "&#34;", `"`, -1)
	*code = strings.Replace(*code, "&#39;", `'`, -1)
	*code = strings.Replace(*code, "&lt;", `<`, -1)
	*code = strings.Replace(*code, "&gt;", `>`, -1)
}
