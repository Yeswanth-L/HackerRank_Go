package main

import (
	"fmt"
	"strings"
    "math"
)
/*
Sample Input 0
3
Ab1

Sample Output 0
3

Sample Input 1
11
#HackerRank

Sample Output 1
1

*/
func minimumNumber(n int, password string) int {
	// Return the minimum number of characters to make the password strong
	res := make(map[string]bool)
	res["numbers"] = false
	res["lower_case"] = false
	res["upper_case"] = false
	res["special_characters"] = false

	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"
	for i := 0; i < len(password); i++ {
		if strings.Contains(numbers, string(password[i])) {
			res["numbers"] = true
		}
		if strings.Contains(lower_case, string(password[i])) {
			res["lower_case"] = true
		}
		if strings.Contains(upper_case, string(password[i])) {
			res["upper_case"] = true
		}
		if strings.Contains(special_characters, string(password[i])) {
			res["special_characters"] = true
		}
	}
	
	missing := 0
	for _, val := range res {
		if !val {
			missing++
		}
	}
    
    return int(math.Max(float64(missing), float64(6-n)))
}

func main() {
	var n int
	var str string
	fmt.Scan(&n)
	fmt.Scan(&str)
	fmt.Println(minimumNumber(n,str))
}