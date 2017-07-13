import random


def randbit():
    """Returns a random bit."""
    return random.randrange(2)


def randrange(a, b):
    """Returns a random integer in the range [a, b]."""
    return a + randupto(b-a+1)


def randnbits(n):
    x = 0
    for i in range(n):
        x = (x << 1) | randbit()
    return x


def randupto(m):
    """Returns a random integer in [0, m)."""
    n = 0
    while 1<<n < m:
        n += 1
    while True:
        x = randnbits(n)
        if x < m:
            return x


def test(m=3, trials=10000):
    counts = [0]*m
    for trial in range(trials):
        x = randupto(m)
        assert x >= 0
        assert x < m
        counts[x] += 1
    print counts


def main():
    test()


if __name__ == '__main__':
    main()
