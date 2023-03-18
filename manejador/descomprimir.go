package manejador

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func descomprimir(origen, destino string) error {
	manejador, manejadorError := zip.OpenReader(origen)
	if manejadorError != nil {
		return manejadorError
	}
	defer manejador.Close()

	for _, v := range manejador.File {
		err := descomprimirArchivo(v, destino)
		if err != nil {
			return err
		}
	}

	return nil
}

func descomprimirArchivo(f *zip.File, destino string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	rutaArchivo := filepath.Join(destino, f.Name)
	if !strings.HasPrefix(rutaArchivo, filepath.Clean(destino)+string(os.PathSeparator)) {
		return fmt.Errorf("ruta inválida: %s", rutaArchivo)
	}

	if f.FileInfo().IsDir() {
		crearCarpetaError := os.MkdirAll(rutaArchivo, 0777)
		if crearCarpetaError != nil {
			return crearCarpetaError
		}
		return nil
	}

	crearCarpetaError := os.MkdirAll(filepath.Dir(rutaArchivo), 0777)
	if crearCarpetaError != nil {
		return crearCarpetaError
	}

	manejadorDestino, manejadorDestinoError := os.OpenFile(rutaArchivo, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if manejadorDestinoError != nil {
		return manejadorDestinoError
	}
	defer manejadorDestino.Close()

	archivoDescomprimido, archivoDescomprimidoError := f.Open()
	if archivoDescomprimidoError != nil {
		return archivoDescomprimidoError
	}
	defer archivoDescomprimido.Close()

	_, copiarError := io.Copy(manejadorDestino, archivoDescomprimido)
	if copiarError != nil {
		return copiarError
	}
	return nil
}
