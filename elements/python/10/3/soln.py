class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r


def is_symmetric(t):
    """
    Returns whether the binary tree t is symmetric.
    """
    return t is None or equal_mirror(t.left, t.right)


def equal_mirror(t, s):
    """
    Returns whether t is the mirror image of s.
    """
    if t is None and s is None:
        return True
    if t is None or s is None:
        return False
    if t.value != s.value:
        return False
    return equal_mirror(t.left, s.right) and equal_mirror(t.right, s.left)


def test():
    a = Node(314,
            l=Node(6, r=Node(2, r=Node(3))),
            r=Node(6, l=Node(2, l=Node(3))),
            )
    assert is_symmetric(a)

    b = Node(314,
            l=Node(6, r=Node(561, r=Node(3))),
            r=Node(6, l=Node(2, l=Node(1))),
            )
    assert not is_symmetric(b)

    c = Node(314,
            l=Node(6, r=Node(561, r=Node(3))),
            r=Node(6, l=Node(561)),
            )
    assert not is_symmetric(c)

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
