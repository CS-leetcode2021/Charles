package com.tmp.struct;

import java.util.Random;

/**
 * @ClassName: SkipList
 * @Description: TODO
 * @Author: jackey
 * @Create: 2021/7/5 下午1:35
 */
public class SkipList<T> {
    private Node<T> head, tail;
    private int nodeNumber;
    private int listLevel;
    private Random random;
    private static final double PROBABILITY = 0.5;

    public SkipList() {
        random = new Random();
        clear();
    }

    // 清空跳表
    public void clear() {
        head = new Node<T>(Node.HEAD_KEY, null);
        tail = new Node<T>(Node.TAIL_KEY, null);
        horizontalLink(head, tail);
        listLevel = 0;
        nodeNumber = 0;
    }

    public boolean isEmpty() {
        return nodeNumber == 0;
    }

    public int size() {
        return nodeNumber;
    }

    //在最下面一层，找到要插入的位置前面的那个 key
    private Node<T> findNode(int key) {
        Node<T> p = head;
        while (true) {
            while (p.right.key != Node.TAIL_KEY && p.right.key <= key) {
                p = p.right;
            }
            if (p.down != null) {
                p = p.down;
            } else {
                break;
            }

        }
        return p;
    }

    // 查找是否存在key
    public Node<T> search(int key) {
        Node<T> p = findNode(key);
        if (key == p.getKey()) {
            return p;
        } else {
            return null;
        }
    }

    // 向跳表中添加KV，这是最复杂的地方
    public void put(int k, T v) {
        Node<T> p = findNode(k);

        if (k == p.getKey()) {
            p.value = v;
            return;
        }

        // 如果key不相同，那么在底层插入该节点，此时仅仅需要调用backLink（）进行插入即可
        // 创建新节点
        Node<T> q = new Node<T>(k, v);
        backLink(p, q);

        int currentLevel = 0;

        while (random.nextDouble() >= PROBABILITY) {
            if (currentLevel >= listLevel) ;
            {
                listLevel++;
                Node<T> p1 = new Node<T>(Node.HEAD_KEY, null);
                Node<T> p2 = new Node<T>(Node.TAIL_KEY, null);
                horizontalLink(p1, p2);
                verticallyLink(p1, head);
                verticallyLink(p2, tail);
                head = p1;
                tail = p2;
                // 新层构建好之后，顶层只有head、tail两个节点
            }

            //???寻找当前层p后面且p.up != null的p元素
            // 如果存在这个元素，说明q节点被选举w为上层节点后，应该置于p.up后
            while (p.up == null) {
                p = p.left;
            }
            p = p.up;
            Node<T> e = new Node<T>(k, null);

            backLink(p, e);
            verticallyLink(e, q);
            q = e;
            currentLevel++;
        }
        nodeNumber++;
    }

    // node1后面插入node2
    private void backLink(Node<T> node1, Node<T> node2) {
        node2.left = node1;
        node2.right = node1.right;
        node1.right.left = node2;
        node1.right = node2;
    }

    // 水平连接
    private void horizontalLink(Node<T> node1, Node<T> node2) {
        node1.right = node2;
        node2.left = node1;
    }

    // 垂直连接
    private void verticallyLink(Node<T> node1, Node<T> node2) {
        node1.down = node2;
        node2.up = node1;
    }

    // 按照key的升序依次打印出原始数据
    @Override
    public String toString() {
        if (isEmpty()) {
            return "跳跃表为空！";
        }

        // string缓存
        StringBuilder builder = new StringBuilder();
        Node<T> p = head;

        // 寻找最底层
        while (p.down != null) {
            p = p.down;
        }

        // 寻找最底层的head
        while (p.left != null) {
            p = p.left;
        }

        // 寻找最底层的第一个节点
        if (p.right != null) {
            p = p.right;
        }

        // 打印
        while (p.right != null) {
            builder.append(p);
            builder.append("\n");
            p = p.right;
        }

        return builder.toString();
    }

    public void layerTraverse() {
        Node<T> headOnCurrentLayer = head;
        Node<T> currentNode = head;
        int i = 0;
        while (headOnCurrentLayer != null) {
            i++;
            System.out.println("第 " + i + " 层");
            while (currentNode != null) {
                System.out.println(currentNode);
                currentNode = currentNode.right;
            }
            headOnCurrentLayer = headOnCurrentLayer.down;
            currentNode = headOnCurrentLayer;
        }
    }
}
