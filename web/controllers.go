package web

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/gold"
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
func loadTemplate(w http.ResponseWriter, data map[string]interface{}, template string) {
	tpl, err := g.ParseFile(template)

	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, data)

	if err != nil {
		panic(err)
	}
}

func GetRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	template_name := "./templates/register.gold"
	data := map[string]interface{}{"Title": "Register"}
	loadTemplate(w, data, template_name)
}

func PostRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	var errors = make(map[string]string)
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	confirmPassword := r.Form.Get("confirmPassword")
	if len(username) <= 0 {
		errors["username"] = "Username is required"
	}
	if len(password) <= 0 {
		errors["password"] = "Password is required"
	}else {
		if len(confirmPassword) <= 0 {
			errors["confirmPassword"] = "Password Confirmation is required"
		}else {
			if password != confirmPassword {
				errors["password"] = "Passwords should match"
			}
		}
	}

	if len(errors) > 0 {
		template_name := "./templates/register.gold"
		data := map[string]interface{}{"Title": "Register", "errors": errors, "username":username}
		loadTemplate(w, data, template_name)
	}else {
		http.Redirect(w, r, "/", 301)
	}
}
