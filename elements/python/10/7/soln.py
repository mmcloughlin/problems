class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r


def has_path_with_sum(n, target):
    if n is None:
        return target == 0
    target -= n.value
    return has_path_with_sum(n.left, target) or has_path_with_sum(n.right, target)


def example():
    """
    Construct example tree in figure 10.4.
    """
    p = Node(28)
    o = Node(271, r=p)
    n = Node(257)
    m = Node(641)
    l = Node(401, r=m)
    k = Node(1, l=l, r=n)
    j = Node(2, r=k)
    i = Node(6, l=j, r=o)
    h = Node(17)
    g = Node(3, l=h)
    f = Node(561, r=g)
    e = Node(0)
    d = Node(28)
    c = Node(271, l=d, r=e)
    b = Node(6, l=c, r=f)
    a = Node(314, l=b, r=i)
    return a


def test():
    n = example()
    assert has_path_with_sum(n, 314+6+271+0)
    assert has_path_with_sum(n, 314+6+2+1+401+641)
    assert not has_path_with_sum(n, 13)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
