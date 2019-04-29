def bsearch(A, pred):
    """
    Assume the boolean function pred returns all False and then all True for
    values in A. Return the index of the first True, or len(A) if that does
    not exist.
    """
    # invariant: last False lies in [l, r) and A[l] is False
    if pred(A[0]):
        return 0
    l = 0
    r = len(A)
    while r-l > 1:
        m = l + (r-l)//2
        result = pred(A[m])
        if result:
            r = m
        else:
            l = m
    return l+1


def first_occurrence(A, k):
    i = bsearch(A, lambda a: a >= k)
    if i >= len(A) or A[i] != k:
        return None
    return i


def test():
    A = [-14, -10, 2, 108, 108, 243, 285, 285, 285, 401]
    assert first_occurrence(A, 108) == 3
    assert first_occurrence(A, 285) == 6
    assert first_occurrence(A, 150) is None
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
