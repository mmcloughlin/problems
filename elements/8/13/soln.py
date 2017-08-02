class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def ziplist(n):
    # find the middle
    m = n
    e = n
    while e is not None:
        e = e.next
        if e is None:
            break
        e = e.next
        m = m.next

    # end the first half
    m.next, m = None, m.next

    # reverse from here
    prev = None
    while m is not None:
        nxt = m.next
        m.next = prev
        prev = m
        m = nxt

    # now we have the head (h) and tail (t)
    # interleave them
    h = n
    t = prev
    while h is not None and t is not None:
        hn = h.next
        tn = t.next
        h.next = t
        t.next = hn
        h = hn
        t = tn

    return n



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
    n = from_list(range(5))
    display(n)
    ziplist(n)
    display(n)
    assert to_list(n) == [0, 4, 1, 3, 2]


def main():
    test()


if __name__ == '__main__':
    main()
