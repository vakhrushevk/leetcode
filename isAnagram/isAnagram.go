package isanagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := make(map[rune]int)

	for _, c := range s {
		chars[c]++
	}

	for _, c := range t {
		chars[c]--
		if chars[c] < 0 {
			return false
		}
	}

	return true
}
