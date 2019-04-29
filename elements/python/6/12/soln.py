import math
import random


def primes_less_than_seive(n):
    """
    Return all the primes less than n.
    """
    is_prime = [True]*n
    is_prime[0], is_prime[1] = False, False
    upper = int(math.sqrt(n))
    for p in range(2, upper+1):
        if not p:
            continue
        for k in range(2*p, n, p):
            is_prime[k] = False
    return [p for p, flag in enumerate(is_prime) if flag]


def is_prime(n):
    return 2 == len([d for d in xrange(1, n+1) if n%d == 0])


def primes_less_than_naive(n):
    return filter(is_prime, xrange(2, n))


def test(trials=10):
    for trial in xrange(trials):
        n = random.randrange(10, 10000)
        expect = primes_less_than_naive(n)
        assert expect == primes_less_than_seive(n)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
