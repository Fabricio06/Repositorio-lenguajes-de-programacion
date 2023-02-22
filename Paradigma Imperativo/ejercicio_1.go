/*
		Solucion del ejercicio 1 a realizar
	 1. Haga un programa que cuente e indique el número de caracteres,
	    el número de palabras y el número de líneas de un texto cualquiera
	    (obtenido de cualquier forma que considere).
	    Considere hacer el cálculo de las tres variables en el mismo proceso.
*/
package main

import "fmt"

func contador(texto string) {
	var numero_caracteres int
	var numero_palabras int
	var numero_lineas int

	//En el caso de que se mande solo un caracter, no cuenta como palabra
	if len(texto) <= 1 {
		numero_caracteres = len(texto)
		numero_palabras = 0
		numero_lineas = 1

	} else {
		var separador bool = false
		numero_palabras++
		for x := 0; x < len(texto); x++ {

			if separador == true {
				separador = false
				numero_palabras++
				numero_caracteres++
			} else if string(texto[x]) != " " {
				numero_caracteres++
			} else if string(texto[x]) == " " {
				separador = true
			}

		}
	}

	fmt.Println("El texto: '"+texto+"'\n"+
		"analizando...\n"+
		"numero de caracteres: ", numero_caracteres, "\n"+
		"numero de palabras: ", numero_palabras, "\n"+
		"numero de lineas: ", numero_lineas)

}

func contador2(texto string) {
	var espacios_blancos int

	for x := 0; x < len(texto); x++ {
		if string(texto[x]) == " " {
			espacios_blancos++
		}
	}

	var numero_caracteres int = len(texto) - espacios_blancos
	var numero_palabras int = espacios_blancos
	var numero_lineas int

	fmt.Println("El texto: '"+texto+"'\n"+
		"analizando...\n"+
		"numero de caracteres: ", numero_caracteres, "\n"+
		"numero de palabras: ", numero_palabras, "\n"+
		"numero de lineas: ", numero_lineas)
}

func main() {
	//contador("A statement can precede conditionals any variables declared in this statement are available in the current and all subsequent branches")
	contador2("A statement can precede conditionals any variables declared in this statement are available in the current and all subsequent branches")
}
