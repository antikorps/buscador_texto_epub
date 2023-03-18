<script lang="ts">
  import Encabezado from "./Encabezado.svelte";
  import Formulario from "./Formulario.svelte";
  import type { RespuestaGUI } from "./modelos";
  import Resultados from "./Resultados.svelte";

  let rutaEpubs: string = "";
  let margen = 15;
  let termino = "";
  let respuestaGUI = {} as RespuestaGUI;
</script>

<main>
  <Encabezado />

  <Formulario bind:rutaEpubs bind:margen bind:termino bind:respuestaGUI />

  {#if respuestaGUI.errorCritico != undefined}
    <Resultados bind:respuestaGUI />
  {/if}
</main>
