package clases

import interfaces "github.com/Billones142/Protoype/Interfaces"

type ReporteMedico struct {
	Paciente         string
	Titulo           string
	Sintomas         string
	Cuerpo           string
	Formato          string
	numeroDelReporte int
}

// Clonar crea y devuelve un clon del reporte medico.
func (t *ReporteMedico) Clonar() interfaces.Reporte {
	return &ReporteMedico{
		Paciente:         t.Paciente,
		Titulo:           t.Titulo,
		Sintomas:         t.Sintomas,
		Cuerpo:           t.Cuerpo,
		Formato:          t.Formato,
		numeroDelReporte: t.numeroDelReporte,
	}
}
