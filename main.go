package main

import (
	"fmt"

	"github.com/Anujsd/BankAPI/util"
)

func main() {
	num := util.GetRandomInt(10, 20)
	fmt.Println(num)

	fmt.Println(util.GetRandomString(1))
	fmt.Println(util.GetRandomCurrency())
	fmt.Println(util.GetRandomEmail())
}
