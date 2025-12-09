package main

func main() {
	// Escenarios oficiales:
	// 1) 10 coches, 10 plazas, 10 mecánicos
	// 2) 20 coches, 5 plazas, 5 mecánicos
	// 3) 5 coches, 5 plazas, 20 mecánicos

	escenarios := []struct {
		numCoches    int
		numPlazas    int
		numMecanicos int
	}{
		{10, 10, 10},
		{20, 5, 5},
		{5, 5, 20},
	}

	for _, e := range escenarios {
		EjecutarEscenarioRW(e.numCoches, e.numPlazas, e.numMecanicos)
		EjecutarEscenarioWG(e.numCoches, e.numPlazas, e.numMecanicos)
	}
}
