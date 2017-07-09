import random


def div(x, y):
    """
    Compute integer division x//y.
    """
    # find largest shift less than x
    i = 0
    s = y
    while s < x:
        s <<= 1
        i += 1
    s >>= 1
    i -= 1

    d = 0
    rem = x
    while i >= 0:
        if s < rem:
            rem -= s
            d += 1<<i
        i -= 1
        s >>= 1

    return d


def test(trials=1000):
    for trial in range(trials):
        x = random.getrandbits(64)
        y = random.getrandbits(48)
        assert div(x, y) == x//y


def main():
    test()
    print 'pass'


if __name__ == '__main__':
    main()
