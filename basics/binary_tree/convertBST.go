package binary_tree

/**
 *  @ClassName:convertBST
 *  @Description: convertBST 将二叉树转化为累加树  538/1038
 *  @Author:jackey
 *  @Create:2021/3/23 下午1:20
 */


/*[-3,-4,0,null,null,-2,1] 没有通过*/
func convertBST01(root *TreeNode) *TreeNode {
	res :=[]int{}

	traverseC(root,&res)
	n := len(res)
	result := 0

	for i := 0; i < n; i++{
		searchC(root,res[i],&result)
		result += res[i]
	}
	return root


}

func traverseC(root *TreeNode,res *[]int)  {
	if root == nil {
		return
	}

	traverseC(root.Right,res)

	*res = append(*res,root.Val)

	traverseC(root.Left,res)

}

func searchC(root *TreeNode, val int,result *int){
	if root == nil {
		return
	}
	if root.Val==val {
		root.Val += *result
	}else if root.Val < val {
		searchC(root.Right,val,result)
	}else {
		searchC(root.Left,val,result)
	}
}

func convertBST(root *TreeNode) *TreeNode {
	sum :=0
	traverseCV(root,&sum)
	return root
}

func traverseCV(root *TreeNode, sum *int) {
	if root ==nil {
		return
	}

	traverseCV(root.Right,sum)
	*sum += root.Val
	root.Val=*sum
	traverseCV(root.Left,sum)
}