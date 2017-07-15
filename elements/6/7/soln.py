import random


def first_missing_positive(a):
    """
    O(n) time but reuses the array for state.
    """
    n = len(a)
    i = 0
    while i < n:
        j = a[i]
        while 0 < j <= n and a[j-1] != j:
            a[j-1], j = j, a[j-1]
        i += 1
    for i, x in enumerate(a):
        if x != i+1:
            return i+1


def first_missing_positive_naive(a):
    """
    O(n) time and O(n) space.
    """
    n = len(a)
    flag = [False]*n
    for x in a:
        if x < 1 or x > n:
            continue
        flag[x-1] = True
    for i, f in enumerate(flag):
        if not f:
            return i+1


def test_example():
    a = [3, 5, 4, -1, 5, 1, -1]
    assert 2 == first_missing_positive(a)


def test_cmp(trials=1000):
    n = 100
    for trial in range(trials):
        a = [random.randrange(-n/4, n+2) for _ in range(n)]
        expect = first_missing_positive_naive(a)
        got = first_missing_positive(a)
        assert expect == got


def test():
    test_example()
    test_cmp()
    print 'pass'


if __name__ == '__main__':
    test()
