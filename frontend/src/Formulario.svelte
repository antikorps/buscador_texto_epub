<script lang="ts">
    import { DevolverResultados, DevolverSeleccionarCarpeta } from "../wailsjs/go/main/App.js";
    import type { RespuestaGUI } from "./modelos.js";
    import cargando from './assets/imagenes/cargando.gif'

    export let rutaEpubs: string;
    export let margen: number;
    export let termino: string;
    export let respuestaGUI: RespuestaGUI;

    let cumplirRequisitos = "";
    let rutaEpubsError = "";
    let deshabilitar = false;
    let deshabilitarEstilo = "";
    let procesando = "";


    async function iniciarBusqueda() {
        // Resetear
        respuestaGUI = {} as RespuestaGUI
        cumplirRequisitos = "";
        rutaEpubsError = "";
        deshabilitar = false;
        deshabilitarEstilo = "";
        procesando = "";

        // Cargando
        deshabilitar = true
        deshabilitarEstilo = `background-image: url(${cargando});background-repeat:no-repeat;background-position:center center`
        procesando = "procesando"


        if (rutaEpubs == "") {
            cumplirRequisitos =
                "Debe seleccionarse la carpeta con los archivos";
        }

        respuestaGUI = await DevolverResultados(
            rutaEpubs,
            termino,
            margen
        );

        console.log(respuestaGUI)
       

        // Fin cargando
        deshabilitar = false
        deshabilitarEstilo = ""
        procesando = ""
    }

    async function seleccionarCarpeta() {
        // Reset
        respuestaGUI = {} as RespuestaGUI
        cumplirRequisitos = "";
        rutaEpubsError = "";
        deshabilitar = false;
        deshabilitarEstilo = "";
        procesando = "";


        try {
            rutaEpubs = await DevolverSeleccionarCarpeta()
            rutaEpubsError = ""
        }
        catch(error) {
            rutaEpubs = ""
            rutaEpubsError = error
        }
    }
</script>



<section>
    <form on:submit|preventDefault={iniciarBusqueda} style={deshabilitarEstilo} class={procesando}>
        <fieldset disabled={deshabilitar}>
            <label>
                Directorio a comprobar:
                <input type="button" value="Seleccionar carpeta" on:click={seleccionarCarpeta}/>
            
            </label>
            {#if rutaEpubsError != ""}
                <code>
                    {rutaEpubsError}
                </code>
            {/if}
            {#if rutaEpubs != ""}
                <p><strong>Ruta escogida</strong>: {rutaEpubs}</p>
            {/if}
        
        <hr>
        <label for="margen">
            Margen caracteres:
            <input type="number" bind:value={margen} required />
        </label>

        <hr>
        <label for="termino">
            Término a buscar: <input type="text" bind:value={termino} required />

        </label>
    </fieldset>
        
        {#if cumplirRequisitos != ""}
            {cumplirRequisitos}
        {/if}
        <button id="buscar">Buscar</button>
    </form>
</section>

<style>
    #buscar {
        display:block;
        margin: 0 auto;
    }
    
    .procesando * {
        opacity: 0.4
    }

</style>


