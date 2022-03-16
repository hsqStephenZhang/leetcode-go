package main

import "fmt"

type Trie struct {
	children []*Trie
	isLeaf   bool
}

func Constructor() Trie {
	return Trie{
		children: make([]*Trie, 26),
		isLeaf:   false,
	}
}

func (root *Trie) Insert(word string) {
	for _, ch := range word[0 : len(word)-1] {
		if root.children[ch-'a'] == nil {
			root.children[ch-'a'] = &Trie{
				children: make([]*Trie, 26),
			}
		}
		root = root.children[ch-'a']
	}
	if root.children[word[len(word)-1]-'a'] == nil {
		root.children[word[len(word)-1]-'a'] = &Trie{
			children: make([]*Trie, 26),
			isLeaf:   true,
		}
	} else {
		root.children[word[len(word)-1]-'a'].isLeaf = true
	}
}

func (root *Trie) Search(word string) bool {
	for _, ch := range word {
		if root.children[ch-'a'] == nil {
			return false
		}
		root = root.children[ch-'a']
	}
	return root.isLeaf
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, ch := range prefix {
		if this.children[ch-'a'] == nil {
			return false
		}
		this = this.children[ch-'a']
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
