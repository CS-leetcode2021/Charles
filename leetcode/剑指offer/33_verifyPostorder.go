package main

/**
 *  @ClassName:33_verifyPostorder
 *  @Description:剑指offer-33 二叉树的后序遍历序列判定
 *  @Author:jackey
 *  @Create:2021/7/17 下午8:04
 */

func verifyPostorder(postorder []int) bool {

	m := len(postorder)
	return judge(postorder,0,m-1 )
}

func judge(array []int, l,r int) bool {
	if l >= r {
		return true
	}

	tag := l

	for array[tag] < array[r] {
		tag++
	}

	tmp := tag
	for array[tmp] > array[r] {
		tmp++
	}

	return tmp == r && judge(array,l,tag-1)&& judge(array,tag,r-1)

}
