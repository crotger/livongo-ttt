package main

import (
	"fmt"
	"time"

	"./ttt"
)

func main() {
	b := ttt.DefaultBoard()
	fmt.Println("Welcome! What's your name?")
	var name string
	fmt.Print("Your name: ")
	fmt.Scanln(&name)

	p := &ttt.HumanPlayer{name, ttt.MARK_X}
	players := []ttt.Player{&ttt.RandomAI{ttt.MARK_O}, p}

	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Welcome, %v, the computer will go first.\n", p.Name())

	gameOver := false
	currentPlayer := 0
	for !gameOver {
		time.Sleep(500 * time.Millisecond)
		cp := players[currentPlayer]
		fmt.Println(cp.Name(), "'s turn.")
		cell, err := cp.Play(b)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(500 * time.Millisecond)
		b.Set(cell, cp.Marker())
		fmt.Println(cp.Name(), "placed an", cp.Marker(), "at", cell+1)
		fmt.Println(b)

		vic := b.CheckForWinner(cell)
		if vic != nil {
			gameOver = true
			fmt.Printf("%v wins with 3 %v's from %v to %v!\n", cp.Name(), vic.Mark, vic.Start+1, vic.End+1)
		}

		if b.IsFull() {
			gameOver = true
			fmt.Println("The board is full. No winner!")
		}

		currentPlayer = (currentPlayer + 1) % 2
	}
}
