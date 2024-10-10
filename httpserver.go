package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Products []Product
	Title    string
	Rows     int
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Получаем имя файла из URL

	tmpl, err := template.ParseFiles("static/index.html", "static/products.html")
	if err != nil {
		http.Error(w, "File not found or unable to load template", http.StatusNotFound)
		log.Println("Error loading template:", err)
		return
	}

	// Данные для передачи в шаблон (при необходимости)
	data := Product{
		Name:        "Test",
		Description: "Testsdvvvvv vvvvvvvv vvvvvvvvv vvvvvvvvvvvvvv vvvvvvv vvvvvvvvvvv vvvvvvvvvvvvvvv vvvvvvvvv vvvvvvvv vvvvvvvv vvvvvvvvvvvvvvv",
		Price:       99,
	}
	prod := []Product{}
	prod = append(prod, data)
	prod = append(prod, data)
	prod = append(prod, data)
	prod = append(prod, data)
	prod = append(prod, data)
	page := Page{
		Products: prod,
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

func startHttpServer() {
	// Определяем обработчик для всех маршрутов
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// Запускаем сервер на порту 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
