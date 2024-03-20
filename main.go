package main

import (
	"fmt"

	"github.com/Richardev0258/GolangDesdeCero/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(123)
	fmt.Println(estado)
	fmt.Println(texto)
}
