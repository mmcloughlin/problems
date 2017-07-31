class Node(object):
    def __init__(self, x, nxt):
        self.x = x
        self.next = nxt


def is_palindrome(h):
    # find the middle
    m = h
    n = h
    while n is not None:
        m = m.next
        n = n.next
        if n is None:
            break
        n = n.next

    # reverse to the end
    prev = None
    while m is not None:
        nxt = m.next
        m.next = prev
        prev = m
        m = nxt
    end = prev

    # check palidromicity (and un-reverse the list)
    pal = True
    s = h
    e = end
    prev = None
    while e is not None:
        if s.x != e.x:
            pal = False
        s = s.next

        nxt = e.next
        e.next = prev
        prev = e
        e = nxt

    return pal


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
    # odd length palindrome
    pal = from_list([1,2,3,4,5,4,3,2,1])
    display(pal)
    assert is_palindrome(pal)
    display(pal)

    # even length
    pal = from_list([1,2,3,3,2,1])
    assert is_palindrome(pal)

    # not palindromes
    assert not is_palindrome(from_list([1,2,3,4,2,1]))
    assert not is_palindrome(from_list([1,2,4,3,2,1]))
    assert not is_palindrome(from_list(range(10)))


def main():
    test()


if __name__ == '__main__':
    main()

