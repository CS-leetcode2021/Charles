# Redis 数据类型

---

1、数据类型

    string、hash、list、set、zset


2、接口

    set、get、del

3、zset底层如何实现的？用score值保证有序，score值可以重复吗？

    Zse原理：有序集合对象是有序的。与列表使用索引下标作为排序依据不同，有序集合为每个元素设置一个分数（score）作为排序依据
    
    Zset底层如何实现：
    
    1、使用ziplist：
        压缩列表（ziplist）是Redis为了节省内存而开发的，是由一系列特殊编码的连续内存块组成的顺序型数据结构，一个压缩列表
        可以包含任意多个节点（entry），每个节点可以保存一个字节数组或者一个整数值。

![](https://upload-images.jianshu.io/upload_images/14654442-9a02fb2da5d9072e.png?imageMogr2/auto-orient/strip|imageView2/2/w/886/format/webp)

![](https://upload-images.jianshu.io/upload_images/14654442-dacfa515b14caabf.png?imageMogr2/auto-orient/strip|imageView2/2/w/498/format/webp)

        entry:每个节点组成如图。previous_entry_length保存前一个节点的长度，遍历时可根据定位到前一个节点。encoding存储
        content的类型和长度。content保存节点的内容

    2、使用字典和跳跃表

```
    typedef struct zset{
         //跳跃表
         zskiplist *zsl;
         //字典
         dict *dice;
    } zset;

```

    字典的键保存元素的值，字典的值则保存元素的分值；跳跃表节点的 object 属性保存元素的值，跳跃表节点的 score 属性保存元素的分值。

    为什么不直接用跳跃表?
        假如我们单独使用 字典，虽然能以 O(1) 的时间复杂度查找成员的分值，但是因为字典是以无序的方式来保存集合元素，所以每次进行
        范围操作的时候都要进行排序；假如我们单独使用跳跃表来实现，虽然能执行范围操作，但是查找操作有 O(1)的复杂度变为了O(logN)。
        因此Redis使用了两种数据结构来共同实现有序集合。

    字典：
        字典底层实现是哈希表，字典有两个哈希表，一个在扩容时使用，哈希表扩容使用渐进式扩容，发送扩容时需要在两个哈希表中进行搜索。

![](https://upload-images.jianshu.io/upload_images/14654442-6d91d250e181280d.png?imageMogr2/auto-orient/strip|imageView2/2/w/804/format/webp)

        发生哈希冲突时使用链地址法解决


    跳跃表:
        跳跃表（skiplist）是一种有序数据结构，它通过在每个节点中维持多个指向其它节点的指针，从而达到快速访问节点的目的。

![](https://upload-images.jianshu.io/upload_images/14654442-4fcade61a3a6193c.png?imageMogr2/auto-orient/strip|imageView2/2/w/940/format/webp)

    由很多层结构组成
        每一层都是一个有序的链表，排列顺序为由高层到底层，都至少包含两个链表节点，分别是前面的head节点和后面的nil节
        最底层的链表包含了所有的元素
        如果一个元素出现在某一层的链表中，那么在该层之下的链表也全都会出现（上一层的元素是当前层的元素的子集）
        链表中的每个节点都包含两个指针，一个指向同一层的下一个链表节点，另一个指向下一层的同一个链表节点；

        