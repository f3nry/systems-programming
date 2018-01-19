package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func swap(array []int, index1 int, index2 int) {
	swap := array[index1]
	array[index1] = array[index2]
	array[index2] = swap
}

func partition(array []int, low int, high int) int {
	if low < high {
		pivot := array[high]
		i := low

		for i < high {
			if array[i] < pivot {
				i++
			} else {
				swap(array, i, high-1)
				swap(array, high-1, high)
				high -= 1
			}
		}
	}

	return high
}

func recurQuickSort(array []int, low int, high int) {
	if low < high {
		p := partition(array, low, high)
		recurQuickSort(array, low, p-1)
		recurQuickSort(array, p+1, high)
	}
}

func quickSort(array []int) {
	recurQuickSort(array, 0, len(array)-1)
}

func makeRandomArray(desiredArrayLength int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	array := make([]int, desiredArrayLength)

	for index, _ := range array {
		array[index] = r.Intn(desiredArrayLength)
	}

	return array
}

func main() {
	desiredArrayLength, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	array := makeRandomArray(desiredArrayLength)

	fmt.Println("Initial Array", array)

	quickSort(array)

	fmt.Println("Sorted Array", array)
}
