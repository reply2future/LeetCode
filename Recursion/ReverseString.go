/**
 * Write a function that reverses a string. The input string is given as an array of characters char[].
 * Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
 * You may assume all the characters consist of printable ascii characters.
 */
package main
func reverseString(s []byte)  {
    _len := len(s)
    if _len == 0 || _len == 1 {
        return
    }
    
    reverseString(s[1 : _len - 1])
    _tmp := s[0]
    s[0] = s[_len - 1]
    s[_len - 1] = _tmp
}
