package main

import "fmt"
import "strings"

func main() {
	fmt.Println("Ingrese una palabra: ")
	var palabra string
	fmt.Scanln((&palabra))

	

	fmt.Println("La cantidad de letras es:", len(palabra))
	letraCompleta := strings.SplitAfter(palabra, "")


	for i := 0; i < len(letraCompleta); i++{
		fmt.Println((letraCompleta[i]))
	}
	
}

