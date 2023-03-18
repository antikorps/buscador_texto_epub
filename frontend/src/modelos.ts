export type { RespuestaGUI };

interface RespuestaGUI {
    errorCritico: boolean;
    errorCriticoMensaje: string;
    resultados: ResultadoGUI[];
  }
  interface ResultadoGUI {
    rutaEpub: string;
    nombreArchivo: string;
    coincidencias: string[];
    error: boolean;
    errorMensaje: string;
  }
  
