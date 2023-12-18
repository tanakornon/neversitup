package quiz

import "regexp"

func CountSmilyFace(texts []string) int {
	result := 0

	r, _ := regexp.Compile(`[:;][-~]?[)D]`)

	for _, text := range texts {
		if r.MatchString(text) {
			result++
		}
	}

	return result
}
