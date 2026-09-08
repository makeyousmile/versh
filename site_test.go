package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExportV2Load(t *testing.T) {
	s := getSiteFromExel("export_v2.xlsx")
	if len(s.Products) == 0 {
		t.Fatalf("Expected products in export_v2.xlsx, got 0")
	}
	t.Logf("Loaded %d products from export_v2.xlsx", len(s.Products))

	if len(s.Categories) == 0 {
		t.Fatalf("Expected categories in export_v2.xlsx, got 0")
	}
	t.Logf("Loaded %d categories from export_v2.xlsx", len(s.Categories))

	if len(s.Articles) < 2 {
		t.Fatalf("Expected at least 2 articles, got %d", len(s.Articles))
	}
	t.Logf("Loaded %d articles", len(s.Articles))

	// Verify required fields on first 100 products
	for i := 0; i < 100 && i < len(s.Products); i++ {
		p := s.Products[i]
		if p.Code == "" {
			t.Errorf("Product %d has empty Code", i)
		}
		if p.Name == "" {
			t.Errorf("Product %d has empty Name", i)
		}
		if p.GroupName == "" {
			t.Errorf("Product %d has empty GroupName", i)
		}
		if p.Category == "" {
			t.Errorf("Product %d has empty Category", i)
		}
	}
}

func TestBackwardCompatibilityExportV1(t *testing.T) {
	s := getSiteFromExel("export.xlsx")
	if len(s.Products) != 68 {
		t.Errorf("Expected 68 products from export.xlsx, got %d", len(s.Products))
	}
	if len(s.Articles) < 2 {
		t.Errorf("Expected at least 2 articles from export.xlsx, got %d", len(s.Articles))
	}
}

func TestHTTPHandlersWithV2(t *testing.T) {
	// Initialize site with export_v2.xlsx
	site = getSiteFromExel("export_v2.xlsx")

	// 1. Test Index handler "/"
	t.Run("IndexHandler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		indexHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("Expected 200, got %d", rec.Code)
		}
		body := rec.Body.String()
		if !strings.Contains(body, "Популярные товары") {
			t.Errorf("Index page missing 'Популярные товары'")
		}
		popular := getPopularProducts(site.Products)
		if len(popular) != 8 {
			t.Errorf("Expected 8 popular products, got %d", len(popular))
		}
		// Also count product-card in body
		cardCount := strings.Count(body, "product-card")
		if cardCount != 8 {
			t.Errorf("Expected 8 product-card on index page, got %d", cardCount)
		}
	})

	// 2. Test Products handler "/products"
	t.Run("ProductsHandler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		rec := httptest.NewRecorder()
		productsHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("Expected 200, got %d", rec.Code)
		}
		body := rec.Body.String()
		if !strings.Contains(body, "Каталог") {
			t.Errorf("Products page missing 'Каталог'")
		}
	})

	// 3. Test Product handler with valid ID "/product/17139"
	t.Run("ProductHandlerValid", func(t *testing.T) {
		firstCode := site.Products[0].Code
		req := httptest.NewRequest(http.MethodGet, "/product/"+firstCode, nil)
		rec := httptest.NewRecorder()
		productHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("Expected 200, got %d", rec.Code)
		}
		body := rec.Body.String()
		if !strings.Contains(body, site.Products[0].Name) {
			t.Errorf("Product page missing product name '%s'", site.Products[0].Name)
		}
	})

	// 4. Test Product handler with invalid ID "/product/999999999"
	t.Run("ProductHandlerNotFound", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/product/999999999", nil)
		rec := httptest.NewRecorder()
		productHandler(rec, req)
		if rec.Code != http.StatusNotFound {
			t.Fatalf("Expected 404, got %d", rec.Code)
		}
	})

	// 5. Test Categories handler
	t.Run("CategoriesHandler", func(t *testing.T) {
		if len(site.Categories) == 0 {
			t.Fatal("No categories available to test")
		}
		catLatin := site.Categories[0].CatLatin
		req := httptest.NewRequest(http.MethodGet, "/categories/"+catLatin, nil)
		rec := httptest.NewRecorder()
		categoriesHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("Expected 200, got %d", rec.Code)
		}
		body := rec.Body.String()
		if !strings.Contains(body, site.Categories[0].CatCyrillic) {
			t.Errorf("Categories page missing category title '%s'", site.Categories[0].CatCyrillic)
		}
	})

	// 6. Test Article handlers
	t.Run("ArticleAboutHandler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article/about", nil)
		rec := httptest.NewRecorder()
		articleHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("Expected 200, got %d", rec.Code)
		}
	})

	t.Run("ArticleShipmentHandler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/article/shipment", nil)
		rec := httptest.NewRecorder()
		articleHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("Expected 200, got %d", rec.Code)
		}
	})

	// 7. Test Filter handler
	t.Run("FilterHandler", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/filter", nil)
		rec := httptest.NewRecorder()
		filterHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("Expected 200, got %d", rec.Code)
		}
	})
}
