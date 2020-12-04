package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewGrado struct {
	Name    string
	IsEdit  bool
	Data    models.Grado
	Widgets []models.Grado
	Niveles []models.Nivel
}

var tmplg = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/grado/index.html", "web/grado/form.html"))

func GradoList(w http.ResponseWriter, req *http.Request) {

	lis := []models.Grado{}
	if err := cfig.DB.Preload("Nivel").Find(&lis).Error; err != nil { // Preload("Nivel") carga los objetos Nivel relacionado
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := ViewGrado{
		Name:    "Grado",
		Widgets: lis,
	}
	err := tmplg.ExecuteTemplate(w, "grado/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GradoForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Grado
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	nivel := models.Nivel{}
	niveles, _ := nivel.GetAll(cfig.DB) // para mostrar los niveles en un combobox

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Nomg = r.FormValue("nomg")
		//n, err := strconv.Atoi(r.FormValue("id"))
		//if err != nil {
		//	log.Printf("Invalid ID: %v - %v\n", n, err)
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		d.NivelId = r.FormValue("nivel_id") //n
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
		http.Redirect(w, r, "/grado/index", 301)
	}

	data := ViewGrado{
		Name:    "Grado",
		Data:    d,
		IsEdit:  IsEdit,
		Niveles: niveles,
	}

	err := tmplg.ExecuteTemplate(w, "grado/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GradoDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var d models.Grado
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		//log.Printf("No save  %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}

	http.Redirect(w, r, "/grado/index", 301)
}
