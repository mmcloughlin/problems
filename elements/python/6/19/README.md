Sampling online data.

You are processing a stream of unknown length. At each new element you need to
be able to return a random sample of size `k`.

---

Suppose we receive our `n`th item `x`. There are `nCk` size `k` subsets of the
items seen so far (including `x`). Among those `(n-1)C(k-1)` contain `x`. Now

```
(n-1)C(k-1) / nCk = ((n-1)! / ((k-1)! * (n-1-k+1)!)) * ((k! * (n-k)!) / n!)
                  = k / n
```

That is, proportion `k/n` of them contain `x`.

Therefore the algorithm is to keep a running sample of size `k`. When we see
the `n`th item we add it to the sample with probability `k/n`. If we add it to
the sample we evict a random element.
