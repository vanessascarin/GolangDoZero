/*
https://www.codewars.com/kata/57f222ce69e09c3630000212

For every good kata idea there seem to be quite a
few bad ones!

In this kata you need to check the provided array
(x) for good ideas 'good' and bad ideas 'bad'. If
there are one or two good ideas, return 'Publish!',
 if there are more than 2 return 'I smell a
 series!'. If there are no good ideas, as is often
 the case, return 'Fail!'.
*/

package main

import (
	"fmt"
)

func main() {

	teste1 := []string{"good", "bad", "good", "good", "bad", "good", "bad", "bad", "good", "bad", "bad"}
	teste2 := []string{"bad", "bad", "bad", "bad", "good", "good", "bad", "bad", "bad"}
	teste3 := []string{"bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "bad", "good", "bad", "bad", "bad"}
	teste4 := []string{"bad", "bad"}
	teste5 := []string{"bad", "bad", "bad", "bad", "bad"}
	teste6 := []string{"bad", "bad", "bad", "good", "bad", "good", "good", "bad", "bad", "bad", "bad", "good", "good"}
	fmt.Println(Well(teste1))
	fmt.Println(Well(teste2))
	fmt.Println(Well(teste3))
	fmt.Println(Well(teste4))
	fmt.Println(Well(teste5))
	fmt.Println(Well(teste6))
}

/*
1-2 good = publish
+ que 2 good = series
0 good = fail
*/

func Well(x []string) string {
	GoodCount := 0
	for _, idea := range x {
		if idea == "good" {
			GoodCount++
		}
	}

	if GoodCount == 0 {
		return "Fail!"
	} else if GoodCount <= 2 {
		return "Publish!"
	} else {
		return "I smell a series!"
	}
}
