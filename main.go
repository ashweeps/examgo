package main

import "fmt"

var productosVendidos []string
var subtotales []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)

	productosVendidos = append(productosVendidos, nombre)
	subtotales = append(subtotales, subtotal)

	fmt.Println("Venta registrada correctamente.")
}

func MostrarEstadisticas() {
	if len(subtotales) == 0 {
		fmt.Println("No existen ventas registradas.")
		return
	}

	total := 0.0

	for i := 0; i < len(subtotales); i++ {
		total += subtotales[i]
	}

	fmt.Println("Total recaudado: $", total)
}

func main() {
	opcion := 0

	for opcion != 3 {
		fmt.Println("MENU")
		fmt.Println("1.Registrar nueva venta")
		fmt.Println("2.Mostrar estadisticas")
		fmt.Println("3.Salir")
		fmt.Print("Seleccione una opcion: ")
		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			fmt.Println("PRODUCTOS")
			fmt.Println("1.Arroz - $1.25")
			fmt.Println("2.Leche - $0.95")
			fmt.Println("3.Pan - $0.50")

			var producto int
			var cantidad int
			var nombre string
			var precio float64

			fmt.Print("Seleccione el producto por el numero del literal: ")
			fmt.Scan(&producto)

			switch producto {
			case 1:
				nombre = "Arroz"
				precio = 1.25

			case 2:
				nombre = "Leche"
				precio = 0.95

			case 3:
				nombre = "Pan"
				precio = 0.50

			default:
				fmt.Println("Producto no valido.")
				continue
			}

			fmt.Print("Ingrese la cantidad: ")
			fmt.Scan(&cantidad)

			RegistrarVenta(nombre, precio, cantidad)

		case 2:
			MostrarEstadisticas()

		case 3:
			fmt.Println("Programa finalizado.")

		default:
			fmt.Println("Opcion no valida.")
		}
	}
}
