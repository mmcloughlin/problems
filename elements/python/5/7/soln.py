import random


def pow(x, y):
    p = 1
    while y != 0:
        if y&1 == 1:
            p *= x
        x = x*x
        y >>= 1
    return p


def test(trials=100):
    for trial in range(trials):
        x = random.getrandbits(32)
        y = random.getrandbits(16)
        assert pow(x, y) == x**y


def main():
    test()
    print 'pass'


if __name__ == '__main__':
    main()
