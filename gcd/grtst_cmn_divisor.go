package gcd

// GreatestCommonDivisor func
func GreatestCommonDivisor(m int, n int) int {
	if m < n {
		return GreatestCommonDivisor(n, m)
	}
	if m%n == 0 {
		return n
	}
	return GreatestCommonDivisor(n, m%n)
}
