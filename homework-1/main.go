package main

import (
	"slices"
	"strconv"
)

func main() {

}

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(arr []int) string {
	var result string

	for _, num := range arr {
		result += strconv.Itoa(num)
	}

	return result
}

func MergeSlices(arrFloat []float32, arrInt []int32) []int {
	var result []int

	for _, num := range arrFloat {
		result = append(result, int(num))
	}

	for _, num := range arrInt {
		result = append(result, int(num))
	}

	return result
}

func GetMapValuesSortedByKey(dataMap map[int]string) []string {
	var result []string
	var keys []int

	for key := range dataMap {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	for _, key := range keys {
		result = append(result, dataMap[key])
	}

	return result
}
