package main
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
// 	pList := []int{}
// 	qList := []int{}
// 	var dfs func(*TreeNode, *[]int)
// 	dfs = func(root *TreeNode, list *[]int) {
// 		if root == nil{
// 			*list = append(*list, -999)
// 			return
// 		}
// 		*list = append(*list, root.Val)
// 		dfs(root.Left,*&list)
// 		dfs(root.Right,*&list)
// 	}
// 	dfs(p,&pList)
// 	dfs(q,&qList)
// 	if len(pList) != len(qList){
// 		return false
// 	}
// 	r := true
// 	for k,_ := range pList{
// 		if pList[k] != qList[k]{
// 			r = false
// 			break
// 		}
// 	}
// 	return r
}