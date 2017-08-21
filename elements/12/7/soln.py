import random
import math


def sqrt(x, tolerance=0.0000001):
    """
    Find a value r with r*r within tolerance of x.
    """
    l = 0.0
    r = float(x)
    while r-l > tolerance:
        m = l + 0.5*(r-l)
        s = m*m
        if s < x:
            l = m
        else:
            r = m
    return l



def test(trials=1000):
    tolerance = 0.000000001
    for _ in xrange(trials):
        x = random.uniform(1, 1000000)
        expect = math.sqrt(x)
        got = sqrt(x, tolerance)
        assert abs(got - expect) <= tolerance
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
