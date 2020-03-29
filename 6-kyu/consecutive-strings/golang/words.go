package words

func combine(words []string, by int) []string {
	badData := by > len(words) || by <= 0
	if badData {
		return []string{}
	}

	newSize := len(words) - by + 1
	arr := make([]string, newSize)

	for i := 0; i < newSize; i++ {
		newWord := ""
		for j := 0; j < by; j++ {
			newWord += words[i + j]
		}
		arr[i] = newWord
	}
	return arr
}

func hello(name string) string {
	return "Hello " + name
}
