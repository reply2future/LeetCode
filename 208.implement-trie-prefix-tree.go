/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */
package main
import (
	utils "reply2future.com/utils"
	"fmt"
)
// @lc code=start
type Trie struct {
    children [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string)  {
	alps := []byte(word)
	cNode := this
	for _, c := range alps {
		index := c - 'a'
		v := cNode.children[index]
		if v == nil {
			nNode := Trie{}
			cNode.children[index] = &nNode
			cNode = &nNode
		} else {
			cNode = v
		}
	}
	cNode.isEnd = true
}


func (this *Trie) Search(word string) bool {
    alps := []byte(word)
	cNode := this
	for _, c := range alps {
		index := c - 'a'
		v := cNode.children[index]
		if v == nil {
			return false
		}
		cNode = v
	}
	return cNode.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	alps := []byte(prefix)
	cNode := this
	for _, c := range alps {
		index := c - 'a'
		v := cNode.children[index]
		if v == nil {
			return false
		}
		cNode = v
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
func main() {
	fmt.Println("Trie 1 ----------------------")
	t1 := Constructor()
	t1.Insert("apple")
	utils.AssertEqual1[bool,string]("apple",true,t1.Search,"apple")
	utils.AssertEqual1[bool,string]("app",false,t1.Search,"app")
	utils.AssertEqual1[bool,string]("app",true,t1.StartsWith,"app")
	t1.Insert("app")
	utils.AssertEqual1[bool,string]("app",true,t1.Search,"app")
}
