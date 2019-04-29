import random


class Queue(object):
    def __init__(self, capacity=4):
        self.items = [None]*capacity
        self.tail = 0
        self.head = 0

    def enqueue(self, x):
        if self.size() == self.cap():
            self._grow()
        idx = self.head % self.cap()
        self.items[idx] = x
        self.head += 1

    def dequeue(self):
        if self.size() == 0:
            return None
        idx = self.tail % self.cap()
        self.tail += 1
        return self.items[idx]

    def _grow(self):
        n = self.cap()
        self.items.extend(self.items)
        self.tail = self.tail % n
        self.head = self.tail + n

    def cap(self):
        return len(self.items)

    def size(self):
        return self.head - self.tail


def test(iterations=1000):
    a = []
    q = Queue()
    for _ in xrange(iterations):
        # push some items
        n = random.randrange(1, 100)
        add = [random.randrange(100000) for _ in xrange(n)]
        a.extend(add)
        for x in add:
            q.enqueue(x)

        # consume some number from the queue
        m = random.randrange(1, len(a)+1)
        for _ in xrange(m):
            v = q.dequeue()
            assert v == a.pop(0)
    print 'pass'


def main():
    test(10000)


if __name__ == '__main__':
    main()
