<script lang="ts">
    import type { RespuestaGUI } from "./modelos";

    export let respuestaGUI: RespuestaGUI;

    function adecuarScrollResultados() {
        const resultados = document.querySelector("#resultados");
        if (resultados === null) {
            return;
        }
        resultados.scrollIntoView({
            behavior: "smooth",
        });
    }

    $: if (respuestaGUI) {
        adecuarScrollResultados();
    }
</script>

<h3 id="resultados">Resultados:</h3>
<hr />
{#if respuestaGUI.errorCritico}
    <h3>
        Se ha producido un error crítico que ha impedido realizar la búsqueda:
    </h3>
    <code>{respuestaGUI.errorCriticoMensaje}</code>
{:else}
    {#each respuestaGUI.resultados as resultado}
        <div class="resultado">
            <p>
                <strong>{resultado.nombreArchivo}</strong>
                ({resultado.rutaEpub})
            </p>
            {#if resultado.error}
                <code>{resultado.errorMensaje}</code>
            {:else}
                <p>
                    <strong>Coincidencias:</strong>
                    {resultado.coincidencias.length}
                </p>
                {#if resultado.coincidencias.length > 0}
                    <details>
                        <summary>Relación coincidencias</summary>
                        <ol>
                            {#each resultado.coincidencias as coincidencia}
                                <li><code>{coincidencia}</code></li>
                            {/each}
                        </ol>
                    </details>
                {/if}
            {/if}
        </div>
        <hr />
    {/each}
{/if}

<style>
    .resultado {
        margin: 40px 0 40px 0;
    }
</style>
