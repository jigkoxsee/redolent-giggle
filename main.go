package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Role struct {
	atk     int
	passive []string
}

type GameMap struct {
}

type Player struct {
	role     Role
	name     string
	hp       int
	focus    int
	location int
	buff     []string
}

func (p *Player) Move(step int) {
	fmt.Println(p.name, " @", p.location)
	p.location = p.location + step
	fmt.Println(p.name, " move to ", p.location)
}

type Board struct {
	players      []Player
	turn         int
	playerNumber int
}

func (b *Board) EndTurn() {
	b.turn++
	if b.turn >= b.playerNumber {
		b.turn = 0
	}
}

func (b *Board) Play(step int) {
	// if >0 number then walk
	// if <= 0 then do other thing - atk,blaa
	b.players[b.turn].Move(step)
}

func dice(rnd *rand.Rand) int {
	return rnd.Intn(6)
}

func main() {
	fmt.Println("Hello")
	rndsrc := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(rndsrc)

	// Initial
	ada := Player{name: "ada"}
	bob := Player{name: "bob"}
	board := Board{playerNumber: 2}
	board.players = make([]Player, 2)
	board.players[0] = ada
	board.players[1] = bob

	fmt.Println(board)

	// Start
	for {
		if board.players[0].location >= 100 || board.players[1].location >= 100 {
			break
		}
		board.Play(dice(rnd))
		board.EndTurn()
	}
}
