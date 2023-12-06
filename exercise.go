package main

import (
	"fmt"
	"math"
	"sync"
)

func squared(val float64, i int, s *[]float64, wg *sync.WaitGroup) {
	defer wg.Done()
	newVal := math.Pow(float64(val), 2)
	(*s)[i] = newVal
}

func newSlice(s []float64) []float64 {
	var wg sync.WaitGroup

	ns := make([]float64, len(s), cap(s))

	for i, v := range s {
		wg.Add(1)
		go squared(v, i, &ns, &wg)
	}
	wg.Wait()
	return ns
}

func main() {
	list := []float64{1, 2, 3, 4, 5, 6, 8.3}
	fmt.Println(list, newSlice(list))
}
