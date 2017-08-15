class Node(object):
    left = None
    content = None
    right = None
    parent = None

    def __init__(self, content):
        self.content = content

    def __str__(self):
        return '({left} {content} {right})'.format(
                left=self.left,
                content=self.content,
                right=self.right,
                )


def example():
    """
    Constructs example binary tree on page 156 of "Elements of
    Programming Interviews".
    """
    edges = {
            'a': ('b',  'i'),
            'b': ('c',  'f'),
            'c': ('d',  'e'),
            'f': (None, 'g'),
            'g': ('h',  None),
            'i': ('j',  'o'),
            'j': (None, 'k'),
            'k': ('l',  'n'),
            'l': (None, 'm'),
            'o': (None, 'p'),
            }
    nodes = {}
    for parent, children in edges.items():
        l, r = children
        parent_node = nodes.setdefault(parent, Node(parent))
        if l is not None:
            parent_node.left = nodes.setdefault(l, Node(l))
            parent_node.left.parent = parent_node
        if r is not None:
            parent_node.right = nodes.setdefault(r, Node(r))
            parent_node.right.parent = parent_node
    return nodes
