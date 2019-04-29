import random


def add(x, y):
    s = 0
    c = 0
    i = 0
    while x != 0 or y != 0 or c != 0:
        b = (x ^ y ^ c) & 1
        c = ((x & y) | (x & c) | (y & c)) & 1
        s |= (b << i)
        x >>= 1
        y >>= 1
        i += 1
    return s


def test_add():
    n = 64
    for trial in range(1000):
        x = random.getrandbits(n)
        y = random.getrandbits(n)
        assert add(x, y) == x+y


def mul(x, y):
    m = 0
    while y != 0:
        if y&1 == 1:
            m = add(m, x)
        y >>= 1
        x <<= 1
    return m


def test_mul():
    n = 32
    for trial in range(1000):
        x = random.getrandbits(n)
        y = random.getrandbits(n)
        assert mul(x, y) == x*y


def main():
    test_add()
    test_mul()
    print 'pass'


if __name__ == '__main__':
    main()
