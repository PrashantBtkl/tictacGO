package main



func main(){
	board := [3][7]string{{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"},{"|"," ","|"," ","|"," ","|"}}
	player := make(map[int]string)

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


func checkWin(b [3][7]int) int{

	res := [8] int {0,0,0,0,0,0,0,0}


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