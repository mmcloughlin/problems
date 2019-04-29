class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt

# From problem 8.5.
def find_cycle(n):
    # fast and slow pointers will collide if there is a cycle
    slow = n
    fast = n.next
    while slow != fast:
        slow = slow.next
        fast = fast.next
        if fast is None: return None
        fast = fast.next
        if fast is None: return None

    # find the length of the cycle
    k = 1
    p = slow.next
    while p != slow:
        p = p.next
        k += 1

    # now set off two pointers k apart and stop when they collide
    behind = n
    ahead = n
    for _ in xrange(k):
        ahead = ahead.next

    while behind != ahead:
        behind = behind.next
        ahead = ahead.next

    return ahead


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


# From problem 8.6.
def lists_overlap_nocycles(l1, l2):
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


def lists_overlap(l1, l2):
    c1 = find_cycle(l1)
    c2 = find_cycle(l2)

    # neither is cyclic then defer to the non-cyclic solution
    if c1 is None and c2 is None:
        return lists_overlap_nocycles(l1, l2)

    # if one is cyclic but not the other, they don't overlap
    if c1 is None and c2 is not None:
        return None

    if c1 is not None and c2 is None:
        return None

    # here we know both are cyclic, but they may still be completely disjoint
    # if they collide in the tail, they'll have the same base of the cycle
    if c1 == c2:
        s1 = 0
        p1 = l1
        while p1 != c1:
            s1 += 1
            p1 = p1.next
        s2 = 0
        p2 = l2
        while p2 != c2:
            s2 += 1
            p2 = p2.next
        if s1 > s2:
            l1, l2 = l2, l1
            s1, s2 = s2, s1
        while s2 > s1:
            l2 = l2.next
            s2 -= 1
        while l1 != l2:
            l1 = l1.next
            l2 = l2.next
        return l1

    # otherwise they collide on the cycle itself
    r = c1.next
    while r is not c1:
        if r == c2:
            return r
        r = r.next

    return None


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
    # no cycle case
    l1 = from_list(range(10))
    l2 = from_list(range(50, 55))
    e2 = last_node(l2)
    e2.next = l1.next.next.next.next
    c = lists_overlap(l1, l2)
    assert c is not None
    assert 4 == c.x

    # one has a cycle
    l1 = from_list(range(10))
    l2 = from_list(range(10, 20))
    last_node(l2).next = l2.next.next
    assert lists_overlap(l1, l2) is None

    # overlap in the tail
    l1 = from_list(range(10))
    last_node(l1).next = l1.next.next.next.next
    l2 = from_list(range(20, 30))
    last_node(l2).next = l1.next.next
    c = lists_overlap(l1, l2)
    assert c is not None
    assert 2 == c.x

    # overlap in the cycle
    l1 = from_list(range(10))
    e1 = last_node(l1)
    e1.next = l1.next.next.next.next
    l2 = from_list(range(30, 37))
    last_node(l2).next = e1
    c = lists_overlap(l1, l2)
    assert c is not None
    assert 9 == c.x

    # disjoint cyclic
    l1 = from_list(range(10))
    last_node(l1).next = l1.next.next
    l2 = from_list(range(10, 33))
    last_node(l2).next = l2.next.next.next.next
    assert lists_overlap(l1, l2) is None

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
