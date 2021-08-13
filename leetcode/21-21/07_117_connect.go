package _1_21

/**
 *  @ClassName:07_117_connect
 *  @Description:TODO
 *  @Author:jackey
 *  @Create:2021/8/13 下午7:59
 */

// 实现一个队列
// 100/60
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{}
	queue = append(queue,root)
	for len(queue) > 0 {
		curLen := len(queue)

		for i := 0; i < curLen; i++ {
			if i+1 < curLen {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				queue = append(queue,queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue,queue[i].Right)
			}
		}
		queue = queue[curLen:]
	}

	return root
}

// 使用递归呢

func connect02(root *Node) *Node {
	maps := make(map[int]*Node)
	maps= dfs(root,0,maps)
	return root
}
func dfs(root *Node,depth int,maps map[int]*Node)(map[int]*Node){
	if root ==nil{
		return maps
	}
	// 递归顺序先右后左是精髓,
	dfs(root.Right,depth+1,maps)
	dfs(root.Left,depth+1,maps)
	if _, ok := maps[depth];ok{
		root.Next = maps[depth]
	}else{
		root.Next = nil
	}
	// 这里会更新迭代一次，每层都最左边的
	maps[depth] =root
	return maps
}