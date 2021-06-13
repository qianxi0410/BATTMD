package hot100

type node struct {
	val rune
	children map[rune]*node
	end bool
}

type Trie struct {
	head *node
}


/** Initialize your data structure here. */
func _Constructor() Trie {
	return Trie{head: &node{
		val: '/',
		children: make(map[rune]*node),
		end: false,
	}}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	p := this.head
	for _, v := range word {
		if p.children == nil {
			p.children = make(map[rune]*node)
			p.children[v] = &node{val: v}
			p = p.children[v]
		} else {
			if next, ok := p.children[v]; ok {
				p = next
			} else {
				p.children[v] = &node{val: v}
				p = p.children[v]
			}
		}
	}
	p.end = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this.head
	for _, v := range word {
		if next, ok := p.children[v]; ok {
			p = next
		} else {
			return false
		}
	}
	return p.end
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this.head
	for _, v := range prefix {
		if next, ok := p.children[v]; ok {
			p = next
		} else {
			return false
		}
	}
	return !p.end
}