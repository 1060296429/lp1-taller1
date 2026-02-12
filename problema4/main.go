package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Objetivo: Implementar Productor–Consumidor con canales.
// Un productor genera N valores y los envía por un canal; varios consumidores los procesan.
// Practicar cierre de canal y uso de WaitGroup.

//completa los pasos marcados.

func productor(n int, out chan<- int) {
	defer close(out) // cerrar el canal cuando no haya más datos
	for i := 1; i <= n; i++ {
		v := rand.Intn(100)
		fmt.Printf("[productor] envía %d\n", v)
		out <- v
		//dormir un poco para ver el flujo
		time.Sleep(time.Duration(rand.Intn(400)+100) * time.Millisecond)
	}
}

func consumidor(id int, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in { // leer hasta que canal se cierre

		fmt.Printf("[consumidor %d] recibe %d\n", id, v)
		//simular trabajo
		// usa Sleep con un valor aleatorio entre 100 y 500 ms
		time.Sleep(time.Duration(rand.Intn(400)+v*50) * time.Millisecond)
	}
	fmt.Printf("[consumidor %d] canal cerrado, termina\n", id)
	wg.Done()
}

func main() {

	valores := 10
	consumidores := 30
	ch := make(chan int, 4)
	var wg sync.WaitGroup
	wg.Add(consumidores)
	//lanzar consumidores
	for i := 1; i <= consumidores; i++ {
		go consumidor(i, ch, &wg)
		wg.Add(1)
	}

	go productor(valores, ch)
	//esperar consumidores
	wg.Wait()
	fmt.Println("Listo: todos los consumidores terminaron.")
}
