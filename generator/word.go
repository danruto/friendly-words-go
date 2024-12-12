package generator

import (
	"log"
	"math/rand"
	"strings"

	w "github.com/danruto/friendly-words-go/generator/word"
)

// Wordlist were extracted from `https://github.com/glitchdotcom/friendly-words/tree/main/words`
// which is MIT licensed
const wordlistSize = 4

func getRandomWordFromWordlist(id int) string {
	switch id {
	// Collection, Predicates, Objects, Teams
	case 0:
		{
			idx := rand.Intn(w.WORDLIST_COLLECTIONS_LEN)
			return w.WORDLIST_COLLECTIONS[idx]
		}
	case 1:
		{
			idx := rand.Intn(w.WORDLIST_PREDICATES_LEN)
			return w.WORDLIST_PREDICATES[idx]
		}
	case 2:
		{
			idx := rand.Intn(w.WORDLIST_OBJECTS_LEN)
			return w.WORDLIST_OBJECTS[idx]
		}
	case 3:
		{
			idx := rand.Intn(w.WORDLIST_TEAMS_LEN)
			return w.WORDLIST_TEAMS[idx]
		}
	}

	// Unreachable
	log.Printf("[friendly-words-go] Invalid wordlist ID %d", id)
	return ""
}

// Sample randomly from each wordlist to generate the string selecting
// `size` words joined with a seperator `sep`
func GenerateUniqueWordListIDWithSep(size int, sep string) string {
	// Nothing to process for a size of 0
	if size <= 0 {
		return ""
	}

	words := make([]string, size)

	for ii := range size {
		// Choose a random wordlist
		wl := rand.Intn(wordlistSize)

		// Choose a random word from the wordlist
		word := getRandomWordFromWordlist(wl)

		// Store it
		words[ii] = word
	}

	// Concat the string with a `-` between each word
	return strings.Join(words, sep)
}

// Sample randomly from each wordlist to generate the string selecting
// `size` words
//
// Example usage:
//
// package main
//
// import (
//
//	"fmt"
//	fw "github.com/danruto/friendly-words-go/generator"
//
// )
//
//	func main() {
//	    s:= fw.GenerateUniqueWordListID(5)
//
//	    // Generates something like `chime-worried-potpourri-phalanx-assembly`
//	    fmt.Println(s)
//	}
func GenerateUniqueWordListID(size int) string {
	return GenerateUniqueWordListIDWithSep(size, "-")
}
