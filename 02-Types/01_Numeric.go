package main

import "fmt"

func main() {

	x := 12
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y := 90.23245
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%.2f", y)

	// %f     default width, default precision
	// %9f    width 9, default precision
	// %.2f   default width, precision 2
	// %9.2f  width 9, precision 2
	// %9.f   width 9, precision 0

	// this is doesnt work -> var z int8 = 128

	//rule of thumb: just use int
	//rule of thumb: just use float64

	// uint8 		-> 	the set of all unsigned  8-bit integers (0 to 255)
	// int8  		-> 	the set of all signed  8-bit integers (-128 to 127)
	//float64		->	the set of all IEEE-754 64-bit floating-point numbers
	//complex128	-> 	the set of all complex numbers with float64 real and imaginary parts
	//byte        	->	alias for uint8
	//rune        	->	alias for int32

	// for more -> https://golang.org/ref/spec#Numeric_types

	//for fmt.printf:

	//integre:
	// %b	base 2
	// %c	the character represented by the corresponding Unicode code point
	// %d	base 10
	// %o	base 8
	// %O	base 8 with 0o prefix
	// %q	a single-quoted character literal safely escaped with Go syntax.
	// %x	base 16, with lower-case letters for a-f
	// %X	base 16, with upper-case letters for A-F
	// %U	Unicode format: U+1234; same as "U+%04X"

	//Floating-point and complex constituents:
	// %b	decimalless scientific notation with exponent a power of two,
	// in the manner of strconv.FormatFloat with the 'b' format,
	// e.g. -123456p-78
	// %e	scientific notation, e.g. -1.234456e+78
	// %E	scientific notation, e.g. -1.234456E+78
	// %f	decimal point but no exponent, e.g. 123.456
	// %F	synonym for %f
	// %g	%e for large exponents, %f otherwise. Precision is discussed below.
	// %G	%E for large exponents, %F otherwise
	// %x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
	// %X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

	// for more -> https://pkg.go.dev/fmt
}
