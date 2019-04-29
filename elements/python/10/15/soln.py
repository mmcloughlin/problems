class Node(object):
    def __init__(self, v, l=None, r=None):
        self.value = v
        self.left = l
        self.right = r
        self.next = None

    def __str__(self):
        return '({} {} {})'.format(
                self.left, self.value, self.right,
                ).replace('None', '')


def perfect(levels, x=1):
    """
    Return a perfect binary tree levels deep.
    """
    if levels == 0:
        return None
    return Node(
            x,
            l=perfect(levels-1, x=2*x),
            r=perfect(levels-1, x=2*x+1),
            )


def connect_right_adjacent_nodes(r):
    level = [r]
    while level[0] is not None:
        for i in xrange(len(level)-1):
            level[i].next = level[i+1]
        below = []
        for n in level:
            below.append(n.left)
            below.append(n.right)
        level = below


def main():
    t = perfect(4)
    print t

    connect_right_adjacent_nodes(t)

    l = t
    while l is not None:
        n = l
        while n is not None:
            print n.value,
            n = n.next
        print
        l = l.left


if __name__ == '__main__':
    main()
