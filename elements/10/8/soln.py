class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r
        lnum = self.left.num if self.left is not None else 0
        rnum = self.right.num if self.right is not None else 0
        self.num = lnum + 1 + rnum


def in_order_kth(n, k):
    if k >= n.num:
        return None
    lnum = n.left.num if n.left is not None else 0
    if k < lnum:
        return in_order_kth(n.left, k)
    if k >= lnum+1:
        return in_order_kth(n.right, k-lnum-1)
    return n


def test():
    r = Node('a',
            l=Node('b',
                l=Node('c', l=Node('d'), r=Node('e')),
                r=Node('f', r=Node('g', l=Node('h'))),
                ),
            r=Node('i',
                l=Node('j', r=Node('k', l=Node('l', r=Node('m')), r=Node('n'))),
                r=Node('o', r=Node('p')),
                ),
            )
    inorder = 'dcebfhgajlmkniop'
    for i, ch in enumerate(inorder):
        assert ch == in_order_kth(r, i).value
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
