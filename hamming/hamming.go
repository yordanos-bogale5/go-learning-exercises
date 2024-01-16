// Package hamming contains a function for computing the distance between DNA strands
package hamming

import "errors"

/*Distance computes the number of differing nucleotides between
two DNA strands, a and b if the strands are unequal in length, an exception is thrown */
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strings a and b are of unequal length")
	} else if len(a) == 0 {
		return 0, nil
	}

	as := []byte(a)
	bs := []byte(b)

	s := 0

	for i := 0; i < len(as); i++ {
		if as[i] != bs[i] {
			s++
		}
	}

	return s, nil
}
