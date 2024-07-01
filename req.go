package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str1 := "1 Double or 2 Twin beds (大床或双床)"
	e1, c1 := splitBedTypeDesc(str1)
	fmt.Println("e1="+e1, "c1="+c1)

	str2 := "1 Double or 2 Twin beds"
	e2, c2 := splitBedTypeDesc(str2)
	fmt.Println("e2="+e2, "c2="+c2)
	// rex := regexp.MustCompile(`((.*?))`)
	// matchArr := rex.FindStringSubmatch(str)
	// fmt.Println(matchArr[1]) // zhname

	// s := strings.Split(str, "(")
	// fmt.Println(s[0]) // enname

	// t := strings.Replace(str, "("+matchArr[1]+")", "", -1)
	// fmt.Println(t) // enname
	// fmt.Println(isContainsNumber("1 Double or 2 Twin beds"))
	// fmt.Println(isContainsNumber("tes1"))
	// fmt.Println(isContainsNumber("test"))
	// fmt.Println(isContainsNumber("te3st"))
	// fmt.Println(isContainsNumber("2test"))

	fmt.Println("a", isSameChar("a"))
	fmt.Println("bb", isSameChar("bb"))
	fmt.Println("aca", isSameChar("aca"))
	fmt.Println("aaab", isSameChar("aaab"))
	fmt.Println("aaaa", isSameChar("aaaa"))
	fmt.Println("baaab", isSameChar("baaab"))
	fmt.Println("AA", isSameChar("AA"))
	fmt.Println("Aa", isSameChar("Aa"))
}

func splitBedTypeDesc(bedTypeDesc string) (enDesc string, cnDesc string) {
	if bedTypeDesc == "" {
		return
	}
	if strings.Contains(bedTypeDesc, "(") {
		fmt.Println("contains")
		rex := regexp.MustCompile(`\((.*?)\)`)
		matchArr := rex.FindStringSubmatch(bedTypeDesc)
		fmt.Println(matchArr)
		if len(matchArr) > 1 {
			cnDesc = matchArr[len(matchArr)-1]
		}
		enDesc = strings.Replace(bedTypeDesc, "("+cnDesc+")", "", -1)
	} else {
		enDesc = bedTypeDesc
	}

	return
}

func isContainsNumber(str string) bool {
	rex := regexp.MustCompile(`\d`)
	return rex.MatchString(str)
}

func isUniquite(str string) bool {
	return regexp.MustCompile(`^(\w)+$`).MatchString(str)
}

func isSameChar(vendorCode string) bool {
	if len(vendorCode) == 0 {
		return false
	}
	base := int(vendorCode[0]) - 'a'
	for _, ch := range vendorCode {
		if int(base)^int(ch-'a') != 0 {
			return false
		}
	}
	return true
}
