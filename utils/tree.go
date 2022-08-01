package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTree(raw []any) *TreeNode {
	rLen := len(raw)
	if rLen == 0 {
		return nil
	}

	root := &TreeNode{Val: raw[0].(int)}
	parentStack := []*TreeNode{root}
	pos := len(parentStack)
	for {
		pLen := len(parentStack)
		for i := 0; i < pLen; i++ {
			parent := parentStack[i]
			if pos >= rLen {
				break
			}

			if raw[pos] != nil {
				parent.Left = &TreeNode{Val: raw[pos].(int)}
				parentStack = append(parentStack, parent.Left)
			}
			pos++
			if raw[pos] != nil {
				parent.Right = &TreeNode{Val: raw[pos].(int)}
				parentStack = append(parentStack, parent.Right)
			}
			pos++
		}

		if pos >= rLen {
			break
		}

		parentStack = parentStack[pLen:]
	}

	return root
}

func TestGenerateTree() {

}
