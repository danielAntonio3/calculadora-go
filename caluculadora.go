package calculadora_go

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calc struct
type Calc struct {
	Value1   int
	Value2   int
	Operator string
}

// GetOperation obtener el resultado de la operación
func (c Calc) GetOperation() {

	switch c.Operator {
	case "+":
		fmt.Println("La suma es: ", c.Value1+c.Value2)
	case "-":
		fmt.Println("La resta es: ", c.Value1-c.Value2)
	case "*":
		fmt.Println("La multiplicación es: ", c.Value1*c.Value2)
	case "/":
		if c.Value2 != 0 {
			fmt.Println("La division es: ", c.Value1/c.Value2)
			break
		}
		fmt.Println("La division no se puedo realizar")
	default:
		fmt.Printf("Aun no sopostamos operaciones con %s", c.Operator)
	}

}

// GetValue obtener el valor de la cadena de entrada
func GetValue(value string) int {
	valor, err := strconv.Atoi(value)

	if err != nil {
		fmt.Println("hay un error en uno de los valores")
	}
	return valor
}

// GetTypeOperation obtenemos el signo de la operación
func GetTypeOperation(operator string) string {

	if strings.Contains(operator, "+") {
		return "+"
	}
	if strings.Contains(operator, "-") {
		return "-"
	}
	if strings.Contains(operator, "*") {
		return "*"
	}
	if strings.Contains(operator, "/") {
		return "/"
	}
	return "0"
}

// LeerEntrada para poder escribir en consola la operación
func LeerEntrada() string {
	fmt.Print("Escribe una operación (+,-,*,/): ")

	//	Recibir input del usuario
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Guardamos el texto que nos escribe
	return scanner.Text()
}

//func main() {
//	operacion := LeerEntrada()
//	// Obtener el operador
//	operator := GetTypeOperation(operacion)
//
//	// Divide la cadena de texto, con un valor como patron de separación
//	valores := strings.Split(operacion, operator)
//
//	// Creamos un objeto de nuestro struct
//	calculadora := Calc{}
//	calculadora.value1 = GetValue(valores[0])
//	calculadora.value2 = GetValue(valores[1])
//	calculadora.operator = operator
//
//	// Realizamos la operación
//	calculadora.GetOperation()
//
//}
