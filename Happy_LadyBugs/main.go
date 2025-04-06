package main

import "fmt"

func isHappy(s string) bool{
	for i := 1; i<len(s)-1;i++{
		if s[i] != s[i+1] && s[i] != s[i-1]{
			return false
		}
	}
	return true
}

func happyLadybugs(s string) string {
	wordCount := make(map[byte]int)
	hasEmpty := false
	for i := 0; i < len(s); i++ {
		if s[i] != '_' {
			wordCount[s[i]]++
		}else{
			hasEmpty = true
		}
	}
	
	for m, _ := range wordCount {
		if wordCount[m] == 1 {
			return "NO"
		}
	}
	

	if !hasEmpty {
		if isHappy(s) {
			return "YES"
		}
		return "NO"
	}
	return "YES"
}

func main() {
	var n int
	var b string
	fmt.Scan(&n)
	fmt.Scan(&b)
	fmt.Println(happyLadybugs(b))
}
