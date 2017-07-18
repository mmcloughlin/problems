import random


def product_excluding(A, i):
    P = 1
    for j, a in enumerate(A):
        if j != i:
            P *= a
    return P


def max_prod_naive(A):
    return max([product_excluding(A, i) for i in xrange(len(A))])


def running_product(A):
    P = 1
    left = []
    for a in A:
        left.append(P)
        P *= a
    return left


def max_prod_running_products(A):
    left = running_product(A)
    right = reversed(running_product(reversed(A)))
    return max([l*r for l, r in zip(left, right)])


def test(n=50, trials=1000):
    for trial in xrange(trials):
        A = [random.randrange(-n, n) for _ in xrange(n)]
        expect = max_prod_naive(A)
        assert expect == max_prod_running_products(A)
    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
