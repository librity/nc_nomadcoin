package main

import "fmt"

func main() {
	foods := [3]string{"potato", "pizza", "tacos"}
	games := []string{"bingo", "scrabble", "yahtzee"}

	printArrayRange(foods)
	fmt.Println("---")

	printArrayLoop(foods)
	fmt.Println("---")

	// [:] transforms an array to a slice
	printSliceRange(foods[:])
	fmt.Println("---")

	printSliceRange(games)
	fmt.Println("---")

	printSliceLoop(games)
	fmt.Println("---")

	printSliceFmt(games)
	fmt.Println("---")

	games = append(games, "tictactoe")
	printSliceFmt(games)
	fmt.Println("---")

	moreGames := append(games, "chess")
	printSliceFmt(moreGames)
	fmt.Println("---")

	printSliceLength(foods[:])
	printSliceLength(games)
	printSliceLength(moreGames)
	fmt.Println("---")

}

func printArrayRange(strings [3]string) {
	for _, str := range strings {
		fmt.Println(str)
	}
}

func printArrayLoop(strings [3]string) {
	for index := 0; index < len(strings); index++ {
		fmt.Println(strings[index])
	}
}

func printSliceRange(strings []string) {
	for _, str := range strings {
		fmt.Println(str)
	}
}

func printSliceLoop(strings []string) {
	for index := 0; index < len(strings); index++ {
		fmt.Println(strings[index])
	}
}

func printSliceFmt(strings []string) {
	fmt.Printf("%v\n", strings)
}

func printSliceLength(strings []string) {
	fmt.Println(len(strings))
}
