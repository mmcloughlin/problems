def gcd(a, b):
    if b == 0:
        return a
    return gcd(b, a%b)


def test():
    assert gcd(21, 6) == 3
    assert gcd(99, 11) == 11
    assert gcd(11, 21) == 1
    print 'pass'


if __name__ == '__main__':
    test()
