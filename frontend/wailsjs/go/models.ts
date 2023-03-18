export namespace manejador {
	
	export class ResultadoGUI {
	    rutaEpub: string;
	    nombreArchivo: string;
	    coincidencias: string[];
	    error: boolean;
	    errorMensaje: string;
	
	    static createFrom(source: any = {}) {
	        return new ResultadoGUI(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.rutaEpub = source["rutaEpub"];
	        this.nombreArchivo = source["nombreArchivo"];
	        this.coincidencias = source["coincidencias"];
	        this.error = source["error"];
	        this.errorMensaje = source["errorMensaje"];
	    }
	}
	export class RespuestaGUI {
	    errorCritico: boolean;
	    errorCriticoMensaje: string;
	    resultados: ResultadoGUI[];
	
	    static createFrom(source: any = {}) {
	        return new RespuestaGUI(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.errorCritico = source["errorCritico"];
	        this.errorCriticoMensaje = source["errorCriticoMensaje"];
	        this.resultados = this.convertValues(source["resultados"], ResultadoGUI);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

