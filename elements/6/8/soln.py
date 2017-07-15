def max_difference(a):
    n = len(a)
    if n == 0:
        return None

    max_after = a[n-1]
    max_delta = 0
    i = n-2
    while i >= 0:
        max_after = max(a[i], max_after)
        max_delta = max(max_after - a[i], max_delta)
        i -= 1
    return max_delta


def test():
    a = [0,1,2,3,2,3,4,5,6,-3,4,5]
    assert 8 == max_difference(a)
    print 'pass'


if __name__ == '__main__':
    test()
