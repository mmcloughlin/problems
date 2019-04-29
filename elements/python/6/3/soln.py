import random


def add(a, d, i=0):
    """
    Add d*10^i to the integer a represented as an array of decimal digits.
    """
    assert 0 <= d < 10
    n = len(a)
    if i+1 > n:
        a.extend([0]*(i+1-n))
        n += i+1-n
    carry = d
    while i < len(a) and carry > 0:
        x = a[i] + carry
        a[i] = x % 10
        carry = x // 10
        i += 1
    if carry > 0:
        a.append(carry)


def mul(A, B):
    m = []
    for i, a in enumerate(A):
        for j, b in enumerate(B):
            k = i+j
            prod = a*b
            lo, hi = prod % 10, prod // 10
            add(m, lo, k)
            add(m, hi, k+1)
    return m


def int_to_digits(x):
    return list(reversed(list(map(int, str(x)))))


def digits_to_int(a):
    return sum([d*10**i for i, d in enumerate(a)])


def test_add(trials=1000):
    for trial in range(trials):
        x = random.randrange(100000)
        a = int_to_digits(x)
        add(a, 3, 2)
        assert digits_to_int(a) == x + 300


def test_mul(trials=1000):
    for trial in range(trials):
        x = random.randrange(100000)
        y = random.randrange(100000)
        A = int_to_digits(x)
        B = int_to_digits(y)
        assert digits_to_int(mul(A, B)) == x*y


def test():
    test_add(10000)
    test_mul(100)
    print 'pass'


if __name__ == '__main__':
    test()
