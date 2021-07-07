package Figure

/**
 *  @ClassName:countComponents
 *  @Description:leetcode-323 无向图中的连通分量的数目 && 涉及到BFS
 *  @Author:jackey
 *  @Create:2021/7/7 下午12:36
 */

func countComponents(n int, edges [][]int) int {
	// 先创建邻接表

	adj := make([][]int,n)

	for i := 0; i < n; i++ {
		adj[i] = make([]int,0)
	}

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]],edge[1])
		adj[edge[1]] = append(adj[edge[1]],edge[0])
	}

	// 构建标记数组
	ans := 0	// 连通分量的数目

	vis := make([]bool,n)
	for i := 0; i < n; i++ {
		if !vis[i] {// 没有访问过
			bfs(adj,i,vis)
			ans++
		}
	}

	return ans
}

func bfs(adj [][]int, u int, vis []bool)  {
	q := new(Queue)
	q.EnQueue(u)
	vis[u] = true
	for q.Size() > 0 {
		n := q.Dequeue()
		if n < 0 {
			continue
		}

		neighbors := adj[n]

		for _, k := range neighbors {
			if !vis[k] {
				q.EnQueue(k)
				vis[k] = true
			}
		}
	}
}


// ---------------------------------
// 并查集

func countComponents2(n int, edges [][]int) int {
	uf := NewUF(n)
	for _, edge := range edges {
		uf.Union(edge[0], edge[1])
	}
	return uf.ConnectedComponents()
}

// Union-Find 并查集
type UF struct {
	// 连通分量个数
	Connects int

	// 每个点的父节点
	Parent []int

	// 每个点的重量
	Size []int
}

func NewUF(size int) *UF {
	uf := UF {
		Connects: size,
		Parent: make([]int, size),
		Size: make([]int, size),
	}
	for i := 0; i < size; i++ {
		uf.Parent[i] = i
		uf.Size[i] = 1
	}
	return &uf
}

func (uf *UF) Union(p, q int) {
	rootP := uf.GetRoot(p)
	rootQ := uf.GetRoot(q)
	if rootP == rootQ {
		return
	}

	// 尽量让小的接到大的后面
	if uf.Size[rootP] > uf.Size[rootQ] {
		uf.Parent[rootQ] = rootP
		uf.Size[rootP] += uf.Size[rootQ]
	} else {
		uf.Parent[rootP] = rootQ
		uf.Size[rootQ] += uf.Size[rootP]
	}
	uf.Connects--
}

func (uf *UF) IsConnected(p, q int) bool {
	rootP := uf.GetRoot(p)
	rootQ := uf.GetRoot(q)
	return rootP == rootQ
}

func (uf *UF) GetRoot(x int) int {
	for uf.Parent[x] != x {
		// 路径压缩
		uf.Parent[x] = uf.Parent[uf.Parent[x]]
		x = uf.Parent[x]
	}
	return x
}

func (uf *UF) ConnectedComponents() int {
	return uf.Connects
}