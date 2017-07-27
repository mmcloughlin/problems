class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def reverse_sublist(n, s, f):
    """
    Reverse the sublist between positions s and f in-place. Indicies are 1-up
    and inclusive.
    """
    start = n

    # advance to the start of the sublist
    # pre is the node before the sublist
    pre = None
    for _ in xrange(s-1):
        pre = n
        n = n.next
    first = n

    # reverse the sublist itself
    prev = None
    for _ in xrange(f - s + 1):
        nxt = n.next
        n.next = prev
        prev = n
        n = nxt

    # wire up the ends
    first.next = n
    if pre is None:
        return prev
    pre.next = prev
    return start


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
    n = from_list(range(1, 11))
    display(n)

    n = reverse_sublist(n, 4, 7)
    display(n)

    n = reverse_sublist(n, 1, 2)
    display(n)


def main():
    test()


if __name__ == '__main__':
    main()
