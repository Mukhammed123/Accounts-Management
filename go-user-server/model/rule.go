package model

import (
	"regexp"
	"strconv"
)

var (
	idNumberRegexp = regexp.MustCompile(`^[A-Z][1289]\d{8}$`)
	rcNumberRegexp = regexp.MustCompile(`^[A-Z][A-D]\d{8}$`)
	tcNumberRegexp = regexp.MustCompile(`^\d{6}[\dA-Z]{4}$`) // temporary case number

	idNumberLatterTable = map[rune]string{
		'A': "10", 'B': "11", 'C': "12", 'D': "13", 'E': "14", 'F': "15", 'G': "16",
		'H': "17", 'I': "34", 'J': "18", 'K': "19", 'L': "20", 'M': "21", 'N': "22",
		'O': "35", 'P': "23", 'Q': "24", 'R': "25", 'S': "26", 'T': "27", 'U': "28",
		'V': "29", 'W': "32", 'X': "30", 'Y': "31", 'Z': "33",
	}
)

func IsCaseNumberValid(caseNumber string) bool {
	if tcNumberRegexp.MatchString(caseNumber) {
		return true
	}
	return IsIDNumberValid(caseNumber)
}

func IsIDNumberValid(idNumber string) bool {
	idNumberToNs := func(v string) string {
		runes := []rune(v)
		if idNumberRegexp.MatchString(v) {
			return idNumberLatterTable[runes[0]] + string(runes[1:])
		} else if rcNumberRegexp.MatchString(v) {
			return idNumberLatterTable[runes[0]] +
				string([]rune(idNumberLatterTable[runes[1]])[1]) + string(runes[2:])
		}
		return ""
	}
	if ns := []rune(idNumberToNs(idNumber)); len(ns) == 0 {
		return false
	} else {
		kA := [11]int{1, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1}
		sum := 0
		for i, r := range ns {
			n, _ := strconv.Atoi(string(r))
			sum += n * kA[i]
		}
		return sum%10 == 0
	}
}
