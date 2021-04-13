package main

import (
	"fmt"
	"math/rand"
)


func main(){
	board := [3][7]string{{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"}}
	player := make(map[int]string)
	turn := 1

	/*
	decidePlayer()

	loop till game over{
		showBoard()

		
		if check_player_turn == true {
			playerMove()
		}esle {
			botMove()
		}

		checkWin()

	}	
	
*/




}

func showBoard(b [3][7]string){

	/*Prints the tic-tac-toe board*/

	for  i := 0; i < 3; i++ {
		fmt.Println()
    	for j := 0; j < 7; j++ {
        	fmt.Printf( b[i][j] )
      }

   }
}


func decidePlayer() int{
	/*
		decides random 'O' or 'X' role to the player
	*/


	player_turn := rand.Intn(2)

	return player_turn
}

func playerCmd(){
	/* takes the command from player
		for eg. move 1,2
				scoreboard
				quit
	*/
}

func playerMove(){
	/*
		takes input commands from playerCmd()
	*/

}



func botMove(){
	/*
		 Random valid moves made by computer
	*/

}




func checkWin(b [3][7]string) int{

	/*
		Checks if the player or the computer won
	*/

	res := [8] string {"","","","","","","",""}


  	res[0]= b[0][5] + b[1][3] + b[2][0]
  	res[1]= b[0][1] + b[1][1] + b[2][1]
  	res[2]= b[0][5] + b[1][5] + b[2][5]
  	res[3]= b[0][1] + b[1][3] + b[2][5]
  	res[4]= b[2][1] + b[2][3] + b[2][5]
  	res[5]= b[1][1] + b[1][3] + b[1][5]
  	res[6]= b[0][1] + b[0][3] + b[0][5]
  	res[7]= b[0][3] + b[1][3] + b[2][3]

  	for _, v := range res {
    	if v == "OOO"{
      		return 1
    	} else if v == "XXX"{
      		return 2
    }
  }
  return 0



}