//https://www.youtube.com/watch?v=G8Du1EOuoLY

package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":5000", nil)
}
func Index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/index.html")                      //template.ParseFiles devuelve un error que si no lo queremos capturar entonces colocamos un guión bajo _
	user := Usuario{"Jorge Loaiza", "Desarrollador Web e Ingeniero de sistemas", 43} //la variable "user" crea la estrucutra de datos con los valores respectivos
	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user) //Ejecuta template
	}

	// fmt.Fprintln(rw,"Hola cliente")
}

// Estructura de datos
type Usuario struct {
	Name   string
	Skills string
	Age    int
}
