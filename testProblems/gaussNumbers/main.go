package main

import (
	"fmt"
)

const errType = "Expected gInt or myInt, got other type instead"

type myInt int

type gInt struct {
	Re myInt
	Im myInt
}

type GaussNumber interface {
	Add(other GaussNumber) GaussNumber
	Sub(other GaussNumber) GaussNumber
	Mul(other GaussNumber) GaussNumber
	Div(other GaussNumber) GaussNumber
	Rem(other GaussNumber) GaussNumber
	Norm() myInt
}

func Complex(number GaussNumber) complex128 {
	switch number := number.(type) {
	case myInt:
		return complex(float64(number), 0)
	case gInt:
		return complex(float64(number.Re), float64(number.Im))
	default:
		panic(errType)
	}
}

func Round(number complex128) GaussNumber {
	a := real(number)
	b := imag(number)
	var (
		c, d myInt = 0, 0
	)
	if a > 0 {
		c = myInt(a + 0.5)
	} else {
		c = myInt(a - 0.5)
	}
	if b > 0 {
		d = myInt(b + 0.5)
	} else {
		d = myInt(b - 0.5)
	}
	return gInt{c, d}.reduce()

}

func Opposite(number GaussNumber) GaussNumber {
	switch number := number.(type) {
	case myInt:
		return -number
	case gInt:
		return gInt{-number.Re, -number.Im}
	default:
		panic(errType)
	}

}

func (a gInt) reduce() GaussNumber {
	if a.Im == 0 {
		return a.Re
	}
	return a
}

func (number myInt) Add(otherNumber GaussNumber) GaussNumber {
	switch otherNumber := otherNumber.(type) {
	case myInt:
		return number + otherNumber
	case gInt:
		return gInt{number + otherNumber.Re, otherNumber.Im}
	default:
		panic(errType)
	}
}

func (number gInt) Add(otherNumber GaussNumber) GaussNumber {
	switch otherNumber := otherNumber.(type) {
	case myInt:
		return otherNumber.Add(number)
	case gInt:
		return gInt{number.Re + otherNumber.Re, number.Im + otherNumber.Im}.reduce()
	default:
		panic(errType)
	}
}

func (number myInt) Sub(otherNumber GaussNumber) GaussNumber {
	switch otherNumber := otherNumber.(type) {
	case myInt:
		return number - otherNumber
	case gInt:
		return gInt{number - otherNumber.Re, otherNumber.Im}
	default:
		panic(errType)
	}
}

func (number gInt) Sub(otherNumber GaussNumber) GaussNumber {
	switch otherNumber := otherNumber.(type) {
	case myInt:
		return Opposite(otherNumber.Sub(number))
	case gInt:
		return gInt{number.Re - otherNumber.Re, number.Im - otherNumber.Im}.reduce()
	default:
		panic(errType)
	}
}

func (number myInt) Mul(otherNumber GaussNumber) GaussNumber {
	switch otherNumber := otherNumber.(type) {
	case myInt:
		return number * otherNumber
	case gInt:
		return gInt{number * otherNumber.Re, number * otherNumber.Im}.reduce()
	default:
		panic(errType)
	}
}

func (number gInt) Mul(otherNumber GaussNumber) GaussNumber {
	switch otherNumber := otherNumber.(type) {
	case myInt:
		return otherNumber.Mul(number)
	case gInt:
		return gInt{number.Re*otherNumber.Re - number.Im*otherNumber.Im, number.Im*otherNumber.Re + number.Re*otherNumber.Im}.reduce()
	default:
		panic(errType)
	}
}

func (number myInt) Div(otherNumber GaussNumber) GaussNumber {
	switch otherNumber := otherNumber.(type) {
	case myInt:
		return number / otherNumber
	case gInt:
		return Round(Complex(number) / Complex(otherNumber))
	default:
		panic(errType)
	}
}

func (number gInt) Div(otherNumber GaussNumber) GaussNumber {
	return Round(Complex(number) / Complex(otherNumber))
}

func (number gInt) Rem(otherNumber GaussNumber) GaussNumber {
	return number.Sub((number.Div(otherNumber)).Mul(otherNumber))
}

func (number myInt) Rem(otherNumber GaussNumber) GaussNumber {
	switch otherNumber := otherNumber.(type) {
	case myInt:
		return number % otherNumber
	case gInt:
		return number.Sub((number.Div(otherNumber)).Mul(otherNumber))
	default:
		panic(errType)
	}
}

func (number myInt) Norm() myInt {
	return number.Mul(number).(myInt)
}

func (number gInt) Norm() myInt {
	return number.Mul(gInt{number.Re, -number.Im}).(myInt)
}

func GCD(a, b GaussNumber) GaussNumber {
	if a == myInt(0) {
		return b
	}
	if b == myInt(0) {
		return a
	}
	if a.Norm() < b.Norm() {
		return GCD(a, b.Rem(a))
	} else {
		return GCD(b, a.Rem(b))
	}
}

func (number gInt) String() string {
	if number.Re == 0 {
		return fmt.Sprintf("%di", number.Im)
	}
	return fmt.Sprintf("%d + %di", number.Re, number.Im)
}

func Gauss(ints ...int) GaussNumber {
	if len(ints) == 0 {
		return myInt(0)
	}
	if len(ints) == 1 {
		return myInt(ints[0])
	}
	if len(ints) > 2 {
		panic("Gauss() needs less than 3 arguments")
	}
	return gInt{myInt(ints[0]), myInt(ints[1])}.reduce()
}

func main() {
	a := Gauss(-1, 3)
	b := Gauss(4)
	fmt.Println(GCD(a, b))
}
