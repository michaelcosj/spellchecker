package main

import "fmt"

type Suggestion struct {
	Word     string
	Distance int
}

func (s Suggestion) String() string {
	return fmt.Sprintf("%s (distance %d)\n", s.Word, s.Distance)
}

type Suggestions []Suggestion

func (s Suggestions) Len() int {
	return len(s)
}

func (s Suggestions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Suggestions) Less(i, j int) bool {
	if s[i].Distance != s[j].Distance {
		return s[i].Distance < s[j].Distance
	}
	return s[i].Word < s[j].Word
}
