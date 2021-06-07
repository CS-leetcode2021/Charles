package union_find

/**
 *  @ClassName:uf_util
 *  @Description:union_find 工具包
 *  @Author:jackey
 *  @Create:2021/4/1 下午7:03
 */

/*关于父节点，是利用一维数组实现的，数组下标索引就是节点的标号，该索引下的数组值，存放的是该节点的父节点索引值*/
/*如果某一个的父节点是他自己，那么它就是根节点*/
func uf(n int) (int,[]int, []int ){
	count := n

	parent := []int{}
	size := []int{}

	for i := 0; i < n; i++ {
		parent = append(parent,i)
		size = append(size,i)
	}
	return count, parent,size
}

func union(p, q, count int, parent, size []int) {
	rootP := find(p, parent)
	rootQ := find(q, parent)
	if rootQ == rootP {
		return
	}

	if size[rootP] > size[rootQ] {
		parent[rootQ] = rootP
		size[rootP] += rootQ
	} else {
		parent[rootP] = rootQ
		size[rootQ] += size[rootP]
	}

	count--
}

/*连接两个图是否联通*/
func connected(p, q int, parent []int) bool {
	rootP := find(p, parent)
	rootQ := find(q, parent)

	return rootP == rootQ
}

/*寻找图的根节点*/
func find(x int, parent []int) int {
	for parent[x] != x {
		parent[x] = parent[parent[x]]
		x = parent[x]
	}
	return x
}

/*统计图的节点权重*/
func count() {

}
