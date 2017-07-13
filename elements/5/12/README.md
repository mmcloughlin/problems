Given a function that can generate a random bit, write a function that returns
a random integer in the range `[a, b]`.

---

I see multiple solutions.

First, use the generator to approximate generating a uniform float in the
range `[0, 1)`. Pick some number of bits `n`, generate an `n`-bit number `x`
and let your uniform float be `u = x/2^n`. Then your integer is `a +
floor((b-a+1)*u)`. You have to pick some precision `n`, the larger the better.
You always have to generate the same number of bits.

A similar second option is to conceptually divide the unit interval into
`b-a+1` equal sized sub-intervals. Then use the bit generator to do a "binary
search" like descent. When the sub-interval your binary search has narrowed to
is completely contained in an interval for one of the sub intervals
corresponding to a result in `[a, b]`, you can stop. Unlike the previous case,
you will often be able to stop before generating as many bits. However it has
infinite worst case performance. For example to pick between three things
there is always a power-of-two interval that contains the boundary `1/3`.

In the third option, we note that we are choosing between `m = b-a+1` things.
Let `n` be minimal with `m <= 2^n`. Now we keep on generating `n`-bit numbers
`x`. If `x < m` we have an answer. Otherwise try again. Again this has
infinite worst case performance. If we're trying to choose between three
things (m=3) then we keep on generating 2-bit numbers, and try again whenever
we get 3. If the generator always gives us ones, we never terminate.

The book seems to use the third option. Presumably we assume the infinite
worst case will never actually happen!
