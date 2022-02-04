package assignment

import (
	"math"
	"strings"
)

//sum two uint32 number and check whether overflows or not
func AddUint32(x, y uint32) (uint32, bool) {
	var z uint32 = x + y
	if z < x {
		return z, true
	}
	return z, false
}

//ceil the float64 to next multiple of 0.25
func CeilNumber(f float64) float64 {

	return math.Ceil(f/0.25) * 0.25
}

//basic implemantation of bubble sort
func AlphabetSoup(s string) string {

	runeArray := []rune(s)
	for i := 0; i < len(runeArray); i++ {

		for j := 0; j < len(runeArray)-i-1; j++ {
			if runeArray[j] > runeArray[j+1] {
				char := runeArray[j+1]
				runeArray[j+1] = runeArray[j]
				runeArray[j] = char
			}

		}
	}
	return string(runeArray)
}

func StringMask(s string, n uint) string {
	if s == "" && n > 0 {
		return "*"
	}
	runeArray := []rune(s)
	if n == 0 || int(n) >= len(runeArray) {
		for i := 0; i < len(runeArray); i++ {
			runeArray[i] = rune('*')
		}
	} else {
		for i := int(n); i < len(runeArray); i++ {
			runeArray[i] = rune('*')
		}
	}

	return string(runeArray)
}

func WordSplit(arr [2]string) string {

	words := strings.Split(arr[1], ",")
	givenWord := arr[0]
	searchedWord := ""
	result := ""
	for i := 0; i < len(givenWord); i++ {
		searchedWord += string(givenWord[i])

		for j := 0; j < len(words); j++ {
			if searchedWord == words[j] {
				result += searchedWord + ","
				searchedWord = ""
			}
		}
	}
	if len(result) == 0 {
		return "not possible"
	}
	return result[:len(result)-1]
}

func VariadicSet(i ...interface{}) []interface{} {
	set := make(map[interface{}]struct{})
	list := []interface{}{}
	for _, v := range i {
		if _, ok := set[v]; !ok {
			set[v] = struct{}{}
			list = append(list, v)

		}
	}

	return list
}
