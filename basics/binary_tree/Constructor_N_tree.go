package binary_tree

/**
 *  @ClassName:Constructor_N_tree
 *  @Description:341 扁平化嵌套列表迭代器
 *  @Author:jackey
 *  @Create:2021/3/29 上午10:38
 */
type NestedIterator struct {
	vals []int
}

func ConstructorN(nestedList []*NestedInteger) *NestedIterator {
	var vals []int

	var dfs func([]*NestedInteger)
	dfs = func(nestList []*NestedInteger) {
		for _, nest := range nestList {
			if nest.IsInteger() {
				vals = append(vals, nest.GetInteger())
			} else {
				dfs(nest.GetList())
			}
		}
	}

	dfs(nestedList)
	return &NestedIterator{vals: vals}
}

/*取值，剩余的结果继续保留在迭代器中*/
func (it *NestedIterator) Next() int {
	val := it.vals[0]
	it.vals = it.vals[1:]
	return val
}

/*判断是否还有数值*/
func (it *NestedIterator) HasNext() bool {
	return len(it.vals) > 0
}

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool {
	return false
}

func (this NestedInteger) GetInteger() int {
	return 0
}

func (n *NestedInteger) SetInteger(value int) {}

func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}
