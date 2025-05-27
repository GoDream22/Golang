package main

import "fmt"

func main() {
	fmt.Println("___Конвертер валют___")
	const usdToEur = 0.8955
	const usdToRub = 80.77
	const eurTorub = usdToRub / usdToEur
	getUserInput()
}

func getUserInput() string {
	var inputValue string
	fmt.Print("Введите операцию, которую хотите выполнить: ")
	fmt.Scan(&inputValue)
	return inputValue
}

func calculateConvert(float64, string, string) float64 {
	var sumToConvert float64
	var currency1 string
	var currency2 string
}
