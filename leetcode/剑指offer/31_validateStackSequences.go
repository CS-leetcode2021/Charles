package main

/**
 *  @ClassName:31_validateStackSequences
 *  @Description:剑指offer-31 栈的压入和弹出顺序 同 leetcode-946
 *  @Author:jackey
 *  @Create:2021/7/15 下午7:52
 */

// 100/12.35
func validateStackSequences(pushed []int, popped []int) bool {
	s := new(SStack)
	tag := 0
	for i := 0; i < len(popped); i++ {
		if s.Top() != popped[i] {
			if tag == len(pushed){
				return false
			}
			for j := tag; j < len(pushed); j++ {
				if pushed[j] != popped[i] {
					s.Push(pushed[j])
				} else if pushed[j] == popped[i] {
					s.Push(pushed[j])
					tag = j+1
					break
				}
				if j == len(pushed)-1 {
					return false
				}
			}
		}
		s.Pop()
	}

	return true
}