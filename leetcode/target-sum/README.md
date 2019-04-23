# [Target Sum](https://leetcode.com/problems/target-sum/)
<img src="https://img.shields.io/badge/difficulty-medium-orange.svg?style=flat-square" alt="Medium" />

<p>
You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols <code>+</code> and <code>-</code>. For each integer, you should choose one from <code>+</code> and <code>-</code> as its new symbol.
</p> 

<p>Find out how many ways to assign symbols to make sum of integers equal to target S.  
</p>

<p><b>Example 1:</b><br />
<pre>
<b>Input:</b> nums is [1, 1, 1, 1, 1], S is 3. 
<b>Output:</b> 5
<b>Explanation:</b> 

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
</pre>
</p>

<p><b>Note:</b><br>
<ol>
<li>The length of the given array is positive and will not exceed 20. </li>
<li>The sum of elements in the given array will not exceed 1000.</li>
<li>Your output answer is guaranteed to be fitted in a 32-bit integer.</li>
</ol>
</p>
