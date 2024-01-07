package main

import (
	"fmt"
	"strings"
)

func doesSubstringDivideString(str string, sub string) bool {
	timesToRepeat := len(str) / len(sub)

	if timesToRepeat == 0 {
		return false
	}

	if str == strings.Repeat(sub, timesToRepeat) {
		return true
	}

	return false
}

func gcdOfStrings(str1 string, str2 string) string {
	str1Length, str2Length := len(str1), len(str2)

	startingWord, otherWord := str2, str1
	maxLength := str2Length
	if str1Length < str2Length {
		startingWord, otherWord = str1, str2
		maxLength = str1Length
	}

	section := ""

	for i := maxLength; i > 0; i -= 1 {
		section = startingWord[0:i]

		if doesSubstringDivideString(startingWord, section) {
			if doesSubstringDivideString(otherWord, section) {
				return section
			}
		}
	}

	return ""
}

func main() {
	stringPairs := [...][2]string{
		{"ABCABC", "ABC"},
		{"ABABAB", "ABAB"},
		{"LEET", "CODE"},
	}

	for _, stringPair := range stringPairs {
		output := gcdOfStrings(stringPair[0], stringPair[1])
		fmt.Println(stringPair)
		fmt.Println(output)
	}
}
