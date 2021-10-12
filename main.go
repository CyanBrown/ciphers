package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var alphabet = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func in(item string, array []string) bool {
	for _, char := range array {
		if item == char {
			return true
		}
	}

	return false
}

func intToArr(i int) []int {
	arr := []int{}
	for {
		arr = append(arr, i%10)
		if i < 10 {
			break
		}

		i = int(math.Floor(float64(i) / 10.0))
	}

	newarr := make([]int, len(arr))
	for i := 0; i < len(arr); i += 1 {
		newarr[i] = arr[len(arr)-1-i]
	}

	return newarr
}

func Amsco(plaintext string, key int) string {
	/*
		Given a n number digit key, seperate the
		plaintext into columns alternating one/two
		letters per row. Then follow the column titles
		(from the key) in ascending order to get the
		new string.
	*/

	nqplaintext := strings.ToLower(plaintext)
	plaintext = ""
	for _, rune := range []rune(nqplaintext) {
		char := string(rune)
		if in(char, alphabet[:]) {
			plaintext += char
		}
	}

	ptrune := []rune(plaintext)
	keyarr := intToArr(key)
	valuemap := make(map[int][]string, len(keyarr))
	chars := make([]string, len(plaintext))
	charperrow := math.Floor(float64(len(keyarr)) / 2.0 * 3.0)
	numrows := int(math.Ceil(float64(len(chars)) / charperrow))

	for i := 0; i < len(plaintext); i++ {
		chars[i] = string(ptrune[i])
	}

	for i := 0; i <= numrows; i++ {
		for idx, key := range keyarr {
			if len(chars) == 1 {
				valuemap[key] = append(valuemap[key], chars[0])
				break
			} else if len(chars) == 0 {
				break
			}

			if (idx%2 == 0 && i%2 == 1) || (idx%2 == 1 && i%2 == 0) {
				valuemap[key] = append(valuemap[key], chars[:2]...)
				chars = chars[2:]
			} else {
				valuemap[key] = append(valuemap[key], chars[0])
				chars = chars[1:]
			}
		}
	}

	keys := make([]int, 0, len(valuemap))
	for k := range valuemap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	ciphertext := ""
	total_index := 0
	for _, key := range keys {
		for _, str := range valuemap[key] {
			ciphertext += str
			if (total_index+1)%5 == 0 && total_index != 0 {
				ciphertext += " "
			}
			total_index++
		}
	}

	return strings.ToUpper(ciphertext)
}

func Ceasar(plaintext string, modulus int) string {
	/*
		This cipher takes in a modulus which is how
		many digits the alphabet is rotated. Once each
		letter is associated with a different letter
		then the letters in the original string are
		replaced with the cipher letter.
	*/

	plaintext = strings.ToLower(plaintext)

	var cipher []string = append(alphabet[modulus:], alphabet[:modulus]...)

	ciphermap := make(map[string]string)
	for i := 0; i < len(alphabet); i++ {
		ciphermap[alphabet[i]] = cipher[i]
	}

	finalstr := ""
	for _, rune := range plaintext {
		char := string(rune)
		if in(char, alphabet[:]) {
			finalstr += ciphermap[char]
		} else {
			finalstr += char
		}
	}

	return finalstr
}

func main() {
	fmt.Println(Amsco("Whoever has made a voyage up the Hudson must remember the Kaatskill mountains.", 35142))
}
