package iterator

// Repeat a character a determinated number of times
func Repeat(character string, times int) string {
	var repeated string

	for range times {
		repeated += character
	}

	return repeated
}
