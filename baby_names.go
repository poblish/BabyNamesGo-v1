package main

import "fmt"
import "unicode"
import "sort"

func main() {
	foreNames := sort.StringSlice{"Ginger", "Hazel", "Lana", "Maya", "Rosemary", "Ruby", "Alana", "Clara", "Cora", "Eloise", "Iris", "Ivy", "Maeve", "Evie", "Joy"}
	middleNames := sort.StringSlice{"Evie", "Ginger", "Hazel", "Lana", "Maya", "Ruby", "Alana", "Bella", "Belle", "Cora", "Eloise", "Fleur", "Iris", "Ivy", "Maeve", "Maud", "Violet", "Daisy", "Grace"}
	count := 0

	/* for _, eachName := range foreNames {
		fmt.Println(eachName, "has", syllableCount(eachName), "syllables")
	} */

	sort.Sort(foreNames)
	sort.Sort(middleNames)

	for _, name1 := range foreNames {
		count++
		fmt.Println(count, "...", name1, "Jordan-Regan")

		syllables1 := syllableCount(name1)

		for _, name2 := range middleNames {
			if name1 != name2 {
				syllables2 := syllableCount(name2)

				if name1[0] == name2[0] || (syllables1 == 1 && syllables2 == 1) || (syllables1 == 1 && syllables2 >= 3) || (syllables1 >= 3 && syllables2 >= 3) {
					continue
				}

				count++
				fmt.Println(count, "...", name1, name2, "Jordan-Regan")
			}
		}
	}
}

func isVowel(inChar rune) bool {
	return (inChar == 'a' || inChar == 'e' || inChar == 'i' || inChar == 'o' || inChar == 'u')
}

func syllableCount(inStr string) int {
	if inStr == "Maya" {
		return 2
	}

	var runes_1 = []rune(inStr)

	syllables := 0
	lastChar := '.'
	wasVowel := false

	for i := 0; i < len(runes_1); i++ {
		c := unicode.ToLower(runes_1[i])

		if wasVowel && ((lastChar == 'u' && c == 'a') || (lastChar == 'i' && c == 'a')) {
			syllables++
		} else if isVowel(c) || c == 'y' {
			if !wasVowel && (!(c == 'e') || i < len(runes_1)-1) {
				syllables++
				wasVowel = true
			}
		} else {
			wasVowel = false
		}

		lastChar = c
	}

	if syllables == 0 {
		return 1
	}

	return syllables
}
