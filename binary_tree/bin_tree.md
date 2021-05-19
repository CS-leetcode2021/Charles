-   先刷二叉树的题目，因为很多经典算法，以及我们前文讲过的所有回溯、动归、分治算法，其实都是树的问题，而树的问题就永远逃不开树的递归遍历框架这几行破代码：
-   **二叉树的难点问题就是，如何把题目的要求细化成每个节点需要做的事情**

```
/* 二叉树遍历框架 */
func traverse(root TreeNode) {
    // 前序遍历
    traverse(root.left)
    // 中序遍历
    traverse(root.right)
    // 后序遍历
}
```

---
### 笔记

-   快速排序就是二叉树的前序遍历
    
```
void sort(int[] nums, int lo, int hi) {
    /****** 前序遍历位置 ******/
    // 通过交换元素构建分界点 p
    int p = partition(nums, lo, hi);
    /************************/

    sort(nums, lo, p - 1);
    sort(nums, p + 1, hi);
}
```
-   归并排序就是二叉树的后序遍历
```
void sort(int[] nums, int lo, int hi) {
    int mid = (lo + hi) / 2;
    sort(nums, lo, mid);
    sort(nums, mid + 1, hi);

    /****** 后序遍历位置 ******/
    // 合并两个排好序的子数组
    merge(nums, lo, mid, hi);
    /************************/
}
```

-   写递归算法的秘诀---计算树的节点

```
// 定义：count(root) 返回以 root 为根的树有多少节点
int count(TreeNode root) {
    // base case
    if (root == null) return 0;
    // 自己加上子树的节点数就是整棵树的节点数
    return 1 + count(root.left) + count(root.right);
}
```

-   [实践_翻转二叉树](./invertTree.go)

---
### 二叉树的四种遍历框架

-   前序遍历

```go
    func traverse(root TreeNode){
    	if root == nil{
    		return 
        }   
        
        res := []int{}
        
        res = append(res,root.val)
        
        traverse(root.left)
        traverse(root.right)
    }

```

-   后序遍历

```go
    func traverse(root TreeNode){
        if root == nil{
          return
        }
        
        res := []int{}
        
        traverse(root.left)
        traverse(root.right)

        res = append(res,root.val)
    }
```

-   中序遍历

```go
    func traverse(root TreeNode){
        if root == nil{
            return
        }
        
        res := []int{}
        
        traverse(root.left)
		res = append(res,root.val)
        traverse(root.right)
    }    
```

-   层次遍历

```go
    func traverse(root TreeNode){
        if root == nil{
            return
        }
        
        /*需要解除队列的形式进行访问，q*/
        q = append(q,root)
        for q !=nil {
            cur := q[0]
            
            /*输出*/
            
            if cur.left!=nil{
            	q = append(q,cur.left)
            }
            
            if cur.right!=nil{
            	q= append(q,cur.right)
            }
            q = q[1:]
            
        }
    }    
```