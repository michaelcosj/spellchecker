package main

import (
	"sort"
	"strings"
)

type Service struct {
	dictionary []string
}

func (s *Service) GetSuggestions(word string, count int) []Suggestion {
	var suggestions []Suggestion
	for _, w := range s.dictionary {
        dictWord := strings.ToLower(w)
		editDistance := levenshteinDistance(dictWord, strings.ToLower(word))
		suggestions = append(suggestions, Suggestion{dictWord, editDistance})
	}
	sort.Sort(Suggestions(suggestions))

	if count >= len(suggestions) {
		count = len(suggestions)
	}
	return suggestions[:count]
}

func levenshteinDistance(src, target string) int {
	srcLen := len(src)
	targetLen := len(target)

	distanceMatrix := make([][]int, srcLen+1)
	for d := range distanceMatrix {
		distanceMatrix[d] = make([]int, targetLen+1)
	}

	for i := 1; i <= srcLen; i++ {
		distanceMatrix[i][0] = i
	}

	for j := 1; j <= targetLen; j++ {
		distanceMatrix[0][j] = j
	}

	for i := 1; i < srcLen+1; i++ {
		for j := 1; j < targetLen+1; j++ {
			deletion := distanceMatrix[i-1][j] + 1
			insertion := distanceMatrix[i][j-1] + 1
			substitution := distanceMatrix[i-1][j-1]

			if src[i-1] != target[j-1] {
				substitution += 1
			}
			distanceMatrix[i][j] = minInt(deletion, insertion, substitution)
		}
	}
	return distanceMatrix[srcLen][targetLen]
}

func minInt(args ...int) int {
	curMin := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < curMin {
			curMin = args[i]
		}
	}
	return curMin
}
