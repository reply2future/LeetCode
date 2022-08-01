/*
 * @lc app=leetcode id=212 lang=golang
 *
 * [212] Word Search II
 */
package main
import (
	utils "reply2future.com/utils"
)
 // @lc code=start
func findWords(board [][]byte, words []string) []string {
	ret := []string{}
	trie := new(TrieNode)
	for _, word := range words {
		trie.Insert(word)
	}

	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, trie, i, j, &ret)
		}
	}

	return ret
}

var boardDirection = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func dfs(board [][]byte, trie *TrieNode, i int, j int, ret *[]string) {
	letter := board[i][j]
	
	node := trie.Children[letter-'a']
	if node == nil {
		return
	}
	if node.Word != "" {
		*ret = append(*ret, node.Word)
		node.Word = ""
	}

	board[i][j] = '.'
	
	m := len(board)
	n := len(board[0])
	for _, v := range boardDirection {
		x, y := i + v[0], j + v[1]
		if x < 0 || y < 0 || x >= m || y >= n || board[x][y] == '.' {
			continue
		}
		dfs(board, node, x, y, ret)
	}
	
	board[i][j] = letter
}

type TrieNode struct {
	Children [26]*TrieNode
	Word string
}

func (this *TrieNode) Insert(word string) {
	cNode := this
	for _, char := range word {
		idx := char - 'a'
		if cNode.Children[idx] == nil {
			cNode.Children[idx] = new(TrieNode)
		}
		cNode = cNode.Children[idx]
	}
	cNode.Word = word
}

// @lc code=end
func main() {
	utils.AssertEqual2[[]string,[][]byte,[]string]("example 1",[]string{"oath", "eat"},findWords,[][]byte{
		{'o','a','a','n'},
		{'e','t','a','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
	}, []string{"oath","pea","eat","rain"})
	utils.AssertEqual2[[]string,[][]byte,[]string]("example 2",[]string{},findWords,[][]byte{
		{'a','b'},
		{'c','d'},
	}, []string{"abcb"})
	utils.AssertEqual2[[]string,[][]byte,[]string]("example 3",[]string{"abcdefg","befa","eaabcdgfa","gfedcbaaa"},findWords,[][]byte{
		{'a','b','c'},
		{'a','e','d'},
		{'a','f','g'},
	}, []string{"abcdefg","gfedcbaaa","eaabcdgfa","befa","dgc","ade"})
	utils.AssertEqual2[[]string,[][]byte,[]string]("example 4",[]string{"abcdeb"},findWords,[][]byte{
		{'a','b','e'},
		{'b','c','d'},
	}, []string{"abcdeb"})
}
