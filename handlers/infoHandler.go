package handlers

import (
	"assignment_1/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type InfoResponse struct {
	/*
		{
			"name": "Norway",
			"continents": ["Europe"],
			"population": 4700000,
			"languages": {"nno":"Norwegian Nynorsk","nob":"Norwegian Bokmål","smi":"Sami"},
			"borders": ["FIN","SWE","RUS"],
			"flag": "https://flagcdn.com/w320/no.png",
			"capital": "Oslo",
			"cities": ["Abelvaer","Adalsbruk","Adland","Agotnes","Agskardet","Aker","Akkarfjord","Akrehamn","Al","Alen","Algard","Almas","Alta","Alvdal","Amli","Amot","Ana-Sira","Andalsnes","Andenes","Angvika","Ankenes","Annstad","Ardal","Ardalstangen","Arendal","Arland","Arneberg","Arnes","Aros","As","Asen","Aseral","Asgardstrand","Ask","Asker","Askim","Aukra","Auli","Aurdal","Aure","Aursmoen","Austbo","Austbygdi","Austevoll","Austmarka","Baerums verk","Bagn","Balestrand","Ballangen","Ballstad","Bangsund","Barkaker","Barstadvik","Batnfjordsora","Batsto","Beisfjord","Beitostolen","Bekkjarvik","Berge","Bergen","Berger","Berkak","Birkeland","Birtavarre","Bjaland","Bjerka","Bjerkvik","Bjoneroa","Bjordal","Bjorke","Bjorkelangen","Bjornevatn","Blaker","Blakset","Bleikvasslia","Bo","Bomlo","Bones","Borge","Borgen","Borhaug","Borkenes","Borregard","Bostad","Bovagen","Boverfjorden","Brandbu","Brandval","Brattholmen","Brattvag","Brekke","Brekstad","Brennasen","Brevik","Bronnoysund","Bru","Bruflat","Brumunddal","Brusand","Bruvik","Bryne","Bud","Burfjord","Buskerud","Buvika","Byglandsfjord","Bygstad","Bykle","Byrknes Nordre","Cavkkus","Dal","Dale","Dalen","Davik","Deknepollen","Digermulen","Dilling","Dimmelsvik","Dirdal","Disena","Dokka","Dolemo","Dovre","Drag","Drammen","Drangedal","Drobak","Dverberg","Dyrvika","Ebru","Egersund","Eggedal","Eggkleiva","Eide","Eidfjord","Eidsa","Eidsberg","Eidsdal","Eidsfoss","Eidsnes","Eidsvag","Eidsvoll","Eidsvoll verk","Eikanger","Eikelandsosen","Eiken","Eina","Eivindvik","Elverum","Enebakkneset","Enga","Engalsvik","Erdal","Erfjord","Ervik","Espeland","Etne","Evanger","Evenskjer","Evje","Eydehavn","Faberg","Faervik","Fagernes","Fagerstrand","Fall","Fardal","Farsund","Fauske","Feda","Fedje","Feiring","Felle","Fenstad","Fetsund","Fevik","Figgjo","Finnoy","Finnsnes","Finsand","Fiska","Fiskum","Fister","Fitjar","Fjellstrand","Fla","Flam","Flateby","Flekke","Flekkefjord","Flemma","Flesberg","Flesnes","Floro","Florvag","Foldereid","Folderoy","Folkestad","Follafoss","Follebu","Follese","Fonnes","Forde","Fornebu","Fosnavag","Fossdalen","Fosser","Fotlandsvag","Fredrikstad","Frekhaug","Fresvik","Frogner","Froland","From","Furnes","Fyrde","Fyresdal","Gan","Gardermoen","Gargan","Garnes","Gasbakken","Gaupen","Geilo","Geithus","Gjerdrum","Gjerstad","Gjolme","Glesvaer","Glomfjord","Godoy","Godvik","Gol","Gran","Gransherad","Granvin","Gratangen","Gravdal","Greaker","Grendi","Gressvik","Grimstad","Groa","Grong","Grua","Gullaug","Gvarv","Haddal","Haegeland","Haerland","Hagan","Hagavik","Hakadal","Halden","Hallingby","Halsa","Haltdalen","Hamar","Hamarvik","Hammerfest","Hansnes","Haram","Hareid","Harstad","Haslum","Hasvik","Hatlestranda","Hauge","Haugesund","Haukeland","Havik","Hebnes","Hedal","Heggedal","Heggenes","Hegra","Heimdal","Helgeland","Helgeroa","Hell","Hellandsjoen","Helleland","Hellesylt","Hellvik","Hemnes","Hemnesberget","Hemnskjela","Hemsedal","Henningsvaer","Herand","Heroysund","Herre","Hersaeter","Hestvika","Hetlevik","Hildre","Hitra","Hjellestad","Hjelmas","Hjelset","Hjorungavag","Hof","Hokkasen","Hokksund","Hol","Hole","Holen","Holmefjord","Holmen","Holmenkollen","Holmestrand","Holsen","Holter","Hommelvik","Hommersak","Honefoss","Hordvik","Hornnes","Horte","Horten","Hov","Hovag","Hovden","Hovet","Hovik verk","Hovin","Hoyanger","Hundven","Hunndalen","Husoy","Hustad","Hvalstad","Hvam","Hvitsten","Hvittingfoss","Hyggen","Hylkje","Hyllestad","Ikornnes","Indre Arna","Indre Billefjord","Indre Klubben","Indre Ulvsvag","Indreby","Innbygda","Inndyr","Innvik","Isdalsto","Ise","Ivgobahta","Jakobselv","Jar","Jaren","Jessheim","Jevnaker","Jomna","Jorpeland","Kabelvag","Kaldfarnes","Kalvag","Kamben","Karasjok","Karlshus","Kaupanger","Kautokeino","Kirkenaer","Kirkenes","Kjeller","Kjellmyra","Kjerstad","Kjollefjord","Kjopsvik","Kleive","Klepp","Kleppe","Kleppesto","Kleppstad","Klofta","Klokkarvik","Knapper","Knappstad","Knarrevik","Knarrlaget","Kolbjornsvik","Kolbotn","Kolbu","Kolltveit","Kolnes","Kolsas","Kolvereid","Kongsberg","Kongshamn","Kongsvika","Kongsvinger","Konsmo","Konsvikosen","Kopervik","Koppang","Korgen","Kornsjo","Korsvegen","Kragero","Krakeroy","Krakstad","Kristiansand","Kristiansund","Kroderen","Krokstadelva","Kval","Kvalsund","Kvam","Kvammen","Kvanne","Kvelde","Kvinesdal","Kvinlog","Kvisvik","Kviteseid","Kyrkjebo","Kyrksaeterora","Lakselv","Laksevag","Laksvatn","Lalm","Land","Langangen","Langesund","Langevag","Langfjordbotn","Langhus","Larkollen","Larvik","Laukvik","Lauvsnes","Lauvstad","Leikang","Leines","Leira","Leirfjord","Leirsund","Leirvik","Leknes","Lena","Lensvik","Lenvik","Lepsoy","Levanger","Lidaladdi","Lier","Lillehammer","Lillesand","Lindas","Loddefjord","Lodingen","Loen","Lofthus","Loken","Lokken Verk","Lom","Lonevag","Longva","Lorenfallet","Loten","Lovund","Lundamo","Lunde","Lunner","Lyngdal","Lyngseidet","Lyngstad","Lysaker","Lysoysundet","Magnor","Malm","Maloy","Malvik","Mandal","Manger","Manndalen","Marheim","Masfjorden","Mathopen","Maura","Mehamn","Meisingset","Melbu","Meldal","Melhus","Melsomvik","Meraker","Mestervik","Midsund","Miland","Minnesund","Mirza Rafi Sauda","Misje","Misvaer","Mjolkeraen","Mjondalen","Mo","Mo i Rana","Modalen","Moelv","Moen","Moi","Molde","Moldjord","Morgedal","Mosby","Mosjoen","Moss","Movik","Myking","Myre","Mysen","Na","Naerbo","Naersnes","Namsos","Namsskogan","Narvik","Naustdal","Nedenes","Nedre Frei","Nesbru","Nesbyen","Nesgrenda","Nesna","Nesoddtangen","Nesttun","Neverdal","Nevlunghamn","Nodeland","Nordby Bruk","Nordfjordeid","Nordkisa","Nordland","Nordstrono","Noresund","Norheimsund","Notodden","Nybergsund","Nyborg","Nydalen","Nygardsjoen","Nyhus","Nykirke","Odda","Odnes","Oksfjord","Oksvoll","Olden","Olderdalen","Olen","Oltedal","Oma","Onarheim","Oppdal","Oppegard","Opphaug","Oresvika","Orje","Orkanger","Ornes","Orre","Os","Oslo","Otta","Otteroy","Ottestad","Oveland","Ovre Ardal","Ovrebo","Oyeren","Oystese","Porsgrunn","Prestfoss","Raholt","Rakkestad","Ramberg","Ramfjordbotn","Ramnes","Rana","Ranasfoss","Randaberg","Ranheim","Raudeberg","Raudsand","Raufoss","Rauland","Re","Reine","Reinsvoll","Reipa","Reistad","Reitan","Rena","Rennebu","Rindal","Ringebu","Ringsaker","Ringstad","Risoyhamn","Rjukan","Roa","Rodberg","Rodoy","Rognan","Rogne","Rokland","Roldal","Rollag","Rolvsoy","Romedal","Rong","Roros","Rorvik","Rosendal","Rossland","Rost","Rovde","Roverud","Royken","Royneberg","Rubbestadneset","Rud","Rygge","Rykene","Rypefjord","Saebo","Saebovik","Saetre","Saevareid","Saeveland","Sagvag","Salhus","Salsbruket","Salsnes","Saltnes","Samuelsberg","Sand","Sandane","Sande","Sandefjord","Sandeid","Sander","Sandnes","Sandnessjoen","Sandshamn","Sandstad","Sandtorg","Sandvika","Sandvoll","Sannidal","Sarpsborg","Saupstad","Selasvatn","Selje","Seljord","Sellebakk","Selva","Selvaer","Sem","Setermoen","Siggerud","Siljan","Silsand","Singsas","Sira","Sirevag","Sistranda","Sjovegan","Skabu","Skage","Skanevik","Skarer","Skarnes","Skatoy","Skaun","Skedsmokorset","Skeie","Ski","Skien","Skjeberg","Skjerstad","Skjervoy","Skjold","Skjoldastraumen","Skjolden","Skodje","Skogn","Skoppum","Skotbu","Skotterud","Skreia","Skudeneshavn","Skulsfjord","Skutvika","Slastad","Slattum","Slemdal","Slemmestad","Sletta","Snaase","Snillfjord","Sogn","Sokna","Sokndal","Soknedal","Sola","Solbergelva","Solvorn","Sommaroy","Somna","Son","Sondeled","Sor-Fron","Sorbo","Soreidgrenda","Sorli","Sortland","Sorum","Sorumsand","Sorvaer","Sorvagen","Sorvik","Spangereid","Sparbu","Sperrebotn","Spillum","Spydeberg","Stabbestad","Stabekk","Stamnes","Stamsund","Stange","Stathelle","Staubo","Stavanger","Stavern","Steigen","Steinberg","Steinkjer","Steinsdalen","Sto","Stokke","Stokmarknes","Stol","Storas","Stordal","Storebo","Storforshei","Storslett","Storsteinnes","Stranda","Straume","Straumen","Strommen","Stronstad","Strusshamn","Stryn","Suldalsosen","Sulisjielmma","Sund","Sundal","Sunde","Sunndalsora","Surnadalsora","Svarstad","Svartskog","Sveio","Svelgen","Svelvik","Svene","Svortland","Sylling","Syvik","Tafjord","Talvik","Tananger","Tanem","Tangen","Tau","Tennevoll","Tennfjord","Tertnes","Tiller","Tingvoll","Tistedal","Tjeldsto","Tjelta","Tjong","Tjorvag","Tjotta","Tofte","Tolga","Tomasjorda","Tomter","Tonstad","Tornes","Torod","Torp","Torpo","Tovik","Trana","Tranby","Trengereid","Tretten","Treungen","Trofors","Trollfjorden","Tromsdalen","Trondheim","Trones","Turoy","Tvedestrand","Tveit","Tynset","Tyristrand","Tysnes","Tysse","Tyssedal","Uggdal","Ulefoss","Ulstein","Ulsteinvik","Ulvagen","Ulvik","Undeim","Uskedalen","Utsira","Utskarpen","Uvdal","Vadheim","Vage","Vagland","Vaksdal","Vale","Valen","Valer","Valestrand","Valestrandfossen","Valldal","Valle","Valsoyfjord","Vangsvika","Vannvag","Vanse","Varangerbotn","Varhaug","Vassenden","Vatne","Vedavagen","Vegarshei","Veggli","Venabygd","Vennesla","Verdal","Vestby","Vestfossen","Vestnes","Vestra Mosterhamn","Vestre Gausdal","Vevang","Vevelstad","Vigrestad","Vikebygd","Vikedal","Vikersund","Vikesa","Vikran","Vingelen","Vinje","Vinstra","Voksa","Volda","Vollen","Vormedal","Vormsund","Voss","Vossestrand","Vraliosen","Ytre Alvik","Olavtoppen","Kapp Valdivia","Kapp Circoncision","Nyrøysa","Kapp Norvegia","Larsøya","Kapp Fie","Cape Lollo","Thompson Island","Åneby","Årnes","Ås","Aurskog-Høland","Bærum","Billingstad","Bjørkelangen","Blakstad","Drøbak","Enebakk","Fet","Fjellfoten","Frogn","Hurdal","Kløfta","Lillestrøm","Lørenskog","Nannestad","Nes","Neskollen","Nesodden","Nittedal","Oppegård","Råholt","Rælingen","Rotnes","Skedsmo","Skui","Sørum","Sørumsand","Ullensaker","Ål","Åros","Flå","Hønefoss","Hurum","Krødsherad","Modum","Nedre Eiker","Nore og Uvdal","Øvre Eiker","Ringerike","Røyken","Sætre","Sigdal","Skoger","Ávanuorri","Båtsfjord","Berlevåg","Bjørnevatn","Gamvik","Honningsvåg","Kárášjohka","Kjøllefjord","Lebesby","Loppa","Måsøy","Nesseby","Nordkapp","Øksfjord","Porsanger","Sør-Varanger","Tana","Vadsø","Vardø","Åmot","Åsnes","Eidskog","Engerdal","Folldal","Grue","Kirkenær","Løten","Nord-Odal","Rendalen","Sør-Odal","Spetalen","Stor-Elvdal","Trysil","Våler","Ågotnes","Askøy","Austrheim","Bømlo","Fjell","Fusa","Jondal","Kinsarvik","Knappskog","Knarvik","Kvinnherad","Lindås","Lonevåg","Meland","Mosterhamn","Osterøy","Øygarden","Øystese","Radøy","Sagvåg","Samnanger","Sandsli","Skogsvågen","Stord","Storebø","Syfteland","Ullensvang","Ytre Arna","Ytrebygda","Ålesund","Åndalsnes","Averøy","Batnfjordsøra","Brattvåg","Eidsvåg","Elnesvågen","Fræna","Giske","Gjemnes","Herøy","Hopen","Larsnes","Nesset","Norddal","Nordstranda","Ørskog","Ørsta","Rauma","Rensvik","Sandøy","Sjøholt","Smøla","Steinshamn","Sula","Sunndal","Sunndalsøra","Surnadal","Sykkylven","Tomra","Ulsteinvik weather pws station","Vanylven","Alstahaug","Andøy","Beiarn","Bindal","Bodø","Bogen","Bø","Brønnøy","Brønnøysund","Dønna","Evenes","Evjen","Flakstad","Gildeskål","Gladstad","Grane","Hadsel","Hamarøy","Hattfjelldal","Hauknes","Kabelvåg","Kjøpsvik","Leland","Løding","Lødingen","Løpsmarka","Lurøy","Meløy","Mosjøen","Moskenes","Øksnes","Ørnes","Rødøy","Røst","Saltdal","Sandnessjøen","Sømna","Sørfold","Sørland","Svolvær","Terråk","Tjeldsund","Træna","Tysfjord","Vågan","Værøy","Vefsn","Vega","Vestvågøy","Vik","Dombås","Etnedal","Fossbergom","Gausdal","Gjøvik","Hundorp","Lesja","Nord-Aurdal","Nord-Fron","Nordre Land","Østre Toten","Øyer","Øystre Slidre","Sel","Skjåk","Søndre Land","Sør-Aurdal","Sør-Fron","Vågå","Vågåmo","Vang","Vestre Slidre","Vestre Toten","Sjølyststranda","Aremark","Fossby","Hobøl","Hvaler","Lervik","Marker","Ørje","Råde","Rømskog","Ryggebyen","Skiptvet","Skjærhalden","Trøgstad","Åkrehamn","Bjerkreim","Bokn","Eigersund","Eike","Finnøy","Forsand","Gjesdal","Hå","Hauge i Dalane","Hjelmeland","Hommersåk","Jørpeland","Judaberg","Karmøy","Kvitsøy","Lund","Lyefjell","Nærbø","Ølen","Rennesøy","Sauda","Sæveland","Strand","Suldal","Time","Tysvær","Vedavågen","Vikeså","Vikevåg","Vindafjord","Årdal","Årdalstangen","Askvoll","Aurland","Bremanger","Eid","Farnes","Fjaler","Flora","Florø","Førde","Gaular","Gaupne","Gloppen","Gulen","Hardbakke","Hermansverk","Hornindal","Høyanger","Jølster","Lærdal","Lærdalsøyri","Leikanger","Luster","Måløy","Sogndal","Solund","Vågsøy","Bamble","Hjartdal","Kragerø","Nissedal","Nome","Prestestranda","Sauherad","Tinn","Tokke","Balsfjord","Bardu","Berg","Dyrøy","Gryllefjord","Ibestad","Kåfjord","Karlsøy","Kvæfjord","Kvænangen","Lavangen","Lyngen","Målselv","Nordreisa","Salangen","Sjøvegan","Skånland","Skjervøy","Sørreisa","Storfjord","Torsken","Tranøy","Tromsø","Å i Åfjord","Åfjord","Agdenes","Berkåk","Bjugn","Botngård","Fillan","Flatanger","Fosnes","Frosta","Frøya","Hemne","Holtålen","Høylandet","Inderøy","Indre Fosen","Klæbu","Kyrksæterøra","Leka","Lierne","Meråker","Midtre Gauldal","Namdalseid","Nærøy","Orkdal","Osen","Overhalla","Ørland","Raarvihke - Røyrvik","Ranemsletta","Roan","Røros","Rørvik","Røyrvik","Selbu","Snåase","Snåase - Snåsa","Stjørdal","Stjørdalshalsen","Tydal","Verran","Vikna","Åseral","Audnedal","Hægebostad","Justvik","Liknes","Lindesnes","Marnardal","Sirdal","Skålevik","Songdalen","Søgne","Strai","Vestbygd","Vigeland","Årøysund","Åsgårdstrand","Barkåker","Færder","Gullhaug","Selvik","Tjøme","Tønsberg"]
		}
	*/
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

// RequestCountries requests info about all countries from CountriesNowAPI, returns a struct of it.
func RequestCountries() CountryResponse {
	resp, err := http.Get(config.COUNTRIES_NOW_API_ENDPOINT + "countries")
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var countries CountryResponse

	err = json.Unmarshal(body, &countries)
	if err != nil {
		panic(err)
	}

	return countries
}

/*
type InfoResponse interface {
	FlagInfoResponse | CapitalInfoResponse | CitiesInfoResponse
}
*/

func RequestCountryInfo[T any](endpoint, param string, paramValue string) T {
	jsonData := map[string]interface{}{
		param: paramValue,
	}

	jsonPayload, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		config.COUNTRIES_NOW_API_ENDPOINT+endpoint,
		"application/json",
		bytes.NewBuffer(jsonPayload),
	)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var info T
	err = json.Unmarshal(body, &info)
	if err != nil {
		panic(err)
	}

	return info
}

var allCountries = RequestCountries().Data

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	countryCode_as_iso2 := getCountryCodeFromPath(r.URL.Path, config.INFO_ENDPOINT)
	limit := getQueryInt(r, "limit", 10)

	var response InfoResponse

	//response := requestCountries()
	country := ""

	for _, _country := range allCountries {
		if strings.ToUpper(_country.Iso2) == strings.ToUpper(countryCode_as_iso2) {
			country = _country.Country

			// Get cities up to limit, check for index overflow

		}
	}

	// Found no country with Iso2 code.
	if country == "" {
		http.Error(w, "Invalid countrycode", http.StatusBadRequest)
		return
	}

	response.Name = country

	flagInfo := RequestCountryInfo[FlagInfoResponse]("countries/flag/images", "iso2", countryCode_as_iso2)
	response.Flag = flagInfo.Data.Flag

	capitalInfo := RequestCountryInfo[CapitalInfoResponse]("countries/capital", "country", country)
	response.Capital = capitalInfo.Data.Capital

	citiesInfo := RequestCountryInfo[CitiesInfoResponse]("countries/cities", "country", country)

	if limit < len(citiesInfo.Data) {
		response.Cities = citiesInfo.Data[:limit]
	} else {
		response.Cities = citiesInfo.Data
	}

	populationInfo := RequestCountryInfo[PopulationInfoResponse]("countries/population", "country", country)
	// Get last value
	n := len(populationInfo.Data.PopulationCounts)

	if n > 0 {
		lastPopulationRecord := populationInfo.Data.PopulationCounts[n-1]
		response.Population = lastPopulationRecord.Value
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	_, err = fmt.Fprintf(w, "%s\n", string(jsonData))
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
