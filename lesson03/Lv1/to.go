//华氏度 = 32 + 摄氏度 * 1.8

package main

type Fahrenheit struct {
	valC float32
}
type Celsius struct {
	valF float32
}

func (c *Fahrenheit) ToFahrenheit(cel float32) {
	cTof := 32.0 + cel*1.8
	c.valC = cTof
}

func (f *Celsius) ToCelsius(fah float32) {
	fToc := (fah - 32.0) / 1.8
	f.valF = fToc
}
