package binary_tree

import "container/list"

/**
 *  @ClassName:levelorder
 *  @Description:leetcode-102 二叉树层次遍历
 *  @Author:jackey
 *  @Create:2021/7/7 上午11:54
 */

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}
	res := [][]int{}

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		currsize := len(queue)
		tmp := []int{}
		for i := 0; i < currsize; i++ {
			tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[currsize:]
		res = append(res, tmp)
		tmp = tmp[0:0]
	}
	return res
}


/**
102. 二叉树的层序遍历
*/
func levelOrder2(root *TreeNode) [][]int {
	res:=[][]int{}
	if root==nil{//防止为空
		return res
	}
	queue:=list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len()>0 {
		length:=queue.Len()//保存当前层的长度，然后处理当前层（十分重要，防止添加下层元素影响判断层中元素的个数）
		for i:=0;i<length;i++{
			node:=queue.Remove(queue.Front()).(*TreeNode)//出队列
			if node.Left!=nil{
				queue.PushBack(node.Left)
			}
			if node.Right!=nil{
				queue.PushBack(node.Right)
			}
			tmpArr=append(tmpArr,node.Val)//将值加入本层切片中
		}
		res=append(res,tmpArr)//放入结果集
		tmpArr=[]int{}//清空层的数据
	}
	return res
}


// DFS
var res [][]int

func levelOrder3(root *TreeNode) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}
func dfs(root *TreeNode, level int) {
	if root != nil {
		if level == len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
}