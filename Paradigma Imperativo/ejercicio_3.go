package main

import (
	"fmt"
)

/* 3)	Escriba una función que permita rotar una secuencia de elementos N posiciones a la izquierda o a la derecha
según sea el caso en función al parámetro que se reciba.
Los parámetros deben ser un puntero a un arreglo previamente creado,
la cantidad de movimiento de rotación y la dirección (0 = izq y 1 = der).

A partir de esta función, escriba un programa que haga varias rotaciones
cualesquiera de una secuencia de elementos previamente creada que sea inmutable.
Al final debe imprimirse el resultado de cada rotación además de la secuencia original.

	Un ejemplo de rotación puede ser el siguiente:
Secuencia Original = [a, b, c, d, e, f, g, h,]
Cantidad de rotaciones = 3
Dirección = izq
Secuencia final rotada = [d, e, f, g, h, a, b, c]   Nótese que cada iteración, el elemento más a la izquierda pasó a
formar parte del final de la secuencia, si la rotación fuera a la derecha, sería lo contrario */

func main() {
	var secuenciaOriginal = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var secuenciaAuxiliar = secuenciaOriginal
	var cantidadDeRotaciones int = 3
	var direccion byte = 1
	var direccionTextual = ""

	var pSecuencia *[]string = &secuenciaOriginal //Puntero de la lista original
	rotarSecuencia(pSecuencia, cantidadDeRotaciones, direccion)

	switch direccion { //switch solo para darle el formato del texto a la direccion como se encuentra en el ejemplo
	case 0:
		direccionTextual = "izq"
	case 1:
		direccionTextual = "der"
	}

	fmt.Println("***************Ejercicio de rotacion***************")
	fmt.Println("\nSecuencia original = ", secuenciaAuxiliar, "\n"+
		"Cantidad de rotaciones = ", cantidadDeRotaciones, "\n"+
		"Direccion = ", direccionTextual, "\n"+
		"Secuencia final rotada = ", *pSecuencia)
	fmt.Println("****************Fin********************************")

}

func rotarSecuencia(pLista *[]string, nPosiciones int, direccion byte) {

	if nPosiciones > len(*pLista) {
		nPosiciones -= len(*pLista)
	}

	if direccion == 0 {
		*pLista = append((*pLista)[nPosiciones:], (*pLista)[:nPosiciones]...)
	} else {
		*pLista = append((*pLista)[len(*pLista)-nPosiciones:], (*pLista)[:len(*pLista)-nPosiciones]...)
	}
}
