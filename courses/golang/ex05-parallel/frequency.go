package letter

type Map map[rune]int

func Frequency(input string) Map {
	result := Map{}

	for _, c := range input {
		result[c]++
	}
	return result
}

func ConcurrentFrequency(input []string) Map {
	result := Map{}

	ch := make(chan Map, len(input))

	for _, text := range input {
		go calculate(text, ch)
	}

	for range input {
		for k, v := range <-ch {
			result[k] += v
		}
	}

	return result
}

func calculate(input string, ch chan Map) {
	ch <- Frequency(input)
}
