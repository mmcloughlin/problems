def bsearch(n, pred):
    """
    Given a boolean function pred that takes index arguments in [0, n).
    Assume the boolean function pred returns all False and then all True for
    values. Return the index of the first True, or n if that does not exist.
    """
    # invariant: last False lies in [l, r) and pred(l) is False
    if pred(0):
        return 0
    l = 0
    r = n
    while r-l > 1:
        m = l + (r-l)//2
        result = pred(m)
        if result:
            r = m
        else:
            l = m
    return l+1


def equal_to_index(A):
    i = bsearch(len(A), lambda i: A[i] >= i)
    return i if A[i] == i else None


def test():
    A = [-2, 0, 2, 3, 6, 7, 9]
    assert equal_to_index(A) == 2
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
