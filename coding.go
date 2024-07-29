package main

func Palindrome(str string) string {
	p1 := 0
	p2 := len(str) - 1
	for p1 < p2 {
		if str[p1] != str[p2] {
			return "NOT PALINDROM"
		}
		p1++
		p2--
	}
	return "Palindrome"
}

func FindCountOfChar(str, char string) int {
	count := 0

	for i := 0; i < len(str); i++ {

		if string(str[i]) == char {
			count++
		}

	}
	return count

}

// Using Maps
func Anagrams(str1, str2 string) string {
	str1Map := make(map[string]int)
	str2Map := make(map[string]int)

	if len(str1) != len(str2) {
		return "NOT ANAGRAM"
	}

	for i := 0; i < len(str1); i++ {
		var temp1 = str1Map[string(str1[i])]
		str1Map[string(str1[i])] = temp1 + 1
		var temp2 = str2Map[string(str2[i])]
		str2Map[string(str2[i])] = temp2 + 1
	}

	for key, _ := range str1Map {
		if str1Map[key] != str2Map[key] {
			return "NOT ANAGRAM"

		}

	}
	return "ANAGRAM"

}

func CovertToLowerCase(str string) string {
	var new string
	for _, value := range str {
		if value >= 'A' && value <= 'Z' {
			new = new + string(value+32)
		} else {
			new = new + string(value)
		}
	}
	return new

}

func CountVowelsAndConsonants(str string) (int, int) {
	vowelCount := 0
	count := 0
	updatedStr := CovertToLowerCase(str)
	for _, v := range updatedStr {
		if string(v) == "a" || string(v) == "e" || string(v) == "i" || string(v) == "o" || string(v) == "u" {
			vowelCount++
		} else {
			count++
		}

	}
	return vowelCount, count

}

func BinarySearch(arr []int, search int) bool {
	// all these are indexes
	low := 0
	high := len(arr) - 1
	mid := (low + high) / 2

	for low <= high {
		if arr[mid] == search {
			return true
		} else if search < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
		mid = (low + high) / 2
	}
	return false
}
