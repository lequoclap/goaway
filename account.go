package main

import (
	"fmt"
	"io/ioutil"
)

type Account struct {
	name   string
	amount float32
}

func (a *Account) setName(name string) {
	a.name = name
}
func (a Account) save() error {
	filename := a.name
	return ioutil.WriteFile(filename, []byte("This is a sample page"), 0600)
}

func (a Account) getAmount() float32 {
	return a.amount
}

func (a *Account) addFund(number float32) {
	a.amount += float32(number)
}

func (a *Account) withdrawMoney(number float32) (bool, float32) {
	var result bool
	if a.amount > 50.0 {
		a.amount -= float32(number)
		result = true
	}
	return result, a.amount
}

func main() {
	lap := Account{"Le quoc Lap", 100}
	fmt.Println(lap.getAmount())
	lap.setName("lap1110")
	fmt.Println(lap.name)

	lap.addFund(100)
	fmt.Println(lap)
	result, balance := lap.withdrawMoney(199)
	fmt.Println(result, balance)

	lap.save()
}
