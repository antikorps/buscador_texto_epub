package manejador

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

var extensionValidas = []string{"html", "htm"}

func esExtensionValida(ruta string) bool {
	for _, v := range extensionValidas {
		if strings.HasSuffix(ruta, v) {
			return true
		}
	}
	return false

}

func cribarArchivos(rutaCarpeta string) []string {
	var archivosValidos []string

	filepath.Walk(rutaCarpeta, func(ruta string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if esExtensionValida(ruta) {
			archivosValidos = append(archivosValidos, ruta)
		}
		return nil
	})

	return archivosValidos

}

func contarCoincidencias(archivosAnalizar []string, termino string, margen int) ([]string, error) {
	var totalCoincidencias []string

	for _, v := range archivosAnalizar {
		leerArchivo, leerArchivoError := os.ReadFile(v)
		if leerArchivoError != nil {
			return totalCoincidencias, leerArchivoError
		}
		expRegSaltos := regexp.MustCompile(`[\n|\r]`)
		contenido := expRegSaltos.ReplaceAllString(string(leerArchivo), " ")
		expRegScript := regexp.MustCompile(`<script.*?</script>`)
		contenido = expRegScript.ReplaceAllString(contenido, "")
		expRegStyle := regexp.MustCompile(`<style.*?</script>`)
		contenido = expRegStyle.ReplaceAllString(contenido, "")
		expRegTags := regexp.MustCompile(`<.*?>`)
		contenido = expRegTags.ReplaceAllString(contenido, "")

		expRegMargen := fmt.Sprintf(".{0,%v}%v.{0,%v}", margen, termino, margen)
		expReg := regexp.MustCompile(expRegMargen)

		coincidencias := expReg.FindAllString(contenido, -1)

		totalCoincidencias = append(totalCoincidencias, coincidencias...)

	}

	return totalCoincidencias, nil
}

func buscarEnEpub(canal chan Resultado, rutaEpub, rutaTemporal, termino string, margen int) {
	defer wg.Done()

	nombreArchivo := strings.TrimSuffix(filepath.Base(rutaEpub), ".epub")
	carpetaDestino := filepath.Join(rutaTemporal, fmt.Sprintf("%v_%v", time.Now().Unix(), nombreArchivo))
	os.Mkdir(carpetaDestino, 0777)

	descomprimirError := descomprimir(rutaEpub, carpetaDestino)
	if descomprimirError != nil {
		canal <- Resultado{
			Error:    descomprimirError,
			RutaEpub: rutaEpub,
		}
		return
	}

	archivosAnalizar := cribarArchivos(carpetaDestino)
	if len(archivosAnalizar) == 0 {
		canal <- Resultado{
			Error:    errors.New("no se han encontrado archivos con contenido en el libro digital"),
			RutaEpub: rutaEpub,
		}
		return
	}

	coincidencias, coincidenciasError := contarCoincidencias(archivosAnalizar, termino, margen)
	if coincidenciasError != nil {
		canal <- Resultado{
			Error:    coincidenciasError,
			RutaEpub: rutaEpub,
		}
		return
	}

	canal <- Resultado{
		Error:         nil,
		RutaEpub:      rutaEpub,
		Coincidencias: coincidencias,
	}

}

func (m *Manejador) RealizarBusqueda() {
	if m.ErrorCritico != "" {
		return
	}

	var coleccionResultados []Resultado

	continuar := true
	indiceColeccion := 0

	for continuar {
		canal := make(chan Resultado)

	controlIndices:
		for i := 0; i < m.Simultaneidad; i++ {
			if indiceColeccion == len(m.Epubs) {
				continuar = false
				break controlIndices
			}

			wg.Add(1)
			go buscarEnEpub(canal, m.Epubs[indiceColeccion], m.RutaTemporal, m.Termino, m.Limite)
			indiceColeccion++
		}
		go func() {
			wg.Wait()
			close(canal)
		}()

		for v := range canal {
			coleccionResultados = append(coleccionResultados, v)
		}

	}
	m.Resultados = coleccionResultados
}
