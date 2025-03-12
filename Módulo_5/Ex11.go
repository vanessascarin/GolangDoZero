/*
https://www.codewars.com/kata/5672a98bdbdd995fad00000f

Let's play! You have to return which player won! In case of a draw return Draw!.

Examples(Input1, Input2 --> Output):

"scissors", "paper" --> "Player 1 won!"
"scissors", "rock" --> "Player 2 won!"
"paper", "paper" --> "Draw!"
*/

package main

import (
	"fmt"
)

func main() {

	teste1a := "rock"
	teste1b := "scissors"
	teste2a := "scissors"
	teste2b := "rock"
	teste3a := "rock"
	teste3b := "rock"
	fmt.Println(Rps(teste1a, teste1b))
	fmt.Println(Rps(teste2a, teste2b))
	fmt.Println(Rps(teste3a, teste3b))

}

/*
func Rps(p1, p2 string) string {
	if p1 == p2 {
		return ("Draw!")
	}
	switch p1 {
	case "rock":
		switch p2 {
		case "paper":
			return ("Player 2 won!")
		case "scissors":
			return ("Player 1 won!")
		}
	case "paper":
		switch p2 {
		case "rock":
			return ("Player 1 won!")
		case "scissors":
			return ("Player 2 won!")
		}
	case "scissors":
		switch p2 {
		case "paper":
			return ("Player 1 won!")
		case "rock":
			return ("Player 2 won!")
		}
	}
	return "Entrada inválida."
}

ALTERNATIVA:*/

func Rps(p1, p2 string) string {
	if p1 == p2 {
		return "Draw!"
	}
	if (p1 == "rock" && p2 == "scissors") || (p1 == "scissors" && p2 == "paper") || (p1 == "paper" && p2 == "rock") {
		return "Player 1 won!"
	} else {
		return "Player 2 won!"
	}
}
