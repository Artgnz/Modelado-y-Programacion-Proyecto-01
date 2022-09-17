function mostrarClimas(elemento){
    var tabla = document.getElementById('tabla');
    var numLinea = elemento.closest('tr').rowIndex;
    let linea = tabla.rows[numLinea].cells;

    var climaO = tabla.rows[numLinea].cells[4];//Referimos a la celda del origen
    var climaD = tabla.rows[numLinea].cells[8]; //Referimos a la celda del destino

    consulta(linea[2].innerHTML,linea[3].innerHTML, numLinea, true);//Consultamos clima origen
    consulta(linea[6].innerHTML,linea[7].innerHTML, numLinea, false);//Consultamos clima destion):
}

async function consulta(lat, lon, fila, origen){
    const url = `http://127.0.0.1:8080/clima?lat=${lat}&lon=${lon}`;
    const response = await fetch(url);
    const data = await response.json();
    const weather = data.weather[0].description;

    switch(origen){
    case true :
	document.getElemenstByClassName('clima')[2*fila].textContent=weather;
	break;
    case false:
	document.getElemenstByClassName('clima')[2*fila+1].textContent=weather;
	break;
    }

}
