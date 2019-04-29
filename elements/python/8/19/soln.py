import random


class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def add_lists(n0, n1):
    carry = 0
    cur = base = Node(0, None)
    while n0 is not None or n1 is not None or carry > 0:
        d0 = n0.x if n0 is not None else 0
        d1 = n1.x if n1 is not None else 0
        s = d0 + d1 + carry

        nxt = Node(s % 10, None)
        cur.next = nxt

        cur = nxt
        carry = s/10

        n0 = n0.next if n0 is not None else None
        n1 = n1.next if n1 is not None else None
    return base.next



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


def int_to_list(x):
    digits = []
    while x != 0:
        digits.append(x%10)
        x /= 10
    return from_list(digits)


def int_from_list(n):
    x = 0
    b = 1
    while n is not None:
        x += b*n.x
        b *= 10
        n = n.next
    return x


def test(trials=1000):
    for _ in xrange(trials):
        x = random.randrange(10000)
        y = random.randrange(100000)
        n0 = int_to_list(x)
        n1 = int_to_list(y)
        s = add_lists(n0, n1)
        assert x+y == int_from_list(s)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
