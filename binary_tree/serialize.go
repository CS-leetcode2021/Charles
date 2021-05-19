package binary_tree

import (
	"strconv"
	"strings"
)

/**
 *  @ClassName:serialize
 *  @Description: 树序列化和反序列化
 *  @Author:jackey
 *  @Create:2021/3/19 上午11:23
 */

type Codec struct {
	l []string
}

func Constructor() Codec {
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

		strbuff += strconv.Itoa(root.Val)+","

		strbuff = serial(root.Left,strbuff)
		strbuff = serial(root.Right,strbuff)
	}
	return strbuff

}



// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
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