class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def reverse_sublists(n, k):
    """
    Reverse sublists of length k.
    """
    head = Node(0, n)
    pre = head
    while n is not None:
        # start of the sublist
        s = n
        # find the element after the sublist
        post = n
        for _ in xrange(k):
            if post is None:
                return head.next
            post = post.next

        # reverse the sublist
        prev = post
        m = s
        for _ in xrange(k):
            nxt = m.next
            m.next = prev
            prev = m
            m = nxt

        pre.next = prev
        pre = s
        n = m

    return head.next


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
    n = reverse_sublists(n, 3)
    display(n)


def main():
    test()


if __name__ == '__main__':
    main()
