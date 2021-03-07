package utils

import "sort"

func BubbleSort(elements []int) {
	for i := 0; i < len(elements); i++ {
		for j := i + 1; j < len(elements); j++ {
			if elements[i] > elements[j] {
				elements[i], elements[j] = elements[j], elements[i]
			}
		}
	}
}

func Sort(elements []int) {
	if len(elements) < 1000 {
		BubbleSort(elements)
		return
	}

	sort.Ints(elements)
}
