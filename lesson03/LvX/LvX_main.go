package main

import (
	"MyLvX/ty"
	"fmt"
)

func main() {
	game := ty.NewCommon(
		"Resident_Evil",
		400,
		10)

	dlc := ty.NewElectronicDevice(
		&game,
		1.1,
		"_DLC")

	//show basic info of game
	fmt.Println(game.GetName())
	fmt.Println(game.GetPrice())
	fmt.Println(game.GetStock())

	//show basic info of dlc
	fmt.Println(dlc.GetModel())
	fmt.Println(dlc.GetBrand())

	//set Stock
	game.SetStock(20)

}
