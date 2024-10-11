package main

import (
	"fmt"
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
}

func ReadExcel(path string) []Product {
	products := make([]Product, 0)
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return products
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Export Products Sheet")
	if err != nil {
		fmt.Println(err)
		return products
	}
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
		products = append(products, product)
	}
	return products
}
func getCategories(products []Product) []string {
	categories := make([]string, 0)
	cat := make(map[string]bool)

	for _, product := range products {
		cat[product.GroupName] = true
	}

	for c, _ := range cat {
		categories = append(categories, c)
	}

	return categories
}
