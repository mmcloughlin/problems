class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


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


def last_node(n):
    while n.next is not None:
        n = n.next
    return n


def test():
    n = from_list(range(10))
    assert find_cycle(n) is None

    init = from_list(range(23))
    cycle = from_list(range(100, 117))

    init_tail = last_node(init)
    cycle_tail = last_node(cycle)
    init_tail.next = cycle
    cycle_tail.next = cycle

    assert 100 == find_cycle(init).x

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
