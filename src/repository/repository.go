package repository

import (
	"prefix-matcher/src/config"
	"strings"
	"sync"
)

var Data *Trie

type Repository struct {
}

func init() {
	Data = NewTrie()
}

// Node represent each character
type Node struct {
	Char     string
	Children [62]*Node
}

func NewNode(char string) *Node {
	node := &Node{Char: char}
	for i := 0; i < 62; i++ {
		node.Children[i] = nil
	}
	return node
}

type Trie struct {
	RootNode *Node
}

func NewTrie() *Trie {
	root := NewNode("\000")
	return &Trie{RootNode: root}
}

func (t *Trie) Insert(PrefixChan chan string, wg *sync.WaitGroup, r *sync.RWMutex) error {
	defer wg.Done()

	for valPrefix := range PrefixChan {
		r.Lock()
		current := Data.RootNode
		strippedWord := strings.ReplaceAll(valPrefix, " ", "")
		for i := 0; i < len(strippedWord); i++ {
			index := config.IndexMap[string(strippedWord[i])]
			if current.Children[index] == nil {
				current.Children[index] = NewNode(string(strippedWord[i]))
			}
			current = current.Children[index]
		}
		r.Unlock()
	}
	return nil
}

func (t *Repository) SearchWord(word string) int {
	strippedWord := strings.ReplaceAll(word, " ", "")

	current := Data.RootNode
	count := 0

	for i := 0; i < len(strippedWord); i++ {
		index := config.IndexMap[string(strippedWord[i])]
		if current != nil && current.Children[index] != nil && current.Children[index].Char == string(strippedWord[i]) {
			count++
			current = current.Children[index]
		} else {
			return count
		}

	}
	return count
}

func NewDatabaseRepository(db *Trie) *Repository {

	return &Repository{}
}
