package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func Printstr(Str string) {
	arrayStr := []rune(Str)

	for i := range arrayStr {
		z01.PrintRune(arrayStr[i])
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func Printint(a int) {
	r := '0'
	if a/10 > 0 {
		Printint(a / 10)
	}
	for i := 0; i < a%10; i++ {
		r++
	}
	z01.PrintRune(r)
}

func main() {
	var points point

	setPoint(&points)

	Printstr("x = ")
	Printint(points.x)
	Printstr(", y = ")
	Printint(points.y)
	z01.PrintRune('\n')
}
