package main

import "fmt"
/*
Sample Input 0

aaabccddd
Sample Output 0

abd
Explanation 0

Perform the following sequence of operations to get the final string:

aaabccddd → abccddd → abddd → abd
Sample Input 1

aa
Sample Output 1

Empty String
Explanation 1

aa → Empty String
Sample Input 2

baab
Sample Output 2

Empty String
Explanation 2

baab → bb → Empty String
*/

func superReducedString(s string) string {
	result := ""

	for _, ch := range s {
		n := len(result)
		if n > 0 && result[n-1] == byte(ch) {
			result = result[:n-1] // remove last character (like a pop)
		} else {
			result += string(ch)
		}
	}

	if result == "" {
		return "Empty String"
	}
	return result
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(superReducedString(str))
}