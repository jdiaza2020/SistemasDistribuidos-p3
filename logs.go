package main

import (
	"fmt"
	"time"
)

// inicioPrograma marca el inicio de cada simulacion (RW o WG).
var inicioPrograma time.Time

// IniciarTiempoPrograma se llama al comenzar un escenario
// para medir el tiempo relativo de los eventos.
func IniciarTiempoPrograma() {
	inicioPrograma = time.Now()
}

// LogEvento imprime la traza de un coche entrando/saliendo de una fase.
// Formato: Tiempo {t} Coche {N} Incidencia {Tipo} Fase {Fase_Actual} Estado {Estado_Fase}
func LogEvento(c *Coche, nombreFase string, estado string) {
	if inicioPrograma.IsZero() || c == nil || c.Incidencia == nil {
		return
	}

	t := time.Since(inicioPrograma).Seconds()

	fmt.Printf(
		"Tiempo %.3f Coche %s Incidencia %s Fase %s Estado %s\n",
		t,
		c.Matricula,
		c.Incidencia.Tipo,
		nombreFase,
		estado,
	)
}
