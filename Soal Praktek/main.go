package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	id    int
	dice  []int
	score int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var numPlayers, numDice int

	// Input jumlah pemain dan dadu
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&numPlayers)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&numDice)

	// Inisialisasi pemain
	players := make([]*Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		players[i] = &Player{id: i + 1, dice: rollDice(numDice), score: 0}
	}

	round := 1
	for countActivePlayers(players) > 1 {
		fmt.Printf("==================\nGiliran %d lempar dadu:\n", round)
		for _, player := range players {
			fmt.Printf("Pemain #%d (%d): ", player.id, player.score)
			printDice(player.dice)
		}

		// Evaluasi dadu setiap pemain
		for i, player := range players {
			if len(player.dice) > 0 { // Pemain hanya mengevaluasi jika masih memiliki dadu
				evaluateDice(player, players, i)
			}
		}

		fmt.Println("Setelah evaluasi:")
		for _, player := range players {
			fmt.Printf("Pemain #%d (%d): ", player.id, player.score)
			printDice(player.dice)
		}

		round++
	}

	fmt.Println("==================")
	declareWinner(players)
}

func rollDice(num int) []int {
	dice := make([]int, num)
	for i := 0; i < num; i++ {
		dice[i] = rand.Intn(6) + 1 // Menghasilkan angka 1 sampai 6
	}
	return dice
}

func printDice(dice []int) {
	if len(dice) == 0 {
		fmt.Print("_")
	} else {
		for i, d := range dice {
			if i != 0 {
				fmt.Print(",")
			}
			fmt.Print(d)
		}
	}
	fmt.Println()
}

func evaluateDice(player *Player, players []*Player, playerIndex int) {
	newDice := []int{}
	for _, d := range player.dice {
		switch d {
		case 6:
			player.score++
		case 1:
			// Berikan dadu 1 ke pemain sebelah jika masih ada pemain lain
			if len(players) > 1 {
				nextPlayer := players[(playerIndex+1)%len(players)]
				if nextPlayer != player { // pastikan tidak memberikan ke diri sendiri
					nextPlayer.dice = append(nextPlayer.dice, 1)
				}
			}
		default:
			newDice = append(newDice, d)
		}
	}
	player.dice = newDice
}

func countActivePlayers(players []*Player) int {
	active := 0
	for _, player := range players {
		if len(player.dice) > 0 {
			active++
		}
	}
	return active
}

func declareWinner(players []*Player) {
	maxScore := -1
	winner := -1
	for _, player := range players {
		if player.score > maxScore {
			maxScore = player.score
			winner = player.id
		}
	}

	fmt.Printf("Game dimenangkan oleh pemain #%d dengan skor %d.\n", winner, maxScore)
}
