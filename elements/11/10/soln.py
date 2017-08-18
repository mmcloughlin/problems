import heapq
import random


class Stack(object):
    def __init__(self):
        self.h = []

    def push(self, x):
        heapq.heappush(self.h, (-len(self.h), x))

    def pop(self):
        if self.empty():
            return None
        _, x = heapq.heappop(self.h)
        return x

    def empty(self):
        return len(self.h) == 0


def test():
    items = [random.randrange(100) for _ in xrange(100)]
    s = Stack()
    for x in items:
        s.push(x)
    for x in reversed(items):
        assert x == s.pop()
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
