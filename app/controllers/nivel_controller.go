package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewNivel struct {
	Name    string
	IsEdit  bool
	Data    models.Nivel
	Widgets []models.Nivel
}

var tmpln = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/nivel/index.html", "web/nivel/form.html"))

func NivelList(w http.ResponseWriter, req *http.Request) {

	nivel := models.Nivel{}

	niveles, _ := nivel.FindAll(cfig.DB)

	for _, lis := range niveles {
		fmt.Println(lis.ToString())
		fmt.Println("Grados: ", len(lis.Grados))
		if len(lis.Grados) > 0 {
			for _, d := range lis.Grados {
				fmt.Println(d.ToString())
				fmt.Println("=============================")
			}
		}
		fmt.Println("--------------------")
	}

	// Create
	//cfig.DB.Create(&models.Nivel{Name: "Juan", City: "Juliaca"})
	lis := []models.Nivel{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewNivel{
		Name:    "Nivel",
		Widgets: lis,
	}

	err := tmpln.ExecuteTemplate(w, "nivel/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NivelForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Nivel
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Nom = r.FormValue("nom")
		d.Desc = r.FormValue("desc")
		if id != "" {
			if err := cfig.DB.Save(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}

		} else {
			if err := cfig.DB.Create(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}
		}
		http.Redirect(w, r, "/nivel/index", 301)
	}

	data := ViewNivel{
		Name:   "Nivel",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmpln.ExecuteTemplate(w, "nivel/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NivelDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.Nivel
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}
	http.Redirect(w, r, "/nivel/index", 301)
}
