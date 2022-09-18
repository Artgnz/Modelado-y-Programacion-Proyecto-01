/**
*Se encarga llamar las consultas de los climas de las ciudades de la fila en la que se dio click. 
*@param {HTMLTableElement} elemento - Elemento HTML
*/
function mostrarClimas(elemento){
    var tabla = document.getElementById('tabla');
    var numLinea = elemento.closest('tr').rowIndex;
    let linea = tabla.rows[numLinea].cells;

    consulta(linea[2].innerHTML,linea[3].innerHTML, numLinea, true, linea[1].innerHTML);//Consultamos clima origen
    consulta(linea[6].innerHTML,linea[7].innerHTML, numLinea, false, linea[5].innerHTML);//Consultamos clima destion):
}

/**
 *Consulta el clima de la ciudad y la imprime.
 *@param {Number} lat -Latitud de la ciudad.
 *@param {Number} lon - Longitud de la ciudad.
 *@param {Number} fila - Numero de fila en la que se encuentra la ciudad.
 *@param {Boolean} origen - Indica si la ciudad es de origen o destino.
 *@param {String} ciudad - Nombre de la ciudad.
 */
async function consulta(lat, lon, fila, origen, ciudad){
    const url = `http://127.0.0.1:8080/clima?lat=${lat}&lon=${lon}`;//Hacemos la consulta
    const response = await fetch(url);
    const data = await response.json();
    const weather = data.weather[0].description + " " + data.main.temp+ "°C";//Seleccionamos la descripcion del clima y lo guardamos en una constante

    switch(origen){
    case true :
	document.getElementsByClassName('clima')[2*(fila-2)].textContent=weather;//Imprimimos el clima de la ciudad de origen.
	break;
    case false:
	document.getElementsByClassName('clima')[2*(fila-2)+1].textContent=weather;//Imprimimos el clima de la ciudad de destino.
	break;
    }

    updateClima(ciudad,weather);
}
/**
 *Actualiza el clima en cada ciudad repetida dentro de la tabla.
 *@param {String} ciudad - Nombre de la ciudad.
 *@param {String} clima - Clima de la ciudad correspondiente.
*/
async function updateClima(ciudad, clima){
    let tabla = document.getElementById('tabla');
    let filas = tabla.rows;

    for(i=2; i<filas.length; i++){
	if(filas[i].cells[1].innerHTML==ciudad){
            filas[i].cells[4].innerHTML = clima;
	}

	if(filas[i].cells[5].innerHTML==ciudad){
            filas[i].cells[8].innerHTML = clima;
	}	
    }
}
