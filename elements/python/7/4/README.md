Reverse the words in a string. You do not need to preseve the original string.

For example "Alice likes Bob" becomes "Bob likes Alice".

Hint: difficult to so this in one pass.

---

Trivial if we can use O(n) extra space. Collect the words in an array. Write
them back out in reverse order.

Perhaps we could reverse the whole string and then reverse individual words.
