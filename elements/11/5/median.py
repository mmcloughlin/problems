import heapq


class Streaming(object):
    """
    Streaming provdes capaility to build up a set of numbers one-by-one,
    and is always able to return the current median.

    The implementation maintains two heaps for the bottom and top halfs of the
    dataset.
    """

    def __init__(self):
        # max heap for bottom half (note heapq library supports min heaps
        # only, so we store negative numbers here to get sorting correctly)
        self.lo = []
        # min heap for top half
        self.hi = []

    def size(self):
        "Return number of points in the dataset."
        return len(self.lo) + len(self.hi)

    def _xfer(self, fr, to):
        "Move an item from one heap to the other"
        y = heapq.heappop(fr)
        heapq.heappush(to, -y)

    def _check_invariants(self):
        # invariant: len(hi)+1 >= len(lo) >= len(hi)
        assert ((len(self.lo) == len(self.hi)) or
                (len(self.lo) == len(self.hi)+1))

        # invariant: everything in lo is <= everything in hi
        if len(self.lo) > 0 and len(self.hi) > 0:
            assert -self.lo[0] <= self.hi[0]

    def add(self, x):
        "Add x to the dataset and return the median."

        self._check_invariants()

        # insert x into the correct half
        if len(self.lo) == 0 or x <= -self.lo[0]:
            heapq.heappush(self.lo, -x)
        else:
            heapq.heappush(self.hi, x)

        # rebalance heaps
        if len(self.lo) < len(self.hi):
            self._xfer(self.hi, self.lo)
        elif len(self.hi) + 2 == len(self.lo):
            self._xfer(self.lo, self.hi)

        # now compute median
        return self.median()

    def median(self):
        "Return median of data set."
        self._check_invariants()

        if self.size() == 0:
            return None

        if len(self.lo) > len(self.hi):
            return -self.lo[0]

        return (-self.lo[0] + self.hi[0]) / 2.0


def naive(X):
    "Compute median via a naive method (for testing)"
    if len(X) == 0:
        return None
    X.sort()
    half = len(X) // 2
    if len(X) % 2 == 1:
        return X[half]
    else:
        return (X[half] + X[half - 1]) / 2.0
