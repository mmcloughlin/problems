import random


class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def insertion_sort(n):
    s = Node(float('-inf'), None)
    while n is not None:
        p = s
        while p.next is not None and p.next.x < n.x:
            p = p.next
        nxt = n.next
        n.next = p.next
        p.next = n
        n = nxt
    return s.next


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


def test(trials=1000):
    for _ in xrange(trials):
        A = [random.randrange(1000) for _ in xrange(16)]
        n = from_list(A)
        s = insertion_sort(n)
        assert sorted(A) == to_list(s)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
