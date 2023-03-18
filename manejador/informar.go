package manejador

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (m *Manejador) InformarResultados() {
	archivoInforme, archivoInformeError := os.Create(m.RutaInforme)
	if archivoInformeError != nil {
		log.Fatalln(archivoInformeError)
	}
	defer archivoInforme.Close()

	var lineasCSV [][]string
	encabezados := []string{"archivo", "numero_coincidencias", "coincidencias", "error"}
	lineasCSV = append(lineasCSV, encabezados)
	for _, v := range m.Resultados {
		var lineaCSV []string
		var celdaFichero, celdaNumeroCoincidencias string
		celdaError := "no"
		if v.Error != nil {
			celdaError = v.Error.Error()
		}
		celdaFichero = v.RutaEpub
		celdaNumeroCoincidencias = fmt.Sprint(len(v.Coincidencias))

		var celdaCoincidencias string
		coincidenciasJSON, coincidenciasJSONError := json.Marshal(v.Coincidencias)
		if coincidenciasJSONError == nil {
			celdaCoincidencias = string(coincidenciasJSON)
		}

		lineaCSV = append(lineaCSV, celdaFichero, celdaNumeroCoincidencias, celdaCoincidencias, celdaError)
		lineasCSV = append(lineasCSV, lineaCSV)
	}

	manejadorCSV := csv.NewWriter(archivoInforme)
	errorEscritura := manejadorCSV.WriteAll(lineasCSV)
	manejadorCSV.Flush()
	if errorEscritura != nil {
		log.Fatalln(errorEscritura)
	}

	os.RemoveAll(m.RutaTemporal)

}

func (m *Manejador) InformarGUI() RespuestaGUI {
	var respuestaGUI RespuestaGUI

	if m.ErrorCritico != "" {
		respuestaGUI.ErrorCritico = true
	}

	respuestaGUI.ErrorCriticoMensaje = m.ErrorCritico

	for _, v := range m.Resultados {
		var resultadoGUI ResultadoGUI
		resultadoGUI.RutaEpub = v.RutaEpub
		resultadoGUI.NombreArchivo = filepath.Base(v.RutaEpub)
		if len(v.Coincidencias) == 0 {
			forzarSlice := make([]string, 0)
			resultadoGUI.Coincidencias = forzarSlice
		} else {
			resultadoGUI.Coincidencias = v.Coincidencias
		}
		if v.Error != nil {
			resultadoGUI.Error = true
			resultadoGUI.ErrorMensaje = v.Error.Error()
		}
		respuestaGUI.Resultados = append(respuestaGUI.Resultados, resultadoGUI)
	}

	os.RemoveAll(m.RutaTemporal)

	return respuestaGUI
}
