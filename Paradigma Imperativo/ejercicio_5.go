package main

import "fmt"

/*5)Implemente un programa que, utilizando pilas creadas a base de slices, realice un evaluador de expresiones
matemáticas expresadas en notación inflija. Se debe probar el programa a partir de al menos 3 expresiones con diferentes
características dadas por el programador (el main tendrá tres llamadas al evaluador con 3 expresiones diferentes)

Como ejemplos de expresiones y de evaluación se adjuntan las siguientes figuras. De ser necesario se puede ahondar más
en las expresiones en infijo y en el algoritmo de evaluación usando pilas.
			---Evaluacion de izquierda a derecha--->

	4		5		6		*				+
	|		|		|		|				|
    |		|		|	 extraer dos	extraer dos
	|		|		|	 veces y hacer	veces y hacer
  incluir incluir incluir aritmetica	aritmetica
    |		|	   [6]		|				|
	|	   [5]     [5]	   [30]				|
   [4]     [4]     [4] 	   [4]			   [34]
*/



Implemente un programa que, utilizando pilas creadas a base de slices, realice un evaluador de expresiones
matemáticas expresadas en notación inflija. Se debe probar el programa a partir de al menos 3 expresiones con diferentes
características dadas por el programador (el main tendrá tres llamadas al evaluador con 3 expresiones diferentes)

func main(){
	fmt.Println()
}