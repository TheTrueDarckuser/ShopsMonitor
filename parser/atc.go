package parser

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// AtcMrporter implements atc for mrporter
func AtcMrporter(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("option").Each(func(i int, s *goquery.Selection) {
		status, _ := s.Attr("data-stock")
		if strings.Contains(status, "In_Stock") || strings.Contains(status, "Low_Stock") {
			size := s.Text()
			id, _ := s.Attr("value")
			link := "https://www.mrporter.com/intl/api/basket/addsku/" + id + ".json"
			atc[size] = link
		}
	})
	return atc
}

// AtcGoodhoodstore implements atc for goodhoodstore
func AtcGoodhoodstore(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("option").Each(func(i int, s *goquery.Selection) {
		id, ex := s.Attr("data-id")
		if ex {
			value, _ := s.Attr("value")
			if !strings.Contains(value, "out_of_stock_email") {
				size := s.Text()
				link := "https://goodhoodstore.com/checkout?callback=Request.JSONP.request_map.request_0&id=" + id + "&action=add&qty=1&action_type=json"
				atc[size] = link
			}
		}
	})
	return atc
}

// AtcNaked implements atc for naked
func AtcNaked(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("option").Each(func(i int, s *goquery.Selection) {
		id, ex := s.Attr("value")
		if ex && len(id) == 5 {
			size := strings.TrimSpace(s.Text())
			link := " https://www.nakedcph.com/cart/add/?id=" + id
			atc[size] = link
		}
	})
	return atc
}

// AtcHbx implements atc for hbx
func AtcHbx(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("label").Each(func(i int, s *goquery.Selection) {
		id, ex := s.Attr("for")
		if ex {
			size := s.Find("span").Text()
			link := "https://hbx.com/cart/add?quantity=1&variant_id=" + id
			atc[size] = link
		}
	})
	return atc
}

// AtcSolebox implements atc for solebox
func AtcSolebox(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.size  > a").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("id")
		size := strings.TrimSpace(s.Text())
		doc.Find("input").Each(func(i int, s *goquery.Selection) {
			cnid, _ := s.Attr("value")
			if len(cnid) == 6 {
				link := "https://www.solebox.com/index.php?cl=details&cnid=" + cnid + "&anid=" + id + "&listtype=list&"
				atc[size] = link
			}
		})
	})
	return atc
}

// AtcNittygritty implements atc for nittygritty
func AtcNittygritty(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.product-picker-size__item").Each(func(i int, s *goquery.Selection) {
		check, _ := s.Attr("class")
		if !strings.Contains(check, "is-soldout") {
			size := strings.TrimSpace(s.Find("label").Text())
			atc[size] = ""
		}
	})
	return atc
}

// AtcSneakers76 implements atc for sneakers76
func AtcSneakers76(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("option").Each(func(i int, s *goquery.Selection) {
		size, _ := s.Attr("title")
		atc[size] = ""
	})
	return atc
}

// Atc43einhalb implements atc for 43einhalb
func Atc43einhalb(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("select.customSelectBox > option").Each(func(i int, s *goquery.Selection) {
		class, ex := s.Attr("class")
		if ex && !strings.Contains(class, "disabled productReminderRevealButto") {
			size := strings.TrimSpace(s.Text())
			atc[size] = ""
		}
	})
	return atc
}

// AtcSize implements atc for size
func AtcSize(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("button.btn.btn-default ").Each(func(i int, s *goquery.Selection) {
		size, _ := (s.Attr("title"))
		size = size[12:]
		atc[size] = ""

	})
	return atc
}

// AtcTitolo implements atc for titolo
func AtcTitolo(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("#attributesize-size_eu > option").Each(func(i int, s *goquery.Selection) {
		_, ex := (s.Attr("value"))
		if ex {
			size := s.Text()
			atc[size] = ""
		}
	})
	return atc
}

// AtcCaliroots implements atc for caliroots
func AtcCaliroots(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.select-product.row > select > option").Each(func(i int, s *goquery.Selection) {
		ex, _ := (s.Attr("value"))
		if len(ex) == 7 {
			size := strings.TrimSpace(s.Text())
			atc[size] = ""
			id, _ := s.Attr("value")
			link := "https://caliroots.com/cart/add/" + id
			atc[size] = link
		}
	})

	return atc
}

// AtcAllikestore implements atc for allikestore
func AtcAllikestore(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("#configurable_swatch_shoe_size > li > a").Each(func(i int, s *goquery.Selection) {
		size, _ := (s.Attr("title"))
		atc[size] = ""
	})
	return atc
}

// AtcThenextdoor implements atc for thenextdoor
func AtcThenextdoor(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.attribute_list > select > option").Each(func(i int, s *goquery.Selection) {
		stock, _ := (s.Attr("data-quantity"))
		if stock == "1" {
			size := s.Text()
			atc[size] = ""
		}
	})
	return atc
}

// AtcAsphaltgold implements atc for asphaltgold
func AtcAsphaltgold(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("ul.clearfix > li > div").Each(func(i int, s *goquery.Selection) {
		size := s.Text()
		if !(len(size) > 5) {
			atc[size] = ""
		}
	})
	return atc
}

// AtcFootshop implements atc for footshop
func AtcFootshop(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.holder > ul > li > div").Each(func(i int, s *goquery.Selection) {
		stock := s.Find("strong")
		if stock.Text() == "US" {
			span := s.Find("span")
			size := span.Text()
			atc[size] = ""
		}
	})
	return atc
}

// AtcChmielna20 implements atc for chmielna20
func AtcChmielna20(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.selector > ul > li").Each(func(i int, s *goquery.Selection) {
		size, _ := s.Attr("data-sizeeu")
		atc[size] = ""
	})
	return atc
}

// AtcShinzo implements atc for shinzo
func AtcShinzo(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.attribute_list > ul > li > label").Each(func(i int, s *goquery.Selection) {
		size := s.Text()
		atc[size] = ""
	})
	return atc
}

// AtcSneakersnstuff implements atc for sneakersnstuff
func AtcSneakersnstuff(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.product-size-container > fieldset > div > div > div > span").Each(func(i int, s *goquery.Selection) {
		size := strings.TrimSpace(s.Text())
		atc[size] = ""
	})
	return atc
}

// AtcSneakerstudio implements atc for sneakerstudio
func AtcSneakerstudio(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.product_section_sub > div.size_cyfra_parent > label").Each(func(i int, s *goquery.Selection) {
		size, _ := s.Attr("data-eu")
		atc[size] = ""
	})
	return atc
}

// AtcTresbien implements atc for tresbien
func AtcTresbien(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("span.size").Each(func(i int, s *goquery.Selection) {
		size := s.Text()
		atc[size] = ""
	})
	return atc
}

// AtcUrbanjunglestore implements atc for urbanjunglestore
func AtcUrbanjunglestore(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.input-box > ul > li > a").Each(func(i int, s *goquery.Selection) {
		size, _ := s.Attr("title")
		atc[size] = ""
	})
	return atc
}

// AtcVooberlin implements atc for vooberlin
func AtcVooberlin(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("select.w_select-item > option").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		if !strings.Contains(t, "no stock") {
			atc[strings.TrimSpace(t)] = ""
		}
	})
	return atc
}

// AtcSsense implements atc for ssense
func AtcSsense(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.add-to-bag-form > div > div > div > select > option").Each(func(i int, s *goquery.Selection) {
		size := s.Text()
		if !(strings.Contains(size, "SELECT A SIZE") || strings.Contains(size, "Sold Out")) {
			atc[strings.TrimSpace(size)] = ""
		}
	})
	return atc
}

// AtcShiekhshoes implements atc for shiekhshoes
func AtcShiekhshoes(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("a.selectSize").Each(func(i int, s *goquery.Selection) {
		size, _ := s.Attr("data-size")
		size = size + "US"
		atc[size] = ""
	})
	return atc
}

// AtcBarneys implements atc for barneys
func AtcBarneys(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("a.atg_store_oneSize.sizePicker ").Each(func(i int, s *goquery.Selection) {
		stock, _ := s.Attr("data-onhand-quantity")
		if !(stock == "0") {
			size, _ := s.Attr("data-sku-size")
			atc[size] = ""
		}
	})
	return atc
}

// AtcIst290 implements atc for ist290
func AtcIst290(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("select.form-control > option").Each(func(i int, s *goquery.Selection) {
		stock, _ := s.Attr("value")
		if len(stock) == 5 {
			size := s.Text()
			size = size[:9]
			atc[size] = ""
		}
	})
	return atc
}

// AtcWoodwood implements atc for woodwood
func AtcWoodwood(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}

	data := doc.Find("script").Next().Text()
	data = data[strings.Index(data, "'pids': ["):strings.Index(data, "], 'cartId'")]

	dataT := data[strings.Index(data, "[")+1 : strings.Index(data, "]")]
	dataIDs := strings.Split(dataT, ",")

	dataT = data[strings.Index(data, "]"):]
	dataSizes := strings.Split(strings.Replace(dataT[strings.Index(dataT, "[")+1:], "'", "", -1), ",")

	atc := make(map[string]string)
	for i := 0; i < len(dataSizes); i++ {
		atc[dataSizes[i]] = fmt.Sprintf("https://www.woodwood.com/cart?action=add&item_pid=%s", dataIDs[i])
	}

	return atc
}

// AtcFootpatrol implements atc for footpatrol
func AtcFootpatrol(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("select.form-control > option").Each(func(i int, s *goquery.Selection) {
		stock, _ := s.Attr("value")
		if len(stock) == 5 {
			size := s.Text()
			size = size[:9]
			atc[size] = ""
		}
	})
	return atc
}

// AtcWellgosh implements atc for wellgosh
func AtcWellgosh(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}

	//pattern, _ := doc.Find("#product_addtocart_form").Attr("action")
	//pattern = strings.Replace(pattern, ",", "", -1)
	//pattern = pattern[:len(pattern)-1]

	raw := doc.Find("#product-options-wrapper script").Text()
	raw = raw[strings.Index(raw, `"options":`)+len(`"options":`):]

	priceID := raw[strings.Index(raw, `"id":"`)+len(`"id":"`):]
	priceID = priceID[:strings.Index(priceID, `"`)]

	raw = raw[strings.Index(raw, `"options":`)+len(`"options":`):]
	raw = raw[:strings.Index(raw, "]}}")] + "]"

	type productInfo struct {
		ID    string
		Label string
	}

	var data []productInfo
	atc := make(map[string]string)
	if err := json.Unmarshal([]byte(raw), &data); err != nil {
		return atc
	}

	//productID, _ := doc.Find("span.no-display input").Attr("value")
	for _, p := range data {
		atc[p.Label] = "" //pattern + fmt.Sprintf("?product=%s&super_attribute[162]=%s&super_attribute[163]=%s", productID, priceID, p.ID)
	}

	return atc
}

// AtcUrbanStaroma implements atc for ist290
func AtcUrbanStaroma(link string) map[string]string {
	doc, err := getDoc(link, false)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.attribute_list > select > option").Each(func(i int, s *goquery.Selection) {
		size := s.Text()
		atc[size] = ""
	})
	return atc
}

// AtcKm20 implements atc for km20
func AtcKm20(link string) map[string]string {
	doc, err := getDoc(link, true)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("#SKU_LIST > option").Each(func(i int, s *goquery.Selection) {
		size, _ := s.Attr("data-product-size")
		if len(size) > 0 {
			atc[size] = ""
		}
	})
	return atc
}

// AtcSummerstore implements atc for summerstore
func AtcSummerstore(link string) map[string]string {
	doc, err := getDoc(link, true)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("#group_5 > option").Each(func(i int, s *goquery.Selection) {
		atc[s.Text()] = ""
	})
	return atc
}

// AtcWiseboutique implements atc for wiseboutique
func AtcWiseboutique(link string) map[string]string {
	doc, err := getDoc(link, true)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.taglia").Each(func(i int, s *goquery.Selection) {
		atc[s.Text()] = ""
	})
	return atc
}

// AtcWalkinmycloset implements atc for walkinmycloset
func AtcWalkinmycloset(link string) map[string]string {
	doc, err := getDoc(link, true)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("#ProductSelect > option").Each(func(i int, s *goquery.Selection) {
		size := strings.Fields(s.Text())[0]
		id, _ := s.Attr("value")
		atc[size] = fmt.Sprintf("https://www.walk-inmycloset.com/cart/add.js?id=%s", id)
	})
	return atc
}

// AtcOffspring implements atc for offspring
func AtcOffspring(link string) map[string]string {
	doc, err := getDoc(link, true)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("#sizeShoe > option").Each(func(i int, s *goquery.Selection) {
		if t, _ := s.Attr("value"); len(t) > 0 {
			atc[s.Text()] = ""
		}
	})
	return atc
}

// AtcOqium implements atc for oqium
func AtcOqium(link string) map[string]string {
	doc, err := getDoc(link, true)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.product-swatches > div.swatch-other > div").Each(func(i int, s *goquery.Selection) {
		size, _ := s.Attr("data-product-swatch-value")
		atc[strings.ToLower(size)] = ""
	})
	return atc
}

/* // AtcBrandShop implements atc for oqium
func AtcBrandShop(link string) map[string]string {
	doc, err := getDoc(link, true)
	if err != nil {
		return nil
	}
	atc := make(map[string]string)
	doc.Find("div.select-box > div.options").Each(func(i int, s *goquery.Selection) {
		size := s.Find("div > span.text").Text()
		atc[size] = ""
	})
	return atc
} */
