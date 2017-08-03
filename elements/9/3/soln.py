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


PAIRS = {
        '}': '{',
        ']': '[',
        ')': '(',
        }


def is_well_formed(s):
    stack = Stack()
    for ch in s:
        if ch in PAIRS:
            expect = PAIRS[ch]
            matching = stack.pop()
            if matching != expect:
                return False
        else:
            stack.push(ch)
    return stack.size() == 0


def test():
    assert is_well_formed('([]){()}')
    assert is_well_formed('[()[]{()()}]')
    assert not is_well_formed('{)')
    assert not is_well_formed('[()[]{()()')
    print 'pass'


if __name__ == '__main__':
    test()
