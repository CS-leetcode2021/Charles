package _1_14

/**
 *  @ClassName:14_190_reverseBits
 *  @Description:颠倒二进制位
 *  @Author:jackey
 *  @Create:2021/8/5 下午6:33
 */

// 假设输入的数字是num
// 通过右移操作将每位移动出来
// （num>>i）&0x01
// 然后定义一个返回值=0,
// 以其为目标把上面移出来的二进制数通过逻辑或逐位左移，
// 就获得结果了

func reverseBits(num uint32) uint32 {
	var ret uint32 = 0

	for i := 0; i < 32; i++ {
		ret |= (num >> i) & 0x01 << (31 - i)
	}
	return ret
}