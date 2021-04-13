package main

import (
	"fmt"
	"math/rand"
	"os"
	"bufio"
	"strings"
	"strconv"
)


func main(){
	board := [3][7]string{{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"}}
	turn := 0
	gameOver,_ := checkWin(board)

	user_role := decidePlayer()
	fmt.Println(user_role)

	for (turn < 9) || (gameOver != true){
		showBoard(board)
		checkWin(board)
		
		if turn % 2 == user_role {
			board = playerCmd(board, user_role)
		}else {
			board = botMove(board, user_role)
		}
		turn += 1
	}	
}




func playerCmd(board [3][7]string, user_role int) [3][7]string{
	/* takes the command from player
		for eg. move 1,2
				scoreboard
				quit
	*/
	in := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Printf("tictactoe>")
    inp, err := in.ReadString('\n')
    if err != nil{
        panic(err)
    }

	cmd := strings.Split(inp, " ")
	fmt.Println(cmd[0])
	fmt.Println(cmd[1])
	points := cmd[1]

	switch cmd[0]{
	case "move":
		i,j := move(points, board, user_role)
		board = playerMove(i,j, board, user_role)
	case "quit":
		quit()
	default:
		panic("Enter valid commands (move, quit, scoreboard):")
		quit()
	}

	return board
}

func move(points string, board [3][7]string, user_role int) (int,int){
	/* evaluates the index of the board from user
		eg:- if input is 1,2 our board indices would be 1,5
	*/
	bp := make(map[int]int)
	bp[0],bp[1],bp[2] = 1,3,5

	i,err := strconv.Atoi(string(points[0]))
	if err != nil {
        panic(err)
    }
	j,err := strconv.Atoi(string(points[2]))
	if err != nil {
        panic(err)
    }
	return i,bp[j]
}


func playerMove(i int, j int, board [3][7]string, user_role int) [3][7]string{
	/*
		takes input value from playerCmd()
	*/

  if board[i][j] != " " {
    fmt.Println("invalid move, place already taken")
    playerCmd(board, user_role)
  }

  if i > 2 || j > 6 {
    fmt.Println("Please enter valid matrix index from 0,0 - 2,2")
    playerCmd(board, user_role)
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
	bot_role := 1 ^ user_role
	
	i,j := rand_idx()
	fmt.Println(i,j)
	for board[i][j] != " "{
		i,j = rand_idx()
	}
	if board[i][j] == " "{
		board[i][j] = role[bot_role]
	}

	return board
}

func rand_idx() (int,int) {
		i := rand.Intn(3)
		j := rand.Intn(7)

		return i,j
	}


func showBoard(b [3][7]string){

	/*Prints the tic-tac-toe board*/
	 fmt.Println("\n")
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



func checkWin(b [3][7]string) (bool,int){

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
    	if v == "OOO" {
      		return true,0
    	} else if v == "XXX"{
      		return true,1
    }
  }
  return false,-1
}


func quit(){
	fmt.Println("bye")
	os.Exit(1)
}