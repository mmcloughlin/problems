def intsqrt(x):
    """
    Return the largest integer whose square is less than or equal to x.
    """
    if x == 0:
        return 0
    # invariant: l^2 <= x and x < r^2
    l = 1
    r = x
    while r-l > 1:
        m = l + (r-l)//2
        s = m*m
        if s <= x:
            l = m
        else:
            r = m
    return l


def test():
    assert intsqrt(16) == 4
    assert intsqrt(300) == 17
    assert intsqrt(1) == 1
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
