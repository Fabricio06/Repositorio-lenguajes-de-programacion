package main

import "fmt"

/*4)	Escriba un programa que utilice una estructura y un arreglo/slice para almacenar en memoria un inventario
de una tienda que vende zapatos. Cada zapato debe contar con la información de su modelo (marca), su precio y la talla
del mismo que debe ir únicamente del 34 al 44. La estructura debe llamarse "calzado". El programa debe poder almacenar
la información “quemada” para 10+ zapatos del inventario y un stock de N cantidad de unidades de dicho zapato
(quiere decir por ejemplo que puede existir en inventario el zapato marca Nike, talla 42 y que cuesta 50000 colones,
pero además que puede existir no solo un par de esos, si no N pares del mismo zapato).

Haga una función que venda zapatos (eliminando del stock cada vez que haya venta e indicando que no se puede
vender porque ya o hay stock) y pruébela dentro del main haciendo las ventas que usted considere para demostrar su
funcionamiento.*/

type calzado struct {
	modelo string
	precio int
	talla  byte
}

// Funcion para restringir la talla del 34 al 44
func (c calzado) rangoValido() bool {
	return c.talla >= 34 && c.talla <= 44
}

// Agregar calzado a al inventario
func (aI *inventario) agregarCalzado(modelo string, precio int, talla byte) {
	nuevoCalzado := calzado{modelo: modelo, precio: precio, talla: talla}
	if nuevoCalzado.rangoValido() {
		*aI = append(*aI, nuevoCalzado)
		fmt.Println("Agregado")
	} else {
		fmt.Println("Esa talla no es valida")
	}
}

type inventario []calzado

var almaceneInventario inventario

func datosQuemados() {
	almaceneInventario.agregarCalzado("Adidas", 60, 34)
	almaceneInventario.agregarCalzado("Nike", 70, 35)
	almaceneInventario.agregarCalzado("Nike", 50, 36)
	almaceneInventario.agregarCalzado("Puma", 70, 37)
	almaceneInventario.agregarCalzado("Crocs", 20, 38)
	almaceneInventario.agregarCalzado("Adidas", 80, 39)
	almaceneInventario.agregarCalzado("Flier", 25, 40)
	almaceneInventario.agregarCalzado("zDesat", 54, 41)
	almaceneInventario.agregarCalzado("pLer", 99, 42)
	almaceneInventario.agregarCalzado("lito", 10, 44)
}

func (l *inventario) buscarZapato(modelo string, precio int, talla byte) bool {
	for _, c := range *l {
		if c.modelo == modelo && c.precio == precio && c.talla == talla {
			return true
		}
	}
	return false
}

func (l *inventario) venderZapatos(modelo string, precio int, talla byte) {
	if l.buscarZapato(modelo, precio, talla) == false {
		fmt.Println("No tenemos el zapato que quiere comprar")
	} else {
		for c := 0; c < len(*l); c++ {
			if (*l)[c].modelo == modelo && (*l)[c].precio == precio && (*l)[c].talla == talla {
				*l = append((*l)[:c], (*l)[c+1:]...)
				break
			}
		}
		fmt.Println("Comprado")
	}
}

func main() {
	datosQuemados()
	fmt.Println(almaceneInventario)
	almaceneInventario.venderZapatos("Adidas", 60, 34)
	fmt.Println(almaceneInventario)
	almaceneInventario.venderZapatos("Adidas", 60, 34)
}
