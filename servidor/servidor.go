package servidor

import (
	"log"
	"net/http"
	"text/template"
)

// EjecutarServidor ejecuta un servidor en el puerto 8080.
func EjecutarServidor() {
	puerto := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", funcionHandle)
	http.ListenAndServe("localhost:"+puerto, mux)
}

// funcionHandle decide que función usar para procesar la petición r y
// para construir una respuesta HTTP con w.
func funcionHandle(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	switch r.URL.Path {
	case "/acceso/":
		acceso(w, r)
	}
}

// acceso procesa peticiones relacionadas con el ingreso de sesión.
func acceso(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/templates/acceso.html"))
	tpl.Execute(w, nil)
}
