package main

import  "fmt"

var (
	player1 = "1"
	player2 = "2"
	emptySpace = "0"
)
var board [3] [3] int

//newboard    create a new board
func newboard ( 
		for i := 0; i<3 ; i++ {
		for j :=0; j<3; j++{
			board [i] [j] =emptyspace
		}
	}
)

//print board - prints the board

func printboard () {
	for i :=0; i<3; i++{
		for j :=0; j<3; j++{
			fmt.Printf(board[i][j])
			if j <2{
				fmt.Printf(" | ")
			}

			if board [i][j]--2{
				fmt.Printf("x |")
			}elase if board [i] [j] ==1 {
				fmt.Printf( "0 |")
			} else {
				fmt.Printf( " |")
			}
		} 
		fmt.Println()
		fmt.Println("-------------")
	}


}

funt