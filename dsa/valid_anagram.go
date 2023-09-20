package dsa

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
