# Buscador de contenido para EPUB

## Descripción
Buscador de coincidencias de términos en el interior (contenido) de EPUB. Cuenta con interfaz gráfica e interacción por línea de comandos.

## Instalación
No tiene ningún requisito ni dependencia especial, simplemente es necesario descargar el binario y ejecutable. Se facilitan opciones compiladas para GNU/Linux y Windows de arquitectura amd64. 

## Funcionamiento

1.  Selecciona la carpeta donde tienes almacenados los libros digitales.
2.  Valora si necesitas recuperar algún número de caracteres previos y posteriores a la coincidencia para contextualizar (límite).
3.  Incorpora el término de búsqueda (admite expresiones regulares).
4.  Consulta los resultados obtenidos 😎.

Uso avanzado (línea de comandos)

La aplicación puede ejecutarse por línea de comandos para automatizar su utilización y obtener un un informe en formato CSV con los resultados. Su ejecución admite los siguientes argumentos o "flags":

*   **\-ruta**: ruta completa donde se encuentran los libros digitales que se revisarán (uso obligatorio)
*   **\-busqueda**: cadena a localizar en el interior del contenido (uso obligatorio).
*   **\-limite**: número que caracteres por delante y por detrás que se recuperarán junto a las coincidencias (uso opcional, de lo contrario se utilizarán 20).
*   **\-informe**: ruta completa donde se generará el informe en formato csv con los resultados (uso opcional, si no es especifica se creará un archivo llamado informe\_resultados.csv en la misma carpeta que el ejecutable).
*   **\-simultaneidad**: número de archivos que se analizarán simultáneamente (uso opcional, si no se especifica se analizarán 3).

Ejemplo en sistemas operativos GNU/Linux:

`./buscardor_texto_epub -ruta /home/nombre_usuario/coleccion_epubs -busqueda "bases de datos" -limite 30 -informe /home/nombre_usuario/informes/resultados_base_datos.csv -simultaneidad 5`

![Captura de pantalla](https://i.imgur.com/xn6K9UH.png)
![Captura de pantalla](https://i.imgur.com/sd9TRby.png)
![Captura de pantalla](https://i.imgur.com/lLOXosF.png)
![Captura de pantalla](https://i.imgur.com/BwlxExP.png)
![Captura de pantalla](https://i.imgur.com/tgFxHmO.png)
![Captura de pantalla](https://i.imgur.com/TkrjeaD.png)
![Captura de pantalla](https://i.imgur.com/G0MEuIY.png)