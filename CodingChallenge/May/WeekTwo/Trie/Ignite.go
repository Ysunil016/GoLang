package main

import "fmt"

func main() {
	trie := Constructor()
	// fmt.Println(trie)

	trie.Insert("hello")
	// fmt.Println(trie.Search("hell"))
	fmt.Println(trie.Search("helloa"))
	// fmt.Println(trie.Search("hello"))
	// fmt.Println(trie.StartsWith("hell"))
	fmt.Println(trie.StartsWith("helloa"))
	// fmt.Println(trie.StartsWith("hello"))
	// fmt.Println(trie.StartsWith("apple"))

	// fmt.Println(trie)
	// fmt.Println(trie.StartsWith("apple"))
	// fmt.Println(trie.Search("app"))

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
	return Trie{
		childrens:   child,
		isEndOfWord: false,
	}
}

// Insert ...
func (trie *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		if len(trie.childrens[word[i]-'a'].childrens) == 0 {
			trie.childrens[word[i]-'a'] = Constructor()
		}
		trie = &trie.childrens[word[i]-'a']
	}
	trie.isEndOfWord = true
}

// Search ...
func (trie *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if len(trie.childrens[word[i]-'a'].childrens) == 0 {
			return false
		}
		trie = &trie.childrens[word[i]-'a']
	}
	if trie.isEndOfWord == true {
		return true
	}

	return false
}

// StartsWith ...
func (trie *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if len(trie.childrens[prefix[i]-'a'].childrens) == 0 {
			return false
		}
		trie = &trie.childrens[prefix[i]-'a']
	}
	return true
}
