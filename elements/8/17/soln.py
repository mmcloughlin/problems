import random


class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def pivot(n, k):
    heads = [None]*3
    tails = [None]*3

    # build sections independently
    while n is not None:
        i = 1+cmp(n.x, k)
        if heads[i] is None:
            heads[i] = n
        if tails[i] is not None:
            tails[i].next = n
        tails[i] = n
        n.next, n = None, n.next

    # stitch them up
    head = None
    tail = None
    for i, h in enumerate(heads):
        if h is not None:
            head = h
            tail = tails[i]
            break
    for j in xrange(i+1, 3):
        if heads[j] is None:
            continue
        tail.next = heads[j]
        tail = tails[j]

    return head


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
    L = [random.randrange(5) for _ in xrange(10)]
    n = from_list(L)
    display(n)
    p = pivot(n, 2)
    display(p)


def main():
    test()


if __name__ == '__main__':
    main()
