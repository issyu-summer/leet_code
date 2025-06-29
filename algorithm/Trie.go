package main

func main() {

}

type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

func ConstructorTrie() Trie {
	return Trie{
		children: map[rune]*Trie{},
		isEnd:    false,
	}
}

func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		child, ok := node.children[rune(word[i])]
		if !ok {
			child = &Trie{children: map[rune]*Trie{}}
			node.children[rune(word[i])] = child
		}
		node = child
	}
	node.isEnd = true
}

func (t *Trie) search(word string) *Trie {
	node := t
	for i := 0; i < len(word); i++ {
		child, ok := node.children[rune(word[i])]
		if !ok {
			return nil
		}
		node = child
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.search(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.search(prefix)
	return node != nil
}
