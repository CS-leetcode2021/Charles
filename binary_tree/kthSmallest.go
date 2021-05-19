package binary_tree

/**
 *  @ClassName:kthSmallest
 *  @Description: kthSmallest 寻找第k小元素 230
 *  @Author:jackey
 *  @Create:2021/3/23 下午12:53
 */

/*借助中序遍历树，实现一个升序数组切片保存结果，然后直接读取第k小的数*/
func kthSmallest01(root *TreeNode,k int) int {
	res := []int{}
	traverseKth01(root,&res)

	return res[k-1]


}

func traverseKth01(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	traverseKth01(root.Left,res)

	*res =append(*res,root.Val)

	traverseKth01(root.Right,res)

}

/*改进*/
/*实际并没有提升*/

func kthSmallest(root *TreeNode,k int) int {
	res := 0

	count := 0

	traverseKth(root,k,&count,&res)
	return res


}


func traverseKth(root *TreeNode, k int,count *int,res *int) {
	if root == nil {
		return
	}

	traverseKth(root.Left,k,count,res)

	*count++
	if k == *count {
		*res = root.Val
		return
	}

	traverseKth(root.Right,k,count,res)

}

/* 可以更新节点结构的定义，每个节点需要记录，以自己为根的这棵二叉树有多少个节点。
	class TreeNode {
		int val;
		// 以该节点为根的树的节点总数
		int size;
		TreeNode left;
		TreeNode right;
	}
*/

