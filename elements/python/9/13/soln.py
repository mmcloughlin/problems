import random


class MaxQueue(object):
    """
    MaxQueue provides a base queue API and tracks the largest item contained
    in the queue.
    """
    def __init__(self):
        self.items = []
        self.peaks = []

    def enqueue(self, x):
        self.items.append(x)
        while len(self.peaks) > 0 and self.peaks[-1] < x:
            self.peaks.pop(-1)
        self.peaks.append(x)

    def dequeue(self):
        if self.empty():
            return None
        x = self.items.pop(0)
        if x == self.peaks[0]:
            self.peaks.pop(0)
        return x

    def max(self):
        return self.peaks[0] if not self.empty() else None

    def empty(self):
        return len(self.items) == 0


def test(iterations=1000, upper=100):
    a = []
    q = MaxQueue()
    for _ in xrange(iterations):
        # push some items
        n = random.randrange(1, 100)
        add = [random.randrange(upper) for _ in xrange(n)]
        a.extend(add)
        for x in add:
            q.enqueue(x)

        assert q.max() == max(a)

        # consume some number from the queue
        m = random.randrange(1, len(a))
        for _ in xrange(m):
            v = q.dequeue()
            assert v == a.pop(0)

        assert q.max() == max(a)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
