package quiz

func Permutations(text string) []string {
	result := []string{}
	visited := map[string]bool{}

	if len(text) == 1 {
		result = append(result, text)
	}

	for i := 0; i < len(text); i++ {
		for _, remain := range Permutations(text[:i] + text[i+1:]) {
			cur := string(text[i]) + remain

			if !visited[cur] {
				visited[cur] = true
				result = append(result, cur)
			}
		}
	}

	return result
}
