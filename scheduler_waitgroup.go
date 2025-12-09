// scheduler_waitgroup.go
package main

import (
	"sync"
	"time"
)

// ordenarPorPrioridad ordena los coches por prioridad de la incidencia:
// menor número = más prioridad (0 alta, 1 media, 2 baja).
func ordenarPorPrioridad(coches []*Coche) {
	n := len(coches)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if coches[j].Incidencia.Prioridad < coches[min].Incidencia.Prioridad {
				min = j
			}
		}
		if min != i {
			coches[i], coches[min] = coches[min], coches[i]
		}
	}
}

// SimularTallerWG lanza la simulación del taller usando WaitGroup.
// En cada fase se procesan los coches por orden de prioridad y en bloques
// de tamaño "Capacidad" para respetar la capacidad de la fase.
func SimularTallerWG(t *Taller, coches []*Coche) {
	if len(coches) == 0 {
		return
	}

	for fase := 0; fase < NumFases; fase++ {
		faseTaller := t.Fases[fase]
		if faseTaller == nil {
			continue
		}

		capacidad := faseTaller.Capacidad
		if capacidad <= 0 {
			capacidad = 1
		}

		// Ordenamos por prioridad antes de procesar esta fase.
		ordenarPorPrioridad(coches)

		// Procesamos en bloques de tamaño "capacidad".
		for i := 0; i < len(coches); i += capacidad {
			fin := i + capacidad
			if fin > len(coches) {
				fin = len(coches)
			}

			var wg sync.WaitGroup
			wg.Add(fin - i)

			// Lanzamos hasta "capacidad" goroutines en este bloque.
			for j := i; j < fin; j++ {
				c := coches[j]

				go func(c *Coche, faseActual int) {
					// Simulamos el trabajo de esta fase.
					tiempo := c.Incidencia.TiempoPorFase
					time.Sleep(time.Duration(tiempo) * time.Second)

					// Actualizamos la fase en la que se encuentra el coche.
					c.FaseActual = faseActual

					wg.Done()
				}(c, fase)
			}

			// Esperamos a que terminen todas las goroutines de este bloque.
			wg.Wait()
		}
	}
}
