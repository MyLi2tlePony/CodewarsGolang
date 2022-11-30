package ordered_count_of_characters

type Tuple struct {
	Char  rune
	Count int
}

func solution(text string) []Tuple {
	frequency := make(map[rune]*Tuple)
	tuples := make([]Tuple, 0)

	for _, char := range text {
		if tuple, ok := frequency[char]; ok {
			tuple.Count++
		} else {
			tuples = append(tuples, Tuple{
				Char:  char,
				Count: 1,
			})

			frequency[char] = &tuples[len(tuples)-1]
		}
	}

	for i := 0; i < len(tuples); i++ {
		tuples[i].Count = frequency[tuples[i].Char].Count
	}

	return tuples
}
