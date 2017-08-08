import random


class Stack(object):
    def __init__(self):
        self.stack = []

    def push(self, x):
        self.stack.append(x)

    def pop(self):
        return self.stack.pop(-1) if not self.empty() else None

    def top(self):
        return self.stack[-1] if not self.empty() else None

    def empty(self):
        return len(self.stack) == 0


def sort_stack(s):
    """
    sort_stack sorts s in descending order.
    """
    if s.empty():
        return
    e = s.pop()
    sort_stack(s)
    insert_sorted_stack(s, e)


def insert_sorted_stack(s, e):
    """
    insert_sorted_stack inserts e into s assuming s is sorted in descending
    order.
    """
    if s.empty() or s.top() <= e:
        s.push(e)
        return
    t = s.pop()
    insert_sorted_stack(s, e)
    s.push(t)


def test(trials=1000):
    for _ in xrange(trials):
        A = [random.randrange(1000) for _ in xrange(10)]
        s = Stack()
        for a in A:
            s.push(a)
        sort_stack(s)
        assert s.stack == sorted(A)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
