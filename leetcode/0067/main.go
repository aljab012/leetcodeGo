package main

import "strconv"

func addBinary(a string, b string) string {
	aPointer := len(a) - 1
	bPointer := len(b) - 1

	result := ""
	carry := 0
	for aPointer >= 0 || bPointer >= 0 || carry > 0 {
		sum := carry

		if aPointer >= 0 {
			sum += int(a[aPointer] - '0')
			aPointer--
		}

		if bPointer >= 0 {
			sum += int(b[bPointer] - '0')
			bPointer--
		}

		result = strconv.Itoa(sum%2) + result
		carry = sum / 2
	}

	return result
}
