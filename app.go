package main

import (
	"changeme/manejador"
	"context"
)

type App struct {
	ctx context.Context
}

func NuevaApp() *App {
	return &App{}
}

func (a *App) iniciar(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DevolverResultados(ruta, termino string, margen int) manejador.RespuestaGUI {
	manejador := manejador.NuevoManejador(ruta, termino, "", 3, margen, "GUI")
	manejador.IdentificarEpubs()
	manejador.RealizarBusqueda()
	return manejador.InformarGUI()
}

func (a *App) DevolverSeleccionarCarpeta() (string, error) {
	return manejador.SeleccionarCarpeta(a.ctx)
}
