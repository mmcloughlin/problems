class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def rev(n):
    """
    rev reverses the linked list n.
    """
    prev = None
    while n is not None:
        nxt = n.next
        n.next = prev # point n backwards
        prev = n # update prev
        n = nxt # update n
    return prev


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
    r = rev(n)
    display(r)


def main():
    test()


if __name__ == '__main__':
    main()
