// Задача: Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

//Алгоритм состоит из трёх шагов:

// 1. Выбрать элемент из массива. Назовём его опорным.
// 2. Разбиение: перераспределение элементов в массиве таким образом, что элементы, меньшие опорного, помещаются перед ним, а большие или равные - после.
// 3. Рекурсивно применить первые два шага к двум подмассивам слева и справа от опорного элемента. Рекурсия не применяется к массиву, в котором только один элемент или отсутствуют элементы.

package main

import "fmt"

func partition(array []int, low, high int) int {
	pivot := array[high]
	i := low

	for j := low; j < high; j++ {
		if array[j] < pivot {
			array[i], array[j] = array[j], array[i]
			i++
		}
	}
	array[i], array[high] = array[high], array[i]
	return i
}

func quicksort(array []int, low, high int) {
	if low < high {
		p := partition(array, low, high)
		quicksort(array, low, p-1)
		quicksort(array, p+1, high)
	}

}

func main() {
	arr := []int{5, 6, 7, 2, 1, 22, 12, 0}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
