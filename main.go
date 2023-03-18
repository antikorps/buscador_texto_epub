package main

import (
	"changeme/manejador"
	"embed"
	"flag"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	rutaEpubs := flag.String("ruta", "", "ruta completa donde se encuentra los epubs")
	terminoBusqueda := flag.String("busqueda", "", "cadena a buscar (admite expresiones regulares)")
	rutaInforme := flag.String("informe", "", "ruta completa donde generar el informe csv con los resultados")
	simultaneidad := flag.Int("simultaneidad", 3, "numero de archivos que se analizaran de manera concurrente")
	limite := flag.Int("limite", 20, "número de caracteres que se recuperarán antes y después del término buscado")
	flag.Parse()

	if *rutaEpubs != "" || *terminoBusqueda != "" {
		manejador := manejador.NuevoManejador(*rutaEpubs, *terminoBusqueda, *rutaInforme, *simultaneidad, *limite, "CLI")
		manejador.IdentificarEpubs()
		manejador.RealizarBusqueda()
		manejador.InformarResultados()
		return
	}

	app := NuevaApp()

	err := wails.Run(&options.App{
		Title:  "buscador_texto_epub",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.iniciar,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
