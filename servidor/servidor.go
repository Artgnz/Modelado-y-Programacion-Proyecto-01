package servidor

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/Artgnz/Modelado-y-Programacion-Proyecto-01/servidor/clima"
)

// climaHandler procesa peticiones sobre datos del clima en
// una latitud y longitud dados.
func climaHandler(cliente *clima.ClienteClima, w http.ResponseWriter, r *http.Request) {
	// Si el cliente es nil, significa que no ha iniciado sesión.
	if cliente == nil {
		http.Redirect(w, r, "/acceso/", http.StatusUnauthorized)
		return
	}
	valoresUrl := r.URL.Query()
	/// Conesguimos latitud de la url de la petición
	lat, error := strconv.ParseFloat(valoresUrl.Get("lat"), 64)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Latitud incorrecta.")
		return
	}
	/// Conseguimos longitud de la url de la petición
	lon, error := strconv.ParseFloat(valoresUrl.Get("lon"), 64)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Longitud incorrecta.")
		return
	}
	// Consguimos datos del clima
	datos, error := cliente.ConseguirDatosClimaPorLatYLong(lat, lon)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, error)
		return
	}
	fmt.Fprint(w, datos)
}

// EjecutarServidor ejecuta el servidor en el puerto 8080.
func EjecutarServidor() {
	puerto := "8080"
	mux := http.NewServeMux()
	var cliente *clima.ClienteClima
	servidorArchivos := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", servidorArchivos))
	mux.HandleFunc("/", funcionHandle(cliente))
	log.Fatal(http.ListenAndServe("localhost:"+puerto, mux))

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
		case "/clima":
			climaHandler(cliente, w, r)
		}
	}
}

// ticketsHandler procesa peticiones relacionadas con mostrar la tabla de tickets.
// Recibe cliente para realizar peticións a openweather.
func ticketsHandler(cliente *clima.ClienteClima, w http.ResponseWriter, r *http.Request) {
	if cliente == nil {
		http.Redirect(w, r, "/acceso/", http.StatusUnauthorized)
		return
	}

	tickets, error := obtenerTicketsCsv()
	if error != nil {
		log.Println("No hay archivo csv con tickets.")
	}
	tpl := template.Must(template.ParseFiles("assets/templates/tickets.html"))
	tpl.Execute(w, tickets)
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
