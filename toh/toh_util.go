package toh

// Util func
func Util(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	Util(num-1, from, temp, to)
	Util(num-1, temp, to, from)
}
