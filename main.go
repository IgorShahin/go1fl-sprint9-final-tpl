package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE     = 100_000_000
	CHUNKS   = 8
	MaxValue = 1_000_000
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	nums := make([]int, size)
	for i := range size {
		nums[i] = rand.Intn(MaxValue) + 1
	}

	return nums
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]

	for i := 1; i < len(data); i++ {
		if data[i] > maxValue {
			maxValue = data[i]
		}
	}

	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunks := min(CHUNKS, len(data))

	maxValues := make([]int, chunks)
	chunkSize := len(data) / chunks

	var wg sync.WaitGroup

	for i := range chunks {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			start := index * chunkSize
			end := start + chunkSize

			if index == chunks-1 {
				end = len(data)
			}

			maxValues[index] = maximum(data[start:end])
		}(i)
	}

	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()

	max := maximum(data)

	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()

	max = maxChunks(data)

	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
