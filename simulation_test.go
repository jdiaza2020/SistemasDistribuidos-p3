package main

import (
	"fmt"
	"testing"
)

const (
	numPlazasTest    = 5
	numMecanicosTest = 5
)

// crearCochesCategorias crea coches con un número fijo por categoría:
//
//	a coches de tipo mecánica (A), b eléctricos (B) y c de carrocería (C).
func crearCochesCategorias(a, b, c int) []*Coche {
	coches := make([]*Coche, 0, a+b+c)
	idInc := 1
	num := 1

	// Categoria A: mecánica
	for i := 0; i < a; i++ {
		inc := NuevaIncidencia(idInc, TipoMecanica, "")
		idInc++
		matricula := printfMatricula(num)
		num++
		coches = append(coches, NuevoCoche(matricula, inc))
	}

	// Categoria B: eléctrica
	for i := 0; i < b; i++ {
		inc := NuevaIncidencia(idInc, TipoElectrica, "")
		idInc++
		matricula := printfMatricula(num)
		num++
		coches = append(coches, NuevoCoche(matricula, inc))
	}

	// Categoria C: carrocería
	for i := 0; i < c; i++ {
		inc := NuevaIncidencia(idInc, TipoCarroceria, "")
		idInc++
		matricula := printfMatricula(num)
		num++
		coches = append(coches, NuevoCoche(matricula, inc))
	}

	return coches
}

// printfMatricula genera matriculas del estilo CAR001, CAR002, ...
func printfMatricula(n int) string {
	return fmt.Sprintf("CAR%03d", n)
}

// comprueba que todos los coches han terminado en la fase de Entrega.
func comprobarFaseFinal(t *testing.T, coches []*Coche, metodo string) {
	for _, c := range coches {
		if c.FaseActual != FaseEntrega {
			t.Fatalf("metodo %s: coche %s terminó en fase %d, se esperaba %d (Entrega)",
				metodo, c.Matricula, c.FaseActual, FaseEntrega)
		}
	}
}

// Tests de los tres escenarios de la tabla

// Test 1: Categoria A:10, B:10, C:10.
func TestEscenario1_RWMutex_y_WaitGroup(t *testing.T) {
	a, b, c := 10, 10, 10

	// RWMutex
	cochesRW := crearCochesCategorias(a, b, c)
	tallerRW := NuevoTaller(len(cochesRW), numPlazasTest, numMecanicosTest)
	IniciarTiempoPrograma()
	SimularTallerRW(tallerRW, cochesRW)
	comprobarFaseFinal(t, cochesRW, "RWMutex")

	// WaitGroup
	cochesWG := crearCochesCategorias(a, b, c)
	tallerWG := NuevoTaller(len(cochesWG), numPlazasTest, numMecanicosTest)
	IniciarTiempoPrograma()
	SimularTallerWG(tallerWG, cochesWG)
	comprobarFaseFinal(t, cochesWG, "WaitGroup")
}

// Test 2: Categoria A:20, B:5, C:5.
func TestEscenario2_RWMutex_y_WaitGroup(t *testing.T) {
	a, b, c := 20, 5, 5

	cochesRW := crearCochesCategorias(a, b, c)
	tallerRW := NuevoTaller(len(cochesRW), numPlazasTest, numMecanicosTest)
	IniciarTiempoPrograma()
	SimularTallerRW(tallerRW, cochesRW)
	comprobarFaseFinal(t, cochesRW, "RWMutex")

	cochesWG := crearCochesCategorias(a, b, c)
	tallerWG := NuevoTaller(len(cochesWG), numPlazasTest, numMecanicosTest)
	IniciarTiempoPrograma()
	SimularTallerWG(tallerWG, cochesWG)
	comprobarFaseFinal(t, cochesWG, "WaitGroup")
}

// Test 3: Categoria A:5, B:5, C:20.
func TestEscenario3_RWMutex_y_WaitGroup(t *testing.T) {
	a, b, c := 5, 5, 20

	cochesRW := crearCochesCategorias(a, b, c)
	tallerRW := NuevoTaller(len(cochesRW), numPlazasTest, numMecanicosTest)
	IniciarTiempoPrograma()
	SimularTallerRW(tallerRW, cochesRW)
	comprobarFaseFinal(t, cochesRW, "RWMutex")

	cochesWG := crearCochesCategorias(a, b, c)
	tallerWG := NuevoTaller(len(cochesWG), numPlazasTest, numMecanicosTest)
	IniciarTiempoPrograma()
	SimularTallerWG(tallerWG, cochesWG)
	comprobarFaseFinal(t, cochesWG, "WaitGroup")
}
