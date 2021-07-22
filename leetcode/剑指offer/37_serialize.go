package main

import (
	"strconv"
	"strings"
)

/**
 *  @ClassName:37_serialize
 *  @Description:剑指offer-37序列化二叉树 同 leetcode-297
 *  @Author:jackey
 *  @Create:2021/7/19 下午2:41
 */


type Codec struct {
	l []string
}

func ConstructorTree() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	str := serial(root,"")

	return str
}

func serial(root *TreeNode,strbuff string) string  {
	if root ==nil {
		strbuff += "null,"
	}else{
		// 前序遍历
		strbuff += strconv.Itoa(root.Val)+","
		strbuff = serial(root.Left,strbuff)
		strbuff = serial(root.Right,strbuff)
	}
	return strbuff

}



// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// 按照','切分字符串数据
	l := strings.Split(data, ",")
	for i := 0; i < len(l); i++ {
		if l[i] != "" {
			this.l = append(this.l, l[i])
		}
	}
	return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
	//
	if this.l[0] == "null" {
		this.l = this.l[1:]
		return nil
	}

	val, _ := strconv.Atoi(this.l[0])
	root := &TreeNode{Val: val}
	this.l = this.l[1:]
	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()
	return root
}