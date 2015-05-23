package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"regexp"
	"time"
	"sort"
)

var differentWords = make(map[string]string)
var countUniqueChars = make(map[string]int)
var reg, _ = regexp.Compile("[^a-zåäö ]+")

// a program to read all words from a file and then get the
// two words containing most unique letters combined
func main() {
	t1 := time.Now()
	readFile("./alastalon_salissa.txt")
	fmt.Println("reading the book took", time.Now().Sub(t1))
	iterateCombinedWords(differentWords)
	for key, value := range countUniqueChars {
		fmt.Println("combined words", key, "had", value, "letters")
	}
	fmt.Println("task took", time.Now().Sub(t1))
}

// read the file and add all different word to differentWords named map
func readFile(path string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var words sort.StringSlice = strings.Split(sanitizeString(strings.ToLower(scanner.Text())), " ")
		wordsFromLine(words)
	}
}

func sanitizeString(line string) string {
	return reg.ReplaceAllString(line, "")
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

func iterateCombinedWords(words map[string]string) {
	var keys = sortedMapKeys(words)
	maxValue, count := 0, 0
	cs := make(chan Result)
	defer close(cs)

	for k := range keys {
		t1 := time.Now()
		go collectUniqueCharWords(cs, &words, maxValue, &countUniqueChars, keys[k], keys)
		returnedMaxValue := <- cs
		if returnedMaxValue.maxValue >= maxValue {
			maxValue = returnedMaxValue.maxValue
		}
		count += 1
		fmt.Println(count, "of", len(words), "(", len(countUniqueChars), "items, for which length is", maxValue, "),",
			"task took", time.Now().Sub(t1))
	}
}

func sortedMapKeys(words map[string]string) []string {
	var keys []string
	for key := range words {
		keys = append(keys, key)
	}
	sort.Sort(ByLength(keys))
	return keys
}

func collectUniqueCharWords(c chan Result, words *map[string]string, maxValue int, countUniqueChars *map[string]int, k string, keys []string) {
	for innerKey := range keys {
		if (len(k + keys[innerKey])) > maxValue {
			twoWordsLength := len(removeDuplicateCharacters(k + keys[innerKey]))
			if twoWordsLength >= maxValue {
				if twoWordsLength > maxValue {
					*countUniqueChars = make(map[string]int)
				}
				maxValue = twoWordsLength
				(*countUniqueChars)[(*words)[k] + "," + (*words)[keys[innerKey]]] = twoWordsLength
			}
		}
	}
	c <- Result{maxValue}
}