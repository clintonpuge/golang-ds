package main

import (
	"fmt"

	"github.com/clintonpuge/go-ds/search"
)

func main() {
	items := []int{1, 3, 2, 5, 4, 10}
	fmt.Println(search.LinearSearch(items, 3))
}
