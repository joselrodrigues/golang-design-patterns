package main

import (
	"fmt"
	. "patterns/builder"
)

func main() {
	// Builder design pattern
	cakeBuilder := GetBuilder("ChocolateLarge")
	director := SetDirector(cakeBuilder)
	bigChocolateCake := director.CookChocalateCake()
	fmt.Println(bigChocolateCake)

	cakeBuilder = GetBuilder("ChocolateSmall")
	director = SetDirector(cakeBuilder)
	smallChocolateCake := director.CookChocalateCake()
	fmt.Println(smallChocolateCake)
}
