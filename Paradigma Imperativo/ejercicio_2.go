package main

import (
	"fmt"
	"math"
)

/*
 2. Escriba el programa más eficiente que pueda para imprimir en pantalla la siguiente figura:

//     *
//   * * *
// * * * * *
//   * * *
//     *

//     *
//   * * *
//     *

//         *
//       * * *
//     * * * * *
//   * * * * * * *
//     * * * * *
//       * * *
//         *
Para dicho fin, escriba y use una función que reciba de parámetro la cantidad de elementos de la línea del centro, la cual debe ser impar positiva.
*/

func main() {
	var cantidad int = 19 //El numero impar positivo del medio
	if cantidad%2 != 0 && cantidad > 0 {
		dibujar(cantidad)
	}
}

func tab(cant_tabs int) {
	if cant_tabs < 0 {
		cant_tabs = -cant_tabs
	}
	for ; cant_tabs != 0; cant_tabs-- {
		print(" ")
	}
}

func dibujar(cant_elementos int) {
	var puntos int = 1
	for x := 0; x < cant_elementos; x++ {
		if x != cant_elementos {
			tab((cant_elementos / 2) - x)
		}

		for y := math.Abs(float64(puntos)); y > 0; y-- {
			fmt.Print("*")

		}
		if puntos >= cant_elementos {
			puntos = -puntos
		}
		puntos += 2

		fmt.Println("")
	}

}
