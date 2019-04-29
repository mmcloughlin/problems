import random
import collections


def weighted_random_choice(weights):
    r = random.randrange(sum(weights))
    for i, w in enumerate(weights):
        if r < w:
            return i
        r -= w


def counts(weights, trials=1000):
    c = collections.Counter()
    for _ in xrange(trials):
        i = weighted_random_choice(weights)
        c[i] += 1
    return c


def main():
    print counts([2, 1, 1])


if __name__ == '__main__':
    main()
