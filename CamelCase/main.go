package main

/*
Sample Input
saveChangesInTheEditor

Sample Output
5

Explanation
String  contains five words:
save
Changes
In
The
Editor

*/

import (
	"fmt"
	"unicode"
)

func camelcase(s string) int32 {
	// Write your code here
	var result []string
	start := 0
	for i, r := range s {
		if i != 0 && unicode.IsUpper(r) {
			result = append(result, s[start:i])
			start = i
		}
	}
	result = append(result, s[start:])
	fmt.Println(result)
	return int32(len(result))
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(camelcase(str))
}
