package com.tmp.struct;

/**
 * @ClassName: Node
 * @Description: TODO
 * @Author: jackey
 * @Create: 2021/7/5 下午1:23
 */
public class Node<T> {
    public int key;
    public T value;
    public Node<T> up,down,left,right;

    public static final int HEAD_KEY = Integer.MIN_VALUE;
    public static final int TAIL_KEY = Integer.MAX_VALUE;

    public Node(int key,T value){
        this.key = key;
        this.value = value;
    }

    public int getKey(){
        return this.key;
    }

    public void setKey(int key){
        this.key = key;
    }

    public T getValue() {
        return value;
    }
    public void setValue(T value) {
        this.value = value;
    }

    public boolean equals(Object o){
        if (this == o){
            return true;
        }

        if (o == null){
            return false;
        }

        if (!(o instanceof Node<?>)){
            return false;
        }

        Node<T> ent;
        try {
            ent = (Node<T>) o;  // 类型检测
        }catch (ClassCastException e){
            return false;
        }
        return (ent.getKey()==key)&&(ent.getValue() == value);
    }

    @Override
    public String toString(){
        return "key-value:"+key+"-"+value;
    }
}
