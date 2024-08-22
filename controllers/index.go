package controllers

import (
	"html/template"
	"main/models"
	"net/http"
	"sort"
)

type Category struct {
	Name   string        `json:"name"`
	Wishes []models.Wish `json:"wishes"`
}

type Wishlist struct {
	Title      string
	Categories []Category
}

var tmpl *template.Template

func init() {
	if tmpl == nil {
		if tmpl == nil {
			tmpl = template.Must(tmpl.ParseGlob("views/*.html"))
		}
	}
}

func groupByCategory(wishes []models.Wish) []Category {
	categories := make(map[string][]models.Wish)
	for _, wish := range wishes {
		categories[wish.Category] = append(categories[wish.Category], wish)
	}

	var result []Category
	for name, wishes := range categories {
		result = append(result, Category{Name: name, Wishes: wishes})
	}
	return result
}

func sortCategories(categories []Category) []Category {
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Name < categories[j].Name
	})
	return categories
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	wishes, err := models.GetWishes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categories := groupByCategory(wishes)
	categories = sortCategories(categories)

	wishlist := Wishlist{
		Title:      "Wishlist",
		Categories: categories,
	}

	tmpl.ExecuteTemplate(w, "index.html", wishlist)
}
