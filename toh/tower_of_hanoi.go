package toh

import "fmt"

// TowersOfHanoi func
func TowersOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are :")
	TowerUtil(num, "A", "C", "B")
}
