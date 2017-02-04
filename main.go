package main

import (
	"fmt"

	"github.com/MarckK/wwg/animals"
)

func main() {

	golden := animals.Dog{}
	golden.SetName("Beau")
	fmt.Println(golden.GetName())
	fmt.Println(golden.Bark())
}
