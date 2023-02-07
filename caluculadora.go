package calculadora_go

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {
	value1   int
	value2   int
	operator string
}

func (c Calc) GetOperation() {

	switch c.operator {
	case "+":
		fmt.Println("La suma es: ", c.value1+c.value2)
	case "-":
		fmt.Println("La resta es: ", c.value1-c.value2)
	case "*":
		fmt.Println("La multiplicación es: ", c.value1*c.value2)
	case "/":
		if c.value2 != 0 {
			fmt.Println("La division es: ", c.value1/c.value2)
			break
		}
		fmt.Println("La division no se puedo realizar")
	default:
		fmt.Printf("Aun no sopostamos operaciones con %s", c.operator)
	}

}

func GetValue(value string) int {
	valor, err := strconv.Atoi(value)

	if err != nil {
		fmt.Println("hay un error en uno de los valores")
	}
	return valor
}

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

func LeerEntrada() string {
	fmt.Print("Escribe una operación (+,-,*,/): ")

	//	Recibir input del usuario
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// Guardamos el texto que nos escribe
	return scanner.Text()
}

func main() {
	operacion := LeerEntrada()
	// Obtener el operador
	operator := GetTypeOperation(operacion)

	// Divide la cadena de texto, con un valor como patron de separación
	valores := strings.Split(operacion, operator)

	// Creamos un objeto de nuestro struct
	calculadora := Calc{}
	calculadora.value1 = GetValue(valores[0])
	calculadora.value2 = GetValue(valores[1])
	calculadora.operator = operator

	// Realizamos la operación
	calculadora.GetOperation()

}
