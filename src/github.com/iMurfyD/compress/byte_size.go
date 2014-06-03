package compress

import "math"

//meant to wrap a byteSize independant of the numberOfTokens
//I was originally going to represent this by the array index, however this seems less ambiguous and prone to a weird screw up
type byteSize struct {
	bs     int
	tokens int
}

func maxByteSize(arr []byteSize) byteSize {
	max := byteSize{0, math.MaxInt16}
	for _, size := range arr {
		if size.tokens < max.tokens {
			max = size
		}
	}
	return max
}
