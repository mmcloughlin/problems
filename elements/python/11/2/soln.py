import heapq
import random


def decompose(A):
    """
    decompose A into a list of sorted lists that contain the same elements.
    """
    if len(A) == 0:
        return [[]]
    acc = [A[0]]
    lists = [acc]
    order = -1
    for a in A[1:]:
        if cmp(acc[-1], a) != order:
            if order == 1:
                acc.reverse()
            acc = []
            lists.append(acc)
            order = -order
        acc.append(a)
    if order == 1:
        acc.reverse()
    return lists


def merge(lists):
    """
    merge the given sorted lists.
    """
    heapq.heapify(lists)
    m = []
    while len(lists) > 0:
        l = heapq.heappop(lists)
        if len(l) == 0:
            continue
        m.append(l.pop(0))
        heapq.heappush(lists, l)
    return m


def incdecsort(A):
    """
    Sort the increasing-decreasing array A.
    """
    return merge(decompose(A))


def test(trials=1000, n=100, m=100000000):
    for _ in xrange(trials):
        A = [random.randrange(m) for _ in xrange(n)]
        assert sorted(A) == incdecsort(A)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
