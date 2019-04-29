import heapq


def largest_max_heap(h, k):
    """
    Returns the k largest items from h, which has the heap property.
    """
    aux = [(-h[0], 0)]
    largest = []
    for _ in xrange(k):
        x, i = heapq.heappop(aux)
        largest.append(-x)
        for j in range(2*i+1, 2*i+3):
            if j < len(h):
                heapq.heappush(aux, (-h[j], j))
    return largest


def test():
    h = [561, 314, 401, 28, 156, 359, 271, 11, 3]
    got = largest_max_heap(h, 4)
    expect = sorted(h, reverse=True)[:4]
    assert expect == got
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
