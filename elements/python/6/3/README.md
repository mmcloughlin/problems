Multiply two integers provided as (possibly negative) strings.

---

```
a = SUM 10^i a_i
b = SUM 10^j b_j

a*b = SUM_{i,j} 10^{i+j} a_i b_j
    = SUM_k 10^k SUM_{k = i+j} a_i b_j
```
