import operator


class Stack(object):
    def __init__(self):
        self.stack = []

    def push(self, x):
        self.stack.append(x)

    def pop(self):
        return self.stack.pop(-1)


OPS = {
        '+': operator.add,
        '-': operator.sub,
        '*': operator.mul,
        'x': operator.mul,
        '/': operator.div,
        }


def evaluate_rpn(expr):
    parts = expr.split(',')
    stack = Stack()
    for part in parts:
        if part not in OPS:
            stack.push(int(part))
            continue
        op = OPS[part]
        b = stack.pop()
        a = stack.pop()
        stack.push(op(a, b))
    return stack.pop()


def test():
    assert 6 == evaluate_rpn('2,3,*')
    assert 15 == evaluate_rpn('3,4,+,2,*,1,+')
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
