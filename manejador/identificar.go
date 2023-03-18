package manejador

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func (m *Manejador) IdentificarEpubs() {
	var epubs []string
	filepath.Walk(m.RutaEpubs, func(ruta string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(ruta, ".epub") {
			epubs = append(epubs, ruta)
		}
		return nil
	})

	if len(epubs) == 0 {
		if m.Ejecucion == "CLI" {
			log.Fatalln("no se han encontrado archivos de tipo .epub en la ruta indicada")
		}
		m.ErrorCritico = "no se han encontrado archivos de tipo .epub en la ruta indicada"
	}

	m.Epubs = epubs
}
