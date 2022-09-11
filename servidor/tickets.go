package servidor

import (
	"encoding/csv"
	"os"
)

// obtenerTicketsCsv lee archivo csv "datos/dataset1.csv" y
// devuelve un arreglo bidimensional con los valores leidos.
func obtenerTicketsCsv() ([][]string, error) {
	direccionArchivo := "datos/dataset1.csv"
	archivo, error := os.Open(direccionArchivo)
	if error != nil {
		return nil, error
	}
	defer archivo.Close()

	lectorCsv := csv.NewReader(archivo)
	lectorCsv.Read()
	valores, error := lectorCsv.ReadAll()
	if error != nil {
		return nil, error
	}
	return valores, nil
}
