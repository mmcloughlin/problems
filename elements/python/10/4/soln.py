class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r


def lca(r, a, b):
    """
    lca returns the least common ancestor of a and b in the tree rooted at r.
    """
    pa = path_to(r, a)
    pb = path_to(r, b)
    if pa is None or pb is None:
        return None
    ancestor = None
    for x, y in zip(pa, pb):
        if x != y:
            break
        ancestor = x
    return ancestor


def path_to(r, n):
    """
    path_to finds a path from the root r to a node n, or None if no such path
    exists.
    """
    if r is None:
        return None
    if r == n:
        return [n]
    pl = path_to(r.left, n)
    if pl is not None:
        return [r] + pl
    pr = path_to(r.right, n)
    if pr is not None:
        return [r] + pr
    return None


def test():
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
    ancestor = lca(r, m, n)
    assert 'k' == ancestor.value
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
