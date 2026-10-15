package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElementsNegative(t *testing.T) {
	sizes := []int{0, -1, -100}

	for _, size := range sizes {
		result := generateRandomElements(size)

		assert.NotNil(t, result)
		assert.Equal(t, 0, len(result))
	}
}

func TestGenerateRandomElementsOk(t *testing.T) {
	sizes := []int{1, 10, 100, 1000}

	for _, size := range sizes {
		result := generateRandomElements(size)

		assert.NotNil(t, result)
		assert.Equal(t, size, len(result))
	}
}

func TestGenerateRandomElementsRandomness(t *testing.T) {
	first := generateRandomElements(100)
	second := generateRandomElements(100)

	assert.Equal(t, len(first), len(second))
	assert.NotEqual(t, first, second)
}

func TestMaximumNegative(t *testing.T) {
	requests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", data: nil, want: 0},
		{name: "empty", data: []int{}, want: 0},
	}

	for _, v := range requests {
		result := maximum(v.data)

		assert.Equal(t, v.want, result)
	}
}

func TestMaximumOk(t *testing.T) {
	requests := []struct {
		data []int
		want int
	}{
		{data: []int{42}, want: 42},
		{data: []int{1, 2, 3, 4, 5}, want: 5},
		{data: []int{5, 4, 3, 2, 1}, want: 5},
		{data: []int{1, 99, 3, 7}, want: 99},
		{data: []int{-10, -5, -1}, want: -1},
		{data: []int{-10, 0, 5, -3, 20}, want: 20},
		{data: []int{math.MinInt, math.MaxInt}, want: math.MaxInt},
	}

	for _, v := range requests {
		result := maximum(v.data)

		assert.Equal(t, v.want, result)
	}
}

func TestMaximumDoesNotMutate(t *testing.T) {
	data := []int{3, 1, 4, 1, 5, 9, 2, 6}
	original := make([]int, len(data))
	copy(original, data)

	maximum(data)

	assert.Equal(t, original, data)
}

func TestMaxChunksNegative(t *testing.T) {
	requests := []struct {
		data []int
		want int
	}{
		{data: nil, want: 0},
		{data: []int{}, want: 0},
	}

	for _, v := range requests {
		result := maxChunks(v.data)

		assert.Equal(t, v.want, result)
	}
}

func TestMaxChunksOk(t *testing.T) {
	requests := []struct {
		data []int
		want int
	}{
		{data: []int{5, 3, 8, 1}, want: 8},
		{data: []int{1, 2, 3, 4, 1337, 6, 7, 8}, want: 1337},
		{data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 67, 228}, want: 228},
	}

	for _, v := range requests {
		result := maxChunks(v.data)

		assert.Equal(t, v.want, result)
	}
}

func TestMaxChunksTail(t *testing.T) {
	data := make([]int, 100)
	for i := range data {
		data[i] = i
	}
	data[99] = 999

	result := maxChunks(data)

	require.Equal(t, 999, result)
}

func TestMaxChunksEqualsMaximum(t *testing.T) {
	data := generateRandomElements(10_000)

	single := maximum(data)
	chunks := maxChunks(data)

	assert.Equal(t, single, chunks)
}
