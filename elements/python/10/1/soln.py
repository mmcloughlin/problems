class Node(object):
    def __init__(self, left=None, right=None):
        self.left = left
        self.right = right


def is_balanced(n):
    mn, mx = depth_range(n)
    return mx - mn <= 1


def depth_range(n):
    if n is None:
        return 0, 0
    mnl, mxl = depth_range(n.left)
    mnr, mxr = depth_range(n.right)
    mn = min(mnl, mnr) + 1
    mx = max(mxl, mxr) + 1
    return mn, mx


def test():
    y = Node(
            left=Node(left=Node(), right=Node()),
            right=Node(left=Node())
            )
    assert is_balanced(y)

    n = Node(
            left=Node(left=Node(), right=Node()),
            right=Node(left=Node(right=Node())),
            )
    assert not is_balanced(n)

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
