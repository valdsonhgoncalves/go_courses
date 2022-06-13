package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordSlice := make(map[string]int)
	words := strings.Fields(s)

	for _, v := range words {
		wordSlice[v]++
	}
	
	return wordSlice
}

func main(){
	wc.Test(WordCount)
}