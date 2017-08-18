Find the `k` largest elements in a max-heap, without modifying the heap.

---

The heap property is

```
A[i] >= A[2i+1], A[2i+2]
```

Example given in the book:

```
0    1    2    3   4    5    6    7   8
561, 314, 401, 28, 156, 359, 271, 11, 3
```




take first item
inspect children, take the larger
