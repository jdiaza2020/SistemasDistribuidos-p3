package main

import (
	"fmt"
	"time"
)

// CrearCoches genera num coches con tipos de incidencia aleatorios
// (mecánica, eléctrica, carrocería) para que haya mezcla de prioridades.
func CrearCoches(num int) []*Coche {
	coches := make([]*Coche, 0, num)
	idInc := 1

	for i := 0; i < num; i++ {
		// Elegimos tipo aleatorio entre las tres categorias.
		var tipo string
		r := rng.Intn(3)
		switch r {
		case 0:
			tipo = TipoMecanica
		case 1:
			tipo = TipoElectrica
		default:
			tipo = TipoCarroceria
		}

		inc := NuevaIncidencia(idInc, tipo, "")
		idInc++

		matricula := fmt.Sprintf("CAR%03d", i+1)
		coche := NuevoCoche(matricula, inc)
		coches = append(coches, coche)
	}

	return coches
}

// EjecutarEscenarioRW crea el taller, genera los coches y lanza la simulación
// usando la versión con RWMutex. Mide el tiempo total y calcula el throughput.
func EjecutarEscenarioRW(numCoches, numPlazas, numMecanicos int) {
	fmt.Println("======================================")
	fmt.Println(" Simulación con RWMutex")
	fmt.Printf(" Coches=%d  Plazas=%d  Mecanicos=%d\n", numCoches, numPlazas, numMecanicos)
	fmt.Println("======================================")

	IniciarTiempoPrograma()

	taller := NuevoTaller(numCoches, numPlazas, numMecanicos)
	coches := CrearCoches(numCoches)

	inicio := time.Now()
	SimularTallerRW(taller, coches)
	fin := time.Now()

	duracion := fin.Sub(inicio)
	fmt.Printf("Tiempo total de simulacion: %v\n", duracion)

	segundos := duracion.Seconds()
	if segundos > 0 {
		throughput := float64(numCoches) / segundos
		fmt.Printf("Throughput aproximado: %.2f coches/segundo\n", throughput)
	}
}

// EjecutarEscenarioWG crea el taller, genera los coches y lanza la simulación
// usando la versión con WaitGroup. Mide el tiempo total y calcula el throughput.
func EjecutarEscenarioWG(numCoches, numPlazas, numMecanicos int) {
	fmt.Println("======================================")
	fmt.Println(" Simulación con WaitGroup")
	fmt.Printf(" Coches=%d  Plazas=%d  Mecanicos=%d\n", numCoches, numPlazas, numMecanicos)
	fmt.Println("======================================")

	IniciarTiempoPrograma()

	taller := NuevoTaller(numCoches, numPlazas, numMecanicos)
	coches := CrearCoches(numCoches)

	inicio := time.Now()
	SimularTallerWG(taller, coches)
	fin := time.Now()

	duracion := fin.Sub(inicio)
	fmt.Printf("Tiempo total de simulacion: %v\n", duracion)

	segundos := duracion.Seconds()
	if segundos > 0 {
		throughput := float64(numCoches) / segundos
		fmt.Printf("Throughput aproximado: %.2f coches/segundo\n", throughput)
	}
}
