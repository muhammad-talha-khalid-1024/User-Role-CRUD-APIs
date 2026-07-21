package index

import (
	"fmt"
	"html/template"
	"math"
	"mustaqel/config"
	"mustaqel/models"
	"net/http"
	"strconv"
)

type PageData struct {
	Users      []models.User
	Page       int
	PrevPage   int
	NextPage   int
	HasPrev    bool
	HasNext    bool
	TotalPages int
	TotalUsers int64
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"sub": func(a, b int) int {
			return a - b
		},
	}

	tmpl := template.Must(
		template.New("index.html").
			Funcs(funcMap).
			ParseFiles("templates/index.html"),
	)

	page := 1
	limit := 1

	if p := r.URL.Query().Get("page"); p != "" {
		if n, err := strconv.Atoi(p); err == nil && n > 0 {
			page = n
		}
	}

	offset := (page - 1) * limit

	var users []models.User
	var total int64

	config.DB.Model(&models.User{}).Count(&total)

	sortBy := r.URL.Query().Get("sort")

	switch sortBy {
	case "id":
		sortBy = "id"
	case "first_name":
		sortBy = "first_name"
	case "email":
		sortBy = "email"
	default:
		sortBy = "id"
	}

	direction := r.URL.Query().Get("direction")
	if direction != "ASC" {
		direction = "DESC"
	}

	order := fmt.Sprintf("%s %s", sortBy, direction)

	config.DB.
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&users)

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	data := PageData{
		Users:      users,
		Page:       page,
		PrevPage:   page - 1,
		NextPage:   page + 1,
		HasPrev:    page > 1,
		HasNext:    page < totalPages,
		TotalUsers: total,
		TotalPages: totalPages,
	}

	tmpl.Execute(w, data)
}
