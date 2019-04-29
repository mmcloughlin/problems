import heapq


def sort_almost_sorted(A, k):
    h = list(A[:k+1])
    heapq.heapify(h)

    result = []
    for a in A[k+1:]:
        x = heapq.heapreplace(h, a)
        result.append(x)

    while len(h) > 0:
        result.append(heapq.heappop(h))

    return result


def test():
    A = [3, -1, 2, 6, 4, 5, 8]
    assert sorted(A) == sort_almost_sorted(A, 2)


def main():
    test()
    print 'pass'


if __name__ == '__main__':
    main()
