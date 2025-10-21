package main

import "fmt"

func main() {
	var c float32
	fmt.Printf("Please input a celsius value:")
	fmt.Scan(&c)

	cIn := Fahrenheit{
		valC: c,
	}
	cIn.ToFahrenheit(c)
	fmt.Printf("After changing into fahrenheit: %f\n", cIn.valC)

	var f float32
	fmt.Printf("Please input a fahrenheit value:")
	fmt.Scan(&f)

	fIn := Celsius{
		valF: f,
	}
	fIn.ToCelsius(f)
	fmt.Printf("After changing into fahrenheit: %f\n", fIn.valF)

	fmt.Printf("AfterFah: %v  AfterCel: %v", cIn, fIn)

}
