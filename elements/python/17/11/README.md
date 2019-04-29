The maximum sub-interval sum problem asks for the interval of an array with
the largest sum.

Problem 17.11 asks for the largest possible sum, allowing for wrap around.

In this solution we consier wrap around by looking at the complement of the
interval, which will be a regular sub-interval. Searching for the largest
wrapped around interval is the same as looking for a minimal sub-interval.
