def find_shift_amount(A):
    """
    find_shift_amount assumes A is a sorted array that has been cyclically
    shifted, and finds the shift amount.
    """
    l = 0
    r = len(A)-1
    if A[l] < A[r]:
        return 0
    while r-l > 1 and A[l] >= A[r]:
        m = l + (r-l)//2
        if A[l] > A[m]:
            r = m
            continue
        if A[m] > A[r]:
            l = m
            continue
        return 0
    return r


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


def search_cyclic_shift(A, t):
    """
    Search for k in the cyclically shifted array A. Return the index if found,
    otherwise None.
    """
    k = find_shift_amount(A)
    n = len(A)
    print cyclic_shift(A, k)
    i = bsearch(n, lambda i: A[(i+k)%n] >= t)
    if i >= n:
        return None
    idx = (i+k)%n
    return idx if A[idx] == t else None


def cyclic_shift(A, k):
    """
    Return the cyclic shift of A left by k.
    """
    n = len(A)
    return [A[(i + k)%n] for i in xrange(n)]


def test():
    A = [-14, -10, 2, 108, 108, 243, 285, 285, 285, 401]
    S = cyclic_shift(A, 3)
    assert 2 == search_cyclic_shift(S, 243)


def main():
    test()


if __name__ == '__main__':
    main()
