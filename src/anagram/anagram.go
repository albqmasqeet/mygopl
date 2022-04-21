package anagram

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// generate letter frquency map of s1
	map1 := make(map[rune]int)

	for _, c := range s1 {
		map1[c]++
	}

	// generate letter frquency map of s1
	map2 := make(map[rune]int)

	for _, c := range s2 {
		map2[c]++
	}

	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}

	for k, v := range map2 {
		if map1[k] != v {
			return false
		}
	}

	return true

}
