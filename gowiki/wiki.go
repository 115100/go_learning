package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl)
	t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit.html", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view.html", p)
}

func main() {
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/view/", viewHandler)

	http.ListenAndServe(":8080", nil)
}
