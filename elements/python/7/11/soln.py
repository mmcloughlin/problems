import random


def elias_encode(a):
    """
    Elias gamma encode the integer a.
    """
    e = ''
    n = 0
    while a != 0:
        e = str(a&1) + e
        a >>= 1
        n += 1
    return '0'*(n-1) + e


def elias_encode_seq(A):
    """
    Elias gamma encode the integer sequence A.
    """
    return ''.join(map(elias_encode, A))


def elias_decode_seq(e):
    """
    Decode an integer sequence represented in Elias gamma encoding.
    """
    A = []
    i = 0
    while i < len(e):
        # how many zeros?
        n = 1
        while e[i] == '0':
            n += 1
            i += 1
        # read n bits
        a = 0
        while n > 0:
            a <<= 1
            a |= ord(e[i])&1
            i += 1
            n -= 1
        A.append(a)
    return A


def test(trials=1000):
    assert elias_encode(1) == '1'
    assert elias_encode(2) == '010'
    assert elias_encode(13) == '0001101'

    for _ in xrange(trials):
        A = [random.randrange(1, 1<<10) for _ in xrange(10)]
        enc = elias_encode_seq(A)
        dec = elias_decode_seq(enc)
        assert dec == A

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
