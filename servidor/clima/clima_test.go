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
	// Se lee archivo .env
	err := godotenv.Load("../../.env")

	if err != nil {
		t.Errorf("Error al leer archivo .env.")
	}

	llaveApi := os.Getenv("LLAVE_API")

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
	// Se lee archivo .env
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Errorf("Error al leer archivo .env.")
	}

	llaveApi := os.Getenv("LLAVE_API")

	if llaveApi == "" {
		t.Errorf("LLAVE_API no definida en archivo .env,")
	}

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
