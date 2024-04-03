package clases

import interfaces "github.com/Billones142/Protoype/Interfaces"

type ReporteMedico struct {
	interfaces.Reporte
	Paciente string
	Titulo   string
	Sintomas string
	Cuerpo   string
	Formato  string
}

// Clonar crea y devuelve un clon del reporte medico.
func (t *ReporteMedico) Clonar() interfaces.Reporte {
	return &ReporteMedico{
		Titulo:   t.Titulo,
		Cuerpo:   t.Cuerpo,
		Sintomas: t.Sintomas,
		Formato:  t.Formato,
	}
}
