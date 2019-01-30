# Definition for binary tree with next pointer.
# class TreeLinkNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
#         self.next = None

class Solution:
    def connect(self, root):
        if root is None:
            return
        level = [root]
        while len(level) > 0:
            # connect next right
            for i in range(len(level)-1):
                level[i].next = level[i+1]
            # build next level
            nxt = []
            for node in level:
                nxt.extend([node.left, node.right])
            level = [node for node in nxt if node is not None]

