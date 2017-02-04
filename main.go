package main

import (
	"fmt"

	"github.com/MarckK/wwg/animals"
)

func main() {

	kitty := animals.Kitten{}
	kitty.SetName("Kit Kit")
	fmt.Println(kitty.GetName())

	kitty2 := &kitty
	kitty2.SetName("Mr Tom")
	fmt.Println(kitty.GetName())
	fmt.Println(kitty2.GetName())

	golden := animals.Dog{}
	golden.SetName("Beau")
	fmt.Println(golden.GetName())
	fmt.Println(golden.Bark())
}
