import random
import collections


def sample(n, k):
    """
    Compute a sample from {0, 1, ..., n-1} of size k.
    """
    a = dict()
    for i in xrange(k):
        j = random.randrange(i, n)
        a[i], a[j] = a.get(j, j), a.get(i, i)
    return [a[i] for i in xrange(k)]


def count_samples(n, k, trials=1000):
    counts = collections.Counter()
    for trial in xrange(trials):
        s = sample(n, k)
        counts[str(sorted(s))] += 1
    for s, count in counts.items():
        print count, s


def main():
    # show a simple case still works the same
    count_samples(4, 3)
    # show we don't consume massive memory for large n
    print sample(100000000000, 10)


if __name__ == '__main__':
    main()
