import random


def atoi(s):
    """
    Convert s to an integer.
    """
    # determine sign
    sgn = 1
    if s[0] == '-':
        sgn = -1
        s = s[1:]
    # accumulate digits
    i = 0
    for d in s:
        x = ord(d) - ord('0')
        if x < 0 or x > 9:
            raise ValueError('bad digit')
        i = 10*i + x
    return sgn*i


DIGITS = '0123456789'

def itoa(i):
    """
    Convert i to a string.
    """
    if i == 0:
        return '0'
    sgn = ''
    if i < 0:
        sgn = '-'
        i = -i
    s = ''
    while i != 0:
        s = DIGITS[i%10] + s
        i /= 10
    return sgn + s



def test(trials=1000):
    assert itoa(103) == '103'
    assert itoa(0) == '0'
    assert itoa(-7) == '-7'
    assert atoi('23') == 23
    assert atoi('0') == 0
    assert atoi('-19') == -19

    for _ in xrange(trials):
        n = random.randrange(-10000, 10000)
        assert atoi(itoa(n)) == n

    print 'pass'


def main():
    test()


if __name__ == '__main__':
    main()
