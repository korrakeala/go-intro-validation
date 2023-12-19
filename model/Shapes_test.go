package model

import (
	"io"
	"math"
	"os"
	"strings"
	"testing"
)

func TestEllipseArea(t *testing.T) {
	ellipse := CreateEllipse(3.0, 4.0)
	expectedArea := math.Pi * 3.0 * 4.0
	actualArea := ellipse.Area()

	if actualArea != expectedArea {
		t.Errorf("Expected area: %f, got: %f", expectedArea, actualArea)
	}
}

func TestRectangleArea(t *testing.T) {
	rect := CreateRectangle(5.0, 8.0)
	expectedArea := 5.0 * 8.0
	actualArea := rect.Area()

	if actualArea != expectedArea {
		t.Errorf("Expected area: %f, got: %f", expectedArea, actualArea)
	}
}

func TestTriangleArea(t *testing.T) {
	tri := CreateTriangle(6.0, 10.0)
	expectedArea := (6.0 * 10.0) / 2
	actualArea := tri.Area()

	if actualArea != expectedArea {
		t.Errorf("Expected area: %f, got: %f", expectedArea, actualArea)
	}
}

func TestPrintEllipse(t *testing.T) {
	ellipse := CreateEllipse(3.0, 4.0)
	expectedOutput := "Ellipse:\n\tId: " + ellipse.Id.String() + "\n\tRadius A: 3.000000\n\tRadius B: 4.000000\n\tArea: 37.699112\n"
	actualOutput := captureOutput(func() {
		ellipse.Print()
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}
}

func TestPrintRectangle(t *testing.T) {
	rect := CreateRectangle(5.0, 8.0)
	expectedOutput := "Rectangle:\n\tId: " + rect.Id.String() + "\n\tWidth: 8.000000\n\tHeight: 5.000000\n\tArea: 40.000000\n"
	actualOutput := captureOutput(func() {
		rect.Print()
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}
}

func TestPrintTriangle(t *testing.T) {
	tri := CreateTriangle(6.0, 10.0)
	expectedOutput := "Triangle:\n\tId: " + tri.Id.String() + "\n\tBase: 6.000000\n\tHeight: 10.000000\n\tArea: 30.000000\n"
	actualOutput := captureOutput(func() {
		tri.Print()
	})

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}
}

func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	err := w.Close()
	if err != nil {
		return err.Error()
	}
	os.Stdout = oldStdout

	var buf strings.Builder
	_, e := io.Copy(&buf, r)
	if e != nil {
		return e.Error()
	}

	return buf.String()
}
