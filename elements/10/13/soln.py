class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r

    def __str__(self):
        return '({} {} {})'.format(
                self.left, self.value, self.right,
                ).replace('None', '')


def reconstruct(preorder):
    stack = []
    for n in reversed(preorder):
        if n is None:
            stack.append(n)
            continue
        left = stack.pop(-1)
        right = stack.pop(-1)
        node = Node(n, l=left, r=right)
        stack.append(node)
    return stack[0]


def main():
    traversal = [
            'H', 'B', 'F', None, None, 'E', 'A', None, None, None,
            'C', None, 'D', None, 'G', 'I', None, None, None,
            ]
    print reconstruct(traversal)


if __name__ == '__main__':
    main()
