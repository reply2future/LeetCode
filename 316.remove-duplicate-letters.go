/*
 * @lc app=leetcode id=316 lang=golang
 *
 * [316] Remove Duplicate Letters
 */
package main

import (
	"sort"
	"strings"

	utils "reply2future.com/utils"
)

// @lc code=start
func removeDuplicateLetters(s string) string {
	// return bruteForce316(s)
	return stack316(s)
}

// #1 direct =================================================================
type Key316 struct {
	s string
	v int
}

func bruteForce316(s string) string {
	letterIndexMaps := make(map[byte][]int)
	letters := []byte(s)
	var queue []byte
	for i, char := range letters {
		if v, ok := letterIndexMaps[char]; ok {
			letterIndexMaps[char] = append(v, i)
		} else {
			letterIndexMaps[char] = []int{i}
			queue = append(queue, char)
		}
	}
	// asc order
	sort.Slice(queue, func(a, b int) bool {
		return queue[a] < queue[b]
	})

	memo := make(map[Key316]string)
	var fn func(string, int) string
	fn = func(s string, offset int) string {
		q := []byte(s)
		qLen := len(q)
		if qLen == 0 {
			return ""
		}
		if val, ok := memo[Key316{s: s, v: offset}]; ok {
			return val
		}

		for i, char := range q {
			indices := letterIndexMaps[char]
			for _, v := range indices {
				if offset > v {
					continue
				} else {
					var next strings.Builder
					next.Write(q[0:i])
					next.Write(q[i+1:])
					var ret string
					if val, ok := memo[Key316{s: next.String(), v: v}]; ok {
						ret = val
					} else {
						ret = fn(next.String(), v)
					}
					if len(next.String()) == len(ret) {
						var sb strings.Builder
						sb.WriteByte(char)
						sb.WriteString(ret)
						return sb.String()
					} else {
						break
					}
				}
			}
		}
		memo[Key316{s: s, v: offset}] = ""
		return ""
	}

	return fn(string(queue), 0)
}
// #1 ========================================================================

// #2 stack ==================================================================
type Stack316 []rune

func (stack Stack316) Len() int { return len(stack) }
func (stack *Stack316) Push(r rune) { 
	(*stack) = append(*stack, r)
}
func (stack *Stack316) Pop() rune {
	sLen := stack.Len()
	r := (*stack)[sLen-1]
	*stack = (*stack)[:sLen-1]
	return r
}
func (stack Stack316) Peek() rune {
	return stack[stack.Len()-1]
}

func stack316(s string) string {
	letterMaxIndexMaps := make(map[rune]int)
	visited := make(map[rune]bool)

	for i, char := range s {
		letterMaxIndexMaps[char] = i
		visited[char] = false 
	}

	var stack Stack316
	for i, char := range s {
		if visited[char] {
			continue
		}

		for stack.Len() > 0 {
			resLastChar := stack.Peek()
			if char >= resLastChar || letterMaxIndexMaps[resLastChar] <= i {
				break
			}
			
			visited[resLastChar] = false 
			stack.Pop()
		}

		stack.Push(char)
		visited[char] = true
	}
	return string(stack)
}
// #2 ========================================================================

// @lc code=end
func main() {
	utils.AssertEqual1[string, string]("example 1", "abc", removeDuplicateLetters, "bcabc")
	utils.AssertEqual1[string, string]("example 2", "acdb", removeDuplicateLetters, "cbacdcbc")
	// time limit
	utils.AssertEqual1[string, string]("example 3", "bfegkuyjorndiqszpcaw", removeDuplicateLetters, "rusrbofeggbbkyuyjsrzornpdguwzizqszpbicdquakqws")
	utils.AssertEqual1[string, string]("example 4", "bhcsdikworfltuzjxaympev", removeDuplicateLetters, "bxshkpdwcsjdbikywvioxrypfzfbppydfilfxxtouzzjxaymjpmdoevv")
}
