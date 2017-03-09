# Problem

How many 10-digit numbers don't start with 0 AND don't contain "911"?

# Approach

Let `a(n, k)` be the number of n-digit numbers that contain "911" k times.

Note: at this point leading digits are allowed to be 0.

So, what can we say about a(n, k)? If we sum over k, we must get all possible
n-digit strings, so:

  SUM{k=0,1,...} a(n, k) = 10^n

What about a(n,k) for k>1? If an n-digit string has at least one "911" in it,
where is it? To avoid double counting problems, focus on the *first
occurance*. Then from left-to-right the digit-string looks like

  (i-digit string not containing "911") + "911" + (j digit string containing k-1 "911"s)

where i + 3 + j = n. The number of these is

  a(i, 0) * a(j, k-1)

Therefore in total

  a(n, k) = SUM{i+j = n-3} a(i,0) * a(j, k-1)
