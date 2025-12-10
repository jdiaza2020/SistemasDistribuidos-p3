package main

import (
	"math/rand"
	"time"
)

// Generador de números aleatorios propio para toda la simulación.
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// TiempoConVariacion devuelve un tiempo en torno al valor base (en segundos).
// Por ejemplo, base=5 -> devuelve 4, 5 o 6 (mínimo 1).
func TiempoConVariacion(base int) time.Duration {
	delta := rng.Intn(3) - 1 // -1, 0, +1
	valor := base + delta
	if valor < 1 {
		valor = 1
	}
	return time.Duration(valor) * time.Second
}
