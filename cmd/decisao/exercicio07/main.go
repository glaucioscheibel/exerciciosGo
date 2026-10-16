package main

import "fmt"

func main() {
	var horasTrabalhadas float64
	var salarioHora float64
	var salarioTotal float64
	horasExtras := 0.0
	var horasNormais float64
	fmt.Print("Salário-hora: ")
	fmt.Scan(&salarioTotal)
	fmt.Print("Horas trabalhadas: ")
	fmt.Scan(&horasTrabalhadas)
	if horasTrabalhadas > 40 {
		horasExtras = horasTrabalhadas - 40
		horasNormais = 40
	} else {
		horasNormais = horasTrabalhadas
	}
	salarioTotal = horasNormais*salarioHora + horasExtras*salarioHora*1.5
	fmt.Printf("Salário total: %1.2f \n", salarioTotal)
}
