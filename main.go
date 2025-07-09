package main

var page = Page{}

func init() {
	products := ReadExcel("export.xlsx")
	page.Products = products
	page.Title = "ОАО ВЕРШ"
	page.Categories = getCats(products)
}
func main() {
	startHttpServer()
}
