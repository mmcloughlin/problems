import tree


def depth_order(n):
    nodes = [n]
    i = 0
    while i < len(nodes):
        n = nodes[i]
        if n.left is not None:
            nodes.append(n.left)
        if n.right is not None:
            nodes.append(n.right)
        i += 1
    return [n.content for n in nodes]


def test():
    t = tree.example()
    keys = depth_order(t)
    s = ''.join(keys)
    assert s == 'abicfjodegkphlnm'
    print s


def main():
    test()


if __name__ == '__main__':
    main()
