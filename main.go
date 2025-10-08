package main

import (
	"fmt"
	"roleplay/item"
	"roleplay/player"
)

func main() {
	items := map[string]string{
		"Portion": "Health",
		"Spear":   "Tool",
		"Jacket":  "Protection",
		"Sword":   "Tool",
	}
	var mapOfItems []item.Item
	for key, value := range items {
		mapOfItems = append(mapOfItems, item.Item{Name: key, Type: value})
	}

	var playerIsmail player.Player
	playerIsmail.Name = "Ismail"
	playerIsmail.PickUpItem(mapOfItems[2])
	playerIsmail.PickUpItem(mapOfItems[0])
	playerIsmail.PickUpItem(mapOfItems[2])
	playerIsmail.UseItem("Sword")

	fmt.Println(playerIsmail)
}
