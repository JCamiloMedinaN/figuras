package figuras

import "fmt"

type Medida interface {
	area() float64
	perimetro() float64
}

func Medidas(medida Medida) {
	fmt.Println("Medidas: ", medida)
	fmt.Println("Area: ", medida.area())
	fmt.Println("Perimetro: ", medida.perimetro())
}
