package main

import "fmt"

type player struct {
	name   string
	atBats int
	hits   int
}

func calcAvg(player *player) float64 {
	if player.atBats == 0 {
		return 0.0
	}
	return float64(player.hits) / float64(player.atBats)
}

func main() {
	players := []player{
		{"Mark", 5, 10},
		{"Anatoly", 7, 12},
		{"Alex", 6, 3},
	}

	for _, player := range players {
		fmt.Printf("Player name: %s, avg: %f\n", player.name, calcAvg(&player))
	}
}
