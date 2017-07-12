def is_palendrome(x):
    if x < 0:
        return False
    if x < 10:
        return True
    n = 0
    while 10**n < x:
        n += 1
    n -= 1
    leading = x / 10**n
    trailing = x % 10
    if leading != trailing:
        return False
    remaining = (x - leading*10**n)/10
    return is_palendrome(remaining)


def test():
    assert all([is_palendrome(n) for n in [0, 1, 7, 11, 121, 333, 2147447412]])
    assert all([not is_palendrome(n) for n in [-1, 12, 100, 2147483647]])


def main():
    test()
    print 'pass'


if __name__ == '__main__':
    main()
