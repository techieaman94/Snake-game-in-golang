package main

import (
	"fmt"
	s "./game"
)

func main() {
	var width ,height int
	fmt.Print("Please enter the height of game board : ")
    fmt.Scanf("%d",&height)
	fmt.Print("Now, Please enter the width of game board : ")
    fmt.Scanf("%d",&width)
    if (height < 30 ){
    	height = 30
    }
    if (width < 30){
		width = 30
    }
	s.StartGame(width,height)
}
