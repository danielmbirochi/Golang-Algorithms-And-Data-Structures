package main

func IsValidSubsequence(array []int, sequence []int) bool {
	var si int
	var subsequence []int
	for _, m := range array {
		if m == sequence[si] {
			subsequence = append(subsequence, sequence[si])
			si++
		}
		if si == len(sequence) {
			return true
		}
	}

	return len(subsequence) > 0 &&
		len(sequence) <= len(array) &&
		len(subsequence) == len(sequence)
}
