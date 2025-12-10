# Jorge Díaz Alcojor

# SistemasDistribuidos-p3 – Taller concurrente con RWMutex y WaitGroup en Go

En esta práctica se modela un **taller de reparación de coches** con 4 fases secuenciales  
(espera de plaza, mecánico, limpieza y entrega) y tres categorías de incidencias:

* A – Mecánica: prioridad alta, 5 s por fase.
* B – Eléctrica: prioridad media, 3 s por fase.
* C – Carrocería: prioridad baja, 1 s por fase.

Cada coche pasa por las 4 fases y el sistema debe respetar la prioridad y la capacidad máxima de cada fase.  
La simulación se ha implementado con **dos métodos de sincronización** del paquete `sync`:

* Una versión con **`RWMutex`** (pipeline con colas por fase).
* Una versión con **`WaitGroup`** (procesamiento por bloques).

Se comparan ambas variantes en varios escenarios de carga y se generan trazas de ejecución con el formato exigido.

---

## Estructura del proyecto

El proyecto está dividido en varios ficheros:

### `model.go`

Contiene el **modelo de datos** del taller:

* Constantes de fases, prioridades y tipos de incidencia.
* Structs:
  * `Incidencia`
  * `Coche`
  * `Fase`
  * `Mecanico`
  * `Taller`
* Funciones de creación:
  * `NuevaIncidencia(...)`
  * `NuevoCoche(...)`
  * `NuevoTaller(...)`
* Cálculo de plazas ocupadas / libres.

Es la capa que define cómo es el taller a nivel lógico.

---

### `queue.go`

Implementa la **cola de prioridad** por fase:

* `EncolarCochePrioridad(...)`  
  Inserta coches en la cola según la prioridad (0 alta, 1 media, 2 baja).
* `SacarSiguienteCoche(...)`  
  Extrae el siguiente coche a ser atendido.
* Funciones auxiliares para saber si la cola está vacía o su tamaño.

---

### `logs.go`

Gestiona el **tiempo global de la simulación** y las trazas:

* `IniciarTiempoPrograma()` inicializa el instante cero para cada escenario.
* `LogEvento(...)` imprime líneas del tipo:

  `Tiempo {t} Coche {N} Incidencia {Tipo} Fase {Fase} Estado {ENTRA/SALE}`

Se llama al entrar y salir de cada fase, para poder seguir el recorrido de cada coche.

---

### `tiempos.go`

Define la **variación de tiempo de uso** en cada fase:

* Generador aleatorio `rng`.
* `TiempoConVariacion(base int)` devuelve `base ± 1` segundo (mínimo 1 s).

Así cada coche no tarda exactamente lo mismo en una fase, cumpliendo el enunciado.

---

### `scheduler_rwmutex.go`

Implementa la simulación con **`RWMutex`**:

* Struct `SchedulerRW` con:
  * Referencia al `Taller`.
  * Un `RWMutex` por fase.
  * Un `WaitGroup` interno para saber cuándo han terminado todos los coches.
* `SimularTallerRW(...)`:
  * Encola los coches en la fase de espera.
  * Lanza varios workers por fase (según `Capacidad`).
  * Cada worker:
    * Saca coches de su cola con prioridad.
    * Simula el tiempo de trabajo con `TiempoConVariacion`.
    * Llama a `LogEvento` (ENTRA/SALE).
    * Pasa el coche a la siguiente fase o marca que ha terminado.

Representa el taller como una **pipeline** de 4 fases concurrentes.

---

### `scheduler_waitgroup.go`

Implementa la simulación con **`WaitGroup`**:

* Función `SimularTallerWG(...)`:
  * Recorre las 4 fases secuencialmente.
  * En cada fase ordena los coches por prioridad.
  * Los procesa en bloques del tamaño `Capacidad` con goroutines y un `WaitGroup`.
  * Dentro de cada goroutine se simula el tiempo de la fase y se loguea ENTRA/SALE.

Es un enfoque más sencillo, basado en “oleadas” por fase, que sirve para comparar rendimiento con la versión RWMutex.

---

### `simulation.go`

Contiene la **lógica de alto nivel de la simulación**:

* `CrearCoches(num int)`  
  Genera coches con matrículas `CAR001`, `CAR002`, … y tipo de incidencia aleatorio (A/B/C).
* `EjecutarEscenarioRW(...)`  
  Crea un taller, genera coches, lanza `SimularTallerRW` y mide:
  * Tiempo total de simulación.
  * Throughput aproximado (coches/segundo).
* `EjecutarEscenarioWG(...)`  
  Hace lo mismo pero usando `SimularTallerWG`.

---

### `simulation_test.go`

Fichero de **tests automáticos** usando el paquete `testing`:

* Genera tres escenarios con las cantidades de coches por categoría que pide el enunciado:
  * Test 1: A=10, B=10, C=10.
  * Test 2: A=20, B=5, C=5.
  * Test 3: A=5, B=5, C=20.
* Para cada test lanza la simulación con RWMutex y con WaitGroup.
* Comprueba que **todos los coches terminan en la fase de Entrega**.

Todos los tests pasan, lo que indica que ambas variantes procesan correctamente los coches en las 4 fases.

---

### `main.go`

Punto de entrada de la práctica:

* Define los tres escenarios globales:
  * 10/10/10
  * 20/5/5
  * 5/5/20
* Para cada uno ejecuta:
  * `EjecutarEscenarioRW(...)`
  * `EjecutarEscenarioWG(...)`

Al ejecutarlo se ven en consola las trazas de los coches, los tiempos totales y el throughput de cada método.

---

### `go.mod`

Archivo de módulo de Go. Indica el nombre del módulo y la versión de Go.  
Permite compilar y ejecutar el proyecto con `go build` y `go run` sin problemas de dependencias.

---

## Cómo compilar y ejecutar la práctica

Desde la carpeta raíz del proyecto (donde está `go.mod`):

### Ejecutar la simulación

```bash
go run .
