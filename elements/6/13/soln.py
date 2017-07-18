import random


def naive(a, pi):
    b = [None]*len(a)
    for i, j in enumerate(pi):
        b[j] = a[i]
    return b


def permute(a, pi):
    n = len(a)
    for i in xrange(n):
        t = a[i]
        j = pi[i]
        while j > i:
            a[j], t = t, a[j]
            pi[j], j = i, pi[j]
        a[i] = t


def random_perm(n):
    pi = range(n)
    for i in xrange(n):
        j = random.randrange(i, n)
        pi[i], pi[j] = pi[j], pi[i]
    return pi


def test(n=100, trials=1000):
    for trial in xrange(trials):
        a = [random.randrange(10000) for _ in xrange(n)]
        pi = random_perm(n)
        expect = naive(a, pi)
        permute(a, pi)
        assert expect == a
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
