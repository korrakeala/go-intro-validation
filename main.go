package main

import (
	"bufio"
	"fmt"
	"go-intro-validation/model"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader
var shapes = make(map[string]model.Shape)

func main() {
	var option string
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("<<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>>")
		fmt.Println("Seleccione la acción a relizar:")
		fmt.Println("a. Listar figuras")
		fmt.Println("b. Obtener figura por su ID")
		fmt.Println("c. Agregar nueva figura")
		fmt.Println("Presione \"q\" para salir.")
		fmt.Println("<<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>>")

		option = readLine()
		if quit(option) {
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

func quit(option string) bool {
	return option == "quit" || option == "q"
}

func readLine() string {
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No se pudo obtener el valor")
	} else {
		return strings.TrimSuffix(option, "\n")
	}
}

func createShape() {
	for {
		fmt.Println("Seleccione el tipo de figura, o presione \"q\" para volver:")
		fmt.Println("a. Elipse")
		fmt.Println("b. Rectángulo")
		fmt.Println("c. Triángulo")

		var option string = readLine()
		if quit(option) {
			return
		}

		switch option {
		case "a", "elipse":
			var radiusA, radiusB string
			fmt.Println("ingrese radio a:")
			radiusA = readLine()
			if quit(radiusA) {
				return
			}

			fmt.Println("ingrese radio b:")
			radiusB = readLine()
			if quit(radiusB) {
				return
			}

			ellipse := model.CreateEllipse(toFloat(radiusA), toFloat(radiusB))
			ellipse.Print()
			shapes[ellipse.Id.String()] = ellipse
		case "b", "rectangulo":
			var height, width string
			fmt.Println("ingrese alto:")
			height = readLine()
			if quit(height) {
				return
			}

			fmt.Println("ingrese ancho:")
			width = readLine()
			if quit(width) {
				return
			}

			rectangle := model.CreateRectangle(toFloat(height), toFloat(width))
			rectangle.Print()
			shapes[rectangle.Id.String()] = rectangle
		case "c", "triangulo":
			var base, height string
			fmt.Println("ingrese base:")
			base = readLine()
			if quit(base) {
				return
			}

			fmt.Println("ingrese altura:")
			height = readLine()
			if quit(height) {
				return
			}

			triangle := model.CreateTriangle(toFloat(base), toFloat(height))
			triangle.Print()
			shapes[triangle.Id.String()] = triangle
		}
	}

}

func toFloat(str string) float64 {
	radiusAfloat, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return radiusAfloat
}

func getShapeById() {
	fmt.Println("ingrese Id:")
}

func listShapes() {
	if len(shapes) == 0 {
		fmt.Println("La lista está vacía. Agregue figuras y vuelva a intentar.")
		return
	} else {
		for _, shape := range shapes {
			shape.Print()
		}
	}
}
