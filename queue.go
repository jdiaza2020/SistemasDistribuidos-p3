// queue.go
package main

// EncolarCochePrioridad mete un coche en la cola de la fase
// respetando la prioridad de la incidencia.
// Prioridad más baja en número = más urgente (0 alta, 1 media, 2 baja).
func EncolarCochePrioridad(f *Fase, c *Coche) {
	if f == nil || c == nil || c.Incidencia == nil {
		return
	}

	prio := c.Incidencia.Prioridad
	insertado := false

	// Recorremos la cola buscando la primera posición con prioridad mayor.
	for i, otro := range f.Cola {
		if otro.Incidencia == nil {
			continue
		}
		if prio < otro.Incidencia.Prioridad {
			// Insertamos en medio: [0..i-1] + c + [i..]
			f.Cola = append(f.Cola, nil)   // hacemos hueco
			copy(f.Cola[i+1:], f.Cola[i:]) // desplazamos a la derecha
			f.Cola[i] = c
			insertado = true
			break
		}
	}

	// Si no hemos encontrado ningún sitio (cola vacía o misma prioridad al final), añadimos al final.
	if !insertado {
		f.Cola = append(f.Cola, c)
	}
}

// SacarSiguienteCoche extrae el coche que debe ser atendido a continuación
// según la prioridad (el que está al principio de la cola).
// Devuelve nil si la cola está vacía.
func SacarSiguienteCoche(f *Fase) *Coche {
	if f == nil {
		return nil
	}
	if len(f.Cola) == 0 {
		return nil
	}

	c := f.Cola[0]
	// Desplazamos la cola hacia la izquierda.
	f.Cola = f.Cola[1:]
	return c
}

// ColaVacia indica si la cola de la fase está vacía.
func ColaVacia(f *Fase) bool {
	if f == nil {
		return true
	}
	return len(f.Cola) == 0
}

// TamCola devuelve el número de coches en la cola de la fase.
func TamCola(f *Fase) int {
	if f == nil {
		return 0
	}
	return len(f.Cola)
}
