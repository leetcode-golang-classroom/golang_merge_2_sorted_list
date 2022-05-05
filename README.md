# golang_merge_2_sorted_list

You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists in a one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

Return *the head of the merged linked list*.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)

```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

```

**Example 2:**

```
Input: list1 = [], list2 = []
Output: []

```

**Example 3:**

```
Input: list1 = [], list2 = [0]
Output: [0]

```

**Constraints:**

- The number of nodes in both lists is in the range `[0, 50]`.
- `100 <= Node.val <= 100`
- Both `list1` and `list2` are sorted in **non-decreasing** order.

## 解析

要把兩個排序過的單向鏈結合成一個單一排序的鏈結串列

直覺的作法就是把每個 l1, l2 的 node 逐步比對

當遇到比較小的值就把值複製到新的結點放到 merged list

每次把 11, 12 比較小的鏈結往後移動

直到 l1 或 l2 走到尾部 

如果 其中一個走完，把沒走完直接接到目前 merged list 的最後

時間複雜度 O(m+n)

## 程式碼

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var merged *ListNode
    var cur *ListNode
    if l1 == nil && l2 == nil {
        return merged
    }
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    for l1 != nil && l2 != nil {
        if merged == nil {
            if l1.Val < l2.Val {
                merged = &ListNode{Val: l1.Val}
                l1 = l1.Next
            } else {
                merged = &ListNode{Val: l2.Val}
                l2 = l2.Next
            }
            cur = merged
        } else {
            if l1.Val < l2.Val {
                cur.Next = &ListNode{Val: l1.Val}
                l1 = l1.Next
            } else {
                cur.Next = &ListNode{Val: l2.Val}
                l2 = l2.Next
            }
            cur = cur.Next
        }
    }
    if l1 != nil {
        cur.Next = l1
    } else {
        cur.Next = l2
    }
    return merged
}
```

## 困難點

1. 必須知到鏈結串列的走訪機制
2. 處理長度不同的處理機制

## Solve Point

- [x]  Understand what problem would like to solve
- [x]  Analysis Complexity