class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def circular_shift_right(n, k):
    """
    Circularly right shift the linked list starting at n by k.
    """
    s = n # s is the start of the input list
    t = n # t will be a "trailing" pointer lagging k steps behind
    # step n ahead k times
    for _ in xrange(k):
        n = n.next
    # step n and k until n is the last node in the list
    while n.next is not None:
        t = t.next
        n = n.next
    # stitch it back up
    new_start = t.next
    n.next = s
    t.next = None
    return new_start


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
    n = from_list(range(10))
    display(n)
    n = circular_shift_right(n, 3)
    display(n)


def main():
    test()


if __name__ == '__main__':
    main()
