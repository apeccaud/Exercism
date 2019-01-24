package hamming

import "errors"

// Distance calculates the Hamming difference between 2 DNA
// strands representated by 2 strings (a and b)
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Length of strings are not equal")
	}
	hammingDist := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDist++
		}
	}
	return hammingDist, nil
}
