// scheduler_rwmutex.go
package main

import (
	"sync"
	"time"
)

// SchedulerRW encapsula la lógica de planificación usando RWMutex
// para proteger las colas de las fases.
type SchedulerRW struct {
	Taller     *Taller
	MutexFases [NumFases]sync.RWMutex // un RWMutex por fase
	wg         sync.WaitGroup         // para esperar a que todos los coches terminen
}

// SimularTallerRW lanza la simulación del taller usando RWMutex.
//
// Parámetros:
//   - t: taller ya inicializado con NuevoTaller.
//   - coches: slice con todos los coches que queremos procesar.
func SimularTallerRW(t *Taller, coches []*Coche) {
	s := &SchedulerRW{
		Taller: t,
	}

	// Todos los coches empiezan en la fase 1 (EsperaPlaza).
	s.enviarCochesAFaseInicial(coches)

	// Añadimos al WaitGroup tantos coches como tengamos.
	s.wg.Add(len(coches))

	// Lanzamos los workers de cada fase.
	for fase := 0; fase < NumFases; fase++ {
		capacidad := t.Fases[fase].Capacidad
		for i := 0; i < capacidad; i++ {
			go s.workerFase(fase)
		}
	}

	// Esperamos a que todos los coches hayan pasado las 4 fases.
	s.wg.Wait()
}

// enviarCochesAFaseInicial encola todos los coches en la cola de la fase de espera de plaza.
func (s *SchedulerRW) enviarCochesAFaseInicial(coches []*Coche) {
	faseInicial := s.Taller.Fases[FaseEsperaPlaza]

	s.MutexFases[FaseEsperaPlaza].Lock()
	defer s.MutexFases[FaseEsperaPlaza].Unlock()

	for _, c := range coches {
		if c == nil {
			continue
		}
		c.FaseActual = FaseEsperaPlaza
		EncolarCochePrioridad(faseInicial, c)
	}
}

// workerFase representa a un "trabajador" de una fase concreta del taller.
// Cada worker:
//  1. Saca el siguiente coche de la cola de su fase.
//  2. Simula el tiempo de trabajo según la prioridad.
//  3. Envía el coche a la siguiente fase o marca como terminado si está en la última.
func (s *SchedulerRW) workerFase(indFase int) {
	f := s.Taller.Fases[indFase]

	for {
		// 1) Sacar siguiente coche de la cola de esta fase.
		s.MutexFases[indFase].Lock()
		var coche *Coche
		if len(f.Cola) > 0 {
			coche = SacarSiguienteCoche(f)
		}
		s.MutexFases[indFase].Unlock()

		// Si no hay coche en este momento, esperamos un poco y volvemos a intentar.
		if coche == nil {
			time.Sleep(10 * time.Millisecond)
			continue
		}

		// 2) Simular el tiempo de trabajo en esta fase.
		// El tiempo por fase viene dado por la incidencia (5, 3 o 1 segundos).
		tiempo := coche.Incidencia.TiempoPorFase
		time.Sleep(time.Duration(tiempo) * time.Second)

		// 3) Pasar a la siguiente fase o marcar como terminado.
		if indFase == FaseEntrega {
			// El coche ha terminado todas las fases.
			s.wg.Done()
		} else {
			// Encolamos el coche en la siguiente fase, manteniendo la prioridad.
			siguiente := indFase + 1
			coche.FaseActual = siguiente

			s.MutexFases[siguiente].Lock()
			EncolarCochePrioridad(s.Taller.Fases[siguiente], coche)
			s.MutexFases[siguiente].Unlock()
		}
	}
}
