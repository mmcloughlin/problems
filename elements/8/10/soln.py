class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def dedup(n):
    """
    Remove duplicates from n, which is assumed to be sorted.
    """
    s = n
    while n is not None:
        if n.x != s.x:
            s.next = n
            s = n
        n = n.next
    s.next = None


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
    n = from_list([1,1,1,2,2,2,3,3,4,5,5,6,7,7])
    display(n)
    dedup(n)
    display(n)


def main():
    test()


if __name__ == '__main__':
    main()
