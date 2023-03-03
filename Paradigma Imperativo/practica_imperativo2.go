package main

import (
	"fmt"
	"sort"
	"time"
)

type producto1 struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto1

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	existeONO := l.buscarProducto(nombre)
	if existeONO == -1 {
		*l = append(*l, producto1{nombre: nombre, cantidad: cantidad, precio: precio})
	} else {
		(*l)[existeONO].cantidad += cantidad
		(*l)[existeONO].precio = precio

	}

	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result int = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cantidad_vender int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 {
		venta := (*l)[prod].cantidad - cantidad_vender
		if venta < 0 {
			println("No hay suficiente stock ", cantidad_vender, " solo le podemos vender: ", (*l)[prod].cantidad)
		} else if venta == 0 {
			*l = append((*l)[:prod], (*l)[prod+1:]...)
		} else {
			(*l)[prod].cantidad = (*l)[prod].cantidad - cantidad_vender
		}
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar
	} else {
		println("No tenemos ese producto ", nombre)
	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("café", 12, 4300)
	lProductos.venderProducto("arroz", 15)
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var listaNueva listaProductos
	for x := 0; x < len(*l); x++ {
		if (*l)[x].cantidad < existenciaMinima {
			valores := (*l)[x]
			listaNueva.agregarProducto(valores.nombre, valores.cantidad, valores.precio)
		}
	}
	return listaNueva
}

// Parte de lo pedido (A lo que yo entendi es de la lista de minimos, aumentar hasta la existencia minima)
// Tambien se podria directamente si es <existenciaMinima, asignar directamente el mismo valor, pero lo hice manual
func (l *listaProductos) aumentarInventarioDeMinimos() {

	for x := 0; x < len(*l); x++ {
		for true {
			if (*l)[x].cantidad < existenciaMinima {
				fmt.Println((*l)[x])
				(*l)[x].cantidad += 1
			} else {
				break
			}
		}

	}

}

// parte de lo pedido
func (l *listaProductos) ordenar() listaProductos {
	sort.Slice(*l, func(i, j int) bool {
		return (*l)[i].precio < (*l)[j].precio
	})
	return *l
}

func main() {
	llenarDatos()
	//fmt.Println(lProductos)
	//fmt.Println(lProductos.listarProductosMínimos())

	// a) Parte de lo pedido
	listaMinimos := lProductos.listarProductosMínimos()
	fmt.Println(listaMinimos)
	listaMinimos.aumentarInventarioDeMinimos() //No entendi porque en el ejemplo pedian un parametro (porque
	// ya haciendolo metodo, va a actuar sobre dicha lista)
	fmt.Println(listaMinimos) //Imprime los minimos con la existencia ya en lo minimo

	// b) parte de lo pedido
	fmt.Println("siguiente parte cargando...")
	time.Sleep(3 * time.Second)
	fmt.Println()
	listaMinimos = listaMinimos.ordenar()
	fmt.Println(listaMinimos)
}
