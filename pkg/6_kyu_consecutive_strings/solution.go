package create_phone_number

func solution(strArr []string, k int) string {
	if len(strArr) < 1 || len(strArr) < k {
		return ""
	}

	bestLen := 0
	bestLenStartIndex := 0

	for globalIndex := 0; globalIndex <= len(strArr)-k; globalIndex++ {
		sequencesLen := 0

		for sequenceIndex := 0; sequenceIndex < k; sequenceIndex++ {
			sequencesLen += len(strArr[globalIndex+sequenceIndex])
		}

		if bestLen < sequencesLen {
			bestLen = sequencesLen
			bestLenStartIndex = globalIndex
		}
	}

	resultString := ""
	bestLenEndIndex := bestLenStartIndex + k

	for ; bestLenStartIndex < bestLenEndIndex && bestLenStartIndex < len(strArr); bestLenStartIndex++ {
		resultString += strArr[bestLenStartIndex]
	}

	return resultString
}
