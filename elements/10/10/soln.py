class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r


def preorder(t):
    stack = [t]
    order = []
    while len(stack) > 0:
        n = stack.pop(-1)
        if n is None:
            continue
        order.append(n.value)
        stack.append(n.right)
        stack.append(n.left)
    return order


def postorder(t):
    stack = [(t, True)]
    order = []
    while len(stack) > 0:
        n, first = stack.pop(-1)
        if n is None:
            continue
        if not first:
            order.append(n.value)
            continue
        stack.append((n, False))
        stack.append((n.right, True))
        stack.append((n.left, True))
    return order


def test():
    r = Node('a',
            l=Node('b',
                l=Node('c', l=Node('d'), r=Node('e')),
                r=Node('f', r=Node('g', l=Node('h'))),
                ),
            r=Node('i',
                l=Node('j', r=Node('k', l=Node('l', r=Node('m')), r=Node('n'))),
                r=Node('o', r=Node('p')),
                ),
            )

    assert 'abcdefghijklmnop' == ''.join(preorder(r))
    assert 'dechgfbmlnkjpoia' == ''.join(postorder(r))

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
