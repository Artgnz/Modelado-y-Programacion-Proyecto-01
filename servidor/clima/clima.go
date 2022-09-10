package clima

import (
	"fmt"
	"log"
	"net/http"
)

// clienteClima es un cliente HTTP con una llave api asociada.
// Tendrá comunicación con la api de openweather.
type ClienteClima struct {
	// Cliente http para realizar peticiones
	http *http.Client
	// Llave api de openweather
	llaveApi string
}

// NuevoClienteClima crea y devuelve un nuevo clienteClima con una llave api asociada.
// En caso de que la llave sea inválida, devuelve un error.
func NuevoClienteClima(clienteHttp *http.Client, llaveApi string) (*ClienteClima, error) {
	if !esLlaveApiValida(llaveApi) {
		return nil, fmt.Errorf("Llave api inválida.")
	}
	return &ClienteClima{clienteHttp, llaveApi}, nil
}

// esLlaveApiValida verifica que la llave api introducida sea válida.
func esLlaveApiValida(llaveApi string) bool {
	url := "https://api.openweathermap.org/data/2.5/weather?lat=44.34&lon=10.99&appid=" + llaveApi
	resp, err := http.Get(url)
	// Si hubo un error al obtener la respuesta.
	if err != nil {
		log.Println("Error:", err)
		return false
	}
	// Si la el código de estatus de la respuesta no es OK.
	if resp.StatusCode != http.StatusOK {
		log.Println("Error:", err)
		return false
	}
	return true
}
