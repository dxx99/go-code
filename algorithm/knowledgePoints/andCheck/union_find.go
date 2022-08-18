package main

// TODO
// andCheck
//
func main() {

}

// UnionFind 并查集接口的定义
type UnionFind interface {
	Find(x int) int				// 查找某个元素的根节点
	Union(x, y int)				// 为x和y建立联系
	Connected(x, y int) bool	// 判断x和y是否相连(在同一颗树也就是连通分量中)
	Count() int					// 返回连通分量的个数，也就是多少棵树
}

type Uf struct {
	parent []int
	ranks []int	// 【按秩合并】新增字段记录树的高度
	count int	// 记录连通量的个数
}

func NewUf(n int) *Uf {
	// 初始化连通分量
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}

	// 初始化树的高度
	rks := make([]int, n)
	for i := 0; i < n; i++ {
		rks[i] = 1
	}

	return &Uf{
		parent: p,
		ranks: rks,
		count:  n,
	}
}

// Find
// 目的是寻找某个元素所在树的根节点
func (u *Uf) Find(x int) int {
	for u.parent[x] != x {
		//1. 隔代压缩
		u.parent[x] = u.parent[u.parent[x]] // 把路径上的每个节点的父节点指向祖父节点，即父节点-->祖父节点
		x = u.parent[x]
	}
	return x
}

// Union
// union 方法顾名思义就是把两个元素联系起来
// 具体的做法先找到各自的根节点，再把其中一个元素的根节点连接到另一个元素的根节点上
func (u *Uf) Union(x, y int) {
	rx, ry := u.Find(x), u.Find(y)
	if rx == ry {	// 根节点相同，直接返回
		return
	}

	//2. 按秩合并
	if u.ranks[rx] > u.ranks[ry] {
		u.parent[ry] = rx
	}else if u.ranks[rx] < u.ranks[ry] {
		u.parent[rx] = ry
	}else {
		u.parent[ry] = rx
		u.ranks[rx]++
	}
	// u.parent[ry] = rx	//连通分量的合并操作
	u.count--
}

func (u *Uf) Connected(x, y int) bool {
	return  u.Find(x) == u.Find(y)
}

func (u *Uf) Count() int {
	return u.count
}
