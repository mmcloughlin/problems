class Stack(object):
    def __init__(self):
        self.stack = []
        self.mx = []
        self.n = 0

    def push(self, x):
        self.stack.append(x)
        mx = self.mx[-1] if self.n > 0 else float('-inf')
        self.mx.append(max(mx, x))
        self.n += 1

    def pop(self):
        if self.n == 0:
            return None
        v = self.stack.pop(-1)
        self.mx = self.mx[:-1]
        self.n -= 1
        return v

    def max(self):
        if self.n == 0:
            return None
        return self.mx[-1]


def test():
    s = Stack()
    assert s.max() is None
    s.push(1)
    assert 1 == s.max()
    s.push(4)
    assert 4 == s.max()
    s.push(2)
    assert 4 == s.max()
    s.pop()
    s.pop()
    s.push(2)
    assert 2 == s.max()
    print 'pass'


if __name__ == '__main__':
    test()
