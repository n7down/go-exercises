# Minimum Depth

Given a binary tree, find its minimum depth. The minimum depth is the total number of nodes along the shortest path from the root node down to the nearest leaf node.

For example the minimum depth of the following binary tree is 3. The shortest path is `1 -> 3 -> 6`

```
        1
      /   \
    /       \
    2       3
  /   \   /   \
 4    5  6    7
  \    \     / \
  8     9   10  11
   \
   12
```

The idea is to traverse the binary tree in a bottom-up manner using postorder traversal and calculate the minimum depth of left and right subtrees for each encountered node. The minimum depth of the subtree rooted at any node is one more then the minimum depth of its left and right subtree. If either left or right subtree does not exist for a node, consider the minimum depth returned by other subtree.
