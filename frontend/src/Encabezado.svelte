
<section>
    <h3 id="encabezado-titulo">Buscador de contenido para EPUB</h3>
    <ol>
        <li>Selecciona la carpeta donde tienes almacenados los libros digitales.</li>
        <li>Valora si necesitas recuperar algún número de caracteres previos y posteriores a la coincidencia para contextualizar (límite).</li>
        <li>Incorpora el término de búsqueda (admite expresiones regulares).</li>
        <li>Consulta los resultados obtenidos 😎.</li>
    </ol>

    <details>
        <summary>Uso avanzado (línea de comandos)</summary>
        <p>La aplicación puede ejecutarse por línea de comandos para automatizar su utilización y obtener un un informe en formato CSV con los resultados. Su ejecución admite los siguientes argumentos o "flags":</p>
        <ul>
            <li><strong>-ruta</strong>: ruta completa donde se encuentran los libros digitales que se revisarán (uso obligatorio)</li>
            <li><strong>-busqueda</strong>: cadena a localizar en el interior del contenido (uso obligatorio).</li>
            <li><strong>-limite</strong>: número que caracteres por delante y por detrás que se recuperarán junto a las coincidencias (uso opcional, de lo contrario se utilizarán 20).</li>
            <li><strong>-informe</strong>: ruta completa donde se generará el informe en formato csv con los resultados (uso opcional, si no es especifica se creará un archivo llamado informe_resultados.csv en la misma carpeta que el ejecutable).</li>
            <li><strong>-simultaneidad</strong>: número de archivos que se analizarán simultáneamente (uso opcional, si no se especifica se analizarán 3).</li>
        </ul>

        <p>Ejemplo en sistemas operativos GNU/Linux:</p>
        <code>./buscardor_texto_epub -ruta /home/nombre_usuario/coleccion_epubs -busqueda "bases de datos" -limite 30 -informe /home/nombre_usuario/informes/resultados_base_datos.csv -simultaneidad 5</code>
    </details>
</section>

<style>
    #encabezado-titulo {
        text-align: center;
    }
</style>