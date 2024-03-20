package variables

import (
	"fmt"
	"strconv"
)

func ConviertoaNumero(numeroStr string) (int, string) {
	var textRet string
	numConv, error := strconv.Atoi(numeroStr)
	if error != nil {
		fmt.Println("Error al convertir la cadena a entero:", error)
		return 0, error.Error()
	}
	if numConv >= 100 {
		textRet = "Es mayor a 100"
	} else {
		textRet = "Es menor a 100"
	}
	return numConv, textRet
}
