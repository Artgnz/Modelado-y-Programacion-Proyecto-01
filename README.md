# Modelado-y-Programación-Proyecto-01
Informes Climáticos AA es un proyecto hecho en go, html y javascript que permite conseguir el informe del clima de ciudades de salida y llegada a partir de un archivo CSV.

## Integrantes
    1. Arturo González Peñaloza
    2. Emilio Arsenio Raudry Rico
    
## Prerequisitos
    1. Contar con sistema operativo Linux, Mac OS o Windows.
    2. [Instalar go v1.19](https://go.dev/doc/install) de acuerdo al sistema operativo que use.
    Para verificar la versión usada, ejecute:
    ```bash
        go version
    ``` 
    3. Contar con un navegador web ([Safari](https://support.apple.com/downloads/safari), [Chrome](https://support.google.com/chrome/answer/95346?hl=en&co=GENIE.Platform%3DDesktop), [Firefox](https://www.mozilla.org/en-US/firefox/new/), etc)
    4. Tener una llave api de [OpenWeather](https://openweathermap.org). Para conseguir la llave seguir los siguientes pasos:
       3.1 [Registrarse](https://home.openweathermap.org/users/sign_up).
       3.2 En la parte superior escoger su nombre de usuario y del menú desplegable seleccionar "My API keys".
       3.3 Su llave API se muestra en el centro de la pagina.
## Instalación
    Clone el repositorio:
    ```bash
        git clone https://github.com/Artgnz/Modelado-y-Programacion-Proyecto-01.git
    ```
## Uso
Primero, colóquese en el directorio Modelado-y-Programacion-Proyecto-01
```bash
    cd Modelado-y-Programacion-Proyecto-01
```
Para compilar ejecute:
```bash
    go build 
```
Para ejecutar el programa, use el comando:
```bash
    ./Modelado-y-Programacion-Proyecto-01
```
Posteriormente, abra el navegador web de su preferencia y busque la dirección http://localhost:8080/acceso/

## Pruebas
    1. Para ejecutar las pruebas es necesario crear un archivo titulado .env en este directorio y poner en el archivo
    ```
    LLAVE_API=TU_LLAVE_API
    ```
    donde sustituyes "TU_LLAVE_API" por tu llave api de openweather.
    2. Ejecutar 
    ```bash
        go test --v ./...
    ```
## Uso de bibliotecas externas.
    1. godotenv
       La biblioteca godotenv la ocupamos para leer variables de entorno de entorno de un archivo .env. De esta biblioteca, únicamente ocupamos la función 
       ```
       godotenv.Load(archivo)
       ```
       para leer las variables del archivo .env, esta función lee el archivo .env que le indicamos y guarda las variables de ambiente que encuentra.
