package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Categories struct {
	CatCyrillic string
	CatLatin    string
}

type FormData struct {
	Name        string
	Phone       string
	Email       string
	City        string
	HousingType string
	WaterSource string
	WaterIssues []string
	TapPoints   string
	Residents   string
	SewerType   string
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {

	tmpl, err := template.ParseFiles("static/index.html", "static/products.gohtml", "static/head.gohtml", "static/footer.gohtml", "static/filter.gohtml")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}

	// Данные для передачи в шаблон (при необходимости)
	tmpSite := site
	tmpSite.Products = getPopularProducts(tmpSite.Products)
	tmpSite.Title = "Популярные товары"

	// Рендерим шаблон с данными
	err = tmpl.Execute(w, tmpSite)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
		return
	}
}
func productsHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("static/products.html", "static/products.gohtml", "static/head.gohtml", "static/footer.gohtml", "static/filter.gohtml")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}

	// Данные для передачи в шаблон (при необходимости)

	tmpSite := site
	tmpSite.Title = "Каталог"
	// Рендерим шаблон с данными
	err = tmpl.Execute(w, tmpSite)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
		return
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем id товара из URL
	id := strings.TrimPrefix(r.URL.Path, "/product/")
	if id == "" {
		http.Error(w, "Product ID is missing", http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("static/product.html", "static/product.gohtml", "static/head.gohtml", "static/footer.gohtml")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}
	tmpSite := site
	tmpSite.Product = FindProductByCode(tmpSite.Products, id)
	if tmpSite.Product.Name == "" {
		http.Error(w, "Product not found", http.StatusNotFound)
	}

	err = tmpl.Execute(w, tmpSite)
	if err != nil {
		log.Println("Error rendering template:", err)
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}

}
func filterHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("static/filter.html", "static/head.gohtml", "static/footer.gohtml", "static/filter.gohtml")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}

	tmpSite := site
	tmpSite.Title = "Фильтр"
	// Рендерим шаблон с данными
	err = tmpl.Execute(w, tmpSite)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
		return
	}
}
func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	cat := strings.TrimPrefix(r.URL.Path, "/categories/")
	if cat == "" {
		http.Error(w, "Category is missing", http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("static/products.html", "static/products.gohtml", "static/head.gohtml", "static/footer.gohtml", "static/filter.gohtml")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}

	// Данные для передачи в шаблон (при необходимости)
	tmpSite := site

	tmpSite.Products = FindProductByCat(tmpSite.Products, cat)
	if len(tmpSite.Products) > 0 {
		tmpSite.Title = tmpSite.Products[0].GroupName
	}
	// Рендерим шаблон с данными
	err = tmpl.Execute(w, tmpSite)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
		return
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			return
		} // 10 MB max file size

		data := FormData{
			Name:        r.FormValue("name"),
			Phone:       r.FormValue("phone"),
			Email:       r.FormValue("email"),
			City:        r.FormValue("city"),
			HousingType: r.FormValue("housing_type"),
			WaterSource: r.FormValue("water_source"),
			WaterIssues: r.Form["water_issues"], // multiple checkboxes
			TapPoints:   r.FormValue("tap_points"),
			Residents:   r.FormValue("residents"),
			SewerType:   r.FormValue("sewer_type"),
		}

		// Handle file upload
		file, handler, err := r.FormFile("water_analysis")
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		f, err := os.Create("./uploads/" + handler.Filename)
		if err != nil {
			log.Println(err)
		}
		_, _ = f.ReadFrom(file)
		defer f.Close()

		fmt.Fprintf(w, "Спасибо! Данные получены:\n%+v", data)
		log.Println(data.SewerType)
		return
	}

	tmpl := template.Must(template.ParseFiles("filter 1.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	cat := strings.TrimPrefix(r.URL.Path, "/article/")
	if cat == "" {
		http.Error(w, "article is missing", http.StatusBadRequest)
		return
	}
	tmpl, err := template.ParseFiles("static/article.html", "static/head.gohtml", "static/footer.gohtml")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}

	// Данные для передачи в шаблон (при необходимости)
	tmpSite := site
	if cat == "about" {
		tmpSite.Text = template.HTML(tmpSite.Articles[0].Text)
		tmpSite.Title = tmpSite.Articles[0].Title
	}
	if cat == "shipment" {
		tmpSite.Text = tmpSite.Articles[1].Text
		tmpSite.Title = tmpSite.Articles[1].Title
	}

	// Рендерим шаблон с данными
	err = tmpl.Execute(w, tmpSite)

	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
		return
	}
}
func reloadHandler(writer http.ResponseWriter, request *http.Request) {
	site = getSiteFromExel(exelFile)
}

func startHttpServer() {
	// Определяем обработчик для всех маршрутов
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/product/", productHandler)
	http.HandleFunc("/filter/", filterHandler)
	http.HandleFunc("/categories/", categoriesHandler)
	http.HandleFunc("/submit", formHandler)
	http.HandleFunc("/article/", articleHandler)
	http.HandleFunc("/reload/", reloadHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// Запускаем сервер на порту 8080
	log.Println("Starting server on :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}

}
