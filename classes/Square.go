package classes

import "fmt"

type Square struct {
	Base float64
}

func (square *Square) GetSquare() float64 {
	return square.Base
}

func (square *Square) SetSquare(base float64) {
	square.Base = base
}

func (square Square) GetArea() float64 {
	return square.Base * square.Base
}

func (square Square) String() string {
	return fmt.Sprintf("this square has base: %f", square.Base)
}
