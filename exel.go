package main

import (
	"fmt"
	"html/template"
	"log"
	"sort"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Product представляет структуру товара с указанными полями
type Product struct {
	Code                string `json:"код_товара"`
	Name                string `json:"название_позиции"`
	SearchQueries       string `json:"поисковые_запросы"`
	Description         string `json:"описание"`
	ProductType         string `json:"тип_товара"`
	Price               string `json:"цена"`
	Currency            string `json:"валюта"`
	UnitOfMeasurement   string `json:"единица_измерения"`
	MinOrderVolume      string `json:"минимальный_объем_заказа"`
	WholesalePrice      string `json:"оптовая_цена"`
	MinWholesaleOrder   string `json:"минимальный_заказ_опт"`
	ImageURL            string `json:"ссылка_изображения"`
	Availability        string `json:"наличие"`
	Quantity            string `json:"количество"`
	GroupID             string `json:"номер_группы"`
	GroupName           string `json:"название_группы"`
	SubsectionURL       string `json:"адрес_подраздела"`
	SupplyCapability    string `json:"возможность_поставки"`
	DeliveryTime        string `json:"срок_поставки"`
	PackagingMethod     string `json:"способ_упаковки"`
	UniqueIdentifier    string `json:"уникальный_идентификатор"`
	ItemID              string `json:"идентификатор_товара"`
	SubsectionID        string `json:"идентификатор_подраздела"`
	GroupIdentifier     string `json:"идентификатор_группы"`
	Manufacturer        string `json:"производитель"`
	WarrantyPeriod      string `json:"гарантийный_срок"`
	CountryOfOrigin     string `json:"страна_производитель"`
	Discount            string `json:"скидка"`
	VariantGroupID      string `json:"id_группы_разновидностей"`
	ManufacturerName    string `json:"название_производителя"`
	ManufacturerAddress string `json:"адрес_производителя"`
	PersonalNotes       string `json:"личные_заметки"`
	ProductOnSite       string `json:"продукт_на_сайте"`
	DiscountStartDate   string `json:"срок_действия_скидки_от"`
	DiscountEndDate     string `json:"срок_действия_скидки_до"`
	PriceFrom           string `json:"цена_от"`
	Label               string `json:"ярлык"`
	HTMLTitle           string `json:"html_заголовок"`
	HTMLDescription     string `json:"html_описание"`
	GTINCode            string `json:"код_маркировки_(gtin)"`
	MPNNumber           string `json:"номер_устройства_(mpn)"`
	SupplierName        string `json:"название_поставщика"`
	SupplierAddress     string `json:"адрес_поставщика"`
	Category            string `json:"категория"`
	CurCategory         string `json:"текущая_категория"`
}

type Site struct {
	Products   []Product
	Product    Product
	Title      string
	Rows       int
	Text       template.HTML
	Categories []Categories
	Articles   []Article
}

type Article struct {
	Title string
	Text  template.HTML
}

// Helper to safely get column by index from a row
func getCell(row []string, idx int) string {
	if idx >= 0 && idx < len(row) {
		return strings.TrimSpace(row[idx])
	}
	return ""
}

// parseV2Products parses sheet formatted like export_v2.xlsx ("Товары")
func parseV2Products(rows [][]string) []Product {
	products := make([]Product, 0, len(rows))
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		id := getCell(row, 0)
		name := getCell(row, 3)
		if id == "" && name == "" {
			continue
		}

		brand := getCell(row, 1)
		cat := getCell(row, 2)
		artikul := getCell(row, 4)
		price := getCell(row, 5)
		avail := getCell(row, 6)
		desc := getCell(row, 9)
		chars := getCell(row, 10)
		brand2 := getCell(row, 11)
		prodType := getCell(row, 12)
		cat2 := getCell(row, 15)
		promo := getCell(row, 20)
		photo := getCell(row, 21)

		groupName := cat
		if groupName == "" {
			groupName = cat2
		}
		if groupName == "" {
			groupName = prodType
		}
		if groupName == "" {
			groupName = "Прочее"
		}

		if brand == "" && brand2 != "" {
			brand = brand2
		}

		product := Product{
			Code:                id,
			Name:                name,
			SearchQueries:       artikul,
			Description:         desc,
			ProductType:         prodType,
			Price:               price,
			Currency:            "BYN",
			UnitOfMeasurement:   "шт.",
			ImageURL:            photo,
			Availability:        avail,
			GroupID:             "",
			GroupName:           groupName,
			SubsectionURL:       "",
			SupplyCapability:    "",
			DeliveryTime:        "",
			PackagingMethod:     "",
			UniqueIdentifier:    id,
			ItemID:              artikul,
			SubsectionID:        "",
			GroupIdentifier:     promo,
			Manufacturer:        brand,
			ManufacturerName:    brand,
			PersonalNotes:       chars,
			Label:               promo,
			Category:            Transliterate(groupName),
			CurCategory:         groupName,
		}
		products = append(products, product)
	}
	return products
}

// parseV1Products parses sheet formatted like export.xlsx ("Export Products Sheet")
func parseV1Products(rows [][]string) []Product {
	products := make([]Product, 0, len(rows))
	for x, row := range rows {
		if x == 0 {
			continue
		}
		product := Product{}
		for y, cell := range row {
			switch y {
			case 0:
				product.Code = cell
			case 1:
				product.Name = cell
			case 2:
				product.SearchQueries = cell
			case 3:
				product.Description = cell
			case 4:
				product.ProductType = cell
			case 5:
				product.Price = cell
			case 6:
				product.Currency = cell
			case 7:
				product.UnitOfMeasurement = cell
			case 8:
				product.MinOrderVolume = cell
			case 9:
				product.WholesalePrice = cell
			case 10:
				product.MinWholesaleOrder = cell
			case 11:
				product.ImageURL = cell
			case 12:
				product.Availability = cell
			case 13:
				product.Quantity = cell
			case 14:
				product.GroupID = cell
			case 15:
				product.GroupName = cell
			case 16:
				product.SubsectionURL = cell
			case 17:
				product.SupplyCapability = cell
			case 18:
				product.DeliveryTime = cell
			case 19:
				product.PackagingMethod = cell
			case 20:
				product.UniqueIdentifier = cell
			case 21:
				product.ItemID = cell
			case 22:
				product.SubsectionID = cell
			case 23:
				product.GroupIdentifier = cell
			case 24:
				product.Manufacturer = cell
			case 25:
				product.WarrantyPeriod = cell
			case 26:
				product.CountryOfOrigin = cell
			case 27:
				product.Discount = cell
			case 28:
				product.VariantGroupID = cell
			case 29:
				product.ManufacturerName = cell
			case 30:
				product.ManufacturerAddress = cell
			case 31:
				product.PersonalNotes = cell
			case 32:
				product.ProductOnSite = cell
			case 33:
				product.DiscountStartDate = cell
			case 34:
				product.DiscountEndDate = cell
			case 35:
				product.PriceFrom = cell
			case 36:
				product.Label = cell
			case 37:
				product.HTMLTitle = cell
			case 38:
				product.HTMLDescription = cell
			case 39:
				product.GTINCode = cell
			case 40:
				product.MPNNumber = cell
			case 41:
				product.SupplierName = cell
			case 42:
				product.SupplierAddress = cell
			}
		}
		if product.GroupName == "" {
			product.GroupName = "Прочее"
		}
		product.Category = Transliterate(product.GroupName)
		product.CurCategory = product.GroupName
		products = append(products, product)
	}
	return products
}

// parseProductsFile inspects sheet names and headers to parse products
func parseProductsFile(f *excelize.File) []Product {
	sheetList := f.GetSheetList()
	// 1. Check for "Товары" (v2 format)
	for _, sheet := range sheetList {
		if sheet == "Товары" {
			rows, err := f.GetRows(sheet)
			if err != nil {
				log.Printf("Error reading sheet %s: %v\n", sheet, err)
				return nil
			}
			return parseV2Products(rows)
		}
	}
	// 2. Check for "Export Products Sheet" (v1 format)
	for _, sheet := range sheetList {
		if sheet == "Export Products Sheet" {
			rows, err := f.GetRows(sheet)
			if err != nil {
				log.Printf("Error reading sheet %s: %v\n", sheet, err)
				return nil
			}
			return parseV1Products(rows)
		}
	}
	// 3. Fallback: inspect first sheet with rows
	for _, sheet := range sheetList {
		rows, err := f.GetRows(sheet)
		if err == nil && len(rows) > 1 {
			if len(rows[0]) > 4 && strings.EqualFold(rows[0][0], "ID") {
				return parseV2Products(rows)
			}
			return parseV1Products(rows)
		}
	}
	return nil
}

// getArticles loads articles from file, falls back to export.xlsx, or uses defaults
func getArticles(f *excelize.File) []Article {
	articles := make([]Article, 0)
	if rows, err := f.GetRows("Articles"); err == nil && len(rows) >= 3 {
		if len(rows[1]) >= 1 && len(rows[2]) >= 1 {
			articles = append(articles, Article{
				Title: rows[1][0],
				Text:  template.HTML(rows[2][0]),
			})
		}
		if len(rows[1]) >= 2 && len(rows[2]) >= 2 {
			articles = append(articles, Article{
				Title: rows[1][1],
				Text:  template.HTML(rows[2][1]),
			})
		}
	}
	if len(articles) < 2 {
		// Fallback to export.xlsx if available
		if fOld, err := excelize.OpenFile("export.xlsx"); err == nil {
			defer fOld.Close()
			if oldRows, err := fOld.GetRows("Articles"); err == nil && len(oldRows) >= 3 {
				if len(articles) == 0 && len(oldRows[1]) >= 1 && len(oldRows[2]) >= 1 {
					articles = append(articles, Article{
						Title: oldRows[1][0],
						Text:  template.HTML(oldRows[2][0]),
					})
				}
				if len(articles) == 1 && len(oldRows[1]) >= 2 && len(oldRows[2]) >= 2 {
					articles = append(articles, Article{
						Title: oldRows[1][1],
						Text:  template.HTML(oldRows[2][1]),
					})
				}
			}
		}
	}
	// Default fallbacks if still missing
	if len(articles) == 0 {
		articles = append(articles, Article{
			Title: "О нас",
			Text:  template.HTML("<p>ООО «Верш» — надежный поставщик оборудования и систем очистки воды.</p>"),
		})
	}
	if len(articles) == 1 {
		articles = append(articles, Article{
			Title: "Оплата и доставка",
			Text:  template.HTML("<p>Доставка осуществляется по всей Беларуси. Оплата безналичным и наличным расчетом.</p>"),
		})
	}
	return articles
}

func getSiteFromExel(path string) Site {
	site := Site{}
	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Printf("getSiteFromExel: failed to open %s: %v\n", path, err)
		return site
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	site.Articles = getArticles(f)
	products := parseProductsFile(f)
	site.Products = products
	site.Categories = getCats(products)
	site.Rows = len(products)
	return site
}

func GetProductsFromExcel(path string) []Product {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return parseProductsFile(f)
}

func getCats(products []Product) []Categories {
	catMap := make(map[string]bool)
	for _, product := range products {
		if product.GroupName != "" {
			catMap[product.GroupName] = true
		}
	}

	catNames := make([]string, 0, len(catMap))
	for c := range catMap {
		catNames = append(catNames, c)
	}
	sort.Strings(catNames)

	return getCatFromNames(catNames)
}

func getCatFromNames(catNames []string) []Categories {
	cats := make([]Categories, 0, len(catNames))
	for _, catName := range catNames {
		cat := Categories{}
		cat.CatCyrillic = catName
		cat.CatLatin = Transliterate(catName)
		cats = append(cats, cat)
	}
	return cats
}

func getPopularProducts(products []Product) []Product {
	tmpProducts := make([]Product, 0)
	for _, product := range products {
		if product.GroupIdentifier != "" {
			tmpProducts = append(tmpProducts, product)
		}
	}
	if len(tmpProducts) == 0 {
		for _, product := range products {
			if product.ImageURL != "" {
				tmpProducts = append(tmpProducts, product)
				if len(tmpProducts) >= 8 {
					break
				}
			}
		}
	}
	if len(tmpProducts) > 8 {
		tmpProducts = tmpProducts[:8]
	}

	return tmpProducts
}

func FindProductByCode(products []Product, code string) Product {
	for _, product := range products {
		if product.Code == code || (product.ItemID != "" && product.ItemID == code) {
			return product
		}
	}
	return Product{}
}

func FindProductByCat(products []Product, cat string) []Product {
	var catProducts []Product
	for _, product := range products {
		if product.Category == cat || Transliterate(product.GroupName) == cat {
			catProducts = append(catProducts, product)
		}
	}
	return catProducts
}
