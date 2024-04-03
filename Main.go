package main

import (
	"fmt"

	clases "github.com/Billones142/Protoype/Clases"
)

func imprimirReporteMedico(reporteMedico *clases.ReporteMedico) {
	fmt.Println("Formato: " + reporteMedico.Formato)
	fmt.Println("Paciente: " + reporteMedico.Paciente)
	fmt.Println("Sintomas: " + reporteMedico.Sintomas)
	fmt.Println("Cuerpo: " + reporteMedico.Cuerpo)

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
	// Se crea un clon del objeto para aprovechar la configuracion anterior ya que en lo
	// unico que difieren es en el nombre del paciente ahorrandose todo el trabajo de crear uno casi igual.
	// Ademas de que las variables privadas no se podian conocer desde afuera asi que el mismo objeto es el que
	// se tenia que encargar de su clonacion

	fmt.Println("Reporte Original:")
	imprimirReporteMedico(reporteOriginal)

	fmt.Println("Nuevo reporte Original:")
	imprimirReporteMedico(nuevoReporte)
}
