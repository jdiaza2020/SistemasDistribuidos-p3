package main

// Constantes para las fases del taller.
const (
	FaseEsperaPlaza = 0 // Fase 1: espera de plaza + documentación
	FaseMecanico    = 1 // Fase 2: espera de mecánico + reparación
	FaseLimpieza    = 2 // Fase 3: limpieza
	FaseEntrega     = 3 // Fase 4: revisión final y entrega
	NumFases        = 4
)

// Constantes de prioridad (cuanto más pequeño, más prioridad).
const (
	PrioridadAlta  = 0 // Categoria A: mecánica
	PrioridadMedia = 1 // Categoria B: eléctrica
	PrioridadBaja  = 2 // Categoria C: carrocería
)

// Constantes para los tipos de incidencia.
const (
	TipoMecanica   = "Mecanica"
	TipoElectrica  = "Electrica"
	TipoCarroceria = "Carroceria"
)

// Incidencia asociada a un coche.
// De aquí sacamos la prioridad y el tiempo por fase.
type Incidencia struct {
	IDIncidencia  int
	Tipo          string // "Mecanica", "Electrica" o "Carroceria"
	Prioridad     int    // 0 = alta, 1 = media, 2 = baja
	TiempoPorFase int    // en segundos: 5, 3 o 1 según la categoría
	Descripcion   string
}

// Coche representa un vehículo que pasa por las 4 fases del taller.
type Coche struct {
	Matricula  string
	Incidencia *Incidencia
	FaseActual int // 0..3 según la fase en la que esté
}

// Fase representa una etapa del taller con una capacidad máxima y una cola.
type Fase struct {
	Nombre    string
	Capacidad int
	Cola      []*Coche // cola de prioridad por Incidencia.Prioridad
}

// Mecanico: en esta práctica todos pueden hacer cualquier tipo de reparación.
type Mecanico struct {
	IDMecanico int
	Ocupado    bool
}

// Taller representa el sistema completo y la configuración principal.
type Taller struct {
	NumCoches    int
	NumPlazas    int
	NumMecanicos int

	Fases     [NumFases]*Fase
	Mecanicos []*Mecanico
}

// Funciones auxiliares

// NuevaIncidencia crea una incidencia asignando prioridad y tiempo según el tipo.
// Categoria A: mecánica   -> prioridad alta  -> 5 s por fase
// Categoria B: eléctrica  -> prioridad media -> 3 s por fase
// Categoria C: carrocería -> prioridad baja  -> 1 s por fase
func NuevaIncidencia(id int, tipo string, descripcion string) *Incidencia {
	inc := &Incidencia{
		IDIncidencia: id,
		Tipo:         tipo,
		Descripcion:  descripcion,
	}

	switch tipo {
	case TipoMecanica:
		inc.Prioridad = PrioridadAlta
		inc.TiempoPorFase = 5
	case TipoElectrica:
		inc.Prioridad = PrioridadMedia
		inc.TiempoPorFase = 3
	case TipoCarroceria:
		inc.Prioridad = PrioridadBaja
		inc.TiempoPorFase = 1
	default:
		// Por defecto consideramos carrocería.
		inc.Prioridad = PrioridadBaja
		inc.TiempoPorFase = 1
	}

	return inc
}

// NuevoCoche crea un coche con matrícula e incidencia asociada.
// El coche siempre empieza en la primera fase (espera de plaza).
func NuevoCoche(matricula string, inc *Incidencia) *Coche {
	return &Coche{
		Matricula:  matricula,
		Incidencia: inc,
		FaseActual: FaseEsperaPlaza,
	}
}

// NuevoTaller crea un taller con los parámetros principales e inicializa
// las 4 fases y la lista de mecánicos.
//
// En prácticas anteriores NumPlazas solía ser 2 * NumMecanicos;
// aquí lo dejamos como parámetro para poder probar distintos escenarios.
func NuevoTaller(numCoches, numPlazas, numMecanicos int) *Taller {
	t := &Taller{
		NumCoches:    numCoches,
		NumPlazas:    numPlazas,
		NumMecanicos: numMecanicos,
	}

	// Inicializamos las fases.
	t.Fases[FaseEsperaPlaza] = &Fase{
		Nombre:    "EsperaPlaza",
		Capacidad: numPlazas, // plazas físicas del taller
		Cola:      []*Coche{},
	}

	t.Fases[FaseMecanico] = &Fase{
		Nombre:    "Mecanico",
		Capacidad: numMecanicos, // mecánicos disponibles
		Cola:      []*Coche{},
	}

	t.Fases[FaseLimpieza] = &Fase{
		Nombre:    "Limpieza",
		Capacidad: numMecanicos, // se puede ajustar si se quiere otra política
		Cola:      []*Coche{},
	}

	t.Fases[FaseEntrega] = &Fase{
		Nombre:    "Entrega",
		Capacidad: numMecanicos,
		Cola:      []*Coche{},
	}

	// Inicializamos la lista de mecánicos.
	t.Mecanicos = make([]*Mecanico, numMecanicos)
	for i := 0; i < numMecanicos; i++ {
		t.Mecanicos[i] = &Mecanico{
			IDMecanico: i + 1,
			Ocupado:    false,
		}
	}

	return t
}

// PlazasOcupadas devuelve cuántos coches están actualmente en la cola de la fase 1.
func (t *Taller) PlazasOcupadas() int {
	fase := t.Fases[FaseEsperaPlaza]
	return len(fase.Cola)
}

// PlazasLibres devuelve cuántas plazas libres quedan en la fase 1.
func (t *Taller) PlazasLibres() int {
	ocupadas := t.PlazasOcupadas()
	libres := t.NumPlazas - ocupadas
	if libres < 0 {
		libres = 0
	}
	return libres
}
