package toh

import "fmt"

// TowerUtil func
func TowerUtil(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	TowerUtil(num-1, from, temp, to)
	fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
	TowerUtil(num-1, temp, to, from)
}
