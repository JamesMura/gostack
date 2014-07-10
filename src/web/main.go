package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/gold"
	"log"
	"net/http"
)

var g = gold.NewGenerator(false)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tpl, err := g.ParseFile("./templates/home.gold")

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{"Title": "Gold"}

	err = tpl.Execute(w, data)

	if err != nil {
		panic(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.ServeFiles("/static/*filepath", http.Dir("public/"))
	log.Fatal(http.ListenAndServe(":8080", router))
}
