package main

import (
	"fmt"
	"math/rand"
	"os"
	"bufio"
	"strings"
	"strconv"
)
var (
	board = [3][7]string{{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"}}
	in = bufio.NewReader(os.Stdin)
	turn = 0
	user_role = decidePlayer()
	gameOver,win_role = checkWin(board, user_role)
	BANNER_HEADER = `
	 _   _      _             _             
	| |_(_) ___| |_ __ _  ___| |_ ___   ___ 
	| __| |/ __| __/ _\ |/ __| __/ _ \ / _ \
	| |_| | (__| || (_| | (__| || (_) |  __/
	 \__|_|\___|\__\__,_|\___|\__\___/ \___|
	`  
)

func main(){

	fmt.Println(BANNER_HEADER)

	for (gameOver != true){
		showBoard(board)
		
		if turn % 2 == user_role {
			board = playerCmd(board, user_role)
		}else {
			board = botMove(board, user_role)
		}

		gameOver,win_role = checkWin(board, user_role)

		turn += 1
		if turn == 9{
			fmt.Println("\nThe Game is a Draw")
			os.Exit(1)
		}

	}
	showBoard(board)
	fmt.Println("\n",win_role," wins!")	
}




func playerCmd(board [3][7]string, user_role int) [3][7]string{
	/* takes the command from player
		for eg. move 1,2
				quit
	*/
	fmt.Println()
	fmt.Printf("tictactoe> ")
	inp, err := in.ReadString('\n')
    if err != nil{
        panic(err)
    }
	cmd := strings.Split(inp, " ")

	switch cmd[0]{
	case "move":
		board = playerMove(cmd[1], board, user_role)
	case "quit\n":
		quit()
	default:
		fmt.Println("Enter valid commands (move, quit)")
		board = playerCmd(board, user_role)
	}
	return board
}


func playerMove(points string, board [3][7]string, user_role int) [3][7]string{
	/*
		takes input co-ordinates and fills the board

		evaluates the index of the board from user
		eg:- if input is 1,2 our board indices would be 1,5
		used in the function playerCmd() for the user command "move"
	*/
	bp := make(map[int]int)
	bp[0],bp[1],bp[2] = 1,3,5

	i,err_i := strconv.Atoi(string(points[0]))
	j,err_j := strconv.Atoi(string(points[2]))
	if err_i != nil || err_j != nil || i>=3 || j>=3{
    	fmt.Println("Enter valid co-ordinates ranging from 0,0 - 2,2")
        os.Exit(1)
    }
    j = bp[j]

  	if board[i][j] == " " {
  		if user_role == 0{
    		board[i][j] = "O"
  	}else{
    	board[i][j] = "X"
  	}
    	
  }else{
  	fmt.Println("invalid move, place already taken")
  	board = playerCmd(board, user_role)
  }
  return board

}




func botMove(board [3][7]string, user_role int) [3][7]string{
	/*
		 Random valid moves made by computer
	*/
	role := make(map[int]string)
	role[0],role[1] = "O","X"
	bot_role := 1 ^ user_role
	
	i,j := rand_idx()
	for board[i][j] != " "{
		i,j = rand_idx()
	}
	if board[i][j] == " "{
		board[i][j] = role[bot_role]
	}

	return board
}

func quit(){
	/*
		used in the function playerCmd() for the user command "quit"
	*/
	fmt.Println("Computer wins")
	os.Exit(1)
}


func rand_idx() (int,int) {
	/* Used to generate random index on the board for 
		called by botMove()
	*/
		i := rand.Intn(3)
		j := rand.Intn(7)

		return i,j
}


func showBoard(b [3][7]string){
	/* Prints the tic-tac-toe board */

	fmt.Println()
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


func checkWin(b [3][7]string , user_role int) (bool,string){

	/*
		Checks if the player or the computer won
	*/
	winner := make(map[int]string)
	winner[user_role] , winner[1^user_role] = "User","Computer"
	res := [8] string {"","","","","","","",""}


  	res[0]= b[0][5] + b[1][3] + b[2][1]
  	res[1]= b[0][1] + b[1][1] + b[2][1]
  	res[2]= b[0][5] + b[1][5] + b[2][5]
  	res[3]= b[0][1] + b[1][3] + b[2][5]
  	res[4]= b[2][1] + b[2][3] + b[2][5]
  	res[5]= b[1][1] + b[1][3] + b[1][5]
  	res[6]= b[0][1] + b[0][3] + b[0][5]
  	res[7]= b[0][3] + b[1][3] + b[2][3]

  	for _, v := range res {
    	if v == "OOO" {
      		return true,winner[0]
    	} else if v == "XXX"{
      		return true,winner[1]
    }
  }
  return false,""
}

