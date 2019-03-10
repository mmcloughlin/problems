# [Continuous Subarray Sum](https://leetcode.com/problems/continuous-subarray-sum/)
<img src="https://img.shields.io/badge/difficulty-medium-orange.svg?style=flat-square" alt="Medium" />

<p>
Given a list of <b>non-negative</b> numbers and a target <b>integer</b> k, write a function to check if the array has a continuous subarray of size at least 2 that sums up to the multiple of <b>k</b>, that is, sums up to n*k where n is also an <b>integer</b>.
</p>


<p><b>Example 1:</b><br />
<pre>
<b>Input:</b> [23, 2, 4, 6, 7],  k=6
<b>Output:</b> True
<b>Explanation:</b> Because [2, 4] is a continuous subarray of size 2 and sums up to 6.
</pre>
</p>


<p><b>Example 2:</b><br />
<pre>
<b>Input:</b> [23, 2, 6, 4, 7],  k=6
<b>Output:</b> True
<b>Explanation:</b> Because [23, 2, 6, 4, 7] is an continuous subarray of size 5 and sums up to 42.
</pre>
</p>

<p><b>Note:</b><br>
<ol>
<li>The length of the array won't exceed 10,000.</li>
<li>You may assume the sum of all the numbers is in the range of a signed 32-bit integer.</li>
</ol>
</p>
