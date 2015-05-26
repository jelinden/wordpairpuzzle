package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"regexp"
	"time"
	"sort"
	"unicode/utf8"
	"strconv"
)

var differentWords = make(map[string]string)
var countUniqueChars = make(map[string]int)
var reg, _ = regexp.Compile("[^a-zåäö ]+")
var regWhiteSpace, _ = regexp.Compile("[\\s]+")
const FILENAME = "alastalon_salissa.txt"

// a program to read all words from a file and then get the
// two words containing most unique letters combined
func main() {
	t1 := time.Now()
	readFile(FILENAME)
	fmt.Println("reading the book took", time.Now().Sub(t1))
	iterateCombinedWords()
	printAnswer()
	fmt.Println("task took", time.Now().Sub(t1))
}

// read the file and add all different word to differentWords named map
func readFile(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		sanitizedString := sanitizeString(strings.ToLower(scanner.Text()))
		if sanitizedString != "" && sanitizedString != " " {
			var words sort.StringSlice = strings.Split(sanitizedString, " ")
			wordsFromLine(words)
		}
	}
}

func sanitizeString(line string) string {
	line = reg.ReplaceAllString(line, "")
	regWhiteSpace.ReplaceAllString(line, " ")
	return line
}

func wordsFromLine(words []string) {
	for _, word := range words {
		_, ok := differentWords[word]
		if !ok && word != "" {
			orderedLetters := removeDuplicateCharacters(word)
			_, isOrderedLettersFound := differentWords[orderedLetters]
			if !isOrderedLettersFound {
				differentWords[orderedLetters] = word
			}
		}
	}
}

// remove duplicate characters from each word and sort them alphabetically
func removeDuplicateCharacters(word string) string {
	found := make(map[string]bool)
	var chars sort.StringSlice = []string{}
	for _, char := range word {
		if !found[string(char)] {
			found[string(char)] = true
			chars = append(chars, string(char))
		}
	}
	chars.Sort()
	return strings.Join(chars, "")
}

// go through every word and compare them to other words
func iterateCombinedWords() {
	var keys = sortedMapKeys(differentWords)
	maxValue := 0
	cs := make(chan Result)
	defer close(cs)
	allKeys := make([]string, len(keys))
	copy(allKeys, keys)
	handleWords(&keys, &differentWords, &maxValue, &cs)
}

// the main loop for going through all words
func handleWords(keys *[]string, words *map[string]string, maxValue *int, cs *chan Result) {
	t1 := time.Now()
	go collectUniqueCharWords(*cs, words, maxValue, &countUniqueChars, &(*keys)[0], keys)
	returnedMaxValue := <-*cs
	if returnedMaxValue.maxValue >= *maxValue {
		*maxValue = returnedMaxValue.maxValue
	}
	fmt.Println(len(*keys), "left (" + strconv.Itoa(len(countUniqueChars)) + " items which have",
		*maxValue, "different letters),", "task took", time.Now().Sub(t1))
	if len(*keys) > 1 {
		*keys = (*keys)[:0+copy((*keys)[0:], (*keys)[1:])]
		handleWords(keys, words, maxValue, cs)
	}
}

// no, you don't have an ordered map in golang
func sortedMapKeys(words map[string]string) []string {
	var keys []string
	for key := range words {
		keys = append(keys, key)
	}
	sort.Sort(ByLength(keys))
	return keys
}

// count the length of combined two words and their unique letters
func collectUniqueCharWords(c chan Result, words *map[string]string, maxValue *int, countUniqueChars *map[string]int, k *string, keys *[]string) {
	for innerKey := range (*keys) {
		if (utf8.RuneCountInString(*k + (*keys)[innerKey])) >= *maxValue {
			twoWordsLength := utf8.RuneCountInString(removeDuplicateCharacters(*k + (*keys)[innerKey]))
			if twoWordsLength >= *maxValue {
				if twoWordsLength > *maxValue {
					*countUniqueChars = make(map[string]int)
				}
				*maxValue = twoWordsLength
				(*countUniqueChars)[(*words)[*k] + ", " + (*words)[(*keys)[innerKey]]] = twoWordsLength
			}
		}
	}
	c <- Result{*maxValue}
}

func printAnswer() {
	for key, value := range countUniqueChars {
		fmt.Println("combined words", removeDuplicateCharacters(sanitizeString(key)), key, "had", value, "different letters")
	}
	fmt.Println("There were", len(countUniqueChars), "different word pairs with same length")
}