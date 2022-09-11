package narytreepostordertraversal

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	var result []int
	for _, node := range root.Children {
		result = append(result, postorder(node)...)
	}

	return append(result, root.Val)
}

func postorder2(root *Node) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	stack := []*Node{root}
	visited := make(map[*Node]bool)

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if len(node.Children) == 0 || visited[node] {
			result = append(result, node.Val)
			stack = stack[:(len(stack) - 1)]
			continue
		}

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}

		visited[node] = true
	}

	return result
}
