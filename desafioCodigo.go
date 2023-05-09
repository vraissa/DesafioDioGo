package main

import "fmt"

func main() {

	var kelvin float64
	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scanf("%f", &kelvin)
	celsius := kelvin - 273.15
	fmt.Printf("%.2fK é igual a %.2f°C\n", kelvin, celsius)

}
