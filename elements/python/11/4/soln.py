import heapq
import random


def largest(X, k):
    """
    Return the k largest elements from X.
    """
    # maintain a min-heap of size k containing the largest elements so far
    h = []
    for x in X:
        if len(h) < k:
            heapq.heappush(h, x)
        elif x > h[0]:
            heapq.heapreplace(h, x)
    return h


def smallest(X, k):
    """
    Return the k smallest elements from X.
    """
    h = []
    for x in X:
        if len(h) < k:
            heapq.heappush(h, -x)
        elif -x > h[0]:
            heapq.heapreplace(h, -x)
    return [-x for x in h]


def test():
    X = range(100)
    random.shuffle(X)

    big = largest(X, 10)
    assert sorted(big) == range(90, 100)

    small = smallest(X, 10)
    assert sorted(small) == range(10)

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
