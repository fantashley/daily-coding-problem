/*
Implement an autocomplete system. That is, given a query string s and a set of
all possible query strings, return all strings in the set that have s as a
prefix.

For example, given the query string de and the set of strings [dog, deer, deal],
return [deer, deal].
*/

package main

import "fmt"

type wordNode struct {
	letters map[byte]*wordNode
	word    string
	isWord  bool
}

func main() {

	fmt.Println(autocomplete("de", map[string]bool{
		"dog":       true,
		"deer":      true,
		"deal":      true,
		"delirious": true,
		"divine":    true,
		"ashley":    true,
		"describe":  true,
	}))
}

func autocomplete(query string, words map[string]bool) []string {

	treeRoot := populateWordTree(words)
	return findMatches(query, treeRoot)
}

func populateWordTree(words map[string]bool) wordNode {

	treeRoot := wordNode{
		letters: make(map[byte]*wordNode),
		word:    "",
	}

	var currentNode *wordNode

	for w := range words {
		currentNode = &treeRoot
		for _, letter := range w {
			newNode, ok := currentNode.letters[byte(letter)]
			if !ok {
				newNode = &wordNode{
					letters: make(map[byte]*wordNode),
					word:    currentNode.word + string(letter),
				}
				currentNode.letters[byte(letter)] = newNode
			}
			currentNode = newNode
		}
		currentNode.isWord = true
	}

	return treeRoot
}

func findMatches(query string, wordTree wordNode) []string {

	currentNode := wordTree

	for _, letter := range query {
		nextNode := currentNode.letters[byte(letter)]
		if nextNode != nil {
			currentNode = *nextNode
		} else {
			return nil
		}
	}

	return findWords(currentNode)
}

func findWords(treeRoot wordNode) []string {

	var wordList []string

	if treeRoot.isWord {
		wordList = append(wordList, treeRoot.word)
	}

	for _, node := range treeRoot.letters {

		wordList = append(wordList, findWords(*node)...)
	}

	return wordList
}
