package main

func NonConstructibleChange(coins []int) int {
	result := 1
	if len(coins) == 0 {
		return result
	}

	coins = InsertionSort(coins)

	var change int
	for i := 0; i < len(coins); i++ {
		if change+1 < coins[i] {
			return change + 1
		}
		change += coins[i]
	}

	return change + 1
}

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
