package main

import "fmt"

func mergeStringAlternately(word1, word2 string) string {
	merge := []byte{}
	remain := []byte{}
	word1Bytes := []byte(word1)
	word2Bytes := []byte(word2)

	if len(word1Bytes) > len(word2Bytes) {
		remain = word1Bytes[len(word2Bytes):]
		word1Bytes = word1Bytes[:len(word2Bytes)]
	} else {
		remain = word2Bytes[len(word1Bytes):]
		word2Bytes = word2Bytes[:len(word1Bytes)]
	}

	for i := range word1Bytes {
		merge = append(merge, word1Bytes[i])
		merge = append(merge, word2Bytes[i])

	}
	merge = append(merge, remain...)
	return string(merge)
}

func main() {
	word1 := "abcd"
	word2 := "pqrstu"

	fmt.Println("res: ", mergeStringAlternately(word1, word2))
}
