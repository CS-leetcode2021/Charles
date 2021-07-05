package com.tmp.struct;

/**
 * @ClassName: Test
 * @Description: TODO
 * @Author: jackey
 * @Create: 2021/7/5 下午3:33
 */
public class Test {
    // 为了减少数据的冗余，只有对底层的节点是有完整数据的，上层的节点无value
    public static void main(String[] args) {
        SkipList<String> list = new SkipList<String>();

        list.put(2,"tmp_cap_2");
        list.put(1,"tmp_cap_1");
        list.put(3,"tmp_cap_3");
        list.put(1,"tmp_cap_1_1");
        list.put(4,"tmp_cap_4");
        list.put(6,"tmp_cap_6");
        list.put(5,"tmp_cap_5");
        System.out.println(list);
        System.out.println("____________________");
        list.layerTraverse();
    }
}
