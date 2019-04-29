def incr(a):
    carry = 1
    i = len(a)-1
    while i >= 0 and carry:
        x = a[i] + carry
        a[i] = x%10
        carry = x/10
        i -= 1
    if carry == 1:
        a.insert(0, 1)


def digits_array(x):
    return list(map(int, str(x)))


def test():
    cases = [129, 999, 23452341, 123999]
    for x in cases:
        a = digits_array(x)
        incr(a)
        assert a == digits_array(x+1)
    print 'pass'


if __name__ == '__main__':
    test()
