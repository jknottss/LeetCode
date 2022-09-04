package main

func reverseNumber(nbr int) int {
	res := 0
	for nbr > 0 {
		reminder := nbr % 10
		res = (res * 10) + reminder
		nbr /= 10
	}
	return res
}

func isPalindrome(x int) bool {
	check := reverseNumber(x)
	return x == check
}
