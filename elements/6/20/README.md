Weighted random chooser.

---

This is the naive solution. If you are making multiple samples, this can be
optimized by saving the array of cumulative weights and using a binary search
to find the bucket the random value falls into.
