package manejador

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func NuevoManejador(ruta, termino, informe string, simultaneidad int, margen int, ejecucion string) Manejador {
	ejecutable, ejecutableError := os.Executable()
	if ejecutableError != nil {
		log.Fatalln(ejecutableError)
	}
	rutaEjecutable := filepath.Dir(ejecutable)

	if informe == "" {
		informe = filepath.Join(rutaEjecutable, "informe_resultados.csv")
	}

	rutaTemporal := filepath.Join(rutaEjecutable, "temp_archivos")
	os.Mkdir(rutaTemporal, 0777)

	return Manejador{
		RutaEpubs:     ruta,
		Termino:       termino,
		RutaTemporal:  rutaTemporal,
		RutaInforme:   informe,
		Simultaneidad: simultaneidad,
		Ejecucion:     ejecucion,
		Limite:        margen,
	}
}

func SeleccionarCarpeta(ctx context.Context) (string, error) {
	carpeta, carpetaError := runtime.OpenDirectoryDialog(ctx, runtime.OpenDialogOptions{})
	if carpetaError != nil {
		return carpeta, carpetaError
	}
	if carpeta == "" {
		return carpeta, errors.New("no se ha seleccionado ninguna carpeta")
	}
	return carpeta, nil

}
