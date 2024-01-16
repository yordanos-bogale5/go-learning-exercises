package luhn

// Valid validates a number using the Luhn algorithm
// The number is pass in as a string, s
// The function returns true is the number is valid
// https://en.wikipedia.org/wiki/Luhn_algorithm
// BenchmarkValid-4        20000000                71.4 ns/op             0 B/op          0 allocs/op
func Valid(number string) bool {
	runes := []rune(number)
	length := len(runes)

	// length must be 2 or more
	if length := len(runes); length < 2 {
		return false
	}

	sum := 0
	digitCount := 0 // only increase if numer is a digit

	for i := length - 1; i >= 0; i-- {
		digit := int(runes[i] - '0')

		if digit < 0 || digit > 9 {
			if runes[i] == ' ' {
				// skip to next iteration if space
				continue
			}
			return false
		}

		if digitCount%2 == 1 {
			// every second from string end
			sum += (digit*2)/10 + (digit*2)%10
		} else {
			sum += digit
		}

		digitCount++
	}

	if digitCount <= 2 {
		return false
	}
	if sum%10 == 0 {
		return true
	}
	return false
}
