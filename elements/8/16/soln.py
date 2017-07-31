class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def find_median_circular(n):
    # find the start: either all nodes are identical, or there is a point
    # where the value drops
    p = n
    s = n.next
    while s != n and s.x >= p.x:
        p = p.next
        s = s.next
    # if we looped all the way around, all points are identical
    if s == n and s.x >= p.x:
        return s.x
    # otherwise s is the first node in the sort order
    # find the middle, with a slow and fast pointer
    m = s
    e = s
    steps = 0
    while True:
        e = e.next
        steps += 1
        if e == s: break
        e = e.next
        steps += 1
        if e == s: break
        m = m.next

    if steps % 2 == 1:
        return m.x

    return (m.x + m.next.x)/2.0


def last_node(n):
    while n.next is not None:
        n = n.next
    return n


def make_circular(n):
    last_node(n).next = n


def to_list(n):
    """
    to_list converts the linked list n to a list.
    """
    L = []
    while n is not None:
        L.append(n.x)
        n = n.next
    return L


def from_list(L):
    """
    from_list builds a linked list from the given list.
    """
    n = None
    for i in xrange(len(L)-1, -1, -1):
        n = Node(x=L[i], nxt=n)
    return n


def display(n):
    """
    display prints a view of the linked list.
    """
    print ' -> '.join(map(str, to_list(n)))


def test():
    # identical case
    n = from_list([5]*13)
    make_circular(n)
    assert 5 == find_median_circular(n)

    n = from_list(range(5))
    make_circular(n)
    assert 2 == find_median_circular(n)

    n = from_list(range(6))
    make_circular(n)
    assert 2.5 == find_median_circular(n)

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
