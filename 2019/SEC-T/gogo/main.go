package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	funcMap := template.FuncMap{
		"Dir":        ioutil.ReadDir,
		"ReadFile":   ioutil.ReadFile,
		"ReplaceAll": strings.ReplaceAll,
		"StrIndex":   strings.Index,
	}
	page := r.FormValue("page")
	if page != "" {
		page = "www" + r.URL.Path + page
	} else {
		page = "www" + r.URL.Path + ".default"
	}

	tpl := template.New("menu").Funcs(funcMap)
	tpl, err := tpl.ParseFiles("menu", page)
	log.Println(page)
	if err != nil {
		log.Println(err)
		tpl = template.New("menu").Funcs(funcMap)
		tpl, err = tpl.ParseFiles("menu", "missing")
		if err != nil {
			log.Println(err)
		}
	}

	tpl.Execute(w, map[string]interface{}{"r": r})

}

func main() {

	log.Fatal(http.ListenAndServe(":8123", &server{}))
}
