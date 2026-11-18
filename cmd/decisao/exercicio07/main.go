package main

import "fmt"

func main() {
	var horasTrabalhadas, salarioHora, salarioTotal, horasNormais float64
	horasExtras := 0.0
	fmt.Print("Salário-hora: ")
	_, _ = fmt.Scan(&salarioTotal)
	fmt.Print("Horas trabalhadas: ")
	_, _ = fmt.Scan(&horasTrabalhadas)
	if horasTrabalhadas > 40 {
		horasExtras = horasTrabalhadas - 40
		horasNormais = 40
	} else {
		horasNormais = horasTrabalhadas
	}
	salarioTotal = horasNormais*salarioHora + horasExtras*salarioHora*1.5
	fmt.Printf("Salário total: %1.2f \n", salarioTotal)
}
