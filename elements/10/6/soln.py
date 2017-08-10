class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r


def paths_sum(n, acc=0):
    if n is None:
        return 0
    acc = (acc << 1) | n.value
    if n.left is None and n.right is None:
        return acc
    return paths_sum(n.left, acc) + paths_sum(n.right, acc)


def example():
    """
    Construct example tree in figure 10.4.
    """
    p = Node(0)
    o = Node(0, r=p)
    n = Node(0)
    m = Node(1)
    l = Node(1, r=m)
    k = Node(0, l=l, r=n)
    j = Node(0, r=k)
    i = Node(1, l=j, r=o)
    h = Node(0)
    g = Node(1, l=h)
    f = Node(1, r=g)
    e = Node(1)
    d = Node(0)
    c = Node(0, l=d, r=e)
    b = Node(0, l=c, r=f)
    a = Node(1, l=b, r=i)
    return a


def test():
    a = example()
    path_numbers = [
        0b1000,
        0b1001,
        0b10110,
        0b110011,
        0b11000,
        0b1100,
    ]
    assert paths_sum(a) == sum(path_numbers)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
