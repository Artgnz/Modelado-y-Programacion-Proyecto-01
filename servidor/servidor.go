package servidor

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/Artgnz/Modelado-y-Programacion-Proyecto-01/servidor/clima"
)

// EjecutarServidor ejecuta el servidor en el puerto 8080.
func EjecutarServidor() {
	puerto := "8080"
	mux := http.NewServeMux()
	var cliente *clima.ClienteClima
	mux.HandleFunc("/", funcionHandle(cliente))
	http.ListenAndServe("localhost:"+puerto, mux)

}

// funcionHandle recibe un cliente de openWeather y devuelve una función
// handler para procesar una petición r y construir una respuesta HTTP con w.
func funcionHandle(cliente *clima.ClienteClima) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/acceso/":
			cliente = accesoHandler(w, r)
		case "/tickets/":
			ticketsHandler(cliente, w, r)
		}
	}
}

// ticketsHandler procesa peticiones relacionadas con mostrar la tabla de tickets.
// Recibe cliente para realizar peticións a openweather.
func ticketsHandler(cliente *clima.ClienteClima, w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/templates/tickets.html"))
	tpl.Execute(w, nil)
}

// accesoHandler procesa peticiones relacionadas con el ingreso de sesión.
func accesoHandler(w http.ResponseWriter, r *http.Request) *clima.ClienteClima {
	switch r.Method {
	// Si el método de la petición es GET
	case "GET":
		tpl := template.Must(template.ParseFiles("assets/templates/acceso.html"))
		tpl.Execute(w, nil)
		// Si el método de la petición es POST
	case "POST":
		// Se lee la llave API de la petición
		llaveApi := r.FormValue("llave-api")
		// Se crea un cliente con la llave.
		cliente, error := clima.NuevoClienteClima(&http.Client{}, llaveApi)
		// Si hubo un error.
		if error != nil {
			// Se crea plantilla para llave incorrecta.
			tpl := template.Must(template.ParseFiles("assets/templates/llave-incorrecta.html"))
			// Se ejecuta la plantilla
			tpl.Execute(w, nil)
		} else {
			// Se redirecciona a url de tickets.
			http.Redirect(w, r, "/tickets/", http.StatusFound)
			// Se regresa el cliente.
			return cliente
		}
	default:
		// La url solo soporta los métodos GET y POST
		fmt.Fprintf(w, "Soporte únicamente para métodos GET y POST.")
	}
	return nil
}
