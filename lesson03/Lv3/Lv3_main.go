package main

import "fmt"

func main() {
	fmt.Printf("Please input the radius of a circle:")
	var cIn float64
	fmt.Scan(&cIn)

	fmt.Printf("Please input the length and width of the rectangle:")
	var rInl, rInw float64
	fmt.Scan(&rInl, &rInw)

	c := Circle{
		radius: cIn,
	}

	r := Rectangle{
		length: rInl,
		width:  rInw,
	}

	cOut := c.Area()
	rOut := r.Area()

	fmt.Printf("The area of circle is: %f, The area of rectangle is: %f\n", cOut, rOut)

}
