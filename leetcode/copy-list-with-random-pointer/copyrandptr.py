# Definition for singly-linked list with a random pointer.
class RandomListNode(object):
    def __init__(self, x):
        self.label = x
        self.next = None
        self.random = None

    def dump(self):
        s = self.label
        if self.random is not None:
            s += ' rand={}'.format(self.random.label)
        print(s)
        if self.next is not None:
            self.next.dump()


class Solution(object):
    def copyRandomList(self, head):
        if head is None:
            return None

        # label->node map
        label_to_node = {}

        node = head
        pre = None

        while node is not None:
            node_copy = label_to_node.setdefault(
                    node.label,
                    RandomListNode(node.label),
                )

            if pre is not None:
                pre.next = node_copy

            if node.random is not None:
                node_copy.random = label_to_node.setdefault(
                    node.random.label,
                    RandomListNode(node.random.label),
                )

            node = node.next
            pre = node_copy

        return label_to_node[head.label]


def test():
    a = RandomListNode('a')
    b = RandomListNode('b')
    c = RandomListNode('c')
    d = RandomListNode('d')

    a.next = b
    b.next = c
    c.next = d

    a.random = d
    b.random = c
    d.random = a

    a.dump()

    cpy = Solution().copyRandomList(a)

    cpy.dump()


if __name__ == '__main__':
    test()
