package manejador

type Manejador struct {
	Simultaneidad int
	RutaEpubs     string
	RutaTemporal  string
	Epubs         []string
	Resultados    []Resultado
	RutaInforme   string
	Termino       string
	Ejecucion     string
	ErrorCritico  string
	Limite        int
}

type Resultado struct {
	RutaEpub      string
	Coincidencias []string
	Error         error
}

type RespuestaGUI struct {
	ErrorCritico        bool           `json:"errorCritico"`
	ErrorCriticoMensaje string         `json:"errorCriticoMensaje"`
	Resultados          []ResultadoGUI `json:"resultados"`
}
type ResultadoGUI struct {
	RutaEpub      string   `json:"rutaEpub"`
	NombreArchivo string   `json:"nombreArchivo"`
	Coincidencias []string `json:"coincidencias"`
	Error         bool     `json:"error"`
	ErrorMensaje  string   `json:"errorMensaje"`
}
