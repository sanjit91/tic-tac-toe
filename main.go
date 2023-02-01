package main

import "fmt"

var ( //create a date structer to hold the board
	// x-x-0
	//0-x-x
	//x-0-x

	board [3][3]int
	x     int = 1
	o     int = 2
)

func main() {
	fmt.Println("tic-tac-toe game")

	//newboard()
	//printboard()

	//get the input from the user
	var row, col int
	fmt.Println("enter the row and colum number e,g 1 1")
	fmt.Scanf("%d %d", &row, &col)
	fmt.Println(row, col)

	//	fmt.Printf("board: %v", board)
}

//newboard() initialize the board
func newboard() {

	for i := 0; i < 3; i++ {
		//for j := 0; j < 3; j++ {
			//board[i][j] = 0

		// }

	}
}

//print board() print the board
func printboard() {

	//	for i := 0; i < 3; i++ {
	//		for j := 0; j < 3; j++ {

	//			if board[i][j] == 1 {
	//				fmt.Printf("x |")

	//			}
	//			fmt.Printf("%v | ", board[i][j])
	//		}
	//		fmt.Println()
	//		fmt.Println("----------")
	//	}
	for i, row := range board {
		for j, col := range row {
			if j == 0 {
				fmt.Printf("%v |", col)
			} else if j == 1 {
				fmt.Printf("%v |", col)
			} else {
				fmt.Printf("%v", col)
			}
			if i != 2 {

				fmt.Println()
				fmt.Println("----------")
			}

		}

	}

}

//A function to check input from the user and update the board
//check if the board is full,meaning no more moves can be made and its draw




var (
player 1 ="x"
player 2="o"
emptyspace =""
)

type board [3][3] string

//newboard creates a new board 
func newnewboard()  {

for i:=0; i<3; j<3 ; j++{
	b[j][j]=emptyspace 
}
return b
}


func main(){
	fmt.Println("please enter your move in row and colum format e,g 1,2")
var row,col int 
n,err:= fmt.scan ("%d %d", &row,%col)
if err !=nil {

fmt.Println("invalid input,please try again")

}
if row<1,|| row >3 ||col <1 || col >3 {

	fmt.Println("invalid input,please try again")
}

fmt.Println("you entered",row,col)



}