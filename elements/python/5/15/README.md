Compute the greatest common divisor of `a` and `b`.

---

Divide `a` by `b` to get divisor and remainder

```
a = d*b + r
```

Observe that any common divisor of a and b is a divisor of r. Likewise a
common divisor of b and r is a divisor of a. In fact

```
(a, b) = (b, r)
```

So the problem can be solved recursively.
