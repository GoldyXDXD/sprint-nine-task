package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	res := make([]int, size)
	for i := range size {
		res[i] = rand.Int()
	}
	return res
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) <= 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	max := data[0]
	for _, v := range data {
		if max < v {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	chunkSize := len(data) / CHUNKS
	localMaxes := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := range CHUNKS {
		start := chunkSize * i
		end := chunkSize + start
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			slice := data[start:end]
			localMaxes[i] = maximum(slice)
		}()
	}
	wg.Wait()

	return maximum(localMaxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
