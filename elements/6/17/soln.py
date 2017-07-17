import random
import collections


def perm(n):
    pi = range(n)
    for i in xrange(n-1, -1, -1):
        j = random.randrange(0, i+1)
        pi[i], pi[j] = pi[j], pi[i]
    return pi


def test(n, trials=1000):
    counts = collections.Counter()
    for i in xrange(trials):
        pi = perm(n)
        counts[str(pi)] += 1
    for pi, count in counts.items():
        print count, pi


def main():
    test(3, 10000)


if __name__ == '__main__':
    main()
