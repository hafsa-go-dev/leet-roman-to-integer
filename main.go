package main

import "strings"

func main() {

}

func romanToInt(s string) int {
	str := strings.ToUpper(s)
	str = strings.TrimSpace(str)

	total := 0

	for i := 0; i < len(str); i++ {
		currentChar := str[i]

		switch currentChar {
		case 'I':
			total += 1
		case 'V':
			total += 5
		case 'X':
			total += 10
		case 'L':
			total += 50
		case 'C':
			total += 100
		case 'D':
			total += 500
		case 'M':
			total += 1000
		default:
			total = -1
		}
	}
	if strings.Contains(str, "IV") || strings.Contains(str, "IX") {
		total -= 2
	}

	if strings.Contains(str, "XL") || strings.Contains(str, "XC") {
		total -= 20
	}

	if strings.Contains(str, "CD") || strings.Contains(str, "CM") {
		total -= 200
	}

	return total
}
