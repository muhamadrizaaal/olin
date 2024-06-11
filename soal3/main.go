package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	concatLen := wordLen * len(words)
	wordCount := make(map[string]int)
	result := []int{}

	for _, word := range words {
		wordCount[word]++
	}

	for i := 0; i <= len(s)-concatLen; i++ {
		seen := make(map[string]int)
		j := 0
		for ; j < len(words); j++ {
			word := s[i+j*wordLen : i+(j+1)*wordLen]
			if _, found := wordCount[word]; found {
				seen[word]++
				if seen[word] > wordCount[word] {
					break
				}
			} else {
				break
			}
		}
		if j == len(words) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "word"}
	fmt.Println(findSubstring(s, words)) 
	// output yang seharusnya keluar adalah [] dan bukan [8]
	// karena tidak ada substring yang digabungkan
}
