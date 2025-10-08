package player

import (
	"errors"
	"fmt"
	"roleplay/item"
)

type Player struct {
	Name      string
	Inventory []item.Item
}

func (p *Player) itemExists(newItem item.Item) (int, error) {
	for index, inventoryItem := range p.Inventory {
		if newItem.Name == inventoryItem.Name {
			return index, nil
		}
	}
	return -1, errors.New("item does not exist in inventory")
}

func (p *Player) PickUpItem(item item.Item) {
	index, _ := p.itemExists(item)
	if index == -1 {
		p.Inventory = append(p.Inventory, item)
		fmt.Printf("New item added - %s\n", item.Name)
		return
	}
	fmt.Printf("Item already exist in inventory - %s\n", item.Name)
	return
}

func (p *Player) DropItem(itemName string) {
	index, err := p.itemExists(item.Item{Name: itemName})
	if err == nil {
		p.Inventory = append(p.Inventory[:index], p.Inventory[index+1:]...)
		fmt.Printf("Item dropped - %s\n", itemName)
		return
	}
	fmt.Println(err)
}

func (p *Player) UseItem(itemName string) {
	index, err := p.itemExists(item.Item{Name: itemName})
	if err == nil {
		itemInInventory := p.Inventory[index]
		fmt.Printf("Now using %s for %s\n", itemInInventory.Name, itemInInventory.Type)
		return
	}
	fmt.Println(err)
}
