package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := TemperatureConverter(25, 'C')
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Printf("Resultado: %.2f F, %.2f K\n", result[0], result[1])

	result, err = TemperatureConverter(25, 'F')
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Printf("Resultado: %.2f C, %.2f K\n", result[0], result[1])

	result, err = TemperatureConverter(25, 'k')
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Printf("Resultado: %.2f C, %.2f F\n", result[0], result[1])
}

func TemperatureConverter(value float64, temperatureType rune) ([2]float64, error) {
	switch temperatureType {
	case 'C', 'c':
		fmt.Println("Celsius para Fahrenheit e Kelvin")
		fah := (value * (9.0 / 5.0)) + 32.0
		kel := value + 273.15
		return [2]float64{fah, kel}, nil

	case 'F', 'f':
		fmt.Println("Fahrenheit para Celsius e Kelvin")
		cel := (value - 32.0) * (5.0 / 9.0)
		kel := cel + 273.15
		return [2]float64{cel, kel}, nil

	case 'K', 'k':
		fmt.Println("Kelvin para Celsius e Fahrenheit")
		cel := value - 273.15
		fah := (cel * (9.0 / 5.0)) + 32.0
		return [2]float64{cel, fah}, nil

	default:
		return [2]float64{}, errors.New("tipo de temperatura inválido (use C, F ou K)")
	}
}
