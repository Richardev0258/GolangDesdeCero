package ejercicios

import (
	"fmt"
	"strconv"
)

func ConviertoaNumero(numeroStr string) (int, string) {
	numConv, error := strconv.Atoi(numeroStr)
	if error != nil {
		fmt.Println("Error al convertir la cadena a entero:", error)
		return 0, "se genero un error al convertir."
	}
	if numConv >= 100 {
		return numConv, "es mayor a 100"
	} else {
		return numConv, "es menor a 100"
	}

}
