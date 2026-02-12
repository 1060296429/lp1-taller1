package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Implementar una versión del problema de los Filósofos Comensales.
// Hay 5 filósofos y 5 tenedores (recursos). Cada filósofo necesita 2 tenedores para comer.
// Estrategia segura: imponer un **orden global** al tomar los tenedores (primero el menor ID, luego el mayor)
// para evitar deadlock. También puedes limitar concurrencia (ej. mayordomo).
//completa la lógica de toma/soltado de tenedores y bucle de pensar/comer.

type tenedor struct {
	mu sync.Mutex
}

func filosofo(id int, izq, der *tenedor, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		//pensar
		pensar(id)
		//orden global
		//tomar primero el de menor direccion de memoria
		primero := izq
		segundo := der
		if fmt.Sprintf("%p", der) < fmt.Sprintf("%p", izq) {
			primero = der
			segundo = izq
		}
		primero.mu.Lock()
		segundo.mu.Lock()
		comer(id)
		//soltar tenedores
		segundo.mu.Unlock()
		primero.mu.Unlock()
	}
	//desarrolla el código para el filósofo {
	fmt.Printf("[filósofo %d] satisfecho\n", id)
}

func pensar(id int) {
	fmt.Printf("[filósofo %d] pensando...\n", id)
	time.Sleep(300 * time.Millisecond)
	//simular tiempo de pensar

}

func comer(id int) {
	fmt.Printf("[filósofo %d] COMIENDO\n", id)
	time.Sleep(300 * time.Millisecond)
	//simular tiempo de comer

}

func main() {
	const n = 5
	var wg sync.WaitGroup
	wg.Add(n)

	// crear tenedores
	tenedores := make([]*tenedor, n)
	for i := 0; i < n; i++ {
		tenedores[i] = &tenedor{}
		//inicializar cada tenedor i

	}

	// crear filosofos (cada uno con su izq y der)
	for i := 0; i < n; i++ {
		izq := tenedores[i]
		der := tenedores[(i+1)%n]
		go filosofo(i, izq, der, &wg)

	}

	wg.Wait()
	fmt.Println("Todos los filósofos han comido sin deadlock.")
}
