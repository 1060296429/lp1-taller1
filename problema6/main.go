package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Provocar un deadlock con dos mutex y dos goroutines que adquieren
// recursos en orden distinto. Luego evitarlo imponiendo un orden global.
// NOTA: La versión con deadlock se quedará bloqueada: ejecútala, observa y luego cambia a la versión segura.
//completa/activa la sección que quieras probar.

func sinDeadlock() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("G1: Lock mu1")
		mu1.Lock()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("G1: Lock mu2")
		mu2.Lock()
		fmt.Println("G1: listo")
		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		defer wg.Done()
		fmt.Println("G2: Lock mu1")
		mu1.Lock()
		time.Sleep(100 * time.Millisecond)

		fmt.Println("G2: Lock mu2")
		mu2.Lock()
		fmt.Println("G2: listo")
		mu1.Unlock()
		mu2.Unlock()
	}()

	// ADVERTENCIA: esto no retornará por el deadlock
	wg.Wait()
}
func main() {
	fmt.Println("=== Elige una sección para ejecutar ===")
	sinDeadlock()
	//comenta/activa la versión que desees probar
	// para probar la vercion correcta:
	//sinDeadlock()
	fmt.Println("fin del programa")
}
