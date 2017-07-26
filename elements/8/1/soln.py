class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


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


def merge(l0, l1):
    """
    merge merges two sorted linked lists.
    """
    if l0 is None:
        return l1
    if l1 is None:
        return l2
    if l0.x > l1.x:
        l0, l1 = l1, l0
    l0.next = merge(l0.next, l1)
    return l0


def test():
    l0 = from_list([2, 5, 7])
    l1 = from_list([3, 11])
    display(l0)
    display(l1)
    m = merge(l0, l1)
    display(m)


def main():
    test()


if __name__ == '__main__':
    main()
