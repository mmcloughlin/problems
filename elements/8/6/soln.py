class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def traverse_list(n):
    """
    Traverse the whole list. Return the last node and the number of steps to
    the last node.
    """
    steps = 0
    while n.next is not None:
        n = n.next
        steps += 1
    return n, steps


def last_node(n):
    last, _ = traverse_list(n)
    return last


def lists_overlap(l1, l2):
    t1, n1 = traverse_list(l1)
    t2, n2 = traverse_list(l2)
    if t1 != t2:
        return None
    # ensure l1 is the shorter
    if n1 > n2:
        t1, t2 = t2, t1
        n1, n2 = n2, n1
    # step l2 until it's as far away from the end as l1
    while n2 > n1:
        l2 = l2.next
        n2 -= 1
    # step them together until they collide
    while l1 != l2:
        l1 = l1.next
        l2 = l2.next
    return l1


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
    l1 = from_list(range(10))
    l2 = from_list(range(50, 55))
    e2 = last_node(l2)
    e2.next = l1.next.next.next.next

    display(l1)
    display(l2)

    c = lists_overlap(l1, l2)
    assert c is not None
    assert 4 == c.x

    l1 = from_list(range(10))
    l2 = from_list(range(10))

    assert not lists_overlap(l1, l2)

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
