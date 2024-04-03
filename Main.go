package main

import (
	"fmt"

	clases "github.com/Billones142/Protoype/Clases"
	interfaces "github.com/Billones142/Protoype/Interfaces"
)

func imprimirReporteMedico(reporte *interfaces.Reporte) {
	fmt.Println("a")

	reporteMedico, ok := (*reporte).(*clases.ReporteMedico)
	if !ok {
		fmt.Println("no es un reporte medico")
		return
	}

	fmt.Println("Formato: " + reporteMedico.Formato)
	fmt.Println("Paciente: " + reporteMedico.Paciente)
	fmt.Println("Sintomas: " + reporteMedico.Sintomas)
	fmt.Println(reporteMedico.Cuerpo)

}

func main() {
	reporteOriginal := &clases.ReporteMedico{
		Paciente: "Lucas",
		Titulo:   "Informe Mensual",
		Cuerpo:   "Este es el cuerpo del informe medico",
		Sintomas: "Dengue",
		Formato:  "PDF",
	}

	nuevoReporte := (reporteOriginal.Clonar()).(*clases.ReporteMedico)

	nuevoReporte.Paciente = "Stiven"

	fmt.Println("Reporte Original:")
	//imprimirReporteMedico(nuevoReporte)

	fmt.Println("Nuevo reporte Original:")
	imprimirReporteMedico(&nuevoReporte.Reporte)
}
