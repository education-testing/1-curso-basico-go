package classes

import "fmt"

type Rectangle struct {
	Base   float64
	Height float64
}

func (rectangle *Rectangle) GetBase() float64 {
	return rectangle.Base
}

func (rectangle *Rectangle) SetBase(base float64) {
	rectangle.Base = base
}

func (rectangle *Rectangle) GetHeight() float64 {
	return rectangle.Height
}

func (rectangle *Rectangle) SetHeight(height float64) {
	rectangle.Height = height
}

func (rectangle Rectangle) GetArea() float64 {
	return rectangle.Height * rectangle.Base
}

func (rectangle Rectangle) String() string {
	return fmt.Sprintf("this rectangle has base: %f, with height: %f", rectangle.Base, rectangle.Height)
}
