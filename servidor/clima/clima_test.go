package clima

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// TestNuevoClienteClima prueba si la función nuevoClienteClima
// crea satisfactoriamente un cliente.
// Para la prueba se necesita que LLAVE_API
// esté definida en el archivo .env en la ráiz del proyecto y que la
// llave sea válida.
func TestNuevoClienteClima(t *testing.T) {
	llaveApi := obtenerLlaveApi(t)

	if llaveApi == "" {
		t.Errorf("LLAVE_API no definida en archivo .env,")
	}

	cliente, error := NuevoClienteClima(&http.Client{}, llaveApi)
	if error != nil {
		t.Errorf("Se esperaba nil, se obtuvo %v", error)
	}
	if cliente == nil {
		t.Errorf("Se esperaba un nuevo, se obtuvo nil")
	}
}

// TestEsLlaveApiValida prueba si la función esLlaveApiValida funciona
// correctamente.
// Para la prueba se necesita que LLAVE_API
// esté definida en el archivo .env en la ráiz del proyecto y que la
// llave sea válida.
func TestEsLlaveApiValida(t *testing.T) {
	llaveApi := obtenerLlaveApi(t)
	// Tabla para realizar pruebas.
	var pruebas = []struct {
		llave    string
		esperado bool
	}{
		{llaveApi, true},
		{"8o80hnrandom", false},
	}
	for _, tt := range pruebas {
		// Nombre de nombrePrueba a ejecutar
		nombrePrueba := fmt.Sprintf("%s,%v", tt.llave, tt.esperado)
		t.Run(nombrePrueba, func(t *testing.T) {
			obtenido := esLlaveApiValida(tt.llave)
			if obtenido != tt.esperado {
				t.Errorf("Se obtuvo %v, se esperaba %v", obtenido, tt.esperado)
			}
		})
	}
}

// TestConseguirDatosClimaPorIdCiudad prueba si la función conseguirDatosClimaPorIdCIudad
// consigue correctamente los datos.
func TestConseguirDatosClimaPorLatYLong(t *testing.T) {
	// Llave con la que se realizarán las peticiones.
	llaveApi := obtenerLlaveApi(t)
	lat := 19.3371
	long := -99.566
	cliente, error := NuevoClienteClima(&http.Client{}, llaveApi)
	if error != nil {
		t.Errorf("Error al crear cliente.")
	}
	datos, error := cliente.conseguirDatosClimaPorLatYLong(lat, long)
	if error != nil {
		t.Errorf("Se esperaba nil, se obtuvo %v", error)
	}
	if datos == "" {
		t.Errorf("Se obtuvo datos vacíos.")
	}
}

// Función auxiliar de las pruebas que obtiene la variable LLAVE_API de archivo .env.
// Recibe t, para mostrar errores en la prueba t si no es posible obtener LLAVE_API.
func obtenerLlaveApi(t *testing.T) string {
	// Se lee archivo .env
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Errorf("Error al leer archivo .env.")
		return ""
	}
	llaveApi := os.Getenv("LLAVE_API")
	if llaveApi == "" {
		t.Errorf("LLAVE_API no definida en archivo .env,")
		return ""
	}
	return llaveApi
}
