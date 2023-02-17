package main

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	score := make(map[string]int)
	var winner string
	for i, round := range competitions {
		if HOME_TEAM_WON == results[i] {
			score[round[0]] = score[round[0]] + 3
		} else {
			score[round[1]] = score[round[1]] + 3
		}
	}

	var higherScore int
	for k, v := range score {
		if v > higherScore {
			higherScore = v
			winner = k
		}
	}

	return winner
}
