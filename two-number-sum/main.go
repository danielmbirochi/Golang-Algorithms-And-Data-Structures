package main

func TwoNumberSum(array []int, target int) []int {
	result := make([]int, 0)

	for _, val := range array {
		for _, innerVal := range array {
			if val != innerVal && val+innerVal == target {
				result = append(result, val, innerVal)
				return result
			}
		}
	}

	return result
}
