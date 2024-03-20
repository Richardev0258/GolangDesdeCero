package main

import (
	"github.com/Richardev0258/GolangDesdeCero/teclado"
)

func main() {
	//LLAMADO A OTRO PACKAGE
	/*estado, texto := variables.ConviertoaTexto(123)
	fmt.Println(estado)
	fmt.Println(texto)*/
	//BLOQUE IF - SWITCH
	/*
		os := runtime.GOOS
		if os == "Linux" || os == "OS X" {
			fmt.Println("Esto no es Windows, es ", os)
		} else {
			fmt.Println("Esto es Windows")
		}

		switch os := runtime.GOOS; os {
		case "Linux":
			fmt.Println("Esto es Linux")
		case "Darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("Esto es -> %s \n", os)

		}*/
	//EJERCICIO UNO
	/*numConv, mensaje := ejercicios.ConviertoaNumero("10")

	fmt.Println("El número convertido es: ", numConv)
	fmt.Println("y", mensaje)*/
	//PACKAGE BUFIO
	teclado.IngresoNumeros()

}
