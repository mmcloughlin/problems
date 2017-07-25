Implement Elias gamma encoding of integer sequences.

A single integer is encoded in binary using `n` bits, preceeded by `n-1` 0
bits. So for example 13 is

```
0001101
```

A sequence of integers `A(0), A(1), ..., a(N-1)` is encoded by concatenating
their Elias encodings.

Implement Elias encoding and decoding for integer sequences.
