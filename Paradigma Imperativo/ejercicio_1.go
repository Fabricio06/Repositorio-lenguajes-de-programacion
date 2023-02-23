/*
					Solucion del ejercicio 1 a realizar
	 1. Haga un programa que cuente e indique el número de caracteres,
	    el número de palabras y el número de líneas de un texto cualquiera
	    (obtenido de cualquier forma que considere).
	    Considere hacer el cálculo de las tres variables en el mismo proceso.
*/
package main

import "fmt"

func main() {
	contador("A statement can precede conditionals any variables declared in this statement are available in the current and all subsequent branches")
	contador("Hola, mi nombre es Fabricio\n me gusta programar\n entre otras cosas")
}

/*
La solucion que implemente, se encarga de tomar el len del texto y se le restan
los espacios, valiendo como los caracteres totales que existen

# No se toman en cuenta si el usuario dijita mas de un espacio entre palabras

# Las vocales se toman como palabras al igual que word

# A los caracteres se le restan el valor de \n palabras encontradas
*/
func contador(texto string) {

	fmt.Println()
	var espacios_blancos int
	var cambios int

	for x := 0; x < len(texto); x++ {
		if string(texto[x]) == " " {
			espacios_blancos++
		} else if string(texto[x]) == "\n" {
			cambios++
		}
	}

	var numero_caracteres int = len(texto) - (espacios_blancos + cambios)
	var numero_palabras int = espacios_blancos + 1 //Por la primera palabra
	var numero_lineas int

	fmt.Println("El texto: '"+texto+"'\n"+
		"analizando...\n"+
		"numero de caracteres: ", numero_caracteres, "\n"+
		"numero de palabras: ", numero_palabras, "\n"+
		"numero de lineas: ", numero_lineas)
	fmt.Println()
}
