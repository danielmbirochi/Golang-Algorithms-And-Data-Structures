package main

func SortedSquaredArray(array []int) []int {
	result := make([]int, len(array))
	for i, v := range array {
		result[i] = v * v
	}
	for j := 1; j < len(result); j++ {
		chave := result[j]
		i := j - 1
		for i > -1 && result[i] > chave {
			result[i+1] = result[i]
			i = i - 1
		}
		result[i+1] = chave
	}
	return result
}
