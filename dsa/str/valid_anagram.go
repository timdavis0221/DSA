package str

// ValidAnagram returns true if t is an anagram of s, and false otherwise.
//
// Anagram is a word or phrase formed by rearranging the letters of a different word or phase, typically
// using all the original letters exactly once.
//
// For example: race -> care, part -> trap.
//
// Constraints: s and t consists of lowercases English letters.
func ValidAnagram(s string, t string) bool {
	sMap := make(map[rune]int, len([]byte(s)))
	tMap := make(map[rune]int, len([]byte(t)))

	for _, v := range s {
		if _, found := sMap[v]; !found {
			sMap[v] = 1
		} else {
			sMap[v] += 1
		}
	}
	for _, v := range t {
		if _, found := tMap[v]; !found {
			tMap[v] = 1
		} else {
			tMap[v] += 1
		}
	}
	return areMapsEqual(sMap, tMap)
}

func areMapsEqual(map1, map2 map[rune]int) bool {
	if len(map1) != len(map2) {
		return false
	}
	for k, v1 := range map1 {
		if v2, found := map2[k]; !found || v2 != v1 {
			return false
		}
	}
	return true
}

// ValidAnagramApproach2 makes the improvements by using only one slice instead of using two hash tables.
func ValidAnagramApproach2(s, t string) bool {
	chars := make([]int, 26)
	for _, v := range s {
		// Use the difference of ascii dec number to represent each of the 26 lowercase English letters.
		// For example, 0  = 'a' - 'a' (97-97), 1 = 'b' - 'a' (98-97), etc.
		i := v - 'a'
		// Incrementing the number when the relevant letter appears.
		chars[i]++
	}
	for _, v := range t {
		i := v - 'a'
		// Decrementing the number when the relevant letter appears again.
		chars[i]--
	}
	for _, v := range chars {
		// Expect all the numbers would be zero if t string is anagram of s string.
		if v != 0 {
			return false
		}
	}
	return true
}

// ValidAnagramApproach3 uses only one hash table to handle the inputs contain Unicode characters.
func ValidAnagramApproach3(s, t string) bool {
	sMap := make(map[rune]int, len([]byte(s)))
	// Increase the number for the appearance of the lowercase English letters as frequency of appearance.
	for _, v := range s {
		sMap[v]++
	}
	// Decrease the number for the appearance of the lowercase English letters.
	for _, v := range t {
		sMap[v]--
	}
	// Expect all the values in map would be zero If t is an anagram of s.
	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}
	return true
}
