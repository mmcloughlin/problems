Sort a stack in descending order (largest at the top). You may use `push`,
`pop`, `top` and `empty`. You may not *explicitly* allocate more memory.

---

I personally found the point about allocating memory to be misleading. The
*explicitly* word is key. The solution does require `O(n)` storage, but they
are hinting at a recursive approach which will *implicitly* allocate that
memory on the stack. The question is contrived. How about instead "Sort a
stack using an auxilliary stack?"
