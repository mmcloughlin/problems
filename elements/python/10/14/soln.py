class ListNode(object):
    def __init__(self, v, n=None):
        self.value = v
        self.next = n


class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r


def leaf_nodes(t):
    f, _ = leaf_nodes_helper(t)
    return f


def leaf_nodes_helper(t):
    """
    Forms a linked list from the leaf nodes in t. Returns the first and last
    nodes.
    """
    if t is None:
        return None
    if t.left is None and t.right is None:
        n = ListNode(t.value)
        return n, n
    if t.left is None:
        return leaf_nodes_helper(t.right)
    if t.right is None:
        return leaf_nodes_helper(t.left)
    lf, ll = leaf_nodes_helper(t.left)
    rf, rl = leaf_nodes_helper(t.right)
    ll.next = rf
    return lf, rl


def linked_list_to_list(n):
    l = []
    while n is not None:
        l.append(n.value)
        n = n.next
    return l


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
    l = leaf_nodes(r)
    expect = ['d', 'e', 'h', 'm', 'n', 'p']
    assert expect == linked_list_to_list(l)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
