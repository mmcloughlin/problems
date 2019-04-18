import struct


class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Codec:
    def serialize(self, root):
        """Encodes a tree to a single string.

        :type root: TreeNode
        :rtype: str
        """
        if root is None:
            return struct.pack('i', 0)
        l = self.serialize(root.left)
        r = self.serialize(root.right)
        return struct.pack('i', len(l)) + struct.pack('i', root.val) + l + r

    def deserialize(self, data):
        """Decodes your encoded data to tree.

        :type data: str
        :rtype: TreeNode
        """
        len_l = struct.unpack('i', data[:4])[0]
        if len_l == 0:
            return None

        data = data[4:]
        # val
        val = struct.unpack('i', data[:4])[0]
        data = data[4:]
        # left
        l = self.deserialize(data[:len_l])
        data = data[len_l:]
        # right
        r = self.deserialize(data)

        # build the node
        node = TreeNode(val)
        node.left = l
        node.right = r
        return node


# Your Codec object will be instantiated and called as such:
# codec = Codec()
# codec.deserialize(codec.serialize(root))

def node(x, l=None, r=None):
    n = TreeNode(x)
    n.left = l
    n.right = r
    return n


def trees_equal(a, b):
    if a is None and b is None:
        return True
    if a is None or b is None:
        return False
    return (
        a.val == b.val and
        trees_equal(a.left, b.left) and
        trees_equal(a.right, b.right)
    )


def test():
    root = node(
        1,
        node(2),
        node(
            3,
            node(4),
            node(5),
        ),
    )
    codec = Codec()
    got = codec.deserialize(codec.serialize(root))
    assert trees_equal(root, got)


if __name__ == '__main__':
    test()
