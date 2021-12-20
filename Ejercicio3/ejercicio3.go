package main

import "fmt"



func main(){

fmt.Println("Ingrese su edad: ")
var edad int
fmt.Scanln((&edad))

if edad > 22 && edad != 0 {
	fmt.Println("¿Esta empleado?: ")
	var empleado bool
	fmt.Scanln((&empleado))
	if empleado == true {
		fmt.Println("Ingrese su antiguedad: ")
	var antiguedad int
	fmt.Scanln((&antiguedad))
	if antiguedad > 1 {
		fmt.Println("Ingrese su sueldo: ")
	var sueldo int
	fmt.Scanln((&sueldo))
		if sueldo > 100 {
			fmt.Println("Puede solicitar un prestamo")
		} else{
			fmt.Println("Se le cobraran intereses por su prestamo")
		}

	} else{
		fmt.Println("Usted no puede solicitar un prestamo")
	}
	} else{
		fmt.Println("Usted no puede solicitar un prestamo")
	}

} else {
	fmt.Println("Usted no puede solicitar un prestamo")
}


















}