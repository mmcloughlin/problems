import tree


def leftmost(n):
    """
    Return the leftmost item in a binary tree.
    """
    while n.left is not None:
        n = n.left
    return n


def successor(n):
    """
    Return the successor of the binary tree node n.
    """
    if n.right is not None:
        return leftmost(n.right)

    while n.parent is not None and n.parent.right == n:
        n = n.parent

    return n.parent


def test():
    nodes = tree.example()
    t = nodes['a']
    assert successor(nodes['g']).content == 'a'
    assert successor(nodes['a']).content == 'j'
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
