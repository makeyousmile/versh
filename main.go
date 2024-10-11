package main

func main() {
	startHttpServer()
	p := ReadExcel("export.xlsx")
	getCategories(p)
	//log.Print(p[0].ImageURL)
}
