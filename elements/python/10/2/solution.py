import tree


def _imbalanced(node, k):
    if node is None:
        return None, 0

    imbalanced_node, num_left = _imbalanced(node.left, k)
    if imbalanced_node:
        return imbalanced_node, 0

    imbalanced_node, num_right = _imbalanced(node.right, k)
    if imbalanced_node:
        return imbalanced_node, 0

    if abs(num_left - num_right) > k:
        return node, 0

    return None, num_left + num_right + 1


def imbalanced(node, k):
    n, _ = _imbalanced(node, k)
    return n


def main():
    t = tree.example()
    n = imbalanced(t, 3)
    if n is not None:
        print n.content
    else:
        print 'no 3-imbalanced node found'


if __name__ == '__main__':
    main()
