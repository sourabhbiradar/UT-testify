package main

import (
	"fmt"
	"testify/player"
)

func main() {
	s := player.Stats{
		Name: "DK",
		Runs: 87,
		Wkts: 0,
	}

	fmt.Println(player.IsMOTM(s))

}
