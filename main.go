package main

import "os"

var (
	site     = Site{}
	exelFile = "export_v2.xlsx"
)

func init() {
	if envFile := os.Getenv("EXCEL_FILE"); envFile != "" {
		exelFile = envFile
	}
	site = getSiteFromExel(exelFile)
}
func main() {
	//getPagesFromExel("export.xlsx")
	startHttpServer()
}
