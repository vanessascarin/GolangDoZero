/*
https://www.codewars.com/kata/5bb904724c47249b10000131

Description:
Our football team has finished the championship.

Our team's match results are recorded in a
collection of strings. Each match is represented by
a string in the format "x:y", where x is our team's
score and y is our opponents score.

For example: ["3:1", "2:2", "0:1", ...]

Points are awarded for each match as follows:

if x > y: 3 points (win)
if x < y: 0 points (loss)
if x = y: 1 point (tie)
We need to write a function that takes this
collection and returns the number of points our
team (x) got in the championship by the rules given
above.

Notes:

our team always plays 10 matches in the championship
0 <= x <= 4
0 <= y <= 4

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	teste1 := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}
	teste2 := []string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}
	teste3 := []string{"0:1", "0:2", "0:3", "0:4", "1:2", "1:3", "1:4", "2:3", "2:4", "3:4"}
	teste4 := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "1:3", "1:4", "2:3", "2:4", "3:4"}
	teste5 := []string{"1:0", "2:0", "3:0", "4:4", "2:2", "3:3", "1:4", "2:3", "2:4", "3:4"}

	fmt.Println(Points(teste1))
	fmt.Println(Points(teste2))
	fmt.Println(Points(teste3))
	fmt.Println(Points(teste4))
	fmt.Println(Points(teste5))
}

func Points(games []string) int {
	var PontosX int
	for _, game := range games {
		placar := strings.Split(game, ":")
		placarX, _ := strconv.Atoi(placar[0])
		placarY, _ := strconv.Atoi(placar[1])

		if placarX > placarY {
			PontosX += 3
		} else if placarX == placarY {
			PontosX++
		}
	}
	return (PontosX)
}

/* TENTATIVA FALHA SEM STRING SPLIT:

func Points(games []string) int {
	var PontosX int
	for _, game := range games {
		placarX, _ := strconv.Atoi(game[:1])
		placarY, _ := strconv.Atoi(game[1:])
		if placarX > placarY {
			PontosX += 3
		} else if placarX == placarY {
			PontosX++
		}
	}
	return (PontosX)
}
*/
