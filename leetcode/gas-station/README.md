# [Gas Station](https://leetcode.com/problems/gas-station/)
<img src="https://img.shields.io/badge/difficulty-medium-orange.svg?style=flat-square" alt="Medium" />

<p>There are <em>N</em> gas stations along a circular route, where the amount of gas at station <em>i</em> is <code>gas[i]</code>.</p>

<p>You have a car with an unlimited gas tank and it costs <code>cost[i]</code> of gas to travel from station <em>i</em> to its next station (<em>i</em>+1). You begin the journey with an empty tank at one of the gas stations.</p>

<p>Return the starting gas station&#39;s index if you can travel around the circuit once in the clockwise direction, otherwise return -1.</p>

<p><strong>Note:</strong></p>

<ul>
	<li>If there exists a&nbsp;solution, it is guaranteed to be unique.</li>
	<li>Both input arrays are non-empty and have the same length.</li>
	<li>Each element in the input arrays is a non-negative integer.</li>
</ul>

<p><strong>Example 1:</strong></p>

<pre>
<strong>Input:</strong> 
gas  = [1,2,3,4,5]
cost = [3,4,5,1,2]

<strong>Output:</strong> 3

<strong>Explanation:
</strong>Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.
</pre>

<p><strong>Example 2:</strong></p>

<pre>
<strong>Input:</strong> 
gas  = [2,3,4]
cost = [3,4,3]

<strong>Output:</strong> -1

<strong>Explanation:
</strong>You can&#39;t start at station 0 or 1, as there is not enough gas to travel to the next station.
Let&#39;s start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can&#39;t travel around the circuit once no matter where you start.
</pre>

