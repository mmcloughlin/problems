import random
import collections


def sample(a, k):
    """
    sample generates a sample of k elements from a.

    After calling this function the sample is the first k elements of a.
    """
    n = len(a)
    if k > n:
        raise RuntimeError('sample size too large')
    for i in xrange(k):
        j = random.randrange(i, n)
        a[i], a[j] = a[j], a[i]


def test(n, k, trials=1000):
    a = range(n)
    counts = collections.Counter()
    for trial in xrange(trials):
        sample(a, k)
        counts[str(sorted(a[:k]))] += 1
    for s, count in counts.items():
        print count, s


def main():
    test(4, 3, 10000)


if __name__ == '__main__':
    main()
