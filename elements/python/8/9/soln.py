class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def delete_kth_last(n, k):
    """
    Delete the kth last element from the linked list n. The index k is 0-up,
    so k=0 will remove the last element.
    """
    # step ahead k+2 times and pick up a "trailing" node
    t = n
    for _ in xrange(k+2):
        n = n.next
    # step ahead to the end, moving the trailing node in lock step
    while n is not None:
        n = n.next
        t = t.next
    # now t points at the node to be removed
    t.next = t.next.next


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

    delete_kth_last(n, 2)
    display(n)


def main():
    test()


if __name__ == '__main__':
    main()
