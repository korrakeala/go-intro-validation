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
	reader = bufio.NewReader(os.Stdin)

	for {
		printMenu()
		option := readLine()

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
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
	fmt.Println("Byee")
}

func printMenu() {
	fmt.Println("<<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>>")
	fmt.Println("Seleccione la acción a relizar:")
	fmt.Println("a. Listar figuras")
	fmt.Println("b. Obtener figura por su ID")
	fmt.Println("c. Agregar nueva figura")
	fmt.Println("Presione \"q\" para salir.")
	fmt.Println("<<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>><<>>")
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
		printShapeMenu()
		option := readLine()

		if quit(option) {
			return
		}

		shapeType, ok := shapeTypeMap[option]
		if !ok {
			fmt.Println("Opción no válida. Intente de nuevo.")
			continue
		}

		dimensions := getShapeDimensions(shapeType)
		if dimensions == nil {
			return
		}

		shape := createAndAddShape(shapeType, dimensions)
		shape.Print()

	}

}

func printShapeMenu() {
	fmt.Println("Seleccione el tipo de figura, o presione \"q\" para volver:")
	fmt.Println("a. Elipse")
	fmt.Println("b. Rectángulo")
	fmt.Println("c. Triángulo")
}

func getShapeDimensions(shapeType string) []string {
	var dimensions []string

	for _, dimension := range shapeDimensions[shapeType] {
		fmt.Printf("Ingrese %s:", dimension)
		value := readLine()

		if quit(value) {
			return nil
		}

		dimensions = append(dimensions, value)
	}

	return dimensions
}

func createAndAddShape(shapeType string, dimensions []string) model.Shape {
	switch shapeType {
	case "elipse":
		radiusA, radiusB := dimensions[0], dimensions[1]
		ellipse := model.CreateEllipse(toFloat(radiusA), toFloat(radiusB))
		shapes[ellipse.Id.String()] = ellipse
		return ellipse
	case "rectangulo":
		height, width := dimensions[0], dimensions[1]
		rectangle := model.CreateRectangle(toFloat(height), toFloat(width))
		shapes[rectangle.Id.String()] = rectangle
		return rectangle
	case "triangulo":
		base, height := dimensions[0], dimensions[1]
		triangle := model.CreateTriangle(toFloat(base), toFloat(height))
		shapes[triangle.Id.String()] = triangle
		return triangle
	default:
		return nil
	}
}

func toFloat(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal("Error al convertir la entrada a número:", err)
	}
	return val
}

func getShapeById() {
	fmt.Println("Ingrese Id:")
	id := readLine()

	shape, exists := shapes[id]
	if !exists {
		fmt.Println("No se encontró una figura con ese Id.")
		return
	}

	shape.Print()
}

func listShapes() {
	if len(shapes) == 0 {
		fmt.Println("La lista está vacía. Agregue figuras y vuelva a intentar.")
		return
	}

	for _, shape := range shapes {
		shape.Print()
	}
}

var shapeTypeMap = map[string]string{
	"a": "elipse",
	"b": "rectangulo",
	"c": "triangulo",
}

var shapeDimensions = map[string][]string{
	"elipse":     {"radio A", "radio B"},
	"rectangulo": {"alto", "ancho"},
	"triangulo":  {"base", "altura"},
}
