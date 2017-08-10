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


def depth(n):
    """
    depth computes the depth of node n.
    """
    d = 0
    while n is not None:
        n = n.parent
        d += 1
    return d


def lca(a, b):
    """
    Computes the least common ancestor of a and b.
    """
    da = depth(a)
    db = depth(b)
    # ensure a is the closest to the root
    if da > db:
        a, b = b, a
        da, db = db, da
    # step b until it is at the same level as a
    while db > da:
        b = b.parent
        db -= 1
    # step them both until collision
    while a != b:
        a = a.parent
        b = b.parent
    return a


def test(trials=1000):
    m = Node('m')
    n = Node('n')
    r = Node('a',
            l=Node('b',
                l=Node('c', l=Node('d'), r=Node('e')),
                r=Node('f', r=Node('g', l=Node('h'))),
                ),
            r=Node('i',
                l=Node('j', r=Node('k', l=Node('l', r=m), r=n)),
                r=Node('o', r=Node('p')),
                ),
            )
    ancestor = lca(m, n)
    assert 'k' == ancestor.value
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
