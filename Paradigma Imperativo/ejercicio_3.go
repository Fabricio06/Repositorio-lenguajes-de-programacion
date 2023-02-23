package main

import (
	"fmt"
)

func main() {
	var secuenciaOriginal = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var pSecuencia *[]string = &secuenciaOriginal
	var cantidadDeRotaciones int = 3
	var direccion byte
	var direccionTextual

	switch direccion {
	case 0:
		direccionTextual = "izq"
	case 1:
		direccionTextual = "der"
	}
	fmt.Println("Secuencia original = ", *pSecuencia, "\n"+
		"Cantidad de rotaciones = ", cantidadDeRotaciones, "\n"+
		"Direccion = ",direccionTextual, "\n" +
		"Secuencia final rotada = ")

}

func rotarSecuencia(nPosiciones int) {

}
