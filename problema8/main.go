package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
//completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		//simular trabajo
		time.Sleep(500 * time.Millisecond)
		ch <- x * x
	}()
	return ch
}
func fanIn(canales ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(canales))

	for _, c := range canales {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func main() {
	//crea varios futuros y recolecta sus resultados:
	f1 := asyncCuadrado(2)
	f2 := asyncCuadrado(3)
	f3 := asyncCuadrado(4)

	// Opción 1:
	fmt.Println("Esperando futuros secuencialmente:")
	fmt.Println("f1 =", <-f1)
	fmt.Println("f2 =", <-f2)
	fmt.Println("f3 =", <-f3)
	fmt.Println()
	f1 = asyncCuadrado(5)
	f2 = asyncCuadrado(6)
	f3 = asyncCuadrado(7)

	// (combinar múltiples canales)
	fmt.Println("Fan-in(todos juntos):")
	for v := range fanIn(f1, f2, f3) {
		fmt.Println("resultado:", v)
	}
	fmt.Println("Fin.")
	// Pista: crea una función fanIn que recibe múltiples <-chan int y retorna un único <-chan int
	// que emita todos los valores. Requiere goroutines y cerrar el canal de salida cuando todas terminen.

}
