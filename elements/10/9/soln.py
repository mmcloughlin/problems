class Node(object):
    def __init__(self, v, p=None, l=None, r=None):
        self.value = v
        self.parent = p
        self.left = l
        if l is not None:
            l.parent = self
        self.right = r
        if r is not None:
            r.parent = self

    def leaf(self):
        return self.left is None and self.right is None


def in_order_recursive(t):
    if t is None:
        return []
    return in_order_recursive(t.left) + [t.value] + in_order_recursive(t.right)


def in_order(t):
    prv = None
    cur = t
    while cur is not None:
        if cur.leaf():
            yield cur.value
            nxt = cur.parent
        elif prv == cur.parent:
            if cur.left is None:
                yield cur.value
                nxt = cur.right
            else:
                nxt = cur.left
        elif prv == cur.left:
            yield cur.value
            nxt = cur.parent if cur.right is None else cur.right
        elif prv == cur.right:
            nxt = cur.parent
        prv, cur = cur, nxt


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
    expect = 'dcebfhgajlmkniop'
    assert expect == ''.join(in_order_recursive(r))
    assert expect == ''.join(in_order(r))
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
