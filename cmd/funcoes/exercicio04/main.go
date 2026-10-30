package main

import "fmt"

func toCelcius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func toFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	var temp float64
	fmt.Print("Informe uma temperatura em celcius: ")
	fmt.Scan(&temp)
	f := toFahrenheit(temp)
	c := toCelcius(f)
	fmt.Printf("A temperatura %1.2f é %1.2f°F e %1.2f°C\n", temp, f, c)
}
