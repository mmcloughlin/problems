Find the first occurence of a substring.

---

Use a hash-based approach. Given a string `s = c(0)c(1)...c(n-1)` define a
running hash by

```
h(0) = 0
h(i+1) = a*h(i) + c(i) (mod m)
```

For some multiplier `a` and modulus `m`.  So the hash of `s` is `h(n)`. We
compute the target hash in advance. Then as we progress through the string we
compute the hash as we go. Note that this hash is really

```
K = H(x(0)x(1)...x(n-1)) = a^(n-1)x(0) + a^(n-2)x(1) + ... + x(n-1)
```

We can step the hash from one to the next with

```
aK - a^n x(0) + x(n)
```
