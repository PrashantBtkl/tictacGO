package main

import (
	"fmt"
	"math/rand"
	"os"
)


func main(){
	board := [3][7]string{{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"}}
	turn := 0
	gameOver := false

	/*
	user_role := decidePlayer()

	loop till (turn > 8) or (checkWin true){
		showBoard(board)
		checkWin(board, user_role)
		
		if turn % 2 == user_role {
			board = playerCmd(board, user_role)
		}esle {
			board = botMove(board, user_role)
		}



		turn += 1

	}	
	
*/

}




func playerCmd(board [3][7]string, user_role){
	/* takes the command from player
		for eg. move 1,2
				scoreboard
				quit
	*/
	var inp string

	fmt.Printf("tictactoe>")
	fmt.Scanf("%s", &inp)

	cmd := strings.Split(inp, " ")

	switch cmd[0]{

	case "move":

		move(cmd[1])

	case "quit":

		quit()

	default:
		panic("Enter valid commands (move, quit, scoreboard):")

	}


	func move(points string){
		/* evaluates the index of the board from user
			eg:- if input is 1,2 our board indices would be 1,5
		*/
		bp := make(map[int]int)
		bp[0],bp[1],bp[2] = 1,3,5

		i := strconv.Atoi(string(points[0]))
		j := strconv.Atoi(string(points[2]))

		playerMove(i,bp[j], board, user_role)
	}

	func quit(){
		fmt.Println("bye")
		os.Exit(1)
	}	
}




func playerMove(i int, j int, board [3][7]string, user_role int) [3][7]string{
	/*
		takes input value from playerCmd()
	*/

  if board[i][j] != " " {
    fmt.Println("Please pick an unoccupied space.")
    playerCmd()
  }

  if i < 3 || j < 7 {
    fmt.Println("Please enter valid matrix index from 0,0 - 2,2")
    playerCmd()
  }

  if user_role == 0{
    board[i][j] = "O"
  }else{
    board[i][j] = "X"
  }
  return board

}




func botMove(board [3][7]string, user_role int) [3][7]string{
	/*
		 Random valid moves made by computer
	*/
	role := make(map[int]string)
	role[0],role[1] = "O","X"
	bot_role = math.Abs(user_role - 1)
	
	i,j := rand_idx()
	if board[i][j] != " "{
		i,j = rand_idx()
	}else{
		board[i][j] = role[bot_role]
	}

	func rand_idx() (int,int) {
		i := rand.Intn(3)
		j := rand.Intn(7)

		return i,j
	}

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
	user_role := rand.Intn(2)

	return user_role
}




func checkWin(b [3][7]string, user_role int) int{

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
      		return 0
    	} else if v == "XXX"{
      		return 1
    }
  }
  return false
}