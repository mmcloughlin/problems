import heapq
import random


def merge(lists):
    heapq.heapify(lists)
    m = []
    while len(lists) > 0:
        l = heapq.heappop(lists)
        if len(l) == 0:
            continue
        m.append(l.pop(0))
        heapq.heappush(lists, l)
    return m


def test(n, k):
    lists = [[] for _ in xrange(k)]
    for i in xrange(n):
        lists[random.randrange(k)].append(i)
    m = merge(lists)
    assert m == range(n)
    print 'pass'


def main():
    test(100000, 50)


if __name__ == '__main__':
    main()
