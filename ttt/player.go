package ttt

import (
	"fmt"
	"math/rand"
)

type Player interface {
	Name() string
	Play(*Board) (int, error)
	Marker() Marker
}

type HumanPlayer struct {
	Nickname string
	Mark     Marker
}

func (player *HumanPlayer) Name() string {
	return player.Nickname
}

func (h *HumanPlayer) Marker() Marker {
	return h.Mark
}

func printOptions(b *Board) {
	for c, m := range b.cells {
		if m == BLANK {
			fmt.Print(c+1, " ")
		} else {
			fmt.Print("  ")
		}

		if c%b.Cols == b.Cols-1 {
			fmt.Println()
		}
	}
}

func (player *HumanPlayer) Play(board *Board) (int, error) {
	fmt.Printf("Your turn, %v! Pick a spot to place your '%v'...\n", player.Nickname, player.Mark)
	printOptions(board)

	var cell int
	var err error = fmt.Errorf("")
	for err != nil {
		fmt.Print("enter a choice: ")
		_, err = fmt.Scanln(&cell)
		if err != nil {
			fmt.Println("Didn't understand your choice.")
			continue
		}
		_, err = board.Get(cell - 1)
		if err != nil {
			fmt.Println("Bad choice:", err)
		}
	}

	return cell - 1, err
}

type RandomAI struct {
	Mark Marker
}

func (r *RandomAI) Name() string {
	return "Random Robot"
}

func (r *RandomAI) Marker() Marker {
	return r.Mark
}

func (r *RandomAI) Play(b *Board) (int, error) {
	choices := b.Find(BLANK)
	cell := choices[rand.Int()%len(choices)]
	return cell, nil
}
