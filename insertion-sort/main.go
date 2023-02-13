package main

func InsertionSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		chave := a[j]
		i := j - 1
		for i > -1 && a[i] > chave {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = chave
	}
	return a
}
