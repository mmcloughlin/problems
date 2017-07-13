The Open Doors Problem.

Suppose there are 500 doors numbered 1, 2, ..., 500.

Somebody comes through and opens every door. Then someone comes through and
closes every other door. Person `i` comes through and toggles the state of the
`i`th door.

Which doors are open after the 500th person has walked through?

---

Door `n` will have been flipped by every divisor of `n`. The number of
divisors of `n` is the divisor function `d(n)`.

Now in terms of the prime factors of n:

  d(p^k) = k+1

and in general

  d(p0^k0 p1^k1 ... pm^km) = (k0 + 1)(k1 + 1)...(km+1)

Since the door starts closed, it will finish open if it has an odd number of
divisors. By the formula above, it finishes open when all the `ki` are even.
In other words, if its square.
