class NestedIterator(object):

    def __init__(self, nestedList):
        """
        Initialize your data structure here.
        :type nestedList: List[NestedInteger]
        """
        self.stack = [nestedList]

    def next(self):
        """
        :rtype: int
        """
        x = self.stack[-1][0].getInteger()
        self.stack[-1] = self.stack[-1][1:]
        return x

    def hasNext(self):
        """
        :rtype: bool
        """
        while len(self.stack) > 0:
            if len(self.stack[-1]) == 0:
                self.stack = self.stack[:-1]
                continue
            if self.stack[-1][0].isInteger():
                break
            l = self.stack[-1][0].getList()
            self.stack[-1] = self.stack[-1][1:]
            self.stack.append(l)
        return len(self.stack) > 0

# """
# This is the interface that allows for creating nested lists.
# You should not implement it, or speculate about its implementation
# """
# class NestedInteger(object):
#    def isInteger(self):
#        """
#        @return True if this NestedInteger holds a single integer, rather than a nested list.
#        :rtype bool
#        """
#
#    def getInteger(self):
#        """
#        @return the single integer that this NestedInteger holds, if it holds a single integer
#        Return None if this NestedInteger holds a nested list
#        :rtype int
#        """
#
#    def getList(self):
#        """
#        @return the nested list that this NestedInteger holds, if it holds a nested list
#        Return None if this NestedInteger holds a single integer
#        :rtype List[NestedInteger]
#        """
