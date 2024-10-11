package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Page struct {
	Products   []Product
	Title      string
	Rows       int
	Categories []string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("static/index.html", "static/products.gohtml", "static/head.gohtml", "static/footer.gohtml", "static/filter.gohtml")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}

	// Данные для передачи в шаблон (при необходимости)
	products := ReadExcel("export.xlsx")
	page := Page{
		Products:   products,
		Title:      "",
		Categories: getCategories(products),
	}
	// Рендерим шаблон с данными
	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
		return
	}
}
func productsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/products.html", "static/products.gohtml", "static/head.gohtml", "static/footer.gohtml", "static/filter.gohtml")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}

	// Данные для передачи в шаблон (при необходимости)

	page := Page{
		Products: ReadExcel("export.xlsx"),
		Title:    "",
	}
	// Рендерим шаблон с данными
	err = tmpl.Execute(w, page)
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

	// Формируем сообщение с названием товара
	productInfo := fmt.Sprintf("Product ID: %s", id)

	// Выводим информацию о товаре
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(productInfo))
}
func filterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/filter.html", "static/head.gohtml", "static/footer.gohtml", "static/filter.gohtml")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}

	page := Page{
		Title: "Фильтр",
	}
	// Рендерим шаблон с данными
	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		log.Println("Error rendering template:", err)
		return
	}
}
func startHttpServer() {
	// Определяем обработчик для всех маршрутов
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/product/", productHandler)
	http.HandleFunc("/filter/", filterHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// Запускаем сервер на порту 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
