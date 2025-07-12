package main

var (
	site     = Site{}
	exelFile = "export.xlsx"
)

func init() {
	site = getSiteFromExel(exelFile)
}
func main() {
	//getPagesFromExel("export.xlsx")
	startHttpServer()
}
