import random


class Stack(object):
    def __init__(self):
        self.stack = []

    def push(self, x):
        self.stack.append(x)

    def pop(self):
        if len(self.stack) == 0:
            return None
        return self.stack.pop(-1)

    def size(self):
        return len(self.stack)


class Queue(object):
    def __init__(self):
        self.inbound = Stack()
        self.outbound = Stack()

    def push(self, x):
        self.inbound.push(x)

    def pop(self):
        v = self.outbound.pop()
        if v is not None:
            return v
        while self.inbound.size() > 0:
            self.outbound.push(self.inbound.pop())
        return self.outbound.pop()


def test(iterations=1000):
    a = []
    q = Queue()
    for _ in xrange(iterations):
        # push some items
        n = random.randrange(1, 100)
        add = [random.randrange(100000) for _ in xrange(n)]
        a.extend(add)
        for x in add:
            q.push(x)

        # consume some number from the queue
        m = random.randrange(1, len(a)+1)
        for _ in xrange(m):
            v = q.pop()
            assert v == a.pop(0)
    print 'pass'


def main():
    test(10000)


if __name__ == '__main__':
    main()
