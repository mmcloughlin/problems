import sys


def cat_cmp(x, y):
    return int(x + y) - int(y + x)


def cat_range(numbers):
    sorted_numbers = sorted(numbers, cmp=cat_cmp)
    mn = ''.join(sorted_numbers)
    mx = ''.join(reversed(sorted_numbers))
    return mn, mx


def test():
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


def main():
    numbers = sys.argv[1:]
    mn, mx = cat_range(numbers)
    print mn, mx


if __name__ == '__main__':
    main()
