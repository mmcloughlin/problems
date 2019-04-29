class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def delete(n):
    """
    Remove the given node from the linked list it belongs to. It is not the
    last, and values may be copied.
    """
    n.x = n.next.x
    n.next = n.next.next


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
    s = from_list(range(10))
    display(s)

    # delete the third
    n = s.next.next.next
    delete(n)

    display(s)


def main():
    test()


if __name__ == '__main__':
    main()
