Implement in-order traversal with `O(1)` space. Nodes have parent pointers.

---

Track previous node in the traversal.

```
at a leaf         visit node
last == parent    traverse left
last == left      visit node, traverse right
last == right     traverse left
```

