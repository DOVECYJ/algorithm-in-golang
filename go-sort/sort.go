//十大排序算法的golang实现
package sort

import (
	"go-heap/heap"
	"strconv"
)

//1.冒泡排序
func BubbleSort(input []int) []int {
	n := len(input)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}

//2.选择排序
func SelectSort(input []int) []int {
	n, min := len(input), 0
	for i := 0; i < n; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if input[j] < input[min] {
				min = j
			}
		}
		input[i], input[min] = input[min], input[i]
	}
	return input
}

//3.插入排序
func InsertSort(input []int) []int {
	n, cur, index := len(input), 0, 0
	for i := 1; i < n; i++ {
		cur, index = input[i], i-1
		for index >= 0 && input[index] > cur {
			input[index+1] = input[index]
			index--
		}
		input[index+1] = cur
	}
	return input
}

//4.希尔排序
func ShellSort(input []int) []int {
	n := len(input)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			cur, j := input[i], i-gap
			for j >= 0 && input[j] > cur {
				input[j+gap] = input[j]
				j -= gap
			}
			input[j+gap] = cur
		}
	}
	return input
}

//5.归并排序
func MergeSort(input []int) []int {
	n := len(input)
	if n < 2 {
		return input
	}
	mid := n / 2
	left, right := input[0:mid], input[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (output []int) {
	for i, j := 0, 0; i < len(left) || j < len(right); {
		if i >= len(left) {
			output = append(output, right[j])
			j++
		} else if j >= len(right) {
			output = append(output, left[i])
			i++
		} else if left[i] < right[j] {
			output = append(output, left[i])
			i++
		} else {
			output = append(output, right[j])
			j++
		}
	}
	return
}

//6.快速排序
func QuickSort(input []int, left, right int) []int {
	if left < right {
		index := left + 1
		for i := index; i <= right; i++ {
			if input[i] < input[left] {
				input[i], input[index] = input[index], input[i]
				index++
			}
		}
		input[left], input[index-1] = input[index-1], input[left]

		QuickSort(input, left, index-2)
		QuickSort(input, index, right)
	}
	return input
}

//7.堆排序
func HeapSort(input []int) (output []int) {
	h := heap.NewMinHeap(input...)
	for !h.Empty() {
		if data, err := h.Get(); err == nil {
			output = append(output, data)
		}
	}
	return
}

//8.计数排序
func CountSort(input []int) (output []int) {
	max := input[0]
	for i := range input {
		if input[i] > max {
			max = input[i]
		}
	}
	bucket := make([]int, max+1, max+1)
	for _, v := range input {
		bucket[v]++
	}
	for i := range bucket {
		for bucket[i] > 0 {
			output = append(output, i)
			bucket[i]--
		}
	}
	return
}

//9.桶排序
func BucketSort(input []int) (output []int) {
	min, max, size := input[0], input[0], len(input)/2
	for _, v := range input {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	gap := (max-min)/size + 1
	bucket := make([][]int, size, size)
	for _, v := range input {
		index := (v - min) / gap
		bucket[index] = append(bucket[index], v)
	}
	for i := 0; i < size; i++ {
		if len(bucket[i]) > 0 {
			bucket[i] = InsertSort(bucket[i])
		}
	}
	for i := 0; i < size; i++ {
		output = append(output, bucket[i]...)
	}
	return
}

//10.基数排序
func RadixSort(input []int) []int {
	max := input[0]
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	bits := len(strconv.Itoa(max))
	bucket := make([][]int, 10, 10)
	for i, mod := 0, 1; i < bits; i, mod = i+1, mod*10 {
		for _, v := range input {
			index := v / mod % 10
			bucket[index] = append(bucket[index], v)
		}
		//fmt.Printf("[%d]  %v\n", i, bucket)
		pos := 0
		for j := 0; j < 10; j++ {
			for len(bucket[j]) > 0 {
				input[pos] = bucket[j][0]
				bucket[j] = bucket[j][1:]
				pos++
			}
		}
		//fmt.Printf("[%d]  %v\n", i, input)
	}
	return input
}
