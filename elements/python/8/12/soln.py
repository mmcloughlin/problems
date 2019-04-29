class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def splitmod(n, k):
    """
    Split n into k lists containing the elements of n in positions i (mod k).
    Return the heads of the lists and the tails.
    """
    heads = [None]*k
    tails = [None]*k
    i = 0
    while n is not None:
        if heads[i] is None:
            heads[i] = n
        if tails[i] is not None:
            tails[i].next = n
        tails[i] = n
        n.next, n = None, n.next
        i = (i+1)%k
    return heads, tails


def evenoddmerge(n):
    """
    Rearrange n so that all even nodes appear first, then all odd nodes.
    """
    heads, tails = splitmod(n, 2)
    tails[0].next = heads[1]
    return heads[0]


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
    n = from_list(range(11))
    display(n)
    evenoddmerge(n)
    display(n)


def main():
    test()


if __name__ == '__main__':
    main()
