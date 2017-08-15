class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r

    def __str__(self):
        return '({} {} {})'.format(
                self.left, self.value, self.right,
                ).replace('None', '')


def reconstruct_from_traversals(inorder, preorder):
    if len(inorder) == 0:
        return None
    root_value = preorder[0]
    # index() is inefficient, could use map to limit to one linear traversal
    n = inorder.index(root_value)
    left = reconstruct_from_traversals(inorder[:n], preorder[1:n+1])
    right = reconstruct_from_traversals(inorder[n+1:], preorder[n+1:])
    return Node(root_value, l=left, r=right)


def main():
    r = reconstruct_from_traversals('fbaehcdig', 'hbfeacdgi')
    print r


if __name__ == '__main__':
    main()
