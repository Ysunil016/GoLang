package main

import "fmt"

func main() {
	trie := Constructor()

	// trie.childrens[10] = Trie{
	// 	childrens:   &[26]Trie{},
	// 	isEndOfWord: false,
	// }

	fmt.Printf("%v,%T", trie.childrens[10], trie.childrens[10])
}

// Trie ...
type Trie struct {
	childrens   []Trie
	isEndOfWord bool
}

// Constructor ...
func Constructor() Trie {
	var child []Trie
	for i := 0; i < 26; i++ {
		child = append(child, Trie{})
	}
	tries := Trie{
		childrens:   child,
		isEndOfWord: false,
	}
	return tries
}
