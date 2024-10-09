package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
)

// Product представляет структуру товара с указанными полями
type Product struct {
	Code                string    `json:"код_товара"`
	Name                string    `json:"название_позиции"`
	SearchQueries       string    `json:"поисковые_запросы"`
	Description         string    `json:"описание"`
	ProductType         string    `json:"тип_товара"`
	Price               float64   `json:"цена"`
	Currency            string    `json:"валюта"`
	UnitOfMeasurement   string    `json:"единица_измерения"`
	MinOrderVolume      int       `json:"минимальный_объем_заказа"`
	WholesalePrice      float64   `json:"оптовая_цена"`
	MinWholesaleOrder   int       `json:"минимальный_заказ_опт"`
	ImageURL            string    `json:"ссылка_изображения"`
	Availability        bool      `json:"наличие"`
	Quantity            int       `json:"количество"`
	GroupID             int       `json:"номер_группы"`
	GroupName           string    `json:"название_группы"`
	SubsectionURL       string    `json:"адрес_подраздела"`
	SupplyCapability    string    `json:"возможность_поставки"`
	DeliveryTime        string    `json:"срок_поставки"`
	PackagingMethod     string    `json:"способ_упаковки"`
	UniqueIdentifier    string    `json:"уникальный_идентификатор"`
	ItemID              string    `json:"идентификатор_товара"`
	SubsectionID        string    `json:"идентификатор_подраздела"`
	GroupIdentifier     string    `json:"идентификатор_группы"`
	Manufacturer        string    `json:"производитель"`
	WarrantyPeriod      string    `json:"гарантийный_срок"`
	CountryOfOrigin     string    `json:"страна_производитель"`
	Discount            float64   `json:"скидка"`
	VariantGroupID      int       `json:"id_группы_разновидностей"`
	ManufacturerName    string    `json:"название_производителя"`
	ManufacturerAddress string    `json:"адрес_производителя"`
	PersonalNotes       string    `json:"личные_заметки"`
	ProductOnSite       bool      `json:"продукт_на_сайте"`
	DiscountStartDate   time.Time `json:"срок_действия_скидки_от"`
	DiscountEndDate     time.Time `json:"срок_действия_скидки_до"`
	PriceFrom           float64   `json:"цена_от"`
	Label               string    `json:"ярлык"`
	HTMLTitle           string    `json:"html_заголовок"`
	HTMLDescription     string    `json:"html_описание"`
	GTINCode            string    `json:"код_маркировки_(gtin)"`
	MPNNumber           string    `json:"номер_устройства_(mpn)"`
	SupplierName        string    `json:"название_поставщика"`
	SupplierAddress     string    `json:"адрес_поставщика"`
}

func ReadExcel(path string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
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
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Println(colCell)

		}
	}
}
