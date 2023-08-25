package parser

import (
	"net/http"
	"testing"
)

func TestKm20(t *testing.T) {
	url := "https://www.km20.ru/catalog/product/104944/"
	status, err := CheckKm20(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.km20.ru/test"
	status, err = CheckKm20(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSneakersnstuff(t *testing.T) {
	url := "https://www.sneakersnstuff.com/en/product/33239/nike-air-vapormax-fk-moc-2-acronym"
	status, err := CheckSneakersnstuff(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.sneakersnstuff.com/test"
	status, err = CheckSneakersnstuff(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSsense(t *testing.T) {
	url := "https://www.ssense.com/en-ru/men/product/loewe/grey-anagram-polo/3479439"
	status, err := CheckSsense(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.ssense.com/test"
	status, err = CheckSsense(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestFootpatrol(t *testing.T) {
	url := "https://www.footpatrol.com/product/white-nike-x-martine-rose-air-monarch-iv-womens/119832_footpatrolcom/"
	status, err := CheckFootpatrol(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.footpatrol.com/test"
	status, err = CheckFootpatrol(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestProdirectsoccer(t *testing.T) {
	url := "https://www.prodirectsoccer.com/products/Nike-Mercurial-Superfly-VI-Elite-FG-Mens-Boots-Firm-Ground-Racer-Blue-Metallic-Silver-Black-192114.aspx"
	status, err := CheckProdirectsoccer(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.prodirectsoccer.com/test"
	status, err = CheckProdirectsoccer(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestThebrokenarm(t *testing.T) {
	url := "https://www.the-broken-arm.com/fr/pantalons/3538-maison-margiela-pantalon.html#/taille-36"
	status, err := CheckThebrokenarm(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.the-broken-arm.com/test"
	status, err = CheckThebrokenarm(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestTitan22(t *testing.T) {
	url := "https://www.titan22.com/catalog/product/view/id/15295/s/nike-lebron-16-gs-i-m-king/"
	status, err := CheckTitan22(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.titan22.com/test"
	status, err = CheckTitan22(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestNosstore(t *testing.T) {
	url := "https://www.nosstore.eu/en/product/nike-air-force-1-utility/a12769/"
	status, err := CheckNosstore(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.nosstore.eu/test"
	status, err = CheckNosstore(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestOqium(t *testing.T) {
	url := "https://oqium.com/collections/footwear/products/legacy-312-wolf-grey?variant=18362075938912"
	status, err := CheckOqium(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://oqium.com/test"
	status, err = CheckOqium(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestOffspring(t *testing.T) {
	url := "https://www.offspring.co.uk/view/product/offspring_catalog/NEWIN/2003112745"
	status, err := CheckOffspring(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.offspring.co.uk/test"
	status, err = CheckOffspring(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestWalkinmycloset(t *testing.T) {
	url := "https://www.walk-inmycloset.com/collections/adidas/products/adidas-original-eqt-support-mid-adv-primeknit-b37513"
	status, err := CheckWalkinmycloset(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.walk-inmycloset.com/test"
	status, err = CheckWalkinmycloset(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestWiseboutique(t *testing.T) {
	url := "https://www.wiseboutique.com/en/black+polyamide+blend+baseball+cap-givenchy-bpz001k0ce001"
	status, err := CheckWiseboutique(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.wiseboutique.com/test"
	status, err = CheckWiseboutique(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSummerstore(t *testing.T) {
	url := "https://www.summer-store.com/en/shoes/5614-converse-x-cdg-play-ct-70-s-white.html"
	status, err := CheckSummerstore(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.summer-store.com/test"
	status, err = CheckSummerstore(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestKicksStore(t *testing.T) {
	url := "https://kicksstore.eu/wmns-air-jordan-4-retro-nrg-aq9128-600.html"
	status, err := CheckKicksStore(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://kicksstore.eu/test"
	status, err = CheckKicksStore(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestDistance(t *testing.T) {
	url := "https://distancewear.co.uk/collections/mens/products/transform-performance-tee-forest-grey"
	status, err := CheckDistance(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://distancewear.co.uk/test"
	status, err = CheckDistance(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestBrownFashion(t *testing.T) {
	url := "https://www.brownsfashion.com/ru/shopping/white-and-red-replica-painted-toes-sneakers-13400107"
	status, err := CheckBrownFashion(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.brownsfashion.com/test"
	status, err = CheckBrownFashion(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSelfridges(t *testing.T) {
	url := "http://www.selfridges.com/US/en/cat/prada-contrast-trim-crewneck-cotton-pique-t-shirt_1106-3005868-SJN235S191/"
	status, err := CheckSelfridges(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "http://www.selfridges.com/test"
	status, err = CheckSelfridges(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSupplyStore(t *testing.T) {
	url := "https://www.supplystore.com.au/brands/new-arrivals/call-me-917-glam-deck-8.25/c-28/c-150/p-186065"
	status, err := CheckSupplyStore(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.supplystore.com.au/test"
	status, err = CheckSupplyStore(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestTike(t *testing.T) {
	url := "https://www.tike.rs/patike/37778-daytona-dmx-mu"
	status, err := CheckTike(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.tike.rs/test"
	status, err = CheckTike(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestTheGoodLifeSpace(t *testing.T) {
	url := "https://www.thegoodlifespace.com/collections/asics-footwear-men/products/asics-tiger-gel-kayano-5-og-1191a099-101"
	status, err := CheckTheGoodLifeSpace(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.thegoodlifespace.com/test"
	status, err = CheckTheGoodLifeSpace(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestBasket4ballers(t *testing.T) {
	url := "https://www.basket4ballers.com/fr/kobe/19050-nike-kobe-ad-black-racer-blue-av3555-003.html"
	status, err := CheckBasket4ballers(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.basket4ballers.com/test"
	status, err = CheckBasket4ballers(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestStorm(t *testing.T) {
	url := "https://www.stormonline.com/product/brush-stroke-silk-top-20570323"
	status, err := CheckStorm(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.stormonline.com/test"
	status, err = CheckStorm(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestKickz(t *testing.T) {
	url := "https://www.kickz.com/us/under-armour-basketball-shoes-ua-anomaly-ash_taupe_gray_flux_black-149374001"
	status, err := CheckKickz(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.kickz.com/test"
	status, err = CheckKickz(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestFeatureSneakerBoutique(t *testing.T) {
	url := "https://www.featuresneakerboutique.com/products/adidas-adilette-slides-adi-blue"
	status, err := CheckFeatureSneakerBoutique(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.featuresneakerboutique.com/test"
	status, err = CheckFeatureSneakerBoutique(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestLeray(t *testing.T) { //TODO
	url := "https://www.featuresneakerboutique.com/products/adidas-adilette-slides-adi-blue"
	status, err := CheckLeray(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.featuresneakerboutique.com/test"
	status, err = CheckLeray(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestUrbanStaroma(t *testing.T) {
	url := "https://www.urbanstaroma.com/en/men-s-sneakers/6422-adidas-sc-premiere-sneakers-bd7583.html"
	status, err := CheckUrbanStaroma(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.urbanstaroma.com/test"
	status, err = CheckUrbanStaroma(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestDoubleClutch(t *testing.T) {
	url := "https://www.doubleclutch.it/it/jordan-xix-flint-melo-aq9213-100"
	status, err := CheckDoubleClutch(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.doubleclutch.it/test"
	status, err = CheckDoubleClutch(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestBaitMe(t *testing.T) {
	url := "http://www.baitme.com/nike-men-nrg-a14-beanie-black-yellow-ochre-white-niav4774-010"
	status, err := CheckBaitMe(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "http://www.baitme.com/test"
	status, err = CheckBaitMe(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestBrandShop(t *testing.T) {
	url := "https://brandshop.ru/goods/150492/krossovki-adidas-originals-yung-1-grey-one-grey-one-white/"
	status, err := CheckBrandShop(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://brandshop.ru/test"
	status, err = CheckBrandShop(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestCruvoir(t *testing.T) {
	url := "https://cruvoir.com/products/check-dress-shirt"
	status, err := CheckCruvoir(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://cruvoir.com/test"
	status, err = CheckCruvoir(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSpace23(t *testing.T) {
	url := "https://www.space23.it/products/nike-aq4137-w-air-max-97-se-scarpe-donna-vast-grey-mtlc"
	status, err := CheckSpace23(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.space23.it/test"
	status, err = CheckSpace23(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestFootish(t *testing.T) {
	url := "https://www.footish.se/sneakers/new-balance/new-balance-x-90-msx90rpa"
	status, err := CheckFootish(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.footish.se/test"
	status, err = CheckFootish(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestExcelsiormilano(t *testing.T) { //TODO
	url := "https://www.footish.se/sneakers/new-balance/new-balance-x-90-msx90rpa"
	status, err := CheckExcelsiormilano(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.footish.se/test"
	status, err = CheckExcelsiormilano(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestBouncewear(t *testing.T) {
	url := "https://bouncewear.com/en/product/wmns-air-jordan-1-jester-xx"
	status, err := CheckBouncewear(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://bouncewear.com/test"
	status, err = CheckBouncewear(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestAwlabr(t *testing.T) {
	url := "https://en.aw-lab.com/shop/men/shoes/nike-sb-nyjah-8896123"
	status, err := CheckAwlab(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://en.aw-lab.com/test"
	status, err = CheckAwlab(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestAntonioli(t *testing.T) {
	url := "https://www.antonioli.eu/en/RU/men/products/cosui18002-green"
	status, err := CheckAntonioli(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.antonioli.eu/test"
	status, err = CheckAntonioli(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestFootlocker(t *testing.T) { //TODO
	url := "https://www.footlocker.co.uk/en/p/converse-erx-260-men-shoes-71405?v=314100724104#!searchCategory=men/shoes/basketball"
	status, err := CheckFootlocker(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.footlocker.co.uk/test"
	status, err = CheckFootlocker(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestFootdistrict(t *testing.T) { //TODO
	url := "https://footdistrict.com/en/nike-air-monarch-iv-martine-rose-at3147-600.html"
	status, err := CheckFootdistrict(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://footdistrict.com/test"
	status, err = CheckFootdistrict(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestNetaporter(t *testing.T) {
	url := "https://www.net-a-porter.com/ru/en/product/1097619/isabel_marant/dolmen-asymmetric-pleated-satin-midi-skirt"
	status, err := CheckNetaporter(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.net-a-porter.com/test"
	status, err = CheckNetaporter(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestStarcowparis(t *testing.T) {
	url := "https://www.starcowparis.com/footwear/1891-asics-gel-kayano-5-og.html"
	status, err := CheckStarcowparis(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.starcowparis.com/test"
	status, err = CheckStarcowparis(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestNordstorm(t *testing.T) { //TODO
	url := "https://shop.nordstrom.com/s/cole-haan-original-grand-wingtip-men/4630799?origin=category-personalizedsort&breadcrumb=Home%2FMen%2FShoes%2FDress%20Shoes&color=chorino%20marine%20matte%20leather"
	status, err := CheckNordstorm(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://shop.nordstrom.com/test"
	status, err = CheckNordstorm(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestJimmyjazz(t *testing.T) { //TODO
	url := "https://www.starcowparis.com/footwear/1891-asics-gel-kayano-5-og.html"
	status, err := CheckJimmyjazz(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.starcowparis.com/test"
	status, err = CheckJimmyjazz(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestYmeuniverse(t *testing.T) {
	url := "https://www.ymeuniverse.com/en/product/8254/thunder"
	status, err := CheckYmeuniverse(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.ymeuniverse.com/test"
	status, err = CheckYmeuniverse(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestBstn(t *testing.T) {
	url := "https://www.bstn.com/p/vans-ua-comfycush-sk8-h-va45jzvny-115427"
	status, err := CheckBstn(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.bstn.com/test"
	status, err = CheckBstn(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func Test290sqm(t *testing.T) {
	url := "https://ist.290sqm.com/Men?product_id=14281"
	status, err := Check290sqm(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://ist.290sqm.com/test"
	status, err = Check290sqm(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestBarneys(t *testing.T) { //TODO
	url := "https://www.barneys.com/product/thom-browne-men-27s-pebbled-leather-wingtip-sneakers-506035353.html"
	status, err := CheckBarneys(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.barneys.com/test"
	status, err = CheckBarneys(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestCitygear(t *testing.T) {
	url := "https://www.citygear.com/nike-air-force-1-07-lv8-4-15.html"
	status, err := CheckCitygear(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.citygear.com/test"
	status, err = CheckCitygear(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestShiek(t *testing.T) { //TODO
	url := "https://www.shiekh.com/adidas-aeroblu-scarlet-conavy-continental-80-aeroblu-scarlet-conavy.html"
	status, err := CheckShiek(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.shiekh.com/test"
	status, err = CheckShiek(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestNorsestore(t *testing.T) {
	url := "https://www.norsestore.com/commodity/28131-nanamica-down-coat"
	status, err := CheckNorsestore(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.norsestore.com/test"
	status, err = CheckNorsestore(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestMrporter(t *testing.T) {
	url := "https://www.mrporter.com/en-ru/mens/saint_laurent/robot-intarsia-wool-blend-sweater/1110935?ppv=2"
	status, err := CheckMrporter(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.mrporter.com/test"
	status, err = CheckMrporter(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func Test1stog(t *testing.T) { //TODO
	url := "https://www.1st-og.ch/air-max-plus-tn-se-av7021-001"
	status, err := Check1stog(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.1st-og.ch/11"
	status, err = Check1stog(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestVooberlin(t *testing.T) {
	url := "https://www.vooberlin.com/detail/index/sArticle/14268"
	status, err := CheckVooberlin(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.vooberlin.com/test"
	status, err = CheckVooberlin(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestUrbanjungle(t *testing.T) {
	url := "https://www.urbanjunglestore.com/it/sneakers/urban-jungle-sneakers/urban-jungle-uj-slider.html"
	status, err := CheckUrbanjungle(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.urbanjunglestore.com/test"
	status, err = CheckUrbanjungle(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestTresbien(t *testing.T) {
	url := "https://tres-bien.com/stone-island-nylon-metal-swim-shorts-grey-mo7015b0943-v0061-ss19"
	status, err := CheckTresbien(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://tres-bien.com/test"
	status, err = CheckTresbien(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestTitolo(t *testing.T) { //TODO
	url := "https://en.titolo.ch/wmns-air-huarache-run-634835-702"
	status, err := CheckTitolo(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://en.titolo.ch/test"
	status, err = CheckTitolo(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSneakerstudio(t *testing.T) {
	url := "https://sneakerstudio.com.ua/product-rus-17273-%D0%9C%D1%83%D0%B6%D1%81%D0%BA%D0%B8%D0%B5-%D0%B1%D0%BE%D1%82%D0%B8%D0%BD%D0%BA%D0%B8-Timberland-Premium-6-IN-A1QXS.html"
	status, err := CheckSneakerstudio(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://sneakerstudio.com.ua/test"
	status, err = CheckSneakerstudio(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestShinzo(t *testing.T) {
	url := "https://www.shinzo.paris/en/sportswear-shoes/6245-the-north-face-base-camp-slide-2-t93fwobk9.html"
	status, err := CheckShinzo(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.shinzo.paris/test"
	status, err = CheckShinzo(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestRuvilla(t *testing.T) {
	url := "https://www.dtlr.com/jordan-air-jordan-retro-8-gs-snowflake-305368-400.html"
	status, err := CheckRuvilla(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.dtlr.com/test"
	status, err = CheckRuvilla(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestPatta(t *testing.T) {
	url := "https://www.patta.nl/nike-x-martine-rose-air-monarch-iv-atmosphere-grey-gunsmoke"
	status, err := CheckPatta(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.patta.nl/test"
	status, err = CheckPatta(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestGoodhoodstore(t *testing.T) {
	url := "https://goodhoodstore.com/store/converse-one-star-ox-leopard-fur-41699"
	status, err := CheckGoodhoodstore(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://goodhoodstore.com/test"
	status, err = CheckGoodhoodstore(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestChmielna20(t *testing.T) {
	url := "https://chmielna20.pl/buty-w-af1-sage-hi-aq2771-001.html"
	status, err := CheckChmielna20(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://chmielna20.pl/test"
	status, err = CheckChmielna20(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestConsortium(t *testing.T) {
	url := "https://www.consortium.co.uk/nike-sb-dunk-low-pro-light-cream-light-cream-light-cream-thunderstorm-bq6817-200.html"
	status, err := CheckConsortium(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.consortium.co.uk/test"
	status, err = CheckConsortium(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestFootshop(t *testing.T) {
	url := "https://www.footshop.eu/en/mens-shoes/37411-reebok-aztrek-white-black-cobalt-yellow.html"
	status, err := CheckFootshop(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.footshop.eu/test"
	status, err = CheckFootshop(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestAsphaltgold(t *testing.T) {
	url := "https://asphaltgold.de/de/nike-air-jordan-4-retro-se-black-laser-black-white-gum-light-brown.html"
	status, err := CheckAsphaltgold(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://asphaltgold.de/test"
	status, err = CheckAsphaltgold(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestCornerstreet(t *testing.T) {
	url := "https://www.cornerstreet.fr/catalog/product/view/id/135661/s/nike-air-safari-se-sp19-university-red/"
	status, err := CheckCornerstreet(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.cornerstreet.fr/test"
	status, err = CheckCornerstreet(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestThenextdoor(t *testing.T) { //not working site
	url := "https://www.cornerstreet.fr/catalog/product/view/id/135661/s/nike-air-safari-se-sp19-university-red/"
	status, err := CheckThenextdoor(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.cornerstreet.fr/test"
	status, err = CheckThenextdoor(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestEnd(t *testing.T) {
	url := "https://www.endclothing.com/us/our-legacy-knitted-hat-1pre199khs.html"
	status, err := CheckEnd(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.endclothing.com/test"
	status, err = CheckEnd(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestWellgosh(t *testing.T) {
	url := "https://wellgosh.com/nike-air-max-2-light-blue-lagoon-white-black-blue-lagoon-laser-orange-ao1741-100.html"
	status, err := CheckWellgosh(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://wellgosh.com/test"
	status, err = CheckWellgosh(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestAllikestore(t *testing.T) {
	url := "https://www.allikestore.com/german/reebok-daytona-dmx-mu-black-purple-cn8386.html"
	status, err := CheckAllikestore(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.allikestore.com/test"
	status, err = CheckAllikestore(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestFinishline(t *testing.T) { //TODO
	url := "https://www.finishline.com/store/product/mens-nike-air-max-97-casual-shoes/prod2770629?styleId=921826&colorId=201"
	status, err := CheckFinishline(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.finishline.com/test"
	status, err = CheckFinishline(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestCaliroots(t *testing.T) { //TODO
	url := "https://caliroots.com/nike-air-force-1-07-premium-2-at4143-400/p/117665"
	status, err := CheckCaliroots(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://caliroots.com/test"
	status, err = CheckCaliroots(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestHbx(t *testing.T) { //TODO the button is loading
	url := "https://hbx.com/brands/common-projects/original-achilles-low-suede-1"
	status, err := CheckHbx(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://hbx.com/test"
	status, err = CheckHbx(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSize(t *testing.T) { //TODO
	url := "https://www.size.co.uk/product/blue-adidas-originals-marathon-tr/066811/"
	status, err := CheckSize(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.size.co.uk/test"
	status, err = CheckSize(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func Test43einhlab(t *testing.T) {
	url := "https://www.43einhalb.com/nike-air-max-2-light-university-gold-270537"
	status, err := Check43einhlab(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.43einhalb.com/test"
	status, err = Check43einhlab(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSneakers76(t *testing.T) {
	url := "https://www.sneakers76.com/it/nike/4172-nike-air-max2-light-ao1741-100.html"
	status, err := CheckSneakers76(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.sneakers76.com/test"
	status, err = CheckSneakers76(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSlamjamsocialism(t *testing.T) {
	url := "https://www.slamjamsocialism.com/low-top/58942-wmns-hypersleek-sneakers.html"
	status, err := CheckSlamjamsocialism(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.slamjamsocialism.com/test"
	status, err = CheckSlamjamsocialism(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestLuisaviaroma(t *testing.T) { //TODO
	url := "https://www.luisaviaroma.com/ru-ru/p/valentino-garavani/%D0%BC%D1%83%D0%B6%D1%87%D0%B8%D0%BD%D1%8B/%D0%B3%D0%BE%D0%BB%D0%BE%D0%B2%D0%BD%D1%8B%D0%B5-%D1%83%D0%B1%D0%BE%D1%80%D1%8B/69I-3GV004?ColorId=WTI40&SubLine=accessories&CategoryId=53&lvrid=_p_dAG1_gm_c53"
	status, err := CheckLuisaviaroma(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.luisaviaroma.com/test"
	status, err = CheckLuisaviaroma(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestWoodwood(t *testing.T) { //TODO
	url := "https://www.woodwood.com/commodity/6460-double-a-ted-shirt"
	status, err := CheckWoodwood(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.woodwood.com/test"
	status, err = CheckWoodwood(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestThewebster(t *testing.T) {
	url := "https://thewebster.us/shop/hemp-leaf-jumper-111837.html"
	status, err := CheckThewebster(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://thewebster.us/test"
	status, err = CheckThewebster(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestNittyGritty(t *testing.T) {
	url := "https://www.nittygrittystore.com/men/footwear/nike-grandstand-ii"
	status, err := CheckNittyGritty(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.nittygrittystore.com/test"
	status, err = CheckNittyGritty(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestNaked(t *testing.T) {
	url := "https://www.nakedcph.com/adidas-originals-falcon-db2714/p/7273"
	status, err := CheckNaked(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.nakedcph.com/test"
	status, err = CheckNaked(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSolebox(t *testing.T) {
	url := "https://www.solebox.com/Footwear/Running/Air-Max-2-Light-variant.html"
	status, err := CheckSolebox(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.solebox.com/test"
	status, err = CheckSolebox(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestSivasdelcalzo(t *testing.T) {
	url := "https://www.sivasdescalzo.com/en/vans-ua-old-skool-vn0a38g1umk"
	status, err := CheckSivasdelcalzo(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.sivasdescalzo.com/test"
	status, err = CheckSivasdelcalzo(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestShoezGalleryEn(t *testing.T) {
	url := "https://www.shoezgallery.com/en/p11528-gsm-asics-d5k2y-100"
	status, err := CheckShoezGalleryEn(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.shoezgallery.com/test"
	status, err = CheckShoezGalleryEn(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestNighshop(t *testing.T) { //TODO
	url := "https://www.shoezgallery.com/en/p11528-gsm-asics-d5k2y-100"
	status, err := CheckNighshop(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.shoezgallery.com/test"
	status, err = CheckNighshop(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestNtrStore(t *testing.T) {
	url := "https://www.ntrstore.com/adidas-originals-deerupt-w-ornge-active-purple-cg6095"
	status, err := CheckNtrStore(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.ntrstore.com/test"
	status, err = CheckNtrStore(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestOkiNi(t *testing.T) {
	url := "https://oki-ni.com/cl-leather-urge-sneaker-in-grey-white"
	status, err := CheckOkiNi(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://oki-ni.com/test"
	status, err = CheckOkiNi(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestOverKill(t *testing.T) {
	url := "https://www.overkillshop.com/en/bradley-theodore-x-puma-thunder-369394-01.html"
	status, err := CheckOverKill(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.overkillshop.com/test"
	status, err = CheckOverKill(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestTriads(t *testing.T) {
	url := "https://www.triads.co.uk/triads-mens-c1/footwear-c24/flip-flop-sandals-c232/rage-bc-slide-ii-black-aztec-blue-p87013"
	status, err := CheckTriads(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.triads.co.uk/test"
	status, err = CheckTriads(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestUebervart(t *testing.T) {
	url := "https://www.uebervart-shop.de/nike-air-max-2-light-white-blue/"
	status, err := CheckUebervart(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://www.uebervart-shop.de/test"
	status, err = CheckUebervart(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}

func TestNowhere(t *testing.T) {
	url := "https://nowhere.ie/products/adidas-by-white-mountaineering-terrex-two-gtx"
	status, err := CheckNowhere(url)
	if err != nil {
		t.Error(err)
	}
	if !status {
		response, err := http.Get(url)
		if err != nil {
			t.Error(err)
		}

		if response.StatusCode != 404 {
			t.Errorf("wrong status code")
		}
	}

	falseURL := "https://nowhere.ie/test"
	status, err = CheckNowhere(falseURL)
	if err != nil {
		t.Error(err)
	}
	if status {
		t.Errorf("found not existing product %t", status)
	}

}
