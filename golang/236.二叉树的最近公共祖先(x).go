package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到了p或者q，就直接返回
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 如果根节点的左右都有值，那么说明查找的两个节点分布在左右节点，就返回根节点
	// 如果一侧有值。那么说明p或者q或者pq都在一侧，把值传递给根节点
	if left != nil && right != nil { // 对应另一侧有
		return root
	}

	if left == nil { // 对应另一侧没有
		return right
	} else {
		return left
	}
}
