import sys


def cat_cmp(x, y):
    if x == y:
        return 0
    for a, b in zip(x, y):
        r = cmp(a, b)
        if r != 0:
            return r
    m = len(x)
    n = len(y)
    if m < n:
        return cat_cmp(x, y[m:])
    else:
        return cat_cmp(x[n:], y)


def test_cat_cmp():
    cases = [
            ('1', '2', -1),
            ('5', '56', -1),
            ('5', '50', 1),
            ('5', '55', 0),
            ('34563456', '345634564', -1),
            ('34563456', '34563456344', 1),
            ]
    for x, y, expect in cases:
        assert cat_cmp(x, y) == expect
        assert cat_cmp(y, x) == -expect


def cat_range(numbers):
    sorted_numbers = sorted(numbers, cmp=cat_cmp)
    mn = ''.join(sorted_numbers)
    mx = ''.join(reversed(sorted_numbers))
    return mn, mx


def test_cat_range():
    cases = [
            ('5 56 50', '50556 56550'),
            ('79 82 34 83 69', '3469798283 8382796934'),
            ('420 34 19 71 341', '193413442071 714203434119'),
            ('17 32 91 7 46', '173246791 917463217'),
            ]
    for numbers, results in cases:
        mn, mx = cat_range(numbers.split())
        expect_mn, expect_mx = results.split()
        assert mn == expect_mn
        assert mx == expect_mx


def test():
    test_cat_cmp()
    test_cat_range()


def main():
    test()

    numbers = sys.argv[1:]
    mn, mx = cat_range(numbers)
    print mn, mx


if __name__ == '__main__':
    main()
