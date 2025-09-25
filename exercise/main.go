package exercise

import "fmt"

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s pickedup %s!\n", p.Name, item.Name)
}

func (p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s!\n", p.Name, item.Name)

		}
	}
	fmt.Printf("%s does not have %s in inventory!\n", p.Name, itemName)
}

func (p *Player) UseItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "portion" {
				fmt.Printf("%s used %s and feels rejuvenated!\n", p.Name, itemName)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			} else {
				fmt.Printf("%s used %s. \n", p.Name, itemName)
			}
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory. \n", p.Name, itemName)

}
