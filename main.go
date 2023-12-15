package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func main() {
	var option string
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("<<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>>")
		fmt.Println("Seleccione la acción a relizar:")
		fmt.Println("a. listar figuras")
		fmt.Println("b. obtener figura por su ID")
		fmt.Println("c. agregar nueva figura")
		fmt.Println("<<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>>")

		option = readLine()

		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "listar":
			listShapes()
		case "b", "obtener":
			getShapeById()
		case "c", "agregar":
			createShape()
		}
	}

	fmt.Println("Byee")

}

func readLine() string {
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No se pudo obtener el valor")
	} else {
		return strings.TrimSuffix(option, "\n")
	}
}

func createShape() {
	fmt.Println("seleccione el tipo de figura:")
	fmt.Println("ingrese valor a:")
	fmt.Println("ingrese valor b:")
}

func getShapeById() {
	fmt.Println("ingrese ID:")
}

func listShapes() {
	fmt.Println("listado de figuras")
}
